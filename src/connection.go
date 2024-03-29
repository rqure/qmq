package qmq

import (
	"context"
	"encoding/base64"
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type QMQConnectionError int

const (
	CONNECTION_FAILED QMQConnectionError = iota
	MARSHAL_FAILED
	UNMARSHAL_FAILED
	SET_FAILED
	TEMPSET_FAILED
	GET_FAILED
	STREAM_ADD_FAILED
	STREAM_READ_FAILED
	DECODE_FAILED
	CAST_FAILED
	STREAM_CONTEXT_FAILED
	STREAM_EMPTY
	UNSET_FAILED
)

func (e QMQConnectionError) Error() string {
	return fmt.Sprintf("QMQConnectionError: %d", e)
}

type QMQConnection struct {
	addr     string
	password string
	redis    *redis.Client
	lock     sync.Mutex
}

func NewReadRequest() *QMQData {
	return &QMQData{
		Data: &anypb.Any{},
	}
}

func NewWriteRequest(m protoreflect.ProtoMessage) *QMQData {
	writeRequest := &QMQData{
		Data: &anypb.Any{},
	}
	writeRequest.Data.MarshalFrom(m)
	return writeRequest
}

func NewQMQConnection(addr string, password string) *QMQConnection {
	return &QMQConnection{
		addr:     addr,
		password: password,
	}
}

func (q *QMQConnection) Connect() error {
	q.Disconnect()

	q.lock.Lock()
	defer q.lock.Unlock()

	opt := &redis.Options{
		Addr:     q.addr,
		Password: q.password,
		DB:       0, // use default DB
	}
	q.redis = redis.NewClient(opt)

	if q.redis.Ping(context.Background()).Err() != nil {
		return CONNECTION_FAILED
	}

	return nil
}

func (q *QMQConnection) Disconnect() {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.redis != nil {
		q.redis.Close()
		q.redis = nil
	}
}

func (q *QMQConnection) Set(k string, d *QMQData) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if d.Writetime == nil {
		d.Writetime = timestamppb.Now()
	}

	writeRequests := make(map[string]interface{})
	v, err := proto.Marshal(d)
	if err != nil {
		return MARSHAL_FAILED
	}
	writeRequests[k] = base64.StdEncoding.EncodeToString(v)

	if q.redis.MSet(context.Background(), writeRequests).Err() != nil {
		return SET_FAILED
	}

	return nil
}

func (q *QMQConnection) TempSet(k string, d *QMQData, timeoutMs int64) (bool, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	if d.Writetime == nil {
		d.Writetime = timestamppb.Now()
	}

	v, err := proto.Marshal(d)
	if err != nil {
		return false, MARSHAL_FAILED
	}

	result, err := q.redis.SetNX(context.Background(),
		k, base64.StdEncoding.EncodeToString(v),
		time.Duration(timeoutMs)*time.Millisecond).Result()
	if err != nil {
		return false, TEMPSET_FAILED
	}

	if !result {
		return false, nil
	}

	return true, nil
}

func (q *QMQConnection) Unset(k string) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	if q.redis.Del(context.Background(), k).Err() != nil {
		return UNSET_FAILED
	}

	return nil
}

func (q *QMQConnection) Get(k string) (*QMQData, error) {
	q.lock.Lock()
	defer q.lock.Unlock()

	result := NewReadRequest()

	val, err := q.redis.Get(context.Background(), k).Result()
	if err != nil {
		return nil, GET_FAILED
	}
	protobytes, err := base64.StdEncoding.DecodeString(val)
	if err != nil {
		return nil, DECODE_FAILED
	}
	err = proto.Unmarshal(protobytes, result)
	if err != nil {
		return nil, UNMARSHAL_FAILED
	}

	return result, nil
}

func (q *QMQConnection) GetSchema(schema interface{}) {
	schemaWrapper := reflect.ValueOf(schema).Elem()
	schemaType := schemaWrapper.Type()

	for i := 0; i < schemaWrapper.NumField(); i++ {
		field := schemaWrapper.Field(i)

		tag := schemaType.Field(i).Tag.Get("qmq")
		if tag == "" {
			continue
		}

		readRequest, err := q.Get(tag)
		if err != nil {
			continue
		}

		fieldType := field.Type()
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		fieldValue := reflect.New(fieldType).Interface()
		err = readRequest.Data.UnmarshalTo(fieldValue.(proto.Message))
		if err != nil {
			continue
		}

		field.Set(reflect.ValueOf(fieldValue))
	}
}

