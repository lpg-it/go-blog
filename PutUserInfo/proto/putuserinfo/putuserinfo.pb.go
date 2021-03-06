// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/putuserinfo/putuserinfo.proto

package go_micro_srv_PutUserInfo

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
	UserName             string   `protobuf:"bytes,2,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_488f4dc530685179, []int{0}
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

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

type Response struct {
	ErrNo                string   `protobuf:"bytes,1,opt,name=ErrNo,proto3" json:"ErrNo,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	UserName             string   `protobuf:"bytes,3,opt,name=UserName,proto3" json:"UserName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_488f4dc530685179, []int{1}
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

func (m *Response) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PutUserInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PutUserInfo.Response")
}

func init() {
	proto.RegisterFile("proto/putuserinfo/putuserinfo.proto", fileDescriptor_488f4dc530685179)
}

var fileDescriptor_488f4dc530685179 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x28, 0x2d, 0x29, 0x2d, 0x4e, 0x2d, 0xca, 0xcc, 0x4b, 0x43, 0x61, 0xeb, 0x81,
	0x65, 0x85, 0x24, 0xd2, 0xf3, 0xf5, 0x72, 0x33, 0x93, 0x8b, 0xf2, 0xf5, 0x8a, 0x8b, 0xca, 0xf4,
	0x02, 0x4a, 0x4b, 0x42, 0x8b, 0x53, 0x8b, 0x3c, 0xf3, 0xd2, 0xf2, 0x95, 0x9c, 0xb9, 0xd8, 0x83,
	0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x64, 0xb8, 0x38, 0x83, 0x53, 0x8b, 0x8b, 0x33, 0xf3,
	0xf3, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x10, 0x02, 0x42, 0x52, 0x5c, 0x1c,
	0x20, 0x4d, 0x7e, 0x89, 0xb9, 0xa9, 0x12, 0x4c, 0x60, 0x49, 0x38, 0x5f, 0x29, 0x84, 0x8b, 0x23,
	0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84, 0x8b, 0xd5, 0xb5, 0xa8, 0xc8, 0x2f,
	0x1f, 0x6a, 0x02, 0x84, 0x23, 0x24, 0xc6, 0xc5, 0xe6, 0x5a, 0x54, 0xe4, 0x5b, 0x9c, 0x0e, 0xd5,
	0x0b, 0xe5, 0xa1, 0x98, 0xca, 0x8c, 0x6a, 0xaa, 0x51, 0x2a, 0x17, 0x37, 0x92, 0x4b, 0x85, 0xc2,
	0x50, 0xb9, 0x8a, 0x7a, 0xb8, 0xfc, 0xa4, 0x07, 0xf5, 0x90, 0x94, 0x12, 0x3e, 0x25, 0x10, 0xe7,
	0x2a, 0x31, 0x24, 0xb1, 0x81, 0x83, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x9c, 0xf5,
	0x16, 0x49, 0x01, 0x00, 0x00,
}
