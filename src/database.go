package qdb

import (
	"context"
	"encoding/base64"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type IDatabase interface {
	Connect()
	Disconnect()
	IsConnected() bool

	CreateSnapshot() *DatabaseSnapshot
	RestoreSnapshot(snapshot *DatabaseSnapshot)

	CreateEntity(entityType, parentId, name string)
	GetEntity(entityId string) *DatabaseEntity
	SetEntity(entityId string, value *DatabaseEntity)
	DeleteEntity(entityId string)

	FindEntities(entityType string) []string
	GetEntityTypes() []string

	EntityExists(entityId string) bool
	FieldExists(fieldName, entityType string) bool

	GetFieldSchemas() []*DatabaseFieldSchema
	GetFieldSchema(fieldName string) *DatabaseFieldSchema
	SetFieldSchema(fieldName string, value *DatabaseFieldSchema)

	GetEntitySchema(entityType string) *DatabaseEntitySchema
	SetEntitySchema(entityType string, value *DatabaseEntitySchema)

	Read(requests []*DatabaseRequest)
	Write(requests []*DatabaseRequest)

	Notify(config *DatabaseNotificationConfig, callback func(*DatabaseNotification)) string
	Unnotify(subscriptionId string)
	ProcessNotifications()
}

type RedisDatabaseConfig struct {
	Address  string
	Password string
}

func (r *DatabaseRequest) FromField(field *DatabaseField) *DatabaseRequest {
	r.Id = field.Id
	r.Field = field.Name
	r.Value = field.Value

	if r.WriteTime == nil {
		r.WriteTime = &Timestamp{Raw: timestamppb.Now()}
	}
	r.WriteTime.Raw = field.WriteTime

	if r.WriterId == nil {
		r.WriterId = &String{Raw: ""}
	}
	r.WriterId.Raw = field.WriterId

	return r
}

func (f *DatabaseField) FromRequest(request *DatabaseRequest) *DatabaseField {
	f.Name = request.Field
	f.Id = request.Id
	f.Value = request.Value

	if request.WriteTime != nil {
		f.WriteTime = request.WriteTime.Raw
	}

	if request.WriterId != nil {
		f.WriterId = request.WriterId.Raw
	}

	return f
}

// schema:entity:<type> -> DatabaseEntitySchema
// schema:field:<name> -> DatabaseFieldSchema
// instance:entity:<entityId> -> DatabaseEntity
// instance:field:<name>:<entityId> -> DatabaseField
// instance:type:<entityType> -> []string{entityId...}
// instance:notification-config:<entityId>:<fieldName> -> []string{subscriptionId...}
// instance:notification-config:<entityType>:<fieldName> -> []string{subscriptionId...}
type RedisDatabaseKeyGenerator struct{}

func (g *RedisDatabaseKeyGenerator) GetEntitySchemaKey(entityType string) string {
	return "schema:entity:" + entityType
}

func (g *RedisDatabaseKeyGenerator) GetFieldSchemaKey(fieldName string) string {
	return "schema:field:" + fieldName
}

func (g *RedisDatabaseKeyGenerator) GetEntityKey(entityId string) string {
	return "instance:entity:" + entityId
}

func (g *RedisDatabaseKeyGenerator) GetFieldKey(fieldName, entityId string) string {
	return "instance:field:" + fieldName + ":" + entityId
}

func (g *RedisDatabaseKeyGenerator) GetEntityTypeKey(entityType string) string {
	return "instance:type:" + entityType
}

func (g *RedisDatabaseKeyGenerator) GetEntityIdNotificationConfigKey(entityId, fieldName string) string {
	return "instance:notification-config:" + entityId + ":" + fieldName
}

func (g *RedisDatabaseKeyGenerator) GetEntityTypeNotificationConfigKey(entityType, fieldName string) string {
	return "instance:notification-config:" + entityType + ":" + fieldName
}

func (g *RedisDatabaseKeyGenerator) GetNotificationChannelKey(marshalledNotificationConfig string) string {
	return "instance:notification:" + marshalledNotificationConfig
}

type RedisDatabase struct {
	client    *redis.Client
	config    RedisDatabaseConfig
	callbacks map[string]func(*DatabaseNotification)
	lastIds   map[string]string
	keygen    RedisDatabaseKeyGenerator
}

func NewRedisDatabase(config RedisDatabaseConfig) IDatabase {
	return &RedisDatabase{
		config:    config,
		callbacks: map[string]func(*DatabaseNotification){},
		lastIds:   map[string]string{},
		keygen:    RedisDatabaseKeyGenerator{},
	}
}

func (db *RedisDatabase) Connect() {
	db.Disconnect()

	Info("[RedisDatabase::Connect] Connecting to %v", db.config.Address)
	db.client = redis.NewClient(&redis.Options{
		Addr:     db.config.Address,
		Password: db.config.Password,
		DB:       0,
	})
}

func (db *RedisDatabase) Disconnect() {
	if db.client == nil {
		return
	}

	db.client.Close()
	db.client = nil
}

func (db *RedisDatabase) IsConnected() bool {
	return db.client != nil && db.client.Ping(context.Background()).Err() == nil
}

func (db *RedisDatabase) CreateSnapshot() *DatabaseSnapshot {
	snapshot := &DatabaseSnapshot{}

	for _, entityType := range db.GetEntityTypes() {
		entitySchema := db.GetEntitySchema(entityType)
		snapshot.EntitySchemas = append(snapshot.EntitySchemas, entitySchema)
		for _, entityId := range db.FindEntities(entityType) {
			snapshot.Entities = append(snapshot.Entities, db.GetEntity(entityId))
			for _, fieldName := range entitySchema.Fields {
				request := &DatabaseRequest{
					Id:    entityId,
					Field: fieldName,
				}
				db.Read([]*DatabaseRequest{request})
				if request.Success {
					snapshot.Fields = append(snapshot.Fields, new(DatabaseField).FromRequest(request))
				}
			}
		}
	}

	snapshot.FieldSchemas = db.GetFieldSchemas()

	return snapshot
}

func (db *RedisDatabase) RestoreSnapshot(snapshot *DatabaseSnapshot) {
	Info("[RedisDatabase::RestoreSnapshot] Restoring snapshot...")

	err := db.client.FlushDB(context.Background()).Err()
	if err != nil {
		Error("[RedisDatabase::RestoreSnapshot] Failed to flush database: %v", err)
		return
	}

	for _, schema := range snapshot.EntitySchemas {
		db.SetEntitySchema(schema.Name, schema)
		Debug("[RedisDatabase::RestoreSnapshot] Restored entity schema: %v", schema)
	}

	for _, schema := range snapshot.FieldSchemas {
		db.SetFieldSchema(schema.Name, schema)
		Debug("[RedisDatabase::RestoreSnapshot] Restored field schema: %v", schema)
	}

	for _, entity := range snapshot.Entities {
		db.SetEntity(entity.Id, entity)
		db.client.SAdd(context.Background(), db.keygen.GetEntityTypeKey(entity.Type), entity.Id)
		Debug("[RedisDatabase::RestoreSnapshot] Restored entity: %v", entity)
	}

	for _, field := range snapshot.Fields {
		db.Write([]*DatabaseRequest{
			{
				Id:        field.Id,
				Field:     field.Name,
				Value:     field.Value,
				WriteTime: &Timestamp{Raw: field.WriteTime},
				WriterId:  &String{Raw: field.WriterId},
			},
		})
		Debug("[RedisDatabase::RestoreSnapshot] Restored field: %v", field)
	}

	Info("[RedisDatabase::RestoreSnapshot] Snapshot restored.")
}

func (db *RedisDatabase) CreateEntity(entityType, parentId, name string) {
	entityId := uuid.New().String()

	requests := []*DatabaseRequest{}
	for _, fieldName := range db.GetEntitySchema(entityType).Fields {
		requests = append(requests, &DatabaseRequest{
			Id:    entityId,
			Field: fieldName,
		})
	}
	db.Write(requests)

	p := &DatabaseEntity{
		Id:       entityId,
		Name:     name,
		Parent:   &EntityReference{Raw: parentId},
		Type:     entityType,
		Children: []*EntityReference{},
	}
	b, err := proto.Marshal(p)
	if err != nil {
		Error("[RedisDatabase::CreateEntity] Failed to marshal entity: %v", err)
		return
	}

	db.client.SAdd(context.Background(), db.keygen.GetEntityTypeKey(entityType), entityId)
	db.client.Set(context.Background(), db.keygen.GetEntityKey(entityId), base64.StdEncoding.EncodeToString(b), 0)

	if parentId != "" {
		parent := db.GetEntity(parentId)
		if parent != nil {
			parent.Children = append(parent.Children, &EntityReference{Raw: entityId})
			db.SetEntity(parentId, parent)
		} else {
			Error("[RedisDatabase::CreateEntity] Failed to get parent entity: %v", parentId)
		}
	}
}

func (db *RedisDatabase) GetEntity(entityId string) *DatabaseEntity {
	e, err := db.client.Get(context.Background(), db.keygen.GetEntityKey(entityId)).Result()
	if err != nil {
		Error("[RedisDatabase::GetEntity] Failed to get entity: %v", err)
		return nil
	}

	b, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		Error("[RedisDatabase::GetEntity] Failed to decode entity: %v", err)
		return nil
	}

	p := &DatabaseEntity{}
	err = proto.Unmarshal(b, p)
	if err != nil {
		Error("[RedisDatabase::GetEntity] Failed to unmarshal entity: %v", err)
		return nil
	}

	return p
}

