// Code generated by protoc-gen-go. DO NOT EDIT.
// source: url_data.proto

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

type Community struct {
	DisplayName          string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	MembersCount         uint32   `protobuf:"varint,3,opt,name=members_count,json=membersCount,proto3" json:"members_count,omitempty"`
	Color                string   `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	TagIndices           []uint32 `protobuf:"varint,5,rep,packed,name=tag_indices,json=tagIndices,proto3" json:"tag_indices,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Community) Reset()         { *m = Community{} }
func (m *Community) String() string { return proto.CompactTextString(m) }
func (*Community) ProtoMessage()    {}
func (*Community) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1e15b5f0115710, []int{0}
}

func (m *Community) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Community.Unmarshal(m, b)
}
func (m *Community) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Community.Marshal(b, m, deterministic)
}
func (m *Community) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Community.Merge(m, src)
}
func (m *Community) XXX_Size() int {
	return xxx_messageInfo_Community.Size(m)
}
func (m *Community) XXX_DiscardUnknown() {
	xxx_messageInfo_Community.DiscardUnknown(m)
}

var xxx_messageInfo_Community proto.InternalMessageInfo

func (m *Community) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Community) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Community) GetMembersCount() uint32 {
	if m != nil {
		return m.MembersCount
	}
	return 0
}

func (m *Community) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Community) GetTagIndices() []uint32 {
	if m != nil {
		return m.TagIndices
	}
	return nil
}

type Channel struct {
	DisplayName          string     `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description          string     `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Emoji                string     `protobuf:"bytes,3,opt,name=emoji,proto3" json:"emoji,omitempty"`
	Color                string     `protobuf:"bytes,4,opt,name=color,proto3" json:"color,omitempty"`
	Community            *Community `protobuf:"bytes,5,opt,name=community,proto3" json:"community,omitempty"`
	Uuid                 string     `protobuf:"bytes,6,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Channel) Reset()         { *m = Channel{} }
func (m *Channel) String() string { return proto.CompactTextString(m) }
func (*Channel) ProtoMessage()    {}
func (*Channel) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1e15b5f0115710, []int{1}
}

func (m *Channel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Channel.Unmarshal(m, b)
}
func (m *Channel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Channel.Marshal(b, m, deterministic)
}
func (m *Channel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Channel.Merge(m, src)
}
func (m *Channel) XXX_Size() int {
	return xxx_messageInfo_Channel.Size(m)
}
func (m *Channel) XXX_DiscardUnknown() {
	xxx_messageInfo_Channel.DiscardUnknown(m)
}

var xxx_messageInfo_Channel proto.InternalMessageInfo

func (m *Channel) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Channel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Channel) GetEmoji() string {
	if m != nil {
		return m.Emoji
	}
	return ""
}

func (m *Channel) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

func (m *Channel) GetCommunity() *Community {
	if m != nil {
		return m.Community
	}
	return nil
}

func (m *Channel) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

type User struct {
	DisplayName          string   `protobuf:"bytes,1,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Color                string   `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1e15b5f0115710, []int{2}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *User) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *User) GetColor() string {
	if m != nil {
		return m.Color
	}
	return ""
}

type URLData struct {
	// Community, Channel, or User
	Content              []byte   `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	Shard                *Shard   `protobuf:"bytes,2,opt,name=shard,proto3" json:"shard,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *URLData) Reset()         { *m = URLData{} }
func (m *URLData) String() string { return proto.CompactTextString(m) }
func (*URLData) ProtoMessage()    {}
func (*URLData) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f1e15b5f0115710, []int{3}
}

func (m *URLData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_URLData.Unmarshal(m, b)
}
func (m *URLData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_URLData.Marshal(b, m, deterministic)
}
func (m *URLData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_URLData.Merge(m, src)
}
func (m *URLData) XXX_Size() int {
	return xxx_messageInfo_URLData.Size(m)
}
func (m *URLData) XXX_DiscardUnknown() {
	xxx_messageInfo_URLData.DiscardUnknown(m)
}

var xxx_messageInfo_URLData proto.InternalMessageInfo

func (m *URLData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *URLData) GetShard() *Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func init() {
	proto.RegisterType((*Community)(nil), "protobuf.Community")
	proto.RegisterType((*Channel)(nil), "protobuf.Channel")
	proto.RegisterType((*User)(nil), "protobuf.User")
	proto.RegisterType((*URLData)(nil), "protobuf.URLData")
}

