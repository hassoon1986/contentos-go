// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app/table/so_blockSummaryObject.proto

package table

import (
	fmt "fmt"
	prototype "github.com/coschain/contentos-go/prototype"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type SoBlockSummaryObject struct {
	Id                   uint32            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BlockId              *prototype.Sha256 `protobuf:"bytes,2,opt,name=block_id,json=blockId,proto3" json:"block_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SoBlockSummaryObject) Reset()         { *m = SoBlockSummaryObject{} }
func (m *SoBlockSummaryObject) String() string { return proto.CompactTextString(m) }
func (*SoBlockSummaryObject) ProtoMessage()    {}
func (*SoBlockSummaryObject) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0052d3a16a16d4f, []int{0}
}

func (m *SoBlockSummaryObject) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoBlockSummaryObject.Unmarshal(m, b)
}
func (m *SoBlockSummaryObject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoBlockSummaryObject.Marshal(b, m, deterministic)
}
func (m *SoBlockSummaryObject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoBlockSummaryObject.Merge(m, src)
}
func (m *SoBlockSummaryObject) XXX_Size() int {
	return xxx_messageInfo_SoBlockSummaryObject.Size(m)
}
func (m *SoBlockSummaryObject) XXX_DiscardUnknown() {
	xxx_messageInfo_SoBlockSummaryObject.DiscardUnknown(m)
}

var xxx_messageInfo_SoBlockSummaryObject proto.InternalMessageInfo

func (m *SoBlockSummaryObject) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SoBlockSummaryObject) GetBlockId() *prototype.Sha256 {
	if m != nil {
		return m.BlockId
	}
	return nil
}

type SoUniqueBlockSummaryObjectById struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SoUniqueBlockSummaryObjectById) Reset()         { *m = SoUniqueBlockSummaryObjectById{} }
func (m *SoUniqueBlockSummaryObjectById) String() string { return proto.CompactTextString(m) }
func (*SoUniqueBlockSummaryObjectById) ProtoMessage()    {}
func (*SoUniqueBlockSummaryObjectById) Descriptor() ([]byte, []int) {
	return fileDescriptor_a0052d3a16a16d4f, []int{1}
}

func (m *SoUniqueBlockSummaryObjectById) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SoUniqueBlockSummaryObjectById.Unmarshal(m, b)
}
func (m *SoUniqueBlockSummaryObjectById) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SoUniqueBlockSummaryObjectById.Marshal(b, m, deterministic)
}
func (m *SoUniqueBlockSummaryObjectById) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SoUniqueBlockSummaryObjectById.Merge(m, src)
}
func (m *SoUniqueBlockSummaryObjectById) XXX_Size() int {
	return xxx_messageInfo_SoUniqueBlockSummaryObjectById.Size(m)
}
func (m *SoUniqueBlockSummaryObjectById) XXX_DiscardUnknown() {
	xxx_messageInfo_SoUniqueBlockSummaryObjectById.DiscardUnknown(m)
}

var xxx_messageInfo_SoUniqueBlockSummaryObjectById proto.InternalMessageInfo

func (m *SoUniqueBlockSummaryObjectById) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*SoBlockSummaryObject)(nil), "table.so_blockSummaryObject")
	proto.RegisterType((*SoUniqueBlockSummaryObjectById)(nil), "table.so_unique_blockSummaryObject_by_id")
}

func init() {
	proto.RegisterFile("app/table/so_blockSummaryObject.proto", fileDescriptor_a0052d3a16a16d4f)
}

var fileDescriptor_a0052d3a16a16d4f = []byte{
	// 202 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0x2c, 0x28, 0xd0,
	0x2f, 0x49, 0x4c, 0xca, 0x49, 0xd5, 0x2f, 0xce, 0x8f, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x0e, 0x2e,
	0xcd, 0xcd, 0x4d, 0x2c, 0xaa, 0xf4, 0x4f, 0xca, 0x4a, 0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x62, 0x05, 0x2b, 0x91, 0x12, 0x01, 0xf3, 0x4a, 0x2a, 0x0b, 0x52, 0xf5, 0x41, 0x04,
	0x44, 0x52, 0x29, 0x94, 0x4b, 0x14, 0xab, 0x5e, 0x21, 0x3e, 0x2e, 0xa6, 0xcc, 0x14, 0x09, 0x46,
	0x05, 0x46, 0x0d, 0xde, 0x20, 0xa6, 0xcc, 0x14, 0x21, 0x1d, 0x2e, 0x0e, 0xb0, 0xaa, 0xf8, 0xcc,
	0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x6e, 0x23, 0x41, 0x3d, 0xb8, 0x89, 0x7a, 0xc5, 0x19, 0x89,
	0x46, 0xa6, 0x66, 0x41, 0xec, 0x60, 0x25, 0x9e, 0x29, 0x4a, 0x26, 0x5c, 0x4a, 0xc5, 0xf9, 0xf1,
	0xa5, 0x79, 0x99, 0x85, 0xa5, 0xa9, 0x58, 0x4c, 0x8f, 0x4f, 0xaa, 0x8c, 0xcf, 0x4c, 0x41, 0xb7,
	0xc3, 0x49, 0x23, 0x4a, 0x2d, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f,
	0x39, 0xbf, 0x38, 0x39, 0x23, 0x31, 0x33, 0x4f, 0x3f, 0x39, 0x3f, 0xaf, 0x24, 0x35, 0xaf, 0x24,
	0xbf, 0x58, 0x37, 0x3d, 0x1f, 0xe2, 0xdf, 0x24, 0x36, 0xb0, 0xd5, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x58, 0x14, 0xdd, 0x03, 0x01, 0x00, 0x00,
}
