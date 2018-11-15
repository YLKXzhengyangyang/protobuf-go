// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comments/comments.proto

// COMMENT: package goproto.protoc.comments;

package comments

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	prototype "github.com/golang/protobuf/v2/reflect/prototype"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
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

// COMMENT: Message1
type Message1 struct {
	// COMMENT: Field1A
	Field1A *string `protobuf:"bytes,1,opt,name=Field1A" json:"Field1A,omitempty"`
	// COMMENT: Oneof1A
	//
	// Types that are valid to be assigned to Oneof1A:
	// COMMENT: Oneof1AField1
	//	*Message1_Oneof1AField1
	Oneof1A              isMessage1_Oneof1A `protobuf_oneof:"Oneof1a"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

type xxx_Message1 struct{ m *Message1 }

func (m *Message1) ProtoReflect() protoreflect.Message {
	return xxx_Message1{m}
}
func (m xxx_Message1) Type() protoreflect.MessageType {
	return xxx_Comments_ProtoFile_MessageTypes[0].Type
}
func (m xxx_Message1) KnownFields() protoreflect.KnownFields {
	return xxx_Comments_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_Message1) UnknownFields() protoreflect.UnknownFields {
	return xxx_Comments_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_Message1) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_Message1) ProtoMutable() {}

func (m *Message1) Reset()         { *m = Message1{} }
func (m *Message1) String() string { return proto.CompactTextString(m) }
func (*Message1) ProtoMessage()    {}
func (*Message1) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{0}
}

func (m *Message1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1.Unmarshal(m, b)
}
func (m *Message1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1.Marshal(b, m, deterministic)
}
func (m *Message1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1.Merge(m, src)
}
func (m *Message1) XXX_Size() int {
	return xxx_messageInfo_Message1.Size(m)
}
func (m *Message1) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1.DiscardUnknown(m)
}

var xxx_messageInfo_Message1 proto.InternalMessageInfo

func (m *Message1) GetField1A() string {
	if m != nil && m.Field1A != nil {
		return *m.Field1A
	}
	return ""
}

type isMessage1_Oneof1A interface {
	isMessage1_Oneof1A()
}

type Message1_Oneof1AField1 struct {
	Oneof1AField1 string `protobuf:"bytes,2,opt,name=Oneof1AField1,oneof"`
}

func (*Message1_Oneof1AField1) isMessage1_Oneof1A() {}

func (m *Message1) GetOneof1A() isMessage1_Oneof1A {
	if m != nil {
		return m.Oneof1A
	}
	return nil
}

func (m *Message1) GetOneof1AField1() string {
	if x, ok := m.GetOneof1A().(*Message1_Oneof1AField1); ok {
		return x.Oneof1AField1
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Message1) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Message1_Oneof1AField1)(nil),
	}
}

// COMMENT: Message1A
type Message1_Message1A struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Message1_Message1A struct{ m *Message1_Message1A }

func (m *Message1_Message1A) ProtoReflect() protoreflect.Message {
	return xxx_Message1_Message1A{m}
}
func (m xxx_Message1_Message1A) Type() protoreflect.MessageType {
	return xxx_Comments_ProtoFile_MessageTypes[2].Type
}
func (m xxx_Message1_Message1A) KnownFields() protoreflect.KnownFields {
	return xxx_Comments_ProtoFile_MessageTypes[2].KnownFieldsOf(m.m)
}
func (m xxx_Message1_Message1A) UnknownFields() protoreflect.UnknownFields {
	return xxx_Comments_ProtoFile_MessageTypes[2].UnknownFieldsOf(m.m)
}
func (m xxx_Message1_Message1A) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_Message1_Message1A) ProtoMutable() {}

func (m *Message1_Message1A) Reset()         { *m = Message1_Message1A{} }
func (m *Message1_Message1A) String() string { return proto.CompactTextString(m) }
func (*Message1_Message1A) ProtoMessage()    {}
func (*Message1_Message1A) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{0, 0}
}

func (m *Message1_Message1A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1_Message1A.Unmarshal(m, b)
}
func (m *Message1_Message1A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1_Message1A.Marshal(b, m, deterministic)
}
func (m *Message1_Message1A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1_Message1A.Merge(m, src)
}
func (m *Message1_Message1A) XXX_Size() int {
	return xxx_messageInfo_Message1_Message1A.Size(m)
}
func (m *Message1_Message1A) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1_Message1A.DiscardUnknown(m)
}

var xxx_messageInfo_Message1_Message1A proto.InternalMessageInfo

// COMMENT: Message1B
type Message1_Message1B struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Message1_Message1B struct{ m *Message1_Message1B }

func (m *Message1_Message1B) ProtoReflect() protoreflect.Message {
	return xxx_Message1_Message1B{m}
}
func (m xxx_Message1_Message1B) Type() protoreflect.MessageType {
	return xxx_Comments_ProtoFile_MessageTypes[3].Type
}
func (m xxx_Message1_Message1B) KnownFields() protoreflect.KnownFields {
	return xxx_Comments_ProtoFile_MessageTypes[3].KnownFieldsOf(m.m)
}
func (m xxx_Message1_Message1B) UnknownFields() protoreflect.UnknownFields {
	return xxx_Comments_ProtoFile_MessageTypes[3].UnknownFieldsOf(m.m)
}
func (m xxx_Message1_Message1B) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_Message1_Message1B) ProtoMutable() {}

func (m *Message1_Message1B) Reset()         { *m = Message1_Message1B{} }
func (m *Message1_Message1B) String() string { return proto.CompactTextString(m) }
func (*Message1_Message1B) ProtoMessage()    {}
func (*Message1_Message1B) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{0, 1}
}

func (m *Message1_Message1B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message1_Message1B.Unmarshal(m, b)
}
func (m *Message1_Message1B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message1_Message1B.Marshal(b, m, deterministic)
}
func (m *Message1_Message1B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message1_Message1B.Merge(m, src)
}
func (m *Message1_Message1B) XXX_Size() int {
	return xxx_messageInfo_Message1_Message1B.Size(m)
}
func (m *Message1_Message1B) XXX_DiscardUnknown() {
	xxx_messageInfo_Message1_Message1B.DiscardUnknown(m)
}

var xxx_messageInfo_Message1_Message1B proto.InternalMessageInfo

// COMMENT: Message2
type Message2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Message2 struct{ m *Message2 }

func (m *Message2) ProtoReflect() protoreflect.Message {
	return xxx_Message2{m}
}
func (m xxx_Message2) Type() protoreflect.MessageType {
	return xxx_Comments_ProtoFile_MessageTypes[1].Type
}
func (m xxx_Message2) KnownFields() protoreflect.KnownFields {
	return xxx_Comments_ProtoFile_MessageTypes[1].KnownFieldsOf(m.m)
}
func (m xxx_Message2) UnknownFields() protoreflect.UnknownFields {
	return xxx_Comments_ProtoFile_MessageTypes[1].UnknownFieldsOf(m.m)
}
func (m xxx_Message2) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_Message2) ProtoMutable() {}

func (m *Message2) Reset()         { *m = Message2{} }
func (m *Message2) String() string { return proto.CompactTextString(m) }
func (*Message2) ProtoMessage()    {}
func (*Message2) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{1}
}

func (m *Message2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2.Unmarshal(m, b)
}
func (m *Message2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2.Marshal(b, m, deterministic)
}
func (m *Message2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2.Merge(m, src)
}
func (m *Message2) XXX_Size() int {
	return xxx_messageInfo_Message2.Size(m)
}
func (m *Message2) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2.DiscardUnknown(m)
}

var xxx_messageInfo_Message2 proto.InternalMessageInfo

// COMMENT: Message2A
type Message2_Message2A struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Message2_Message2A struct{ m *Message2_Message2A }

func (m *Message2_Message2A) ProtoReflect() protoreflect.Message {
	return xxx_Message2_Message2A{m}
}
func (m xxx_Message2_Message2A) Type() protoreflect.MessageType {
	return xxx_Comments_ProtoFile_MessageTypes[4].Type
}
func (m xxx_Message2_Message2A) KnownFields() protoreflect.KnownFields {
	return xxx_Comments_ProtoFile_MessageTypes[4].KnownFieldsOf(m.m)
}
func (m xxx_Message2_Message2A) UnknownFields() protoreflect.UnknownFields {
	return xxx_Comments_ProtoFile_MessageTypes[4].UnknownFieldsOf(m.m)
}
func (m xxx_Message2_Message2A) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_Message2_Message2A) ProtoMutable() {}

func (m *Message2_Message2A) Reset()         { *m = Message2_Message2A{} }
func (m *Message2_Message2A) String() string { return proto.CompactTextString(m) }
func (*Message2_Message2A) ProtoMessage()    {}
func (*Message2_Message2A) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{1, 0}
}

func (m *Message2_Message2A) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2_Message2A.Unmarshal(m, b)
}
func (m *Message2_Message2A) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2_Message2A.Marshal(b, m, deterministic)
}
func (m *Message2_Message2A) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2_Message2A.Merge(m, src)
}
func (m *Message2_Message2A) XXX_Size() int {
	return xxx_messageInfo_Message2_Message2A.Size(m)
}
func (m *Message2_Message2A) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2_Message2A.DiscardUnknown(m)
}

var xxx_messageInfo_Message2_Message2A proto.InternalMessageInfo

// COMMENT: Message2B
type Message2_Message2B struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_Message2_Message2B struct{ m *Message2_Message2B }

func (m *Message2_Message2B) ProtoReflect() protoreflect.Message {
	return xxx_Message2_Message2B{m}
}
func (m xxx_Message2_Message2B) Type() protoreflect.MessageType {
	return xxx_Comments_ProtoFile_MessageTypes[5].Type
}
func (m xxx_Message2_Message2B) KnownFields() protoreflect.KnownFields {
	return xxx_Comments_ProtoFile_MessageTypes[5].KnownFieldsOf(m.m)
}
func (m xxx_Message2_Message2B) UnknownFields() protoreflect.UnknownFields {
	return xxx_Comments_ProtoFile_MessageTypes[5].UnknownFieldsOf(m.m)
}
func (m xxx_Message2_Message2B) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_Message2_Message2B) ProtoMutable() {}

func (m *Message2_Message2B) Reset()         { *m = Message2_Message2B{} }
func (m *Message2_Message2B) String() string { return proto.CompactTextString(m) }
func (*Message2_Message2B) ProtoMessage()    {}
func (*Message2_Message2B) Descriptor() ([]byte, []int) {
	return fileDescriptor_885e8293f1fab554, []int{1, 1}
}

func (m *Message2_Message2B) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message2_Message2B.Unmarshal(m, b)
}
func (m *Message2_Message2B) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message2_Message2B.Marshal(b, m, deterministic)
}
func (m *Message2_Message2B) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message2_Message2B.Merge(m, src)
}
func (m *Message2_Message2B) XXX_Size() int {
	return xxx_messageInfo_Message2_Message2B.Size(m)
}
func (m *Message2_Message2B) XXX_DiscardUnknown() {
	xxx_messageInfo_Message2_Message2B.DiscardUnknown(m)
}

var xxx_messageInfo_Message2_Message2B proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Message1)(nil), "goproto.protoc.comments.Message1")
	proto.RegisterType((*Message1_Message1A)(nil), "goproto.protoc.comments.Message1.Message1A")
	proto.RegisterType((*Message1_Message1B)(nil), "goproto.protoc.comments.Message1.Message1B")
	proto.RegisterType((*Message2)(nil), "goproto.protoc.comments.Message2")
	proto.RegisterType((*Message2_Message2A)(nil), "goproto.protoc.comments.Message2.Message2A")
	proto.RegisterType((*Message2_Message2B)(nil), "goproto.protoc.comments.Message2.Message2B")
}

func init() { proto.RegisterFile("comments/comments.proto", fileDescriptor_885e8293f1fab554) }

var fileDescriptor_885e8293f1fab554 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xce, 0xcf, 0xcd,
	0x4d, 0xcd, 0x2b, 0x29, 0xd6, 0x87, 0x31, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xc4, 0xd3,
	0xf3, 0xc1, 0x0c, 0x08, 0x37, 0x59, 0x0f, 0x26, 0xad, 0x54, 0xc8, 0xc5, 0xe1, 0x9b, 0x5a, 0x5c,
	0x9c, 0x98, 0x9e, 0x6a, 0x28, 0x24, 0xc1, 0xc5, 0xee, 0x96, 0x99, 0x9a, 0x93, 0x62, 0xe8, 0x28,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0xa9, 0x71, 0xf1, 0xfa, 0xe7, 0xa5, 0xe6,
	0xa7, 0x19, 0x3a, 0x42, 0x44, 0x24, 0x98, 0x40, 0xf2, 0x1e, 0x0c, 0x41, 0xa8, 0xc2, 0x52, 0xdc,
	0x5c, 0x9c, 0x30, 0xd3, 0x1c, 0x91, 0x39, 0x4e, 0x4e, 0x9c, 0x5c, 0xec, 0x10, 0xa5, 0x89, 0x4a,
	0x2a, 0x70, 0x2b, 0x8d, 0x90, 0xd4, 0x18, 0x21, 0x6b, 0x30, 0x72, 0x72, 0x72, 0x8e, 0x72, 0x4c,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0x02, 0xb9, 0x55, 0x3f, 0x3d, 0x3f, 0x27, 0x31, 0x2f, 0x5d, 0x1f,
	0xec, 0xfa, 0xa4, 0xd2, 0x34, 0xfd, 0x32, 0x23, 0xfd, 0xe4, 0xdc, 0x14, 0x08, 0x3f, 0x59, 0x37,
	0x3d, 0x35, 0x4f, 0x37, 0x3d, 0x5f, 0xbf, 0x24, 0xb5, 0xb8, 0x24, 0x25, 0xb1, 0x24, 0x11, 0xee,
	0x79, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0xeb, 0x57, 0xe0, 0x10, 0x01, 0x00, 0x00,
}

func init() {
	xxx_Comments_ProtoFile_FileDesc.Messages = xxx_Comments_ProtoFile_MessageDescs[0:2]
	xxx_Comments_ProtoFile_MessageDescs[0].Messages = xxx_Comments_ProtoFile_MessageDescs[2:4]
	xxx_Comments_ProtoFile_MessageDescs[1].Messages = xxx_Comments_ProtoFile_MessageDescs[4:6]
	var err error
	Comments_ProtoFile, err = prototype.NewFile(&xxx_Comments_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var Comments_ProtoFile protoreflect.FileDescriptor

var xxx_Comments_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto2,
	Path:    "comments/comments.proto",
	Package: "goproto.protoc.comments",
}
var xxx_Comments_ProtoFile_MessageTypes = [6]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_Comments_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(Message1)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Comments_ProtoFile_MessageDescs[1].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(Message2)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Comments_ProtoFile_MessageDescs[2].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(Message1_Message1A)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Comments_ProtoFile_MessageDescs[3].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(Message1_Message1B)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Comments_ProtoFile_MessageDescs[4].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(Message2_Message2A)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_Comments_ProtoFile_MessageDescs[5].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(Message2_Message2B)
		},
	)},
}
var xxx_Comments_ProtoFile_MessageDescs = [6]prototype.Message{
	{
		Name: "Message1",
		Fields: []prototype.Field{
			{
				Name:        "Field1A",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "Field1A",
			},
			{
				Name:        "Oneof1AField1",
				Number:      2,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.StringKind,
				JSONName:    "Oneof1AField1",
				OneofName:   "Oneof1a",
			},
		},
		Oneofs: []prototype.Oneof{
			{Name: "Oneof1a"},
		},
	},
	{
		Name: "Message2",
	},
	{
		Name: "Message1A",
	},
	{
		Name: "Message1B",
	},
	{
		Name: "Message2A",
	},
	{
		Name: "Message2B",
	},
}
