// Code generated by protoc-gen-go. DO NOT EDIT.
// source: imports/test_a_1/m2.proto

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

type M2 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

type xxx_M2 struct{ m *M2 }

func (m *M2) ProtoReflect() protoreflect.Message {
	return xxx_M2{m}
}
func (m xxx_M2) Type() protoreflect.MessageType {
	return xxx_M2_ProtoFile_MessageTypes[0].Type
}
func (m xxx_M2) KnownFields() protoreflect.KnownFields {
	return xxx_M2_ProtoFile_MessageTypes[0].KnownFieldsOf(m.m)
}
func (m xxx_M2) UnknownFields() protoreflect.UnknownFields {
	return xxx_M2_ProtoFile_MessageTypes[0].UnknownFieldsOf(m.m)
}
func (m xxx_M2) Interface() protoreflect.ProtoMessage {
	return m.m
}
func (m xxx_M2) ProtoMutable() {}

func (m *M2) Reset()         { *m = M2{} }
func (m *M2) String() string { return proto.CompactTextString(m) }
func (*M2) ProtoMessage()    {}
func (*M2) Descriptor() ([]byte, []int) {
	return fileDescriptor_20cf27515c0d621c, []int{0}
}

func (m *M2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2.Unmarshal(m, b)
}
func (m *M2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2.Marshal(b, m, deterministic)
}
func (m *M2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2.Merge(m, src)
}
func (m *M2) XXX_Size() int {
	return xxx_messageInfo_M2.Size(m)
}
func (m *M2) XXX_DiscardUnknown() {
	xxx_messageInfo_M2.DiscardUnknown(m)
}

var xxx_messageInfo_M2 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M2)(nil), "test.a.M2")
}

func init() { proto.RegisterFile("imports/test_a_1/m2.proto", fileDescriptor_20cf27515c0d621c) }

var fileDescriptor_20cf27515c0d621c = []byte{
	// 114 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8c, 0x37, 0xd4, 0xcf, 0x35, 0xd2,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x09, 0xe9, 0x25, 0x2a, 0xb1, 0x70, 0x31, 0xf9,
	0x1a, 0x39, 0xb9, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xa7, 0xe7, 0xe7, 0x24, 0xe6, 0xa5, 0xeb, 0x83, 0x15, 0x26, 0x95, 0xa6, 0x41, 0x18, 0xc9, 0xba,
	0xe9, 0xa9, 0x79, 0xba, 0xe9, 0xf9, 0x60, 0xb3, 0x52, 0x12, 0x4b, 0x12, 0xf5, 0xd1, 0x0d, 0x4f,
	0x62, 0x03, 0x2b, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xe0, 0x7e, 0xc0, 0x77, 0x00,
	0x00, 0x00,
}

func init() {
	xxx_M2_ProtoFile_FileDesc.Messages = xxx_M2_ProtoFile_MessageDescs[0:1]
	var err error
	M2_ProtoFile, err = prototype.NewFile(&xxx_M2_ProtoFile_FileDesc)
	if err != nil {
		panic(err)
	}
}

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var M2_ProtoFile protoreflect.FileDescriptor

var xxx_M2_ProtoFile_FileDesc = prototype.File{
	Syntax:  protoreflect.Proto3,
	Path:    "imports/test_a_1/m2.proto",
	Package: "test.a",
}
var xxx_M2_ProtoFile_MessageTypes = [1]protoimpl.MessageType{
	{Type: prototype.GoMessage(
		xxx_M2_ProtoFile_MessageDescs[0].Reference(),
		func(protoreflect.MessageType) protoreflect.ProtoMessage {
			return new(M2)
		},
	)},
}
var xxx_M2_ProtoFile_MessageDescs = [1]prototype.Message{
	{
		Name: "M2",
	},
}