func (db *RedisDatabase) SetEntity(entityId string, value *DatabaseEntity) {
	b, err := proto.Marshal(value)
	if err != nil {
		Error("[RedisDatabase::SetEntity] Failed to marshal entity: %v", err)
		return
	}

	err = db.client.Set(context.Background(), db.keygen.GetEntityKey(entityId), base64.StdEncoding.EncodeToString(b), 0).Err()
	if err != nil {
		Error("[RedisDatabase::SetEntity] Failed to set entity '%s': %v", entityId, err)
		return
	}
}

func (db *RedisDatabase) DeleteEntity(entityId string) {
	p := db.GetEntity(entityId)
	if p == nil {
		Error("[RedisDatabase::DeleteEntity] Failed to get entity: %v", entityId)
		return
	}

	parent := db.GetEntity(p.Parent.Raw)
	if parent != nil {
		newChildren := []*EntityReference{}
		for _, child := range parent.Children {
			if child.Raw != entityId {
				newChildren = append(newChildren, child)
			}
		}
		parent.Children = newChildren
		db.SetEntity(p.Parent.Raw, parent)
	}

	for _, child := range p.Children {
		db.DeleteEntity(child.Raw)
	}

	for _, fieldName := range db.GetEntitySchema(p.Type).Fields {
		db.client.Del(context.Background(), db.keygen.GetFieldKey(fieldName, entityId))
	}

	db.client.SRem(context.Background(), db.keygen.GetEntityTypeKey(p.Type), entityId)
	db.client.Del(context.Background(), db.keygen.GetEntityKey(entityId))
}

