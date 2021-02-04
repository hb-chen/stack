// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/stack-labs/stack-rpc/plugin/service/stackway/test/proto/test.proto

package stack_rpc_stackway_test

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/stack-labs/stack-rpc/api/proto"
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

type Request struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc0c9e6e8b9a187f, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type Response struct {
	Msg                  string   `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_cc0c9e6e8b9a187f, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "stack.rpc.stackway.test.Request")
	proto.RegisterType((*Response)(nil), "stack.rpc.stackway.test.Response")
}

func init() {
	proto.RegisterFile("github.com/stack-labs/stack-rpc/plugin/service/stackway/test/proto/test.proto", fileDescriptor_cc0c9e6e8b9a187f)
}

var fileDescriptor_cc0c9e6e8b9a187f = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x8f, 0xbf, 0x4a, 0xc0, 0x30,
	0x10, 0xc6, 0x2d, 0x15, 0xff, 0x64, 0xb1, 0x64, 0x51, 0xaa, 0x43, 0xed, 0x24, 0x42, 0x2f, 0x60,
	0x9f, 0xc0, 0x5d, 0x1c, 0x82, 0x2f, 0x90, 0x86, 0x10, 0x83, 0x6d, 0x73, 0xe6, 0x52, 0xc5, 0xcd,
	0x47, 0x97, 0xa4, 0x2d, 0x0e, 0xa2, 0x6e, 0x3f, 0xf2, 0xfb, 0xee, 0xbe, 0x0b, 0x7b, 0xb4, 0x2e,
	0x3e, 0x2f, 0x03, 0x68, 0x3f, 0x09, 0x8a, 0x4a, 0xbf, 0x74, 0xa3, 0x1a, 0x68, 0xc3, 0x80, 0xba,
	0xc3, 0x71, 0xb1, 0x6e, 0x26, 0x41, 0x26, 0xbc, 0x39, 0x6d, 0x56, 0xf3, 0xae, 0x3e, 0x44, 0x34,
	0x14, 0x05, 0x06, 0x1f, 0x7d, 0x46, 0xc8, 0xc8, 0xcf, 0xb3, 0x87, 0x80, 0x1a, 0xf6, 0x24, 0x24,
	0x5d, 0xf7, 0xff, 0x14, 0x09, 0x85, 0x6e, 0x5b, 0xa7, 0xd0, 0xad, 0xdb, 0xda, 0x4b, 0x76, 0x2c,
	0xcd, 0xeb, 0x62, 0x28, 0xf2, 0x8a, 0x95, 0x13, 0xd9, 0x8b, 0xa2, 0x29, 0x6e, 0x4e, 0x65, 0xc2,
	0xf6, 0x8a, 0x9d, 0x48, 0x43, 0xe8, 0x67, 0x32, 0x3f, 0xed, 0xdd, 0x67, 0xc1, 0x0e, 0x9f, 0xd2,
	0xe0, 0x03, 0x2b, 0x25, 0x6a, 0xde, 0xc0, 0x2f, 0x97, 0xc1, 0xd6, 0x50, 0x5f, 0xff, 0x91, 0x58,
	0x6b, 0xda, 0x03, 0x7e, 0xcb, 0xca, 0x7b, 0x74, 0xfc, 0x0c, 0xac, 0x87, 0x74, 0xe7, 0x3e, 0x5c,
	0x7d, 0x3f, 0xec, 0xd9, 0xe1, 0x28, 0x7f, 0xa2, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xa8,
	0x48, 0x38, 0x64, 0x01, 0x00, 0x00,
}