func (q *QMQConnection) SetSchema(schema interface{}) {
	schemaWrapper := reflect.ValueOf(schema).Elem()
	schemaType := schemaWrapper.Type()

	for i := 0; i < schemaWrapper.NumField(); i++ {
		field := schemaWrapper.Field(i)

		tag := schemaType.Field(i).Tag.Get("qmq")
		if tag == "" {
			continue
		}

		writeRequest := NewWriteRequest(field.Interface().(proto.Message))
		q.Set(tag, writeRequest)
	}
}

func (q *QMQConnection) StreamAdd(s *QMQStream, m proto.Message) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	b, err := proto.Marshal(m)
	if err != nil {
		return MARSHAL_FAILED
	}

	_, err = q.redis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: s.Key(),
		Values: []string{"data", base64.StdEncoding.EncodeToString(b)},
		MaxLen: s.Length,
		Approx: true,
	}).Result()

	if err != nil {
		return STREAM_ADD_FAILED
	}

	return nil
}

func (q *QMQConnection) StreamAddRaw(s *QMQStream, d string) error {
	q.lock.Lock()
	defer q.lock.Unlock()

	_, err := q.redis.XAdd(context.Background(), &redis.XAddArgs{
		Stream: s.Key(),
		Values: []string{"data", d},
		MaxLen: s.Length,
		Approx: true,
	}).Result()

	if err != nil {
		return STREAM_ADD_FAILED
	}

	return nil
}

func (q *QMQConnection) StreamRead(s *QMQStream, m protoreflect.ProtoMessage) error {
	gResult, err := q.Get(s.ContextKey())
	if err == nil {
		gResult.Data.UnmarshalTo(&s.Context)
	} else {
		writeRequest := NewWriteRequest(&s.Context)
		q.Set(s.ContextKey(), writeRequest)
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	xResult, err := q.redis.XRead(context.Background(), &redis.XReadArgs{
		Streams: []string{s.Key(), s.Context.LastConsumedId},
		Block:   -1,
	}).Result()

	if err != nil {
		return STREAM_READ_FAILED
	}

	for _, xMessage := range xResult {
		for _, message := range xMessage.Messages {
			decodedMessage := make(map[string]string)

			for key, value := range message.Values {
				if castedValue, ok := value.(string); ok {
					decodedMessage[key] = castedValue
				} else {
					return CAST_FAILED
				}
			}

			if data, ok := decodedMessage["data"]; ok {
				protobytes, err := base64.StdEncoding.DecodeString(data)
				if err != nil {
					return DECODE_FAILED
				}
				err = proto.Unmarshal(protobytes, m)
				if err != nil {
					return UNMARSHAL_FAILED
				}
				s.Context.LastConsumedId = message.ID
				return nil
			}
		}
	}

	return STREAM_EMPTY
}

func (q *QMQConnection) StreamReadRaw(s *QMQStream) (string, error) {
	gResult, err := q.Get(s.ContextKey())
	if err == nil {
		gResult.Data.UnmarshalTo(&s.Context)
	} else {
		writeRequest := NewWriteRequest(&s.Context)
		q.Set(s.ContextKey(), writeRequest)
	}

	q.lock.Lock()
	defer q.lock.Unlock()

	xResult, err := q.redis.XRead(context.Background(), &redis.XReadArgs{
		Streams: []string{s.Key(), s.Context.LastConsumedId},
		Block:   -1,
	}).Result()

	if err != nil {
		return "", STREAM_READ_FAILED
	}

	for _, xMessage := range xResult {
		for _, message := range xMessage.Messages {
			decodedMessage := make(map[string]string)

			for key, value := range message.Values {
				if castedValue, ok := value.(string); ok {
					decodedMessage[key] = castedValue
				} else {
					return "", CAST_FAILED
				}
			}

			if data, ok := decodedMessage["data"]; ok {
				s.Context.LastConsumedId = message.ID
				return data, nil
			}
		}
	}

	return "", STREAM_EMPTY
}
