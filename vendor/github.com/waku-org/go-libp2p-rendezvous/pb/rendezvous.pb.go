// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.12
// source: rendezvous.proto

package rendezvous_pb

import (
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

type Message_MessageType int32

const (
	Message_REGISTER          Message_MessageType = 0
	Message_REGISTER_RESPONSE Message_MessageType = 1
	Message_UNREGISTER        Message_MessageType = 2
	Message_DISCOVER          Message_MessageType = 3
	Message_DISCOVER_RESPONSE Message_MessageType = 4
)

// Enum value maps for Message_MessageType.
var (
	Message_MessageType_name = map[int32]string{
		0: "REGISTER",
		1: "REGISTER_RESPONSE",
		2: "UNREGISTER",
		3: "DISCOVER",
		4: "DISCOVER_RESPONSE",
	}
	Message_MessageType_value = map[string]int32{
		"REGISTER":          0,
		"REGISTER_RESPONSE": 1,
		"UNREGISTER":        2,
		"DISCOVER":          3,
		"DISCOVER_RESPONSE": 4,
	}
)

func (x Message_MessageType) Enum() *Message_MessageType {
	p := new(Message_MessageType)
	*p = x
	return p
}

func (x Message_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_rendezvous_proto_enumTypes[0].Descriptor()
}

func (Message_MessageType) Type() protoreflect.EnumType {
	return &file_rendezvous_proto_enumTypes[0]
}

func (x Message_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Message_MessageType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Message_MessageType(num)
	return nil
}

// Deprecated: Use Message_MessageType.Descriptor instead.
func (Message_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0, 0}
}

type Message_ResponseStatus int32

const (
	Message_OK                           Message_ResponseStatus = 0
	Message_E_INVALID_NAMESPACE          Message_ResponseStatus = 100
	Message_E_INVALID_SIGNED_PEER_RECORD Message_ResponseStatus = 101
	Message_E_INVALID_TTL                Message_ResponseStatus = 102
	Message_E_INVALID_COOKIE             Message_ResponseStatus = 103
	Message_E_NOT_AUTHORIZED             Message_ResponseStatus = 200
	Message_E_INTERNAL_ERROR             Message_ResponseStatus = 300
	Message_E_UNAVAILABLE                Message_ResponseStatus = 400
)

// Enum value maps for Message_ResponseStatus.
var (
	Message_ResponseStatus_name = map[int32]string{
		0:   "OK",
		100: "E_INVALID_NAMESPACE",
		101: "E_INVALID_SIGNED_PEER_RECORD",
		102: "E_INVALID_TTL",
		103: "E_INVALID_COOKIE",
		200: "E_NOT_AUTHORIZED",
		300: "E_INTERNAL_ERROR",
		400: "E_UNAVAILABLE",
	}
	Message_ResponseStatus_value = map[string]int32{
		"OK":                           0,
		"E_INVALID_NAMESPACE":          100,
		"E_INVALID_SIGNED_PEER_RECORD": 101,
		"E_INVALID_TTL":                102,
		"E_INVALID_COOKIE":             103,
		"E_NOT_AUTHORIZED":             200,
		"E_INTERNAL_ERROR":             300,
		"E_UNAVAILABLE":                400,
	}
)

func (x Message_ResponseStatus) Enum() *Message_ResponseStatus {
	p := new(Message_ResponseStatus)
	*p = x
	return p
}

func (x Message_ResponseStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Message_ResponseStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_rendezvous_proto_enumTypes[1].Descriptor()
}

func (Message_ResponseStatus) Type() protoreflect.EnumType {
	return &file_rendezvous_proto_enumTypes[1]
}

func (x Message_ResponseStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Message_ResponseStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Message_ResponseStatus(num)
	return nil
}