func (db *RedisDatabase) FindEntities(entityType string) []string {
	return db.client.SMembers(context.Background(), db.keygen.GetEntityTypeKey(entityType)).Val()
}

func (db *RedisDatabase) EntityExists(entityId string) bool {
	e, err := db.client.Get(context.Background(), db.keygen.GetEntityKey(entityId)).Result()
	if err != nil {
		return false
	}

	return e != ""
}

func (db *RedisDatabase) FieldExists(fieldName, entityTypeOrId string) bool {
	if !strings.Contains(entityTypeOrId, "-") {
		schema := db.GetEntitySchema(entityTypeOrId)
		if schema != nil {
			for _, field := range schema.Fields {
				if field == fieldName {
					return true
				}
			}

			return false
		}
	}

	request := &DatabaseRequest{
		Id:    entityTypeOrId,
		Field: fieldName,
	}
	db.Read([]*DatabaseRequest{request})

	return request.Success
}

func (db *RedisDatabase) GetFieldSchemas() []*DatabaseFieldSchema {
	it := db.client.Scan(context.Background(), 0, db.keygen.GetFieldSchemaKey("*"), 0).Iterator()
	schemas := []*DatabaseFieldSchema{}

	for it.Next(context.Background()) {
		e, err := db.client.Get(context.Background(), it.Val()).Result()
		if err != nil {
			Error("[RedisDatabase::GetFieldSchemas] Failed to get field schema: %v", err)
			continue
		}

		b, err := base64.StdEncoding.DecodeString(e)
		if err != nil {
			Error("[RedisDatabase::GetFieldSchemas] Failed to decode field schema: %v", err)
			continue
		}

		p := &DatabaseFieldSchema{}
		err = proto.Unmarshal(b, p)
		if err != nil {
			Error("[RedisDatabase::GetFieldSchemas] Failed to unmarshal field schema: %v", err)
			continue
		}

		schemas = append(schemas, p)
	}

	return schemas

}

