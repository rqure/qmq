// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.12.4
// source: src/data.proto

package qmq

import (
	any1 "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type QMQLogLevelEnum int32

const (
	QMQLogLevelEnum_LOG_LEVEL_UNSPECIFIED QMQLogLevelEnum = 0
	QMQLogLevelEnum_LOG_LEVEL_TRACE       QMQLogLevelEnum = 1
	QMQLogLevelEnum_LOG_LEVEL_DEBUG       QMQLogLevelEnum = 2
	QMQLogLevelEnum_LOG_LEVEL_ADVISE      QMQLogLevelEnum = 3
	QMQLogLevelEnum_LOG_LEVEL_WARN        QMQLogLevelEnum = 4
	QMQLogLevelEnum_LOG_LEVEL_ERROR       QMQLogLevelEnum = 5
	QMQLogLevelEnum_LOG_LEVEL_PANIC       QMQLogLevelEnum = 6
)

// Enum value maps for QMQLogLevelEnum.
var (
	QMQLogLevelEnum_name = map[int32]string{
		0: "LOG_LEVEL_UNSPECIFIED",
		1: "LOG_LEVEL_TRACE",
		2: "LOG_LEVEL_DEBUG",
		3: "LOG_LEVEL_ADVISE",
		4: "LOG_LEVEL_WARN",
		5: "LOG_LEVEL_ERROR",
		6: "LOG_LEVEL_PANIC",
	}
	QMQLogLevelEnum_value = map[string]int32{
		"LOG_LEVEL_UNSPECIFIED": 0,
		"LOG_LEVEL_TRACE":       1,
		"LOG_LEVEL_DEBUG":       2,
		"LOG_LEVEL_ADVISE":      3,
		"LOG_LEVEL_WARN":        4,
		"LOG_LEVEL_ERROR":       5,
		"LOG_LEVEL_PANIC":       6,
	}
)

func (x QMQLogLevelEnum) Enum() *QMQLogLevelEnum {
	p := new(QMQLogLevelEnum)
	*p = x
	return p
}

func (x QMQLogLevelEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QMQLogLevelEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_src_data_proto_enumTypes[0].Descriptor()
}

func (QMQLogLevelEnum) Type() protoreflect.EnumType {
	return &file_src_data_proto_enumTypes[0]
}

func (x QMQLogLevelEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QMQLogLevelEnum.Descriptor instead.
func (QMQLogLevelEnum) EnumDescriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{0}
}

type QMQLightBulbStateEnum int32

const (
	QMQLightBulbStateEnum_LIGHT_BULB_STATE_UNSPECIFIED QMQLightBulbStateEnum = 0
	QMQLightBulbStateEnum_LIGHT_BULB_STATE_UNKNOWN     QMQLightBulbStateEnum = 1
	QMQLightBulbStateEnum_LIGHT_BULB_STATE_ON          QMQLightBulbStateEnum = 2
	QMQLightBulbStateEnum_LIGHT_BULB_STATE_OFF         QMQLightBulbStateEnum = 3
)

// Enum value maps for QMQLightBulbStateEnum.
var (
	QMQLightBulbStateEnum_name = map[int32]string{
		0: "LIGHT_BULB_STATE_UNSPECIFIED",
		1: "LIGHT_BULB_STATE_UNKNOWN",
		2: "LIGHT_BULB_STATE_ON",
		3: "LIGHT_BULB_STATE_OFF",
	}
	QMQLightBulbStateEnum_value = map[string]int32{
		"LIGHT_BULB_STATE_UNSPECIFIED": 0,
		"LIGHT_BULB_STATE_UNKNOWN":     1,
		"LIGHT_BULB_STATE_ON":          2,
		"LIGHT_BULB_STATE_OFF":         3,
	}
)

func (x QMQLightBulbStateEnum) Enum() *QMQLightBulbStateEnum {
	p := new(QMQLightBulbStateEnum)
	*p = x
	return p
}

func (x QMQLightBulbStateEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QMQLightBulbStateEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_src_data_proto_enumTypes[1].Descriptor()
}

func (QMQLightBulbStateEnum) Type() protoreflect.EnumType {
	return &file_src_data_proto_enumTypes[1]
}

func (x QMQLightBulbStateEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QMQLightBulbStateEnum.Descriptor instead.
func (QMQLightBulbStateEnum) EnumDescriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{1}
}

type QMQDoorLockStateEnum int32

const (
	QMQDoorLockStateEnum_DOOR_LOCK_STATE_UNSPECIFIED QMQDoorLockStateEnum = 0
	QMQDoorLockStateEnum_DOOR_LOCK_STATE_UNKNOWN     QMQDoorLockStateEnum = 1
	QMQDoorLockStateEnum_DOOR_LOCK_STATE_LOCKED      QMQDoorLockStateEnum = 2
	QMQDoorLockStateEnum_DOOR_LOCK_STATE_UNLOCKED    QMQDoorLockStateEnum = 3
)

// Enum value maps for QMQDoorLockStateEnum.
var (
	QMQDoorLockStateEnum_name = map[int32]string{
		0: "DOOR_LOCK_STATE_UNSPECIFIED",
		1: "DOOR_LOCK_STATE_UNKNOWN",
		2: "DOOR_LOCK_STATE_LOCKED",
		3: "DOOR_LOCK_STATE_UNLOCKED",
	}
	QMQDoorLockStateEnum_value = map[string]int32{
		"DOOR_LOCK_STATE_UNSPECIFIED": 0,
		"DOOR_LOCK_STATE_UNKNOWN":     1,
		"DOOR_LOCK_STATE_LOCKED":      2,
		"DOOR_LOCK_STATE_UNLOCKED":    3,
	}
)

func (x QMQDoorLockStateEnum) Enum() *QMQDoorLockStateEnum {
	p := new(QMQDoorLockStateEnum)
	*p = x
	return p
}

func (x QMQDoorLockStateEnum) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (QMQDoorLockStateEnum) Descriptor() protoreflect.EnumDescriptor {
	return file_src_data_proto_enumTypes[2].Descriptor()
}

func (QMQDoorLockStateEnum) Type() protoreflect.EnumType {
	return &file_src_data_proto_enumTypes[2]
}

func (x QMQDoorLockStateEnum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QMQDoorLockStateEnum.Descriptor instead.
func (QMQDoorLockStateEnum) EnumDescriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{2}
}

type QMQData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      *any1.Any            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Writer    string               `protobuf:"bytes,2,opt,name=writer,proto3" json:"writer,omitempty"`
	Writetime *timestamp.Timestamp `protobuf:"bytes,3,opt,name=writetime,proto3" json:"writetime,omitempty"`
}

