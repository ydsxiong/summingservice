// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/ydsxiong/summingservice/gRPC/proto-files/domain/numbers.proto

package domain

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type SumRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Numbers              []int64  `protobuf:"varint,2,rep,packed,name=numbers,proto3" json:"numbers,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumRequest) Reset()         { *m = SumRequest{} }
func (m *SumRequest) String() string { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()    {}
func (*SumRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc76abb10aa063a8, []int{0}
}

func (m *SumRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumRequest.Unmarshal(m, b)
}
func (m *SumRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumRequest.Marshal(b, m, deterministic)
}
func (m *SumRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumRequest.Merge(m, src)
}
func (m *SumRequest) XXX_Size() int {
	return xxx_messageInfo_SumRequest.Size(m)
}
func (m *SumRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SumRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SumRequest proto.InternalMessageInfo

func (m *SumRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SumRequest) GetNumbers() []int64 {
	if m != nil {
		return m.Numbers
	}
	return nil
}

type SumResponse struct {
	Input                *SumRequest `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Sum                  int64       `protobuf:"varint,2,opt,name=sum,proto3" json:"sum,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *SumResponse) Reset()         { *m = SumResponse{} }
func (m *SumResponse) String() string { return proto.CompactTextString(m) }
func (*SumResponse) ProtoMessage()    {}
func (*SumResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc76abb10aa063a8, []int{1}
}

func (m *SumResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumResponse.Unmarshal(m, b)
}
func (m *SumResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumResponse.Marshal(b, m, deterministic)
}
func (m *SumResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumResponse.Merge(m, src)
}
func (m *SumResponse) XXX_Size() int {
	return xxx_messageInfo_SumResponse.Size(m)
}
func (m *SumResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SumResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SumResponse proto.InternalMessageInfo

func (m *SumResponse) GetInput() *SumRequest {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *SumResponse) GetSum() int64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

type SumFilter struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SumFilter) Reset()         { *m = SumFilter{} }
func (m *SumFilter) String() string { return proto.CompactTextString(m) }
func (*SumFilter) ProtoMessage()    {}
func (*SumFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_bc76abb10aa063a8, []int{2}
}

func (m *SumFilter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SumFilter.Unmarshal(m, b)
}
func (m *SumFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SumFilter.Marshal(b, m, deterministic)
}
func (m *SumFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SumFilter.Merge(m, src)
}
func (m *SumFilter) XXX_Size() int {
	return xxx_messageInfo_SumFilter.Size(m)
}
func (m *SumFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_SumFilter.DiscardUnknown(m)
}

var xxx_messageInfo_SumFilter proto.InternalMessageInfo

func (m *SumFilter) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

func init() {
	proto.RegisterType((*SumRequest)(nil), "domain.SumRequest")
	proto.RegisterType((*SumResponse)(nil), "domain.SumResponse")
	proto.RegisterType((*SumFilter)(nil), "domain.SumFilter")
}

func init() {
	proto.RegisterFile("github.com/ydsxiong/summingservice/internal/proto-files/domain/numbers.proto", fileDescriptor_bc76abb10aa063a8)
}

var fileDescriptor_bc76abb10aa063a8 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x31, 0x4f, 0xc3, 0x30,
	0x10, 0x85, 0xd5, 0x06, 0x8a, 0xea, 0x2e, 0xc8, 0x53, 0xc6, 0x28, 0xa8, 0x52, 0x16, 0x62, 0x09,
	0x36, 0x46, 0x90, 0x90, 0x90, 0x18, 0x90, 0xbb, 0xb1, 0x25, 0xcd, 0x11, 0x4e, 0xe4, 0xce, 0xc1,
	0x67, 0x03, 0xfd, 0xf7, 0xa8, 0x4e, 0x2a, 0x26, 0xb6, 0xf7, 0xec, 0x7b, 0xf7, 0xee, 0x53, 0xcf,
	0x3d, 0x86, 0xf7, 0xd8, 0xd6, 0x7b, 0x47, 0xe6, 0xd0, 0xc9, 0x0f, 0x3a, 0xee, 0x8d, 0x44, 0x22,
	0xe4, 0x5e, 0xc0, 0x7f, 0xe1, 0x1e, 0x0c, 0x72, 0x00, 0xcf, 0xcd, 0x60, 0x46, 0xef, 0x82, 0xbb,
	0x7e, 0xc3, 0x01, 0xc4, 0x74, 0x8e, 0x1a, 0x64, 0xc3, 0x91, 0x5a, 0xf0, 0x52, 0xa7, 0x2f, 0xbd,
	0x9a, 0x5e, 0xcb, 0x3b, 0xa5, 0x76, 0x91, 0x2c, 0x7c, 0x46, 0x90, 0xa0, 0xb5, 0x3a, 0xe3, 0x86,
	0x20, 0x5f, 0x14, 0x8b, 0x6a, 0x6d, 0x93, 0xd6, 0xb9, 0xba, 0x98, 0xa3, 0xf9, 0xb2, 0xc8, 0xaa,
	0xcc, 0x9e, 0x6c, 0xf9, 0xa4, 0x36, 0x29, 0x2b, 0xa3, 0x63, 0x01, 0x5d, 0xa9, 0x73, 0xe4, 0x31,
	0x86, 0x94, 0xde, 0xdc, 0xe8, 0x7a, 0xaa, 0xa8, 0xff, 0xf6, 0xdb, 0x69, 0x40, 0x5f, 0xaa, 0x4c,
	0x22, 0xe5, 0xcb, 0x62, 0x51, 0x65, 0xf6, 0x28, 0xcb, 0xad, 0x5a, 0xef, 0x22, 0x3d, 0xe2, 0x10,
	0xc0, 0x1f, 0x1b, 0x3f, 0xe0, 0xf0, 0xed, 0x7c, 0x37, 0x1f, 0x72, 0xb2, 0xf7, 0xdb, 0xd7, 0xab,
	0xff, 0x88, 0x7b, 0xfb, 0xf2, 0x30, 0xa3, 0xb6, 0xab, 0xc4, 0x78, 0xfb, 0x1b, 0x00, 0x00, 0xff,
	0xff, 0x40, 0x8a, 0x2a, 0x3a, 0x33, 0x01, 0x00, 0x00,
}