func (db *RedisDatabase) GetFieldSchema(fieldName string) *DatabaseFieldSchema {
	e, err := db.client.Get(context.Background(), db.keygen.GetFieldSchemaKey(fieldName)).Result()
	if err != nil {
		Error("[RedisDatabase::GetFieldSchema] Failed to get field schema: %v", err)
		return nil
	}

	b, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		Error("[RedisDatabase::GetFieldSchema] Failed to decode field schema: %v", err)
		return nil
	}

	a := &DatabaseFieldSchema{}
	err = proto.Unmarshal(b, a)
	if err != nil {
		Error("[RedisDatabase::GetFieldSchema] Failed to unmarshal field schema: %v", err)
		return nil
	}

	return a
}

func (db *RedisDatabase) SetFieldSchema(fieldName string, value *DatabaseFieldSchema) {
	b, err := proto.Marshal(value)
	if err != nil {
		Error("[RedisDatabase::SetFieldSchema] Failed to marshal field schema: %v", err)
		return
	}

	db.client.Set(context.Background(), db.keygen.GetFieldSchemaKey(fieldName), base64.StdEncoding.EncodeToString(b), 0)
}

func (db *RedisDatabase) GetEntityTypes() []string {
	it := db.client.Scan(context.Background(), 0, db.keygen.GetEntitySchemaKey("*"), 0).Iterator()
	types := []string{}

	for it.Next(context.Background()) {
		types = append(types, strings.ReplaceAll(it.Val(), db.keygen.GetEntitySchemaKey(""), ""))
	}

	return types
}

func (db *RedisDatabase) GetEntitySchema(entityType string) *DatabaseEntitySchema {
	e, err := db.client.Get(context.Background(), db.keygen.GetEntitySchemaKey(entityType)).Result()
	if err != nil {
		Error("[RedisDatabase::GetEntitySchema] Failed to get entity schema: %v", err)
		return nil
	}

	b, err := base64.StdEncoding.DecodeString(e)
	if err != nil {
		Error("[RedisDatabase::GetEntitySchema] Failed to decode entity schema: %v", err)
		return nil
	}

	p := &DatabaseEntitySchema{}
	err = proto.Unmarshal(b, p)
	if err != nil {
		Error("[RedisDatabase::GetEntitySchema] Failed to unmarshal entity schema: %v", err)
		return nil
	}

	return p
}

func (db *RedisDatabase) SetEntitySchema(entityType string, value *DatabaseEntitySchema) {
	b, err := proto.Marshal(value)
	if err != nil {
		Error("[RedisDatabase::SetEntitySchema] Failed to marshal entity schema: %v", err)
		return
	}

	oldSchema := db.GetEntitySchema(entityType)
	if oldSchema != nil {
		removedFields := []string{}
		newFields := []string{}

		for _, field := range oldSchema.Fields {
			if !slices.Contains(value.Fields, field) {
				removedFields = append(removedFields, field)
			}
		}

		for _, field := range value.Fields {
			if !slices.Contains(oldSchema.Fields, field) {
				newFields = append(newFields, field)
			}
		}

		for _, entityId := range db.FindEntities(entityType) {
			for _, field := range removedFields {
				db.client.Del(context.Background(), db.keygen.GetFieldKey(field, entityId))
			}

			for _, field := range newFields {
				request := &DatabaseRequest{
					Id:    entityId,
					Field: field,
				}
				db.Write([]*DatabaseRequest{request})
			}
		}
	}

	db.client.Set(context.Background(), db.keygen.GetEntitySchemaKey(entityType), base64.StdEncoding.EncodeToString(b), 0)
}

