// Code generated by protoc-gen-go. DO NOT EDIT.
// source: feast/specs/StorageSpec.proto

package specs // import "github.com/gojektech/feast/go-feast-proto/feast/specs"

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

type StorageSpec struct {
	// unique identifier for this instance
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type should define what sort of store it is
	// e.g. redis, bigquery, etc
	Type string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	// options contain (but are not restricted to) options like
	// connection information.
	Options              map[string]string `protobuf:"bytes,3,rep,name=options,proto3" json:"options,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *StorageSpec) Reset()         { *m = StorageSpec{} }
func (m *StorageSpec) String() string { return proto.CompactTextString(m) }
func (*StorageSpec) ProtoMessage()    {}
func (*StorageSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_StorageSpec_e36eea24415a551a, []int{0}
}
func (m *StorageSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StorageSpec.Unmarshal(m, b)
}
func (m *StorageSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StorageSpec.Marshal(b, m, deterministic)
}
func (dst *StorageSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StorageSpec.Merge(dst, src)
}
func (m *StorageSpec) XXX_Size() int {
	return xxx_messageInfo_StorageSpec.Size(m)
}
func (m *StorageSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_StorageSpec.DiscardUnknown(m)
}

var xxx_messageInfo_StorageSpec proto.InternalMessageInfo

func (m *StorageSpec) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *StorageSpec) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *StorageSpec) GetOptions() map[string]string {
	if m != nil {
		return m.Options
	}
	return nil
}

func init() {
	proto.RegisterType((*StorageSpec)(nil), "feast.specs.StorageSpec")
	proto.RegisterMapType((map[string]string)(nil), "feast.specs.StorageSpec.OptionsEntry")
}

func init() {
	proto.RegisterFile("feast/specs/StorageSpec.proto", fileDescriptor_StorageSpec_e36eea24415a551a)
}

var fileDescriptor_StorageSpec_e36eea24415a551a = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0x4b, 0x4d, 0x2c,
	0x2e, 0xd1, 0x2f, 0x2e, 0x48, 0x4d, 0x2e, 0xd6, 0x0f, 0x2e, 0xc9, 0x2f, 0x4a, 0x4c, 0x4f, 0x0d,
	0x2e, 0x48, 0x4d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x4b, 0xeb, 0x81, 0xa5,
	0x95, 0xd6, 0x31, 0x72, 0x71, 0x23, 0x29, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe2, 0x62, 0x29, 0xa9, 0x2c, 0x48, 0x95,
	0x60, 0x02, 0x8b, 0x80, 0xd9, 0x42, 0xf6, 0x5c, 0xec, 0xf9, 0x05, 0x25, 0x99, 0xf9, 0x79, 0xc5,
	0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xaa, 0x7a, 0x48, 0x46, 0xea, 0x21, 0xdb, 0xe8, 0x0f,
	0x51, 0xe7, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x04, 0xd3, 0x25, 0x65, 0xc5, 0xc5, 0x83, 0x2c, 0x21,
	0x24, 0xc0, 0xc5, 0x9c, 0x9d, 0x5a, 0x09, 0xb5, 0x15, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b,
	0xcc, 0x29, 0x85, 0xd9, 0x0b, 0xe1, 0x58, 0x31, 0x59, 0x30, 0x3a, 0x85, 0x71, 0x21, 0xbb, 0xdf,
	0x49, 0x00, 0xc9, 0xb6, 0x00, 0x90, 0xf7, 0xa2, 0x4c, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4,
	0x92, 0xf3, 0x73, 0xf5, 0xd3, 0xf3, 0xb3, 0x52, 0xb3, 0x4b, 0x52, 0x93, 0x33, 0xf4, 0x21, 0x41,
	0x92, 0x9e, 0xaf, 0x0b, 0x66, 0xe8, 0x82, 0x43, 0x42, 0x1f, 0x29, 0x9c, 0x92, 0xd8, 0xc0, 0x42,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf0, 0x9e, 0x30, 0x12, 0x3d, 0x01, 0x00, 0x00,
}