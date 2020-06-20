// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/getsession/getsession.proto

package go_micro_srv_GetSession

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	SessionId            string   `protobuf:"bytes,1,opt,name=SessionId,proto3" json:"SessionId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb548ff9f91aaed, []int{0}
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

func (m *Request) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

type Response struct {
	ErrNo                string   `protobuf:"bytes,1,opt,name=ErrNo,proto3" json:"ErrNo,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	Data                 string   `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_acb548ff9f91aaed, []int{1}
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

func (m *Response) GetErrNo() string {
	if m != nil {
		return m.ErrNo
	}
	return ""
}

func (m *Response) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetSession.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetSession.Response")
}

func init() { proto.RegisterFile("proto/getsession/getsession.proto", fileDescriptor_acb548ff9f91aaed) }

var fileDescriptor_acb548ff9f91aaed = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x41, 0xcb, 0x82, 0x40,
	0x10, 0x86, 0x3f, 0xbf, 0x4a, 0x73, 0x8e, 0x43, 0x94, 0x44, 0x07, 0xf5, 0x52, 0xa7, 0x0d, 0xea,
	0x37, 0x48, 0x04, 0xd5, 0x41, 0xef, 0x81, 0xe5, 0x20, 0x42, 0xba, 0xb6, 0xb3, 0x45, 0x3f, 0x3f,
	0x58, 0x17, 0xec, 0x52, 0xb7, 0x79, 0x9f, 0xf7, 0x3d, 0x3c, 0x03, 0x51, 0xab, 0xa4, 0x96, 0xeb,
	0x92, 0x34, 0x13, 0x73, 0x25, 0x9b, 0x8f, 0x53, 0x98, 0x0e, 0x67, 0xa5, 0x14, 0x75, 0x75, 0x55,
	0x52, 0xb0, 0x7a, 0x8a, 0x1d, 0xe9, 0xac, 0xab, 0xe3, 0x25, 0x78, 0x29, 0xdd, 0x1f, 0xc4, 0x1a,
	0x17, 0xe0, 0x5b, 0xba, 0x2f, 0x02, 0x27, 0x74, 0x56, 0x7e, 0xda, 0x83, 0xf8, 0x00, 0xe3, 0x94,
	0xb8, 0x95, 0x0d, 0x13, 0x4e, 0x60, 0x94, 0x28, 0x75, 0x92, 0x76, 0xd5, 0x05, 0x9c, 0x82, 0x9b,
	0x28, 0x75, 0xe4, 0x32, 0xf8, 0x37, 0xd8, 0x26, 0x44, 0x18, 0x16, 0xb9, 0xce, 0x83, 0x81, 0xa1,
	0xe6, 0xde, 0x9c, 0xc1, 0x4b, 0x5e, 0x79, 0xdd, 0xde, 0x08, 0x33, 0x80, 0xde, 0x07, 0x43, 0xf1,
	0xc5, 0x54, 0x58, 0xcd, 0x79, 0xf4, 0x63, 0xd1, 0xf9, 0xc5, 0x7f, 0x17, 0xd7, 0xbc, 0xbd, 0x7d,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x0e, 0x1e, 0x21, 0x1b, 0x01, 0x00, 0x00,
}