func (x *QMQData) Reset() {
	*x = QMQData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQData) ProtoMessage() {}

func (x *QMQData) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQData.ProtoReflect.Descriptor instead.
func (*QMQData) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{0}
}

func (x *QMQData) GetData() *any1.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *QMQData) GetWriter() string {
	if x != nil {
		return x.Writer
	}
	return ""
}

func (x *QMQData) GetWritetime() *timestamp.Timestamp {
	if x != nil {
		return x.Writetime
	}
	return nil
}

type QMQInt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QMQInt) Reset() {
	*x = QMQInt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQInt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQInt) ProtoMessage() {}

func (x *QMQInt) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQInt.ProtoReflect.Descriptor instead.
func (*QMQInt) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{1}
}

func (x *QMQInt) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type QMQString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value int64 `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QMQString) Reset() {
	*x = QMQString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQString) ProtoMessage() {}

func (x *QMQString) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQString.ProtoReflect.Descriptor instead.
func (*QMQString) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{2}
}

func (x *QMQString) GetValue() int64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type QMQTimestamp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value *timestamp.Timestamp `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QMQTimestamp) Reset() {
	*x = QMQTimestamp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQTimestamp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQTimestamp) ProtoMessage() {}

func (x *QMQTimestamp) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQTimestamp.ProtoReflect.Descriptor instead.
func (*QMQTimestamp) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{3}
}

