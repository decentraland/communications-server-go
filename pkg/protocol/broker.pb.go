// Code generated by protoc-gen-go. DO NOT EDIT.
// source: broker.proto

package protocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type MessageType int32

const (
	MessageType_UNKNOWN_MESSAGE_TYPE MessageType = 0
	MessageType_WELCOME              MessageType = 1
	MessageType_CONNECT              MessageType = 2
	MessageType_WEBRTC_OFFER         MessageType = 4
	MessageType_WEBRTC_ANSWER        MessageType = 5
	MessageType_WEBRTC_ICE_CANDIDATE MessageType = 6
	MessageType_PING                 MessageType = 7
	MessageType_TOPIC_SUBSCRIPTION   MessageType = 8
	MessageType_TOPIC                MessageType = 9
	MessageType_DATA                 MessageType = 10
	MessageType_AUTH                 MessageType = 11
)

var MessageType_name = map[int32]string{
	0:  "UNKNOWN_MESSAGE_TYPE",
	1:  "WELCOME",
	2:  "CONNECT",
	4:  "WEBRTC_OFFER",
	5:  "WEBRTC_ANSWER",
	6:  "WEBRTC_ICE_CANDIDATE",
	7:  "PING",
	8:  "TOPIC_SUBSCRIPTION",
	9:  "TOPIC",
	10: "DATA",
	11: "AUTH",
}

var MessageType_value = map[string]int32{
	"UNKNOWN_MESSAGE_TYPE": 0,
	"WELCOME":              1,
	"CONNECT":              2,
	"WEBRTC_OFFER":         4,
	"WEBRTC_ANSWER":        5,
	"WEBRTC_ICE_CANDIDATE": 6,
	"PING":                 7,
	"TOPIC_SUBSCRIPTION":   8,
	"TOPIC":                9,
	"DATA":                 10,
	"AUTH":                 11,
}

func (x MessageType) String() string {
	return proto.EnumName(MessageType_name, int32(x))
}

func (MessageType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{0}
}

type Role int32

const (
	Role_UNKNOWN_ROLE         Role = 0
	Role_CLIENT               Role = 1
	Role_COMMUNICATION_SERVER Role = 2
)

var Role_name = map[int32]string{
	0: "UNKNOWN_ROLE",
	1: "CLIENT",
	2: "COMMUNICATION_SERVER",
}

var Role_value = map[string]int32{
	"UNKNOWN_ROLE":         0,
	"CLIENT":               1,
	"COMMUNICATION_SERVER": 2,
}

func (x Role) String() string {
	return proto.EnumName(Role_name, int32(x))
}

func (Role) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{1}
}

type Format int32

const (
	Format_UNKNOWN_FORMAT Format = 0
	Format_PLAIN          Format = 1
	Format_GZIP           Format = 2
)

var Format_name = map[int32]string{
	0: "UNKNOWN_FORMAT",
	1: "PLAIN",
	2: "GZIP",
}

var Format_value = map[string]int32{
	"UNKNOWN_FORMAT": 0,
	"PLAIN":          1,
	"GZIP":           2,
}

func (x Format) String() string {
	return proto.EnumName(Format_name, int32(x))
}

func (Format) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{2}
}

type CoordinatorMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *CoordinatorMessage) Reset()         { *m = CoordinatorMessage{} }
func (m *CoordinatorMessage) String() string { return proto.CompactTextString(m) }
func (*CoordinatorMessage) ProtoMessage()    {}
func (*CoordinatorMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{0}
}

