// Code generated by protoc-gen-go. DO NOT EDIT.
// source: keyref.proto

/*
Package schema is a generated protocol buffer package.

It is generated from these files:
	keyref.proto
	list.proto

It has these top-level messages:
	KeyRefId
	KeyRef
	List
*/
package schema

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// KeyRefId  id part of key reference
type KeyRefId struct {
	// entity type
	Schema string `protobuf:"bytes,1,opt,name=schema" json:"schema,omitempty"`
	// idx name from entity type
	Idx string `protobuf:"bytes,2,opt,name=idx" json:"idx,omitempty"`
	// referred key
	RefKey []string `protobuf:"bytes,3,rep,name=ref_key,json=refKey" json:"ref_key,omitempty"`
}

func (m *KeyRefId) Reset()                    { *m = KeyRefId{} }
func (m *KeyRefId) String() string            { return proto.CompactTextString(m) }
func (*KeyRefId) ProtoMessage()               {}
func (*KeyRefId) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *KeyRefId) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *KeyRefId) GetIdx() string {
	if m != nil {
		return m.Idx
	}
	return ""
}

func (m *KeyRefId) GetRefKey() []string {
	if m != nil {
		return m.RefKey
	}
	return nil
}

type KeyRef struct {
	// entity type
	Schema string `protobuf:"bytes,1,opt,name=schema" json:"schema,omitempty"`
	// idx name from entity type
	Idx string `protobuf:"bytes,2,opt,name=idx" json:"idx,omitempty"`
	// referred key
	RefKey []string `protobuf:"bytes,3,rep,name=ref_key,json=refKey" json:"ref_key,omitempty"`
	// primary key instance linked to
	PKey []string `protobuf:"bytes,4,rep,name=p_key,json=pKey" json:"p_key,omitempty"`
}

func (m *KeyRef) Reset()                    { *m = KeyRef{} }
func (m *KeyRef) String() string            { return proto.CompactTextString(m) }
func (*KeyRef) ProtoMessage()               {}
func (*KeyRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *KeyRef) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

func (m *KeyRef) GetIdx() string {
	if m != nil {
		return m.Idx
	}
	return ""
}

func (m *KeyRef) GetRefKey() []string {
	if m != nil {
		return m.RefKey
	}
	return nil
}

func (m *KeyRef) GetPKey() []string {
	if m != nil {
		return m.PKey
	}
	return nil
}

func init() {
	proto.RegisterType((*KeyRefId)(nil), "cckit.state.schema.KeyRefId")
	proto.RegisterType((*KeyRef)(nil), "cckit.state.schema.KeyRef")
}

func init() { proto.RegisterFile("keyref.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4e, 0xad, 0x2c,
	0x4a, 0x4d, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4a, 0x4e, 0xce, 0xce, 0x2c, 0xd1,
	0x2b, 0x2e, 0x49, 0x2c, 0x49, 0xd5, 0x2b, 0x4e, 0xce, 0x48, 0xcd, 0x4d, 0x54, 0xf2, 0xe5, 0xe2,
	0xf0, 0x4e, 0xad, 0x0c, 0x4a, 0x4d, 0xf3, 0x4c, 0x11, 0x12, 0xe3, 0x62, 0x83, 0x88, 0x4a, 0x30,
	0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41, 0x79, 0x42, 0x02, 0x5c, 0xcc, 0x99, 0x29, 0x15, 0x12, 0x4c,
	0x60, 0x41, 0x10, 0x53, 0x48, 0x9c, 0x8b, 0xbd, 0x28, 0x35, 0x2d, 0x3e, 0x3b, 0xb5, 0x52, 0x82,
	0x59, 0x81, 0x19, 0xa4, 0xb4, 0x28, 0x35, 0xcd, 0x3b, 0xb5, 0x52, 0x29, 0x81, 0x8b, 0x0d, 0x62,
	0x1c, 0x15, 0x0c, 0x13, 0x12, 0xe6, 0x62, 0x2d, 0x00, 0x0b, 0xb3, 0x80, 0x85, 0x59, 0x0a, 0xbc,
	0x53, 0x2b, 0x9d, 0x38, 0xa2, 0xa0, 0x26, 0x25, 0xb1, 0x81, 0x7d, 0x65, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0x12, 0xf3, 0xec, 0xb6, 0xe5, 0x00, 0x00, 0x00,
}