func (db *RedisDatabase) Read(requests []*DatabaseRequest) {
	for _, request := range requests {
		request.Success = false

		indirectField, indirectEntity := db.ResolveIndirection(request.Field, request.Id)

		if indirectField == "" || indirectEntity == "" {
			Error("[RedisDatabase::Read] Failed to resolve indirection: %v", request)
			continue
		}

		e, err := db.client.Get(context.Background(), db.keygen.GetFieldKey(indirectField, indirectEntity)).Result()
		if err != nil {
			Error("[RedisDatabase::Read] Failed to read field: %v", err)
			continue
		}

		b, err := base64.StdEncoding.DecodeString(e)
		if err != nil {
			Error("[RedisDatabase::Read] Failed to decode field: %v", err)
			continue
		}

		p := &DatabaseField{}
		err = proto.Unmarshal(b, p)
		if err != nil {
			Error("[RedisDatabase::Read] Failed to unmarshal field: %v", err)
			continue
		}

		request.Value = p.Value

		if request.WriteTime == nil {
			request.WriteTime = &Timestamp{Raw: timestamppb.Now()}
		}
		request.WriteTime.Raw = p.WriteTime

		if request.WriterId == nil {
			request.WriterId = &String{Raw: ""}
		}
		request.WriterId.Raw = p.WriterId

		request.Success = true
	}
}

func (db *RedisDatabase) Write(requests []*DatabaseRequest) {
	for _, request := range requests {
		request.Success = false

		indirectField, indirectEntity := db.ResolveIndirection(request.Field, request.Id)

		if indirectField == "" || indirectEntity == "" {
			Error("[RedisDatabase::Write] Failed to resolve indirection: %v", request)
			continue
		}

		schema := db.GetFieldSchema(indirectField)
		sampleType, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(schema.Type))
		if err != nil {
			Error("[RedisDatabase::Write] Failed to find message type: %v", err)
			continue
		}

		if request.Value == nil {
			if request.Value, err = anypb.New(sampleType.New().Interface()); err != nil {
				Error("[RedisDatabase::Write] Failed to create anypb: %v", err)
				continue
			}
		} else {
			sampleAnyType, err := anypb.New(sampleType.New().Interface())
			if err != nil {
				Error("[RedisDatabase::Write] Failed to create anypb: %v", err)
				continue
			}

			if request.Value.TypeUrl != sampleAnyType.TypeUrl {
				Warn("[RedisDatabase::Write] Field type mismatch: %v != %v. Overwritting...", request.Value.TypeUrl, sampleAnyType.TypeUrl)
				request.Value = sampleAnyType
			}
		}

		if request.WriteTime == nil {
			request.WriteTime = &Timestamp{Raw: timestamppb.Now()}
		}

		if request.WriterId == nil {
			request.WriterId = &String{Raw: ""}
		}

		p := new(DatabaseField).FromRequest(request)

		b, err := proto.Marshal(p)
		if err != nil {
			Error("[RedisDatabase::Write] Failed to marshal field: %v", err)
			continue
		}

		p.Id = indirectEntity
		p.Name = indirectField

		db.triggerNotifications(request)

		_, err = db.client.Set(context.Background(), db.keygen.GetFieldKey(indirectField, indirectEntity), base64.StdEncoding.EncodeToString(b), 0).Result()
		if err != nil {
			Error("[RedisDatabase::Write] Failed to write field: %v", err)
			continue
		}
		request.Success = true
	}
}

func (db *RedisDatabase) Notify(notification *DatabaseNotificationConfig, callback func(*DatabaseNotification)) string {
	b, err := proto.Marshal(notification)
	if err != nil {
		Error("[RedisDatabase::Notify] Failed to marshal notification config: %v", err)
		return ""
	}

	e := base64.StdEncoding.EncodeToString(b)

	if db.FieldExists(notification.Field, notification.Id) {
		r, err := db.client.XInfoStream(context.Background(), db.keygen.GetNotificationChannelKey(e)).Result()
		if err != nil {
			Warn("[RedisDatabase::Notify] Failed to find stream for: %v (%v)", e, err)
			db.lastIds[e] = "0"
		} else {
			db.lastIds[e] = r.LastGeneratedID
		}

		db.client.SAdd(context.Background(), db.keygen.GetEntityIdNotificationConfigKey(notification.Id, notification.Field), e)
		db.callbacks[e] = callback
		return e
	}

	if db.FieldExists(notification.Field, notification.Type) {
		r, err := db.client.XInfoStream(context.Background(), db.keygen.GetNotificationChannelKey(e)).Result()
		if err != nil {
			Warn("[RedisDatabase::Notify] Failed to find stream for: %v (%v)", e, err)
			db.lastIds[e] = "0"
		} else {
			db.lastIds[e] = r.LastGeneratedID
		}

		db.client.SAdd(context.Background(), db.keygen.GetEntityTypeNotificationConfigKey(notification.Type, notification.Field), e)
		db.callbacks[e] = callback
		return e
	}

	Warn("[RedisDatabase::Notify] Failed to find field: %v", notification)
	return ""
}

