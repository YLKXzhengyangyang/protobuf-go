// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_1/m1.proto

package test_a_1

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

type E1 int32

const (
	E1_E1_ZERO E1 = 0
)

type xxx_E1 E1

func (e E1) ProtoReflect() protoreflect.Enum {
	return (xxx_E1)(e)
}
func (e xxx_E1) Type() protoreflect.EnumType {
	return xxx_M1_ProtoFile_EnumTypes[0]
}
func (e xxx_E1) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

var E1_name = map[int32]string{
	0: "E1_ZERO",
}

var E1_value = map[string]int32{
	"E1_ZERO": 0,
}

func (x E1) String() string {
	return proto.EnumName(E1_name, int32(x))
}

func (E1) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c1091de3fa870a14, []int{0}
}

type M1 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_M1 struct{ m *M1 }

func (m *M1) ProtoReflect() protoreflect.Message {
	return xxx_M1{m}
}
func (m xxx_M1) Type() protoreflect.MessageType {
	return xxx_M1_ProtoFile_MessageTypes[0].Type
}
func (m xxx_M1) KnownFields() protoreflect.KnownFields {
	return xxx_M1_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_M1) UnknownFields() protoreflect.UnknownFields {
	return xxx_M1_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_M1) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_M1) ProtoMutable() {}

func (m *M1) Reset()         { *m = M1{} }
func (m *M1) String() string { return proto.CompactTextString(m) }
func (*M1) ProtoMessage()    {}
func (*M1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1091de3fa870a14, []int{0}
}

func (m *M1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M1.Unmarshal(m, b)
}
func (m *M1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M1.Marshal(b, m, deterministic)
}
func (m *M1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M1.Merge(m, src)
}
func (m *M1) XXX_Size() int {
	return xxx_messageInfo_M1.Size(m)
}
func (m *M1) XXX_DiscardUnknown() {
	xxx_messageInfo_M1.DiscardUnknown(m)
}

var xxx_messageInfo_M1 proto.InternalMessageInfo

type M1_1 struct {
	M1                   *M1      `protobuf:"bytes,1,opt,name=m1,proto3" json:"m1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_M1_1 struct{ m *M1_1 }

func (m *M1_1) ProtoReflect() protoreflect.Message {
	return xxx_M1_1{m}
}
func (m xxx_M1_1) Type() protoreflect.MessageType {
	return xxx_M1_ProtoFile_MessageTypes[1].Type
}
func (m xxx_M1_1) KnownFields() protoreflect.KnownFields {
	return xxx_M1_ProtoFile_MessageTypes[1].KnownFieldsOf(m.m)
}
func (m xxx_M1_1) UnknownFields() protoreflect.UnknownFields {
	return xxx_M1_ProtoFile_MessageTypes[1].UnknownFieldsOf(m.m)
}
func (m xxx_M1_1) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_M1_1) ProtoMutable() {}

func (m *M1_1) Reset()         { *m = M1_1{} }
func (m *M1_1) String() string { return proto.CompactTextString(m) }
func (*M1_1) ProtoMessage()    {}
func (*M1_1) Descriptor() ([]byte, []int) {
	return fileDescriptor_c1091de3fa870a14, []int{1}
}

func (m *M1_1) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M1_1.Unmarshal(m, b)
}
func (m *M1_1) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M1_1.Marshal(b, m, deterministic)
}
func (m *M1_1) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M1_1.Merge(m, src)
}
func (m *M1_1) XXX_Size() int {
	return xxx_messageInfo_M1_1.Size(m)
}
func (m *M1_1) XXX_DiscardUnknown() {
	xxx_messageInfo_M1_1.DiscardUnknown(m)
}

var xxx_messageInfo_M1_1 proto.InternalMessageInfo

func (m *M1_1) GetM1() *M1 {
	if m != nil {
		return m.M1
	}
	return nil
}

func init() {
	proto.RegisterEnum("test.a.E1", E1_name, E1_value)
	proto.RegisterType((*M1)(nil), "test.a.M1")
	proto.RegisterType((*M1_1)(nil), "test.a.M1_1")
}

func init() { proto.RegisterFile("imports/test_a_1/m1.proto", fileDescriptor_c1091de3fa870a14) }

var fileDescriptor_c1091de3fa870a14 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8c, 0x37, 0xd4, 0xcf, 0x35, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x09, 0xe9, 0x25, 0x2a, 0xb1, 0x70, 0x31, 0xf9,
	0x1a, 0x2a, 0x29, 0x71, 0xb1, 0xf8, 0x1a, 0xc6, 0x1b, 0x0a, 0x49, 0x71, 0x31, 0xe5, 0x1a, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x71, 0xe9, 0x41, 0x94, 0xe8, 0xf9, 0x1a, 0x06, 0x31, 0xe5,
	0x1a, 0x6a, 0x09, 0x72, 0x31, 0xb9, 0x1a, 0x0a, 0x71, 0x73, 0xb1, 0xbb, 0x1a, 0xc6, 0x47, 0xb9,
	0x06, 0xf9, 0x0b, 0x30, 0x38, 0xb9, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25,
	0xe7, 0xe7, 0xea, 0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x83, 0xcd, 0x4f, 0x2a, 0x4d, 0x83,
	0x30, 0x92, 0x75, 0xd3, 0x53, 0xf3, 0x74, 0xd3, 0xf3, 0xc1, 0x4e, 0x48, 0x49, 0x2c, 0x49, 0xd4,
	0x47, 0x77, 0x53, 0x12, 0x1b, 0x58, 0xa1, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xcc, 0xae, 0xc9,
	0xcd, 0xae, 0x00, 0x00, 0x00,
}

func init() {
	xxx_M1_ProtoFile_FileDesc.Enums = xxx_M1_ProtoFile_EnumDescs[0:1]
	xxx_M1_ProtoFile_FileDesc.Messages = xxx_M1_ProtoFile_MessageDescs[0:2]
	xxx_M1_ProtoFile_MessageDescs[1].Fields[0].MessageType = xxx_M1_ProtoFile_MessageTypes[0].Type
	var err error
	M1_ProtoFile, err = prototype.NewFile(&xxx_M1_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var M1_ProtoFile protoreflect.FileDescriptor

var xxx_M1_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto3,
	Path:    "imports/test_a_1/m1.proto",
	Package: "test.a",
}
var xxx_M1_ProtoFile_EnumTypes = [1]protoreflect.EnumType{
	prototype.GoEnum(
		xxx_M1_ProtoFile_EnumDescs[0].Reference(),
		func(_ protoreflect.EnumType, n protoreflect.EnumNumber) protoreflect.ProtoEnum {
			return E1(n)
		},
	),
}
var xxx_M1_ProtoFile_EnumDescs = [1]prototype.Enum{
	{
		Name: "E1",
		Values: []prototype.EnumValue{
			{Name: "E1_ZERO", Number: 0},
		},
	},
}
var xxx_M1_ProtoFile_MessageTypes = [2]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_M1_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(M1)
		},
	)},
	{Type: prototype.GoMessage(
		xxx_M1_ProtoFile_MessageDescs[1].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(M1_1)
		},
	)},
}
var xxx_M1_ProtoFile_MessageDescs = [2]prototype.Message{
	{
		Name: "M1",
	},
	{
		Name: "M1_1",
		Fields: []prototype.Field{
			{
				Name:        "m1",
				Number:      1,
				Cardinality: protoreflect.Optional,
				Kind:        protoreflect.MessageKind,
				JSONName:    "m1",
			},
		},
	},
}
