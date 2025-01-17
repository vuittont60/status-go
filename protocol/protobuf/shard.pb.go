// Code generated by protoc-gen-go. DO NOT EDIT.
// source: shard.proto

package protobuf

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

type Shard struct {
	Cluster              int32    `protobuf:"varint,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	Index                int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Shard) Reset()         { *m = Shard{} }
func (m *Shard) String() string { return proto.CompactTextString(m) }
func (*Shard) ProtoMessage()    {}
func (*Shard) Descriptor() ([]byte, []int) {
	return fileDescriptor_319ea41e44cdc364, []int{0}
}

func (m *Shard) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Shard.Unmarshal(m, b)
}
func (m *Shard) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Shard.Marshal(b, m, deterministic)
}
func (m *Shard) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Shard.Merge(m, src)
}
func (m *Shard) XXX_Size() int {
	return xxx_messageInfo_Shard.Size(m)
}
func (m *Shard) XXX_DiscardUnknown() {
	xxx_messageInfo_Shard.DiscardUnknown(m)
}

var xxx_messageInfo_Shard proto.InternalMessageInfo

func (m *Shard) GetCluster() int32 {
	if m != nil {
		return m.Cluster
	}
	return 0
}

func (m *Shard) GetIndex() int32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func init() {
	proto.RegisterType((*Shard)(nil), "protobuf.Shard")
}

func init() {
	proto.RegisterFile("shard.proto", fileDescriptor_319ea41e44cdc364)
}

var fileDescriptor_319ea41e44cdc364 = []byte{
	// 99 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xce, 0x48, 0x2c,
	0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x53, 0x49, 0xa5, 0x69, 0x4a, 0xe6,
	0x5c, 0xac, 0xc1, 0x20, 0x09, 0x21, 0x09, 0x2e, 0xf6, 0xe4, 0x9c, 0xd2, 0xe2, 0x92, 0xd4, 0x22,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x18, 0x57, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25,
	0xb5, 0x42, 0x82, 0x09, 0x2c, 0x0e, 0xe1, 0x38, 0xf1, 0x46, 0x71, 0xeb, 0xe9, 0x5b, 0xc3, 0xcc,
	0x49, 0x62, 0x03, 0xb3, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xef, 0xcf, 0x9d, 0x21, 0x67,
	0x00, 0x00, 0x00,
}
