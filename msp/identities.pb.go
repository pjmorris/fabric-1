// Code generated by protoc-gen-go.
// source: identities.proto
// DO NOT EDIT!

/*
Package msp is a generated protocol buffer package.

It is generated from these files:
    identities.proto

It has these top-level messages:
    SerializedIdentity
*/
package msp

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

// This struct represents an Identity
// (with its MSP identifier) to be used
// to serialize it and deserialize it
type SerializedIdentity struct {
	// The identifier of the associated membership service provider
	Mspid string `protobuf:"bytes,1,opt,name=Mspid" json:"Mspid,omitempty"`
	// the Identity, serialized according to the rules of its MPS
	IdBytes []byte `protobuf:"bytes,2,opt,name=IdBytes,proto3" json:"IdBytes,omitempty"`
}

func (m *SerializedIdentity) Reset()                    { *m = SerializedIdentity{} }
func (m *SerializedIdentity) String() string            { return proto.CompactTextString(m) }
func (*SerializedIdentity) ProtoMessage()               {}
func (*SerializedIdentity) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func init() {
	proto.RegisterType((*SerializedIdentity)(nil), "msp.SerializedIdentity")
}

func init() { proto.RegisterFile("identities.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 143 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x4c, 0x49, 0xcd,
	0x2b, 0xc9, 0x2c, 0xc9, 0x4c, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0x2d,
	0x2e, 0x50, 0x72, 0xe1, 0x12, 0x0a, 0x4e, 0x2d, 0xca, 0x4c, 0xcc, 0xc9, 0xac, 0x4a, 0x4d, 0xf1,
	0x84, 0x28, 0xa9, 0x14, 0x12, 0xe1, 0x62, 0xf5, 0x2d, 0x2e, 0xc8, 0x4c, 0x91, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x24, 0xb8, 0xd8, 0x3d, 0x53, 0x9c, 0x2a, 0x4b, 0x52, 0x8b,
	0x25, 0x98, 0x80, 0xe2, 0x3c, 0x41, 0x30, 0xae, 0x93, 0x72, 0x94, 0x62, 0x7a, 0x66, 0x49, 0x46,
	0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x46, 0x65, 0x41, 0x6a, 0x51, 0x4e, 0x6a, 0x4a, 0x7a,
	0x6a, 0x91, 0x7e, 0x5a, 0x62, 0x52, 0x51, 0x66, 0xb2, 0x3e, 0xd0, 0xaa, 0x24, 0x36, 0xb0, 0xb5,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xdc, 0x7a, 0xe7, 0x45, 0x8a, 0x00, 0x00, 0x00,
}