// Deprecated: Use Message_ResponseStatus.Descriptor instead.
func (Message_ResponseStatus) EnumDescriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0, 1}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type             *Message_MessageType      `protobuf:"varint,1,opt,name=type,enum=rendezvous.pb.Message_MessageType" json:"type,omitempty"`
	Register         *Message_Register         `protobuf:"bytes,2,opt,name=register" json:"register,omitempty"`
	RegisterResponse *Message_RegisterResponse `protobuf:"bytes,3,opt,name=registerResponse" json:"registerResponse,omitempty"`
	Unregister       *Message_Unregister       `protobuf:"bytes,4,opt,name=unregister" json:"unregister,omitempty"`
	Discover         *Message_Discover         `protobuf:"bytes,5,opt,name=discover" json:"discover,omitempty"`
	DiscoverResponse *Message_DiscoverResponse `protobuf:"bytes,6,opt,name=discoverResponse" json:"discoverResponse,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rendezvous_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_rendezvous_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetType() Message_MessageType {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Message_REGISTER
}

func (x *Message) GetRegister() *Message_Register {
	if x != nil {
		return x.Register
	}
	return nil
}

func (x *Message) GetRegisterResponse() *Message_RegisterResponse {
	if x != nil {
		return x.RegisterResponse
	}
	return nil
}

func (x *Message) GetUnregister() *Message_Unregister {
	if x != nil {
		return x.Unregister
	}
	return nil
}

func (x *Message) GetDiscover() *Message_Discover {
	if x != nil {
		return x.Discover
	}
	return nil
}

func (x *Message) GetDiscoverResponse() *Message_DiscoverResponse {
	if x != nil {
		return x.DiscoverResponse
	}
	return nil
}

type Message_Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns               *string `protobuf:"bytes,1,opt,name=ns" json:"ns,omitempty"`
	SignedPeerRecord []byte  `protobuf:"bytes,2,opt,name=signedPeerRecord" json:"signedPeerRecord,omitempty"`
	Ttl              *uint64 `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"` // in seconds
}

func (x *Message_Register) Reset() {
	*x = Message_Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rendezvous_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Register) ProtoMessage() {}

func (x *Message_Register) ProtoReflect() protoreflect.Message {
	mi := &file_rendezvous_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Register.ProtoReflect.Descriptor instead.
func (*Message_Register) Descriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Message_Register) GetNs() string {
	if x != nil && x.Ns != nil {
		return *x.Ns
	}
	return ""
}

func (x *Message_Register) GetSignedPeerRecord() []byte {
	if x != nil {
		return x.SignedPeerRecord
	}
	return nil
}

func (x *Message_Register) GetTtl() uint64 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

type Message_RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *Message_ResponseStatus `protobuf:"varint,1,opt,name=status,enum=rendezvous.pb.Message_ResponseStatus" json:"status,omitempty"`
	StatusText *string                 `protobuf:"bytes,2,opt,name=statusText" json:"statusText,omitempty"`
	Ttl        *uint64                 `protobuf:"varint,3,opt,name=ttl" json:"ttl,omitempty"` // in seconds
}

func (x *Message_RegisterResponse) Reset() {
	*x = Message_RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rendezvous_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_RegisterResponse) ProtoMessage() {}

func (x *Message_RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rendezvous_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_RegisterResponse.ProtoReflect.Descriptor instead.
func (*Message_RegisterResponse) Descriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Message_RegisterResponse) GetStatus() Message_ResponseStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Message_OK
}

func (x *Message_RegisterResponse) GetStatusText() string {
	if x != nil && x.StatusText != nil {
		return *x.StatusText
	}
	return ""
}

func (x *Message_RegisterResponse) GetTtl() uint64 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

type Message_Unregister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns *string `protobuf:"bytes,1,opt,name=ns" json:"ns,omitempty"` // optional bytes id = 2; deprecated as per https://github.com/libp2p/specs/issues/335
}

func (x *Message_Unregister) Reset() {
	*x = Message_Unregister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rendezvous_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_Unregister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Unregister) ProtoMessage() {}

func (x *Message_Unregister) ProtoReflect() protoreflect.Message {
	mi := &file_rendezvous_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Unregister.ProtoReflect.Descriptor instead.
func (*Message_Unregister) Descriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Message_Unregister) GetNs() string {
	if x != nil && x.Ns != nil {
		return *x.Ns
	}
	return ""
}