func (x *QMQTimestamp) GetValue() *timestamp.Timestamp {
	if x != nil {
		return x.Value
	}
	return nil
}

type QMQFloat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QMQFloat) Reset() {
	*x = QMQFloat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQFloat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQFloat) ProtoMessage() {}

func (x *QMQFloat) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQFloat.ProtoReflect.Descriptor instead.
func (*QMQFloat) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{4}
}

func (x *QMQFloat) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

type QMQBool struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value bool `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *QMQBool) Reset() {
	*x = QMQBool{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQBool) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQBool) ProtoMessage() {}

func (x *QMQBool) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQBool.ProtoReflect.Descriptor instead.
func (*QMQBool) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{5}
}

func (x *QMQBool) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type QMQLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Application string               `protobuf:"bytes,1,opt,name=application,proto3" json:"application,omitempty"`
	Level       QMQLogLevelEnum      `protobuf:"varint,2,opt,name=level,proto3,enum=qmq.QMQLogLevelEnum" json:"level,omitempty"`
	Message     string               `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp   *timestamp.Timestamp `protobuf:"bytes,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *QMQLog) Reset() {
	*x = QMQLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQLog) ProtoMessage() {}

func (x *QMQLog) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQLog.ProtoReflect.Descriptor instead.
func (*QMQLog) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{6}
}

func (x *QMQLog) GetApplication() string {
	if x != nil {
		return x.Application
	}
	return ""
}

func (x *QMQLog) GetLevel() QMQLogLevelEnum {
	if x != nil {
		return x.Level
	}
	return QMQLogLevelEnum_LOG_LEVEL_UNSPECIFIED
}

func (x *QMQLog) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *QMQLog) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type QMQStreamContext struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastConsumedId string `protobuf:"bytes,1,opt,name=last_consumed_id,json=lastConsumedId,proto3" json:"last_consumed_id,omitempty"`
	Length         int64  `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *QMQStreamContext) Reset() {
	*x = QMQStreamContext{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQStreamContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQStreamContext) ProtoMessage() {}

func (x *QMQStreamContext) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQStreamContext.ProtoReflect.Descriptor instead.
func (*QMQStreamContext) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{7}
}

func (x *QMQStreamContext) GetLastConsumedId() string {
	if x != nil {
		return x.LastConsumedId
	}
	return ""
}

func (x *QMQStreamContext) GetLength() int64 {
	if x != nil {
		return x.Length
	}
	return 0
}

type QMQPrayer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string               `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Time *timestamp.Timestamp `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *QMQPrayer) Reset() {
	*x = QMQPrayer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQPrayer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQPrayer) ProtoMessage() {}

func (x *QMQPrayer) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQPrayer.ProtoReflect.Descriptor instead.
func (*QMQPrayer) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{8}
}

func (x *QMQPrayer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *QMQPrayer) GetTime() *timestamp.Timestamp {
	if x != nil {
		return x.Time
	}
	return nil
}

type QMQAudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filename string `protobuf:"bytes,1,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *QMQAudioRequest) Reset() {
	*x = QMQAudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQAudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQAudioRequest) ProtoMessage() {}

func (x *QMQAudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQAudioRequest.ProtoReflect.Descriptor instead.
func (*QMQAudioRequest) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{9}
}

func (x *QMQAudioRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type QMQLightBulb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentState   QMQLightBulbStateEnum `protobuf:"varint,1,opt,name=current_state,json=currentState,proto3,enum=qmq.QMQLightBulbStateEnum" json:"current_state,omitempty"`
	RequestedState QMQLightBulbStateEnum `protobuf:"varint,2,opt,name=requested_state,json=requestedState,proto3,enum=qmq.QMQLightBulbStateEnum" json:"requested_state,omitempty"`
}

func (x *QMQLightBulb) Reset() {
	*x = QMQLightBulb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQLightBulb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQLightBulb) ProtoMessage() {}

func (x *QMQLightBulb) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQLightBulb.ProtoReflect.Descriptor instead.
func (*QMQLightBulb) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{10}
}

func (x *QMQLightBulb) GetCurrentState() QMQLightBulbStateEnum {
	if x != nil {
		return x.CurrentState
	}
	return QMQLightBulbStateEnum_LIGHT_BULB_STATE_UNSPECIFIED
}

func (x *QMQLightBulb) GetRequestedState() QMQLightBulbStateEnum {
	if x != nil {
		return x.RequestedState
	}
	return QMQLightBulbStateEnum_LIGHT_BULB_STATE_UNSPECIFIED
}

type QMQDoorLock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentState   QMQDoorLockStateEnum `protobuf:"varint,1,opt,name=current_state,json=currentState,proto3,enum=qmq.QMQDoorLockStateEnum" json:"current_state,omitempty"`
	RequestedState QMQDoorLockStateEnum `protobuf:"varint,2,opt,name=requested_state,json=requestedState,proto3,enum=qmq.QMQDoorLockStateEnum" json:"requested_state,omitempty"`
}

func (x *QMQDoorLock) Reset() {
	*x = QMQDoorLock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_data_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QMQDoorLock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QMQDoorLock) ProtoMessage() {}

func (x *QMQDoorLock) ProtoReflect() protoreflect.Message {
	mi := &file_src_data_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QMQDoorLock.ProtoReflect.Descriptor instead.
func (*QMQDoorLock) Descriptor() ([]byte, []int) {
	return file_src_data_proto_rawDescGZIP(), []int{11}
}

func (x *QMQDoorLock) GetCurrentState() QMQDoorLockStateEnum {
	if x != nil {
		return x.CurrentState
	}
	return QMQDoorLockStateEnum_DOOR_LOCK_STATE_UNSPECIFIED
}

func (x *QMQDoorLock) GetRequestedState() QMQDoorLockStateEnum {
	if x != nil {
		return x.RequestedState
	}
	return QMQDoorLockStateEnum_DOOR_LOCK_STATE_UNSPECIFIED
}

var File_src_data_proto protoreflect.FileDescriptor

var file_src_data_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x72, 0x63, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x71, 0x6d, 0x71, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x85, 0x01, 0x0a, 0x07, 0x51, 0x4d, 0x51, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x12,
	0x38, 0x0a, 0x09, 0x77, 0x72, 0x69, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x77, 0x72, 0x69, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x06, 0x51, 0x4d, 0x51,
	0x49, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x09, 0x51, 0x4d, 0x51,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x40, 0x0a, 0x0c,
	0x51, 0x4d, 0x51, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x30, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x20,
	0x0a, 0x08, 0x51, 0x4d, 0x51, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x1f, 0x0a, 0x07, 0x51, 0x4d, 0x51, 0x42, 0x6f, 0x6f, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0xaa, 0x01, 0x0a, 0x06, 0x51, 0x4d, 0x51, 0x4c, 0x6f, 0x67, 0x12, 0x20, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2a,
	0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e,
	0x71, 0x6d, 0x71, 0x2e, 0x51, 0x4d, 0x51, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45,
	0x6e, 0x75, 0x6d, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x54,
	0x0a, 0x10, 0x51, 0x4d, 0x51, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x28, 0x0a, 0x10, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6d, 0x65, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6c, 0x61,
	0x73, 0x74, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6c, 0x65,
	0x6e, 0x67, 0x74, 0x68, 0x22, 0x4f, 0x0a, 0x09, 0x51, 0x4d, 0x51, 0x50, 0x72, 0x61, 0x79, 0x65,
	0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x2d, 0x0a, 0x0f, 0x51, 0x4d, 0x51, 0x41, 0x75, 0x64, 0x69,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x94, 0x01, 0x0a, 0x0c, 0x51, 0x4d, 0x51, 0x4c, 0x69, 0x67, 0x68,
	0x74, 0x42, 0x75, 0x6c, 0x62, 0x12, 0x3f, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x71,
	0x6d, 0x71, 0x2e, 0x51, 0x4d, 0x51, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x75, 0x6c, 0x62, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x1a, 0x2e, 0x71, 0x6d, 0x71, 0x2e, 0x51, 0x4d, 0x51, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x75,
	0x6c, 0x62, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0e, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x91, 0x01, 0x0a, 0x0b,
	0x51, 0x4d, 0x51, 0x44, 0x6f, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x3e, 0x0a, 0x0d, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x71, 0x6d, 0x71, 0x2e, 0x51, 0x4d, 0x51, 0x44, 0x6f, 0x6f, 0x72,
	0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52, 0x0c, 0x63,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x42, 0x0a, 0x0f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x71, 0x6d, 0x71, 0x2e, 0x51, 0x4d, 0x51, 0x44, 0x6f,
	0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x52,
	0x0e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65, 0x2a,
	0xaa, 0x01, 0x0a, 0x0f, 0x51, 0x4d, 0x51, 0x4c, 0x6f, 0x67, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x45,
	0x6e, 0x75, 0x6d, 0x12, 0x19, 0x0a, 0x15, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13,
	0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x54, 0x52, 0x41, 0x43,
	0x45, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c,
	0x5f, 0x44, 0x45, 0x42, 0x55, 0x47, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x4f, 0x47, 0x5f,
	0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x41, 0x44, 0x56, 0x49, 0x53, 0x45, 0x10, 0x03, 0x12, 0x12,
	0x0a, 0x0e, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f, 0x57, 0x41, 0x52, 0x4e,
	0x10, 0x04, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c, 0x45, 0x56, 0x45, 0x4c, 0x5f,
	0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x13, 0x0a, 0x0f, 0x4c, 0x4f, 0x47, 0x5f, 0x4c,
	0x45, 0x56, 0x45, 0x4c, 0x5f, 0x50, 0x41, 0x4e, 0x49, 0x43, 0x10, 0x06, 0x2a, 0x8a, 0x01, 0x0a,
	0x15, 0x51, 0x4d, 0x51, 0x4c, 0x69, 0x67, 0x68, 0x74, 0x42, 0x75, 0x6c, 0x62, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x45, 0x6e, 0x75, 0x6d, 0x12, 0x20, 0x0a, 0x1c, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x42, 0x55, 0x4c, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x4c, 0x49, 0x47, 0x48,
	0x54, 0x5f, 0x42, 0x55, 0x4c, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01, 0x12, 0x17, 0x0a, 0x13, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f,
	0x42, 0x55, 0x4c, 0x42, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4f, 0x4e, 0x10, 0x02, 0x12,
	0x18, 0x0a, 0x14, 0x4c, 0x49, 0x47, 0x48, 0x54, 0x5f, 0x42, 0x55, 0x4c, 0x42, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x4f, 0x46, 0x46, 0x10, 0x03, 0x2a, 0x8e, 0x01, 0x0a, 0x14, 0x51, 0x4d,
	0x51, 0x44, 0x6f, 0x6f, 0x72, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x45, 0x6e,
	0x75, 0x6d, 0x12, 0x1f, 0x0a, 0x1b, 0x44, 0x4f, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x44, 0x4f, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x4b,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x01,
	0x12, 0x1a, 0x0a, 0x16, 0x44, 0x4f, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18,
	0x44, 0x4f, 0x4f, 0x52, 0x5f, 0x4c, 0x4f, 0x43, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x55, 0x4e, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x42, 0x09, 0x5a, 0x07, 0x71, 0x6d,
	0x71, 0x2f, 0x71, 0x6d, 0x71, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_data_proto_rawDescOnce sync.Once
	file_src_data_proto_rawDescData = file_src_data_proto_rawDesc
)

func file_src_data_proto_rawDescGZIP() []byte {
	file_src_data_proto_rawDescOnce.Do(func() {
		file_src_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_data_proto_rawDescData)
	})
	return file_src_data_proto_rawDescData
}

var file_src_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_src_data_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_src_data_proto_goTypes = []interface{}{
	(QMQLogLevelEnum)(0),        // 0: qmq.QMQLogLevelEnum
	(QMQLightBulbStateEnum)(0),  // 1: qmq.QMQLightBulbStateEnum
	(QMQDoorLockStateEnum)(0),   // 2: qmq.QMQDoorLockStateEnum
	(*QMQData)(nil),             // 3: qmq.QMQData
	(*QMQInt)(nil),              // 4: qmq.QMQInt
	(*QMQString)(nil),           // 5: qmq.QMQString
	(*QMQTimestamp)(nil),        // 6: qmq.QMQTimestamp
	(*QMQFloat)(nil),            // 7: qmq.QMQFloat
	(*QMQBool)(nil),             // 8: qmq.QMQBool
	(*QMQLog)(nil),              // 9: qmq.QMQLog
	(*QMQStreamContext)(nil),    // 10: qmq.QMQStreamContext
	(*QMQPrayer)(nil),           // 11: qmq.QMQPrayer
	(*QMQAudioRequest)(nil),     // 12: qmq.QMQAudioRequest
	(*QMQLightBulb)(nil),        // 13: qmq.QMQLightBulb
	(*QMQDoorLock)(nil),         // 14: qmq.QMQDoorLock
	(*any1.Any)(nil),            // 15: google.protobuf.Any
	(*timestamp.Timestamp)(nil), // 16: google.protobuf.Timestamp
}
var file_src_data_proto_depIdxs = []int32{
	15, // 0: qmq.QMQData.data:type_name -> google.protobuf.Any
	16, // 1: qmq.QMQData.writetime:type_name -> google.protobuf.Timestamp
	16, // 2: qmq.QMQTimestamp.value:type_name -> google.protobuf.Timestamp
	0,  // 3: qmq.QMQLog.level:type_name -> qmq.QMQLogLevelEnum
	16, // 4: qmq.QMQLog.timestamp:type_name -> google.protobuf.Timestamp
	16, // 5: qmq.QMQPrayer.time:type_name -> google.protobuf.Timestamp
	1,  // 6: qmq.QMQLightBulb.current_state:type_name -> qmq.QMQLightBulbStateEnum
	1,  // 7: qmq.QMQLightBulb.requested_state:type_name -> qmq.QMQLightBulbStateEnum
	2,  // 8: qmq.QMQDoorLock.current_state:type_name -> qmq.QMQDoorLockStateEnum
	2,  // 9: qmq.QMQDoorLock.requested_state:type_name -> qmq.QMQDoorLockStateEnum
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_src_data_proto_init() }
func file_src_data_proto_init() {
	if File_src_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_data_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQInt); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQString); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQTimestamp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQFloat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQBool); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQLog); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQStreamContext); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQPrayer); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQAudioRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQLightBulb); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_src_data_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QMQDoorLock); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_data_proto_goTypes,
		DependencyIndexes: file_src_data_proto_depIdxs,
		EnumInfos:         file_src_data_proto_enumTypes,
		MessageInfos:      file_src_data_proto_msgTypes,
	}.Build()
	File_src_data_proto = out.File
	file_src_data_proto_rawDesc = nil
	file_src_data_proto_goTypes = nil
	file_src_data_proto_depIdxs = nil
}
