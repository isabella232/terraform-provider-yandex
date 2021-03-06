// Code generated by protoc-gen-go. DO NOT EDIT.
// source: yandex/cloud/dataproc/v1/resource_preset.proto

package dataproc

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

// A ResourcePreset resource for describing hardware configuration presets.
type ResourcePreset struct {
	// ID of the ResourcePreset resource.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// IDs of availability zones where the resource preset is available.
	ZoneIds []string `protobuf:"bytes,2,rep,name=zone_ids,json=zoneIds,proto3" json:"zone_ids,omitempty"`
	// Number of CPU cores for a Data Proc host created with the preset.
	Cores int64 `protobuf:"varint,3,opt,name=cores,proto3" json:"cores,omitempty"`
	// RAM volume for a Data Proc host created with the preset, in bytes.
	Memory               int64    `protobuf:"varint,4,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResourcePreset) Reset()         { *m = ResourcePreset{} }
func (m *ResourcePreset) String() string { return proto.CompactTextString(m) }
func (*ResourcePreset) ProtoMessage()    {}
func (*ResourcePreset) Descriptor() ([]byte, []int) {
	return fileDescriptor_c06789b855564c6d, []int{0}
}

func (m *ResourcePreset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResourcePreset.Unmarshal(m, b)
}
func (m *ResourcePreset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResourcePreset.Marshal(b, m, deterministic)
}
func (m *ResourcePreset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResourcePreset.Merge(m, src)
}
func (m *ResourcePreset) XXX_Size() int {
	return xxx_messageInfo_ResourcePreset.Size(m)
}
func (m *ResourcePreset) XXX_DiscardUnknown() {
	xxx_messageInfo_ResourcePreset.DiscardUnknown(m)
}

var xxx_messageInfo_ResourcePreset proto.InternalMessageInfo

func (m *ResourcePreset) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ResourcePreset) GetZoneIds() []string {
	if m != nil {
		return m.ZoneIds
	}
	return nil
}

func (m *ResourcePreset) GetCores() int64 {
	if m != nil {
		return m.Cores
	}
	return 0
}

func (m *ResourcePreset) GetMemory() int64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func init() {
	proto.RegisterType((*ResourcePreset)(nil), "yandex.cloud.dataproc.v1.ResourcePreset")
}

func init() {
	proto.RegisterFile("yandex/cloud/dataproc/v1/resource_preset.proto", fileDescriptor_c06789b855564c6d)
}

var fileDescriptor_c06789b855564c6d = []byte{
	// 216 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0xb1, 0x4b, 0xc4, 0x30,
	0x14, 0xc6, 0x69, 0xab, 0xa7, 0x97, 0xe1, 0x86, 0x20, 0x12, 0xc1, 0xa1, 0x38, 0x75, 0xb9, 0x84,
	0xc3, 0xd1, 0x4d, 0x70, 0x70, 0x93, 0x8c, 0x2e, 0x25, 0x4d, 0x1e, 0x35, 0x60, 0xfb, 0x42, 0x92,
	0x16, 0xeb, 0x5f, 0x2f, 0x26, 0x2d, 0xe8, 0x70, 0xdb, 0xfb, 0x1e, 0xbf, 0x0f, 0xbe, 0x1f, 0xe1,
	0x8b, 0x1a, 0x0d, 0x7c, 0x09, 0xfd, 0x89, 0x93, 0x11, 0x46, 0x45, 0xe5, 0x3c, 0x6a, 0x31, 0x9f,
	0x84, 0x87, 0x80, 0x93, 0xd7, 0xd0, 0x3a, 0x0f, 0x01, 0x22, 0x77, 0x1e, 0x23, 0x52, 0x96, 0x79,
	0x9e, 0x78, 0xbe, 0xf1, 0x7c, 0x3e, 0x3d, 0x58, 0x72, 0x90, 0x6b, 0xe5, 0x2d, 0x35, 0xe8, 0x81,
	0x94, 0xd6, 0xb0, 0xa2, 0x2e, 0x9a, 0xbd, 0x2c, 0xad, 0xa1, 0x77, 0xe4, 0xfa, 0x1b, 0x47, 0x68,
	0xad, 0x09, 0xac, 0xac, 0xab, 0x66, 0x2f, 0xaf, 0x7e, 0xf3, 0xab, 0x09, 0xf4, 0x86, 0x5c, 0x6a,
	0xf4, 0x10, 0x58, 0x55, 0x17, 0x4d, 0x25, 0x73, 0xa0, 0xb7, 0x64, 0x37, 0xc0, 0x80, 0x7e, 0x61,
	0x17, 0xe9, 0xbd, 0xa6, 0x67, 0x20, 0xf7, 0xff, 0x66, 0x28, 0x67, 0xff, 0x4e, 0x79, 0x7f, 0xe9,
	0x6d, 0xfc, 0x98, 0x3a, 0xae, 0x71, 0x10, 0x19, 0x3c, 0x66, 0xbf, 0x1e, 0x8f, 0x3d, 0x8c, 0xc9,
	0x44, 0x9c, 0x13, 0x7f, 0xda, 0xee, 0x6e, 0x97, 0xc0, 0xc7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x70, 0x3e, 0x3e, 0x36, 0x24, 0x01, 0x00, 0x00,
}