func (m *CoordinatorMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CoordinatorMessage.Unmarshal(m, b)
}
func (m *CoordinatorMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CoordinatorMessage.Marshal(b, m, deterministic)
}
func (m *CoordinatorMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinatorMessage.Merge(m, src)
}
func (m *CoordinatorMessage) XXX_Size() int {
	return xxx_messageInfo_CoordinatorMessage.Size(m)
}
func (m *CoordinatorMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinatorMessage.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinatorMessage proto.InternalMessageInfo

func (m *CoordinatorMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

type WelcomeMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	Alias                uint64      `protobuf:"varint,2,opt,name=alias,proto3" json:"alias,omitempty"`
	AvailableServers     []uint64    `protobuf:"varint,3,rep,packed,name=available_servers,json=availableServers,proto3" json:"available_servers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WelcomeMessage) Reset()         { *m = WelcomeMessage{} }
func (m *WelcomeMessage) String() string { return proto.CompactTextString(m) }
func (*WelcomeMessage) ProtoMessage()    {}
func (*WelcomeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{1}
}

func (m *WelcomeMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WelcomeMessage.Unmarshal(m, b)
}
func (m *WelcomeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WelcomeMessage.Marshal(b, m, deterministic)
}
func (m *WelcomeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WelcomeMessage.Merge(m, src)
}
func (m *WelcomeMessage) XXX_Size() int {
	return xxx_messageInfo_WelcomeMessage.Size(m)
}
func (m *WelcomeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WelcomeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WelcomeMessage proto.InternalMessageInfo

func (m *WelcomeMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *WelcomeMessage) GetAlias() uint64 {
	if m != nil {
		return m.Alias
	}
	return 0
}

func (m *WelcomeMessage) GetAvailableServers() []uint64 {
	if m != nil {
		return m.AvailableServers
	}
	return nil
}

type ConnectMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	FromAlias            uint64      `protobuf:"varint,2,opt,name=from_alias,json=fromAlias,proto3" json:"from_alias,omitempty"`
	ToAlias              uint64      `protobuf:"varint,3,opt,name=to_alias,json=toAlias,proto3" json:"to_alias,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConnectMessage) Reset()         { *m = ConnectMessage{} }
func (m *ConnectMessage) String() string { return proto.CompactTextString(m) }
func (*ConnectMessage) ProtoMessage()    {}
func (*ConnectMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{2}
}

func (m *ConnectMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConnectMessage.Unmarshal(m, b)
}
func (m *ConnectMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConnectMessage.Marshal(b, m, deterministic)
}
func (m *ConnectMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConnectMessage.Merge(m, src)
}
func (m *ConnectMessage) XXX_Size() int {
	return xxx_messageInfo_ConnectMessage.Size(m)
}
func (m *ConnectMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ConnectMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ConnectMessage proto.InternalMessageInfo

func (m *ConnectMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *ConnectMessage) GetFromAlias() uint64 {
	if m != nil {
		return m.FromAlias
	}
	return 0
}

func (m *ConnectMessage) GetToAlias() uint64 {
	if m != nil {
		return m.ToAlias
	}
	return 0
}

type WebRtcMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	FromAlias            uint64      `protobuf:"varint,2,opt,name=from_alias,json=fromAlias,proto3" json:"from_alias,omitempty"`
	ToAlias              uint64      `protobuf:"varint,3,opt,name=to_alias,json=toAlias,proto3" json:"to_alias,omitempty"`
	Sdp                  string      `protobuf:"bytes,4,opt,name=sdp,proto3" json:"sdp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *WebRtcMessage) Reset()         { *m = WebRtcMessage{} }
func (m *WebRtcMessage) String() string { return proto.CompactTextString(m) }
func (*WebRtcMessage) ProtoMessage()    {}
func (*WebRtcMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{3}
}

func (m *WebRtcMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WebRtcMessage.Unmarshal(m, b)
}
func (m *WebRtcMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WebRtcMessage.Marshal(b, m, deterministic)
}
func (m *WebRtcMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WebRtcMessage.Merge(m, src)
}
func (m *WebRtcMessage) XXX_Size() int {
	return xxx_messageInfo_WebRtcMessage.Size(m)
}
func (m *WebRtcMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_WebRtcMessage.DiscardUnknown(m)
}

var xxx_messageInfo_WebRtcMessage proto.InternalMessageInfo

func (m *WebRtcMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *WebRtcMessage) GetFromAlias() uint64 {
	if m != nil {
		return m.FromAlias
	}
	return 0
}

func (m *WebRtcMessage) GetToAlias() uint64 {
	if m != nil {
		return m.ToAlias
	}
	return 0
}

func (m *WebRtcMessage) GetSdp() string {
	if m != nil {
		return m.Sdp
	}
	return ""
}

type MessageHeader struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MessageHeader) Reset()         { *m = MessageHeader{} }
func (m *MessageHeader) String() string { return proto.CompactTextString(m) }
func (*MessageHeader) ProtoMessage()    {}
func (*MessageHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{4}
}

func (m *MessageHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageHeader.Unmarshal(m, b)
}
func (m *MessageHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageHeader.Marshal(b, m, deterministic)
}
func (m *MessageHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageHeader.Merge(m, src)
}
func (m *MessageHeader) XXX_Size() int {
	return xxx_messageInfo_MessageHeader.Size(m)
}
func (m *MessageHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageHeader.DiscardUnknown(m)
}

var xxx_messageInfo_MessageHeader proto.InternalMessageInfo

func (m *MessageHeader) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

type PingMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	Time                 float64     `protobuf:"fixed64,2,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{5}
}

func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (m *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(m, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *PingMessage) GetTime() float64 {
	if m != nil {
		return m.Time
	}
	return 0
}

// NOTE: topics is a space separated string in the format specified by Format
type TopicSubscriptionMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	Format               Format      `protobuf:"varint,2,opt,name=format,proto3,enum=protocol.Format" json:"format,omitempty"`
	Topics               []byte      `protobuf:"bytes,3,opt,name=topics,proto3" json:"topics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TopicSubscriptionMessage) Reset()         { *m = TopicSubscriptionMessage{} }
func (m *TopicSubscriptionMessage) String() string { return proto.CompactTextString(m) }
func (*TopicSubscriptionMessage) ProtoMessage()    {}
func (*TopicSubscriptionMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{6}
}

func (m *TopicSubscriptionMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicSubscriptionMessage.Unmarshal(m, b)
}
func (m *TopicSubscriptionMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicSubscriptionMessage.Marshal(b, m, deterministic)
}
func (m *TopicSubscriptionMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicSubscriptionMessage.Merge(m, src)
}
func (m *TopicSubscriptionMessage) XXX_Size() int {
	return xxx_messageInfo_TopicSubscriptionMessage.Size(m)
}
func (m *TopicSubscriptionMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicSubscriptionMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TopicSubscriptionMessage proto.InternalMessageInfo

func (m *TopicSubscriptionMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *TopicSubscriptionMessage) GetFormat() Format {
	if m != nil {
		return m.Format
	}
	return Format_UNKNOWN_FORMAT
}

func (m *TopicSubscriptionMessage) GetTopics() []byte {
	if m != nil {
		return m.Topics
	}
	return nil
}

type TopicMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	FromAlias            uint64      `protobuf:"varint,2,opt,name=from_alias,json=fromAlias,proto3" json:"from_alias,omitempty"`
	Topic                string      `protobuf:"bytes,3,opt,name=topic,proto3" json:"topic,omitempty"`
	Body                 []byte      `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *TopicMessage) Reset()         { *m = TopicMessage{} }
func (m *TopicMessage) String() string { return proto.CompactTextString(m) }
func (*TopicMessage) ProtoMessage()    {}
func (*TopicMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{7}
}

func (m *TopicMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TopicMessage.Unmarshal(m, b)
}
func (m *TopicMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TopicMessage.Marshal(b, m, deterministic)
}
func (m *TopicMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopicMessage.Merge(m, src)
}
func (m *TopicMessage) XXX_Size() int {
	return xxx_messageInfo_TopicMessage.Size(m)
}
func (m *TopicMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_TopicMessage.DiscardUnknown(m)
}

var xxx_messageInfo_TopicMessage proto.InternalMessageInfo

func (m *TopicMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *TopicMessage) GetFromAlias() uint64 {
	if m != nil {
		return m.FromAlias
	}
	return 0
}

func (m *TopicMessage) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *TopicMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type DataMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	FromAlias            uint64      `protobuf:"varint,2,opt,name=from_alias,json=fromAlias,proto3" json:"from_alias,omitempty"`
	Body                 []byte      `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *DataMessage) Reset()         { *m = DataMessage{} }
func (m *DataMessage) String() string { return proto.CompactTextString(m) }
func (*DataMessage) ProtoMessage()    {}
func (*DataMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{8}
}

func (m *DataMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataMessage.Unmarshal(m, b)
}
func (m *DataMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataMessage.Marshal(b, m, deterministic)
}
func (m *DataMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataMessage.Merge(m, src)
}
func (m *DataMessage) XXX_Size() int {
	return xxx_messageInfo_DataMessage.Size(m)
}
func (m *DataMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_DataMessage.DiscardUnknown(m)
}

var xxx_messageInfo_DataMessage proto.InternalMessageInfo

func (m *DataMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *DataMessage) GetFromAlias() uint64 {
	if m != nil {
		return m.FromAlias
	}
	return 0
}

func (m *DataMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

type AuthMessage struct {
	Type                 MessageType `protobuf:"varint,1,opt,name=type,proto3,enum=protocol.MessageType" json:"type,omitempty"`
	Role                 Role        `protobuf:"varint,2,opt,name=role,proto3,enum=protocol.Role" json:"role,omitempty"`
	Method               string      `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	Body                 []byte      `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *AuthMessage) Reset()         { *m = AuthMessage{} }
func (m *AuthMessage) String() string { return proto.CompactTextString(m) }
func (*AuthMessage) ProtoMessage()    {}
func (*AuthMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_f209535e190f2bed, []int{9}
}

func (m *AuthMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthMessage.Unmarshal(m, b)
}
func (m *AuthMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthMessage.Marshal(b, m, deterministic)
}
func (m *AuthMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthMessage.Merge(m, src)
}
func (m *AuthMessage) XXX_Size() int {
	return xxx_messageInfo_AuthMessage.Size(m)
}
func (m *AuthMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AuthMessage proto.InternalMessageInfo

func (m *AuthMessage) GetType() MessageType {
	if m != nil {
		return m.Type
	}
	return MessageType_UNKNOWN_MESSAGE_TYPE
}

func (m *AuthMessage) GetRole() Role {
	if m != nil {
		return m.Role
	}
	return Role_UNKNOWN_ROLE
}

func (m *AuthMessage) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *AuthMessage) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func init() {
	proto.RegisterEnum("protocol.MessageType", MessageType_name, MessageType_value)
	proto.RegisterEnum("protocol.Role", Role_name, Role_value)
	proto.RegisterEnum("protocol.Format", Format_name, Format_value)
	proto.RegisterType((*CoordinatorMessage)(nil), "protocol.CoordinatorMessage")
	proto.RegisterType((*WelcomeMessage)(nil), "protocol.WelcomeMessage")
	proto.RegisterType((*ConnectMessage)(nil), "protocol.ConnectMessage")
	proto.RegisterType((*WebRtcMessage)(nil), "protocol.WebRtcMessage")
	proto.RegisterType((*MessageHeader)(nil), "protocol.MessageHeader")
	proto.RegisterType((*PingMessage)(nil), "protocol.PingMessage")
	proto.RegisterType((*TopicSubscriptionMessage)(nil), "protocol.TopicSubscriptionMessage")
	proto.RegisterType((*TopicMessage)(nil), "protocol.TopicMessage")
	proto.RegisterType((*DataMessage)(nil), "protocol.DataMessage")
	proto.RegisterType((*AuthMessage)(nil), "protocol.AuthMessage")
}

func init() { proto.RegisterFile("broker.proto", fileDescriptor_f209535e190f2bed) }

var fileDescriptor_f209535e190f2bed = []byte{
	// 616 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xcf, 0x6e, 0xda, 0x4a,
	0x14, 0xc6, 0x33, 0x60, 0x08, 0x1c, 0x08, 0x9a, 0x8c, 0x72, 0x23, 0xee, 0xe2, 0x4a, 0xc8, 0x2b,
	0x6e, 0x2a, 0x45, 0x6a, 0xba, 0xeb, 0xa2, 0x95, 0x33, 0x0c, 0x89, 0x55, 0xb0, 0xad, 0xb1, 0x29,
	0x6a, 0x37, 0x96, 0x81, 0x49, 0x62, 0xc5, 0x30, 0xc8, 0x9e, 0x44, 0xca, 0xa6, 0x8b, 0x2e, 0xda,
	0xbe, 0x56, 0xdf, 0xac, 0x9a, 0xc1, 0x34, 0x59, 0x74, 0x13, 0x14, 0x75, 0xc5, 0xf9, 0xf3, 0x71,
	0x7e, 0xdf, 0x39, 0x63, 0x68, 0xcf, 0x72, 0x79, 0x2b, 0xf2, 0xd3, 0x75, 0x2e, 0x95, 0x24, 0x0d,
	0xf3, 0x33, 0x97, 0x99, 0xfd, 0x1e, 0x08, 0x95, 0x32, 0x5f, 0xa4, 0xab, 0x44, 0xc9, 0x7c, 0x2c,
	0x8a, 0x22, 0xb9, 0x16, 0xe4, 0x7f, 0xb0, 0xd4, 0xc3, 0x5a, 0x74, 0x51, 0x0f, 0xf5, 0x3b, 0x67,
	0xff, 0x9c, 0x6e, 0xe5, 0xa7, 0xa5, 0x20, 0x7a, 0x58, 0x0b, 0x6e, 0x24, 0xf6, 0x17, 0xe8, 0x4c,
	0x45, 0x36, 0x97, 0x4b, 0xf1, 0xfc, 0x3f, 0x93, 0x23, 0xa8, 0x25, 0x59, 0x9a, 0x14, 0xdd, 0x4a,
	0x0f, 0xf5, 0x2d, 0xbe, 0x49, 0xc8, 0x2b, 0x38, 0x4c, 0xee, 0x93, 0x34, 0x4b, 0x66, 0x99, 0x88,
	0x0b, 0x91, 0xdf, 0x8b, 0xbc, 0xe8, 0x56, 0x7b, 0xd5, 0xbe, 0xc5, 0xf1, 0xef, 0x46, 0xb8, 0xa9,
	0xdb, 0x77, 0xd0, 0xa1, 0x72, 0xb5, 0x12, 0x73, 0xb5, 0x03, 0xff, 0x3f, 0x80, 0xab, 0x5c, 0x2e,
	0xe3, 0xa7, 0x26, 0x9a, 0xba, 0xe2, 0x18, 0x23, 0xff, 0x42, 0x43, 0xc9, 0xb2, 0x59, 0x35, 0xcd,
	0x7d, 0x25, 0x4d, 0xcb, 0xfe, 0x86, 0xe0, 0x60, 0x2a, 0x66, 0x5c, 0xcd, 0xff, 0x26, 0x96, 0x60,
	0xa8, 0x16, 0x8b, 0x75, 0xd7, 0xea, 0xa1, 0x7e, 0x93, 0xeb, 0xd0, 0x7e, 0x0b, 0x07, 0x25, 0xe0,
	0x52, 0x24, 0x0b, 0x91, 0x3f, 0xe7, 0xed, 0x46, 0xd0, 0x0a, 0xd2, 0xd5, 0xf5, 0x0e, 0x1b, 0x10,
	0xb0, 0x54, 0xba, 0x14, 0xc6, 0x3b, 0xe2, 0x26, 0xb6, 0xbf, 0x23, 0xe8, 0x46, 0x72, 0x9d, 0xce,
	0xc3, 0xbb, 0x59, 0x31, 0xcf, 0xd3, 0xb5, 0x4a, 0xe5, 0x6a, 0x87, 0xd9, 0x7d, 0xa8, 0x5f, 0xc9,
	0x7c, 0x99, 0x28, 0x33, 0xbd, 0x73, 0x86, 0x1f, 0xc5, 0x43, 0x53, 0xe7, 0x65, 0x9f, 0x1c, 0x43,
	0x5d, 0x69, 0xe0, 0xe6, 0x4c, 0x6d, 0x5e, 0x66, 0xf6, 0x57, 0x04, 0x6d, 0xe3, 0xe4, 0xe5, 0xdf,
	0xe6, 0x08, 0x6a, 0x06, 0x62, 0x88, 0x4d, 0xbe, 0x49, 0xf4, 0x39, 0x66, 0x72, 0xf1, 0x60, 0xde,
	0xa5, 0xcd, 0x4d, 0x6c, 0xdf, 0x42, 0x6b, 0x90, 0xa8, 0xe4, 0xe5, 0x2d, 0x6c, 0x61, 0xd5, 0x27,
	0xb0, 0x1f, 0x08, 0x5a, 0xce, 0x9d, 0xba, 0xd9, 0x81, 0x66, 0x83, 0x95, 0xcb, 0x4c, 0x94, 0xc7,
	0xee, 0x3c, 0x4a, 0xb9, 0xcc, 0x04, 0x37, 0x3d, 0x7d, 0xe8, 0xa5, 0x50, 0x37, 0x72, 0x51, 0xae,
	0x5d, 0x66, 0x7f, 0xda, 0xfb, 0xe4, 0x27, 0x82, 0xd6, 0x13, 0x0a, 0xe9, 0xc2, 0xd1, 0xc4, 0xfb,
	0xe0, 0xf9, 0x53, 0x2f, 0x1e, 0xb3, 0x30, 0x74, 0x2e, 0x58, 0x1c, 0x7d, 0x0a, 0x18, 0xde, 0x23,
	0x2d, 0xd8, 0x9f, 0xb2, 0x11, 0xf5, 0xc7, 0x0c, 0x23, 0x9d, 0x50, 0xdf, 0xf3, 0x18, 0x8d, 0x70,
	0x85, 0x60, 0x68, 0x4f, 0xd9, 0x39, 0x8f, 0x68, 0xec, 0x0f, 0x87, 0x8c, 0x63, 0x8b, 0x1c, 0xc2,
	0x41, 0x59, 0x71, 0xbc, 0x70, 0xca, 0x38, 0xae, 0xe9, 0xc1, 0x65, 0xc9, 0xa5, 0x2c, 0xa6, 0x8e,
	0x37, 0x70, 0x07, 0x4e, 0xc4, 0x70, 0x9d, 0x34, 0xc0, 0x0a, 0x5c, 0xef, 0x02, 0xef, 0x93, 0x63,
	0x20, 0x91, 0x1f, 0xb8, 0x34, 0x0e, 0x27, 0xe7, 0x21, 0xe5, 0x6e, 0x10, 0xb9, 0xbe, 0x87, 0x1b,
	0xa4, 0x09, 0x35, 0x53, 0xc7, 0x4d, 0x2d, 0x1e, 0x38, 0x91, 0x83, 0x41, 0x47, 0xce, 0x24, 0xba,
	0xc4, 0xad, 0x93, 0x77, 0x60, 0xe9, 0xed, 0xb5, 0x8f, 0xad, 0x77, 0xee, 0x8f, 0xb4, 0x67, 0x80,
	0x3a, 0x1d, 0xb9, 0xcc, 0x8b, 0x30, 0xd2, 0x06, 0xa8, 0x3f, 0x1e, 0x4f, 0x3c, 0x97, 0x3a, 0x7a,
	0x6e, 0x1c, 0x32, 0xfe, 0x91, 0x71, 0x5c, 0x39, 0x79, 0x0d, 0xf5, 0xcd, 0xa7, 0x4a, 0x08, 0x74,
	0xb6, 0x13, 0x86, 0x3e, 0x1f, 0x3b, 0x11, 0xde, 0xd3, 0xf0, 0x60, 0xe4, 0xb8, 0x1e, 0x46, 0x1a,
	0x79, 0xf1, 0xd9, 0x0d, 0x70, 0x65, 0x56, 0x37, 0x77, 0x7f, 0xf3, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0xa9, 0xc4, 0x6d, 0x3c, 0xa9, 0x05, 0x00, 0x00,
}