func (db *RedisDatabase) Unnotify(e string) {
	if db.callbacks[e] == nil {
		Warn("[RedisDatabase::Unnotify] Failed to find callback: %v", e)
		return
	}

	delete(db.callbacks, e)
	delete(db.lastIds, e)
}

func (db *RedisDatabase) ProcessNotifications() {
	for e := range db.callbacks {
		r, err := db.client.XRead(context.Background(), &redis.XReadArgs{
			Streams: []string{db.keygen.GetNotificationChannelKey(e), db.lastIds[e]},
			Count:   100,
			Block:   -1,
		}).Result()

		if err != nil && err != redis.Nil {
			Error("[RedisDatabase::ProcessNotifications] Failed to read stream %v: %v", e, err)
			continue
		}

		for _, x := range r {
			for _, m := range x.Messages {
				db.lastIds[e] = m.ID
				decodedMessage := make(map[string]string)

				for key, value := range m.Values {
					if castedValue, ok := value.(string); ok {
						decodedMessage[key] = castedValue
					} else {
						Error("[RedisDatabase::ProcessNotifications] Failed to cast value: %v", value)
						continue
					}
				}

				if data, ok := decodedMessage["data"]; ok {
					p, err := base64.StdEncoding.DecodeString(data)
					if err != nil {
						Error("[RedisDatabase::ProcessNotifications] Failed to decode notification: %v", err)
						continue
					}

					n := &DatabaseNotification{}
					err = proto.Unmarshal(p, n)
					if err != nil {
						Error("[RedisDatabase::ProcessNotifications] Failed to unmarshal notification: %v", err)
						continue
					}

					db.callbacks[e](n)
				}
			}
		}
	}
}

func (db *RedisDatabase) ResolveIndirection(indirectField, entityId string) (string, string) {
	fields := strings.Split(indirectField, "->")
	if len(fields) == 1 {
		return indirectField, entityId
	}

	for _, field := range fields {
		request := &DatabaseRequest{
			Id:    entityId,
			Field: field,
		}

		db.Read([]*DatabaseRequest{request})

		if request.Success {
			entityReference := &EntityReference{}
			if request.Value.MessageIs(entityReference) {
				err := request.Value.UnmarshalTo(entityReference)
				if err != nil {
					Error("[RedisDatabase::ResolveIndirection] Failed to unmarshal entity reference: %v", err)
					return "", ""
				}

				entityId = entityReference.Raw
				continue
			}

			Error("[RedisDatabase::ResolveIndirection] Field is not an entity reference: %v", request)
			return "", ""
		}

		// Fallback to parent entity reference by name
		entity := db.GetEntity(entityId)
		if entity == nil {
			Error("[RedisDatabase::ResolveIndirection] Failed to get entity: %v", entityId)
			return "", ""
		}

		if entity.Parent != nil && entity.Parent.Raw != "" {
			parentEntity := db.GetEntity(entity.Parent.Raw)

			if parentEntity != nil && parentEntity.Name == field {
				entityId = entity.Parent.Raw
				continue
			}
		}

		// Fallback to child entity reference by name
		foundChild := false
		for _, child := range entity.Children {
			childEntity := db.GetEntity(child.Raw)
			if childEntity == nil {
				Error("[RedisDatabase::ResolveIndirection] Failed to get child entity: %v", child.Raw)
				continue
			}

			if childEntity.Name == field {
				entityId = child.Raw
				foundChild = true
				break
			}
		}

		if !foundChild {
			return "", ""
		}
	}

	return fields[len(fields)-1], entityId
}

func (db *RedisDatabase) triggerNotifications(request *DatabaseRequest) {
	oldRequest := &DatabaseRequest{
		Id:    request.Id,
		Field: request.Field,
	}
	db.Read([]*DatabaseRequest{oldRequest})

	// failed to read old value (it may not exist initially)
	if !oldRequest.Success {
		Warn("[RedisDatabase::triggerNotifications] Failed to read old value: %v", oldRequest)
		return
	}

	changed := proto.Equal(oldRequest.Value, request.Value)

	indirectField, indirectEntity := db.ResolveIndirection(request.Field, request.Id)

	if indirectField == "" || indirectEntity == "" {
		Error("[RedisDatabase::triggerNotifications] Failed to resolve indirection: %v", request)
		return
	}

	m, err := db.client.SMembers(context.Background(), db.keygen.GetEntityIdNotificationConfigKey(indirectEntity, indirectField)).Result()
	if err != nil {
		Error("[RedisDatabase::triggerNotifications] Failed to get notification config: %v", err)
		return
	}

	for _, e := range m {
		b, err := base64.StdEncoding.DecodeString(e)
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to decode notification config: %v", err)
			continue
		}

		p := &DatabaseNotificationConfig{}
		err = proto.Unmarshal(b, p)
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to unmarshal notification config: %v", err)
			continue
		}

		if p.NotifyOnChange && !changed {
			continue
		}

		n := &DatabaseNotification{
			Token:    e,
			Current:  new(DatabaseField).FromRequest(request),
			Previous: new(DatabaseField).FromRequest(oldRequest),
			Context:  []*DatabaseField{},
		}

		for _, context := range p.ContextFields {
			contextRequest := &DatabaseRequest{
				Id:    request.Id,
				Field: context,
			}
			db.Read([]*DatabaseRequest{contextRequest})
			if contextRequest.Success {
				n.Context = append(n.Context, new(DatabaseField).FromRequest(contextRequest))
			}
		}

		b, err = proto.Marshal(n)
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to marshal notification: %v", err)
			continue
		}

		_, err = db.client.XAdd(context.Background(), &redis.XAddArgs{
			Stream: db.keygen.GetNotificationChannelKey(e),
			Values: []string{"data", base64.StdEncoding.EncodeToString(b)},
			MaxLen: 100,
			Approx: true,
		}).Result()
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to add notification: %v", err)
			continue
		}
	}

	entity := db.GetEntity(indirectEntity)
	if entity == nil {
		Error("[RedisDatabase::triggerNotifications] Failed to get entity: %v (indirect=%v)", request.Id, indirectEntity)
		return
	}

	m, err = db.client.SMembers(context.Background(), db.keygen.GetEntityTypeNotificationConfigKey(entity.Type, indirectField)).Result()
	if err != nil {
		Error("[RedisDatabase::triggerNotifications] Failed to get notification config: %v", err)
		return
	}

	for _, e := range m {
		b, err := base64.StdEncoding.DecodeString(e)
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to decode notification config: %v", err)
			continue
		}

		p := &DatabaseNotificationConfig{}
		err = proto.Unmarshal(b, p)
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to unmarshal notification config: %v", err)
			continue
		}

		if p.NotifyOnChange && !changed {
			continue
		}

		n := &DatabaseNotification{
			Token:    e,
			Current:  new(DatabaseField).FromRequest(request),
			Previous: new(DatabaseField).FromRequest(oldRequest),
			Context:  []*DatabaseField{},
		}

		for _, context := range p.ContextFields {
			contextRequest := &DatabaseRequest{
				Id:    request.Id,
				Field: context,
			}
			db.Read([]*DatabaseRequest{contextRequest})
			if contextRequest.Success {
				n.Context = append(n.Context, new(DatabaseField).FromRequest(contextRequest))
			}
		}

		b, err = proto.Marshal(n)
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to marshal notification: %v", err)
			continue
		}

		_, err = db.client.XAdd(context.Background(), &redis.XAddArgs{
			Stream: db.keygen.GetNotificationChannelKey(e),
			Values: []string{"data", base64.StdEncoding.EncodeToString(b)},
			MaxLen: 100,
			Approx: true,
		}).Result()
		if err != nil {
			Error("[RedisDatabase::triggerNotifications] Failed to add notification: %v", err)
			continue
		}
	}
}