type Message_Discover struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ns     *string `protobuf:"bytes,1,opt,name=ns" json:"ns,omitempty"`
	Limit  *uint64 `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Cookie []byte  `protobuf:"bytes,3,opt,name=cookie" json:"cookie,omitempty"`
}

func (x *Message_Discover) Reset() {
	*x = Message_Discover{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rendezvous_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_Discover) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_Discover) ProtoMessage() {}

func (x *Message_Discover) ProtoReflect() protoreflect.Message {
	mi := &file_rendezvous_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_Discover.ProtoReflect.Descriptor instead.
func (*Message_Discover) Descriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0, 3}
}

func (x *Message_Discover) GetNs() string {
	if x != nil && x.Ns != nil {
		return *x.Ns
	}
	return ""
}

func (x *Message_Discover) GetLimit() uint64 {
	if x != nil && x.Limit != nil {
		return *x.Limit
	}
	return 0
}

func (x *Message_Discover) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

type Message_DiscoverResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Registrations []*Message_Register     `protobuf:"bytes,1,rep,name=registrations" json:"registrations,omitempty"`
	Cookie        []byte                  `protobuf:"bytes,2,opt,name=cookie" json:"cookie,omitempty"`
	Status        *Message_ResponseStatus `protobuf:"varint,3,opt,name=status,enum=rendezvous.pb.Message_ResponseStatus" json:"status,omitempty"`
	StatusText    *string                 `protobuf:"bytes,4,opt,name=statusText" json:"statusText,omitempty"`
}

func (x *Message_DiscoverResponse) Reset() {
	*x = Message_DiscoverResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rendezvous_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message_DiscoverResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message_DiscoverResponse) ProtoMessage() {}

func (x *Message_DiscoverResponse) ProtoReflect() protoreflect.Message {
	mi := &file_rendezvous_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message_DiscoverResponse.ProtoReflect.Descriptor instead.
func (*Message_DiscoverResponse) Descriptor() ([]byte, []int) {
	return file_rendezvous_proto_rawDescGZIP(), []int{0, 4}
}

func (x *Message_DiscoverResponse) GetRegistrations() []*Message_Register {
	if x != nil {
		return x.Registrations
	}
	return nil
}

func (x *Message_DiscoverResponse) GetCookie() []byte {
	if x != nil {
		return x.Cookie
	}
	return nil
}

func (x *Message_DiscoverResponse) GetStatus() Message_ResponseStatus {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return Message_OK
}

func (x *Message_DiscoverResponse) GetStatusText() string {
	if x != nil && x.StatusText != nil {
		return *x.StatusText
	}
	return ""
}

var File_rendezvous_proto protoreflect.FileDescriptor

var file_rendezvous_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70,
	0x62, 0x22, 0xed, 0x09, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3b, 0x0a, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a,
	0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x12, 0x53, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x0a, 0x75, 0x6e, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x0a,
	0x75, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x3b, 0x0a, 0x08, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x72,
	0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x08, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x53, 0x0a, 0x10, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x27, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70,
	0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x10, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x58, 0x0a, 0x08,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x69, 0x67, 0x6e,
	0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x10, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x64, 0x50, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x1a, 0x83, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x72, 0x65,
	0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x74,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x03, 0x74, 0x74, 0x6c, 0x1a, 0x1c, 0x0a, 0x0a,
	0x55, 0x6e, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x1a, 0x48, 0x0a, 0x08, 0x44, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f,
	0x6f, 0x6b, 0x69, 0x65, 0x1a, 0xd0, 0x01, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0d, 0x72, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1f, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65, 0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x52, 0x0d, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x63, 0x6f, 0x6f, 0x6b, 0x69, 0x65, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x72, 0x65, 0x6e, 0x64, 0x65,
	0x7a, 0x76, 0x6f, 0x75, 0x73, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x54, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x54, 0x65, 0x78, 0x74, 0x22, 0x67, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54,
	0x45, 0x52, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52,
	0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x55,
	0x4e, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x44,
	0x49, 0x53, 0x43, 0x4f, 0x56, 0x45, 0x52, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x44, 0x49, 0x53,
	0x43, 0x4f, 0x56, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x04,
	0x22, 0xbe, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41,
	0x43, 0x45, 0x10, 0x64, 0x12, 0x20, 0x0a, 0x1c, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49,
	0x44, 0x5f, 0x53, 0x49, 0x47, 0x4e, 0x45, 0x44, 0x5f, 0x50, 0x45, 0x45, 0x52, 0x5f, 0x52, 0x45,
	0x43, 0x4f, 0x52, 0x44, 0x10, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x45, 0x5f, 0x49, 0x4e, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x5f, 0x54, 0x54, 0x4c, 0x10, 0x66, 0x12, 0x14, 0x0a, 0x10, 0x45, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x5f, 0x43, 0x4f, 0x4f, 0x4b, 0x49, 0x45, 0x10, 0x67, 0x12,
	0x15, 0x0a, 0x10, 0x45, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49,
	0x5a, 0x45, 0x44, 0x10, 0xc8, 0x01, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45,
	0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xac, 0x02, 0x12, 0x12, 0x0a,
	0x0d, 0x45, 0x5f, 0x55, 0x4e, 0x41, 0x56, 0x41, 0x49, 0x4c, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x90,
	0x03,
}

var (
	file_rendezvous_proto_rawDescOnce sync.Once
	file_rendezvous_proto_rawDescData = file_rendezvous_proto_rawDesc
)

func file_rendezvous_proto_rawDescGZIP() []byte {
	file_rendezvous_proto_rawDescOnce.Do(func() {
		file_rendezvous_proto_rawDescData = protoimpl.X.CompressGZIP(file_rendezvous_proto_rawDescData)
	})
	return file_rendezvous_proto_rawDescData
}

var file_rendezvous_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_rendezvous_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_rendezvous_proto_goTypes = []interface{}{
	(Message_MessageType)(0),         // 0: rendezvous.pb.Message.MessageType
	(Message_ResponseStatus)(0),      // 1: rendezvous.pb.Message.ResponseStatus
	(*Message)(nil),                  // 2: rendezvous.pb.Message
	(*Message_Register)(nil),         // 3: rendezvous.pb.Message.Register
	(*Message_RegisterResponse)(nil), // 4: rendezvous.pb.Message.RegisterResponse
	(*Message_Unregister)(nil),       // 5: rendezvous.pb.Message.Unregister
	(*Message_Discover)(nil),         // 6: rendezvous.pb.Message.Discover
	(*Message_DiscoverResponse)(nil), // 7: rendezvous.pb.Message.DiscoverResponse
}
var file_rendezvous_proto_depIdxs = []int32{
	0, // 0: rendezvous.pb.Message.type:type_name -> rendezvous.pb.Message.MessageType
	3, // 1: rendezvous.pb.Message.register:type_name -> rendezvous.pb.Message.Register
	4, // 2: rendezvous.pb.Message.registerResponse:type_name -> rendezvous.pb.Message.RegisterResponse
	5, // 3: rendezvous.pb.Message.unregister:type_name -> rendezvous.pb.Message.Unregister
	6, // 4: rendezvous.pb.Message.discover:type_name -> rendezvous.pb.Message.Discover
	7, // 5: rendezvous.pb.Message.discoverResponse:type_name -> rendezvous.pb.Message.DiscoverResponse
	1, // 6: rendezvous.pb.Message.RegisterResponse.status:type_name -> rendezvous.pb.Message.ResponseStatus
	3, // 7: rendezvous.pb.Message.DiscoverResponse.registrations:type_name -> rendezvous.pb.Message.Register
	1, // 8: rendezvous.pb.Message.DiscoverResponse.status:type_name -> rendezvous.pb.Message.ResponseStatus
	9, // [9:9] is the sub-list for method output_type
	9, // [9:9] is the sub-list for method input_type
	9, // [9:9] is the sub-list for extension type_name
	9, // [9:9] is the sub-list for extension extendee
	0, // [0:9] is the sub-list for field type_name
}

func init() { file_rendezvous_proto_init() }
func file_rendezvous_proto_init() {
	if File_rendezvous_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rendezvous_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_rendezvous_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_Register); i {
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
		file_rendezvous_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_RegisterResponse); i {
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
		file_rendezvous_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_Unregister); i {
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
		file_rendezvous_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_Discover); i {
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
		file_rendezvous_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message_DiscoverResponse); i {
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
			RawDescriptor: file_rendezvous_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rendezvous_proto_goTypes,
		DependencyIndexes: file_rendezvous_proto_depIdxs,
		EnumInfos:         file_rendezvous_proto_enumTypes,
		MessageInfos:      file_rendezvous_proto_msgTypes,
	}.Build()
	File_rendezvous_proto = out.File
	file_rendezvous_proto_rawDesc = nil
	file_rendezvous_proto_goTypes = nil
	file_rendezvous_proto_depIdxs = nil
}