func init() {
	proto.RegisterFile("url_data.proto", fileDescriptor_5f1e15b5f0115710)
}

var fileDescriptor_5f1e15b5f0115710 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x40, 0x15, 0xda, 0xb4, 0xe4, 0xdc, 0x80, 0x64, 0x3a, 0x58, 0x2c, 0x84, 0x20, 0xa4, 0x4c,
	0x41, 0x94, 0x91, 0x8d, 0xb2, 0x80, 0x10, 0x83, 0x51, 0x17, 0x96, 0xc8, 0x4d, 0x4c, 0x6b, 0x14,
	0xdb, 0x95, 0xe3, 0x0c, 0xfd, 0x27, 0x7e, 0x82, 0x3f, 0x43, 0xb5, 0x1b, 0xc2, 0xc2, 0xd6, 0xc9,
	0x77, 0xef, 0x4e, 0xe7, 0xe7, 0x93, 0xe1, 0xa4, 0x35, 0x75, 0x51, 0x31, 0xcb, 0xf2, 0x8d, 0xd1,
	0x56, 0xe3, 0x63, 0x77, 0x2c, 0xdb, 0x8f, 0x73, 0xd4, 0xac, 0x99, 0xa9, 0x3c, 0x4e, 0xbf, 0x02,
	0x88, 0xe6, 0x5a, 0xca, 0x56, 0x09, 0xbb, 0xc5, 0x97, 0x30, 0xa9, 0x44, 0xb3, 0xa9, 0xd9, 0xb6,
	0x50, 0x4c, 0x72, 0x12, 0x24, 0x41, 0x16, 0x51, 0xb4, 0x67, 0xaf, 0x4c, 0x72, 0x9c, 0x00, 0xaa,
	0x78, 0x53, 0x1a, 0xb1, 0xb1, 0x42, 0x2b, 0x72, 0xb4, 0xef, 0xe8, 0x11, 0xbe, 0x82, 0x58, 0x72,
	0xb9, 0xe4, 0xa6, 0x29, 0x4a, 0xdd, 0x2a, 0x4b, 0x06, 0x49, 0x90, 0xc5, 0x74, 0xb2, 0x87, 0xf3,
	0x1d, 0xc3, 0x53, 0x08, 0x4b, 0x5d, 0x6b, 0x43, 0x86, 0x6e, 0x80, 0x4f, 0xf0, 0x05, 0x20, 0xcb,
	0x56, 0x85, 0x50, 0x95, 0x28, 0x79, 0x43, 0xc2, 0x64, 0x90, 0xc5, 0x14, 0x2c, 0x5b, 0x3d, 0x79,
	0x92, 0x7e, 0x07, 0x30, 0x9e, 0xaf, 0x99, 0x52, 0xbc, 0x3e, 0x8c, 0xec, 0x14, 0x42, 0x2e, 0xf5,
	0xa7, 0x70, 0x92, 0x11, 0xf5, 0xc9, 0x3f, 0x76, 0xb7, 0x10, 0x95, 0xdd, 0xaa, 0x48, 0x98, 0x04,
	0x19, 0x9a, 0x9d, 0xe5, 0xdd, 0x5a, 0xf3, 0xdf, 0x2d, 0xd2, 0xbe, 0x0b, 0x63, 0x18, 0xb6, 0xad,
	0xa8, 0xc8, 0xc8, 0xcd, 0x71, 0x71, 0xca, 0x60, 0xb8, 0x68, 0xb8, 0x39, 0x98, 0xbf, 0x37, 0x1d,
	0xfc, 0x31, 0x4d, 0x9f, 0x61, 0xbc, 0xa0, 0x2f, 0x8f, 0xcc, 0x32, 0x4c, 0x60, 0x5c, 0x6a, 0x65,
	0xb9, 0xb2, 0xee, 0x82, 0x09, 0xed, 0x52, 0x7c, 0x0d, 0xa1, 0xfb, 0x09, 0x6e, 0x2c, 0x9a, 0x9d,
	0xf6, 0x4f, 0x79, 0xdb, 0x61, 0xea, 0xab, 0x0f, 0xf1, 0x3b, 0xca, 0x6f, 0xee, 0xbb, 0xda, 0x72,
	0xe4, 0xa2, 0xbb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71, 0x2e, 0xf3, 0xde, 0x60, 0x02, 0x00,
	0x00,
}
