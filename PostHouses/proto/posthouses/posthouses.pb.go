// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/posthouses/posthouses.proto

package go_micro_srv_PostHouses

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
	Data                 []byte   `protobuf:"bytes,2,opt,name=Data,proto3" json:"Data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_741c3cf0f93ca077, []int{0}
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

func (m *Request) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type Response struct {
	ErrNo                string   `protobuf:"bytes,1,opt,name=ErrNo,proto3" json:"ErrNo,omitempty"`
	ErrMsg               string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
	HouseId              string   `protobuf:"bytes,3,opt,name=HouseId,proto3" json:"HouseId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_741c3cf0f93ca077, []int{1}
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

func (m *Response) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostHouses.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostHouses.Response")
}

func init() { proto.RegisterFile("proto/posthouses/posthouses.proto", fileDescriptor_741c3cf0f93ca077) }

var fileDescriptor_741c3cf0f93ca077 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xc8, 0x2f, 0x2e, 0xc9, 0xc8, 0x2f, 0x2d, 0x4e, 0x2d, 0x46, 0x62, 0xea, 0x81,
	0xe5, 0x84, 0xc4, 0xd3, 0xf3, 0xf5, 0x72, 0x33, 0x93, 0x8b, 0xf2, 0xf5, 0x8a, 0x8b, 0xca, 0xf4,
	0x02, 0xf2, 0x8b, 0x4b, 0x3c, 0xc0, 0xd2, 0x4a, 0xd6, 0x5c, 0xec, 0x41, 0xa9, 0x85, 0xa5, 0xa9,
	0xc5, 0x25, 0x42, 0x32, 0x5c, 0x9c, 0xc1, 0xa9, 0xc5, 0xc5, 0x99, 0xf9, 0x79, 0x9e, 0x29, 0x12,
	0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x08, 0x01, 0x21, 0x21, 0x2e, 0x16, 0x97, 0xc4, 0x92, 0x44,
	0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20, 0x30, 0x5b, 0x29, 0x88, 0x8b, 0x23, 0x28, 0xb5, 0xb8,
	0x20, 0x3f, 0xaf, 0x38, 0x55, 0x48, 0x84, 0x8b, 0xd5, 0xb5, 0xa8, 0xc8, 0x2f, 0x1f, 0xaa, 0x13,
	0xc2, 0x11, 0x12, 0xe3, 0x62, 0x73, 0x2d, 0x2a, 0xf2, 0x2d, 0x4e, 0x07, 0xeb, 0xe3, 0x0c, 0x82,
	0xf2, 0x84, 0x24, 0xb8, 0xd8, 0xc1, 0x0e, 0xf0, 0x4c, 0x91, 0x60, 0x06, 0x4b, 0xc0, 0xb8, 0x46,
	0x89, 0x5c, 0x5c, 0x08, 0xe7, 0x09, 0x05, 0xa3, 0xf0, 0x14, 0xf4, 0x70, 0x78, 0x43, 0x0f, 0xea,
	0x07, 0x29, 0x45, 0x3c, 0x2a, 0x20, 0x0e, 0x55, 0x62, 0x48, 0x62, 0x03, 0x87, 0x89, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x07, 0x4e, 0xc9, 0xf4, 0x38, 0x01, 0x00, 0x00,
}