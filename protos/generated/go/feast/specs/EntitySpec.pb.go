// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/specs/EntitySpec.proto

package specs // import "github.com/gojektech/feast/protos/generated/go/feast/specs"

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

type EntitySpec struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Tags                 []string `protobuf:"bytes,3,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntitySpec) Reset()         { *m = EntitySpec{} }
func (m *EntitySpec) String() string { return proto.CompactTextString(m) }
func (*EntitySpec) ProtoMessage()    {}
func (*EntitySpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_EntitySpec_5830cf373dc122f9, []int{0}
}
func (m *EntitySpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntitySpec.Unmarshal(m, b)
}
func (m *EntitySpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntitySpec.Marshal(b, m, deterministic)
}
func (dst *EntitySpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntitySpec.Merge(dst, src)
}
func (m *EntitySpec) XXX_Size() int {
	return xxx_messageInfo_EntitySpec.Size(m)
}
func (m *EntitySpec) XXX_DiscardUnknown() {
	xxx_messageInfo_EntitySpec.DiscardUnknown(m)
}

var xxx_messageInfo_EntitySpec proto.InternalMessageInfo

func (m *EntitySpec) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *EntitySpec) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *EntitySpec) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*EntitySpec)(nil), "feast.specs.EntitySpec")
}

func init() {
	proto.RegisterFile("feast/specs/EntitySpec.proto", fileDescriptor_EntitySpec_5830cf373dc122f9)
}

var fileDescriptor_EntitySpec_5830cf373dc122f9 = []byte{
	// 179 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x31, 0x0b, 0xc2, 0x30,
	0x10, 0x85, 0xa9, 0x15, 0xc1, 0x74, 0x10, 0x32, 0x75, 0x70, 0x28, 0x4e, 0x4e, 0xbd, 0xc1, 0xcd,
	0xb1, 0xe0, 0x2e, 0x0a, 0x0e, 0xdd, 0xd2, 0xf4, 0x4c, 0xa3, 0x34, 0x09, 0xcd, 0x39, 0xf8, 0xef,
	0x25, 0xe7, 0xd0, 0x6e, 0x8f, 0x7b, 0xef, 0xf1, 0xbe, 0x13, 0xfb, 0x27, 0xaa, 0x48, 0x10, 0x03,
	0xea, 0x08, 0x17, 0x47, 0x96, 0xbe, 0xf7, 0x80, 0xba, 0x0e, 0x93, 0x27, 0x2f, 0x0b, 0x76, 0x6b,
	0x76, 0x0f, 0x0f, 0x21, 0xe6, 0x80, 0x94, 0x62, 0xed, 0xd4, 0x88, 0x65, 0x56, 0x65, 0xc7, 0xed,
	0x8d, 0xb5, 0xac, 0x44, 0xd1, 0x63, 0xd4, 0x93, 0x0d, 0x64, 0xbd, 0x2b, 0x57, 0x6c, 0x2d, 0x4f,
	0xa9, 0x45, 0xca, 0xc4, 0x32, 0xaf, 0xf2, 0xd4, 0x4a, 0xba, 0x69, 0xc5, 0x72, 0xa6, 0xd9, 0xcd,
	0x23, 0xd7, 0x04, 0xd1, 0x9e, 0x8d, 0xa5, 0xe1, 0xd3, 0xd5, 0xda, 0x8f, 0x60, 0xfc, 0x0b, 0xdf,
	0x84, 0x7a, 0x80, 0x3f, 0x37, 0x63, 0x46, 0x30, 0xe8, 0x70, 0x52, 0x84, 0x3d, 0x18, 0x0f, 0x8b,
	0x8f, 0xba, 0x0d, 0x07, 0x4e, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xd4, 0x89, 0xe8, 0xe7,
	0x00, 0x00, 0x00,
}
