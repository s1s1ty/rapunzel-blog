// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/session.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Session struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=AccessToken" json:"AccessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=RefreshToken" json:"RefreshToken,omitempty"`
	CreatedTime  string `protobuf:"bytes,3,opt,name=CreatedTime" json:"CreatedTime,omitempty"`
	ExpireTime   string `protobuf:"bytes,4,opt,name=ExpireTime" json:"ExpireTime,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *Session) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *Session) GetRefreshToken() string {
	if m != nil {
		return m.RefreshToken
	}
	return ""
}

func (m *Session) GetCreatedTime() string {
	if m != nil {
		return m.CreatedTime
	}
	return ""
}

func (m *Session) GetExpireTime() string {
	if m != nil {
		return m.ExpireTime
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "protos.Session")
}

func init() { proto.RegisterFile("protos/session.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x03, 0x73, 0x85, 0xd8, 0x20,
	0xa2, 0x4a, 0x13, 0x19, 0xb9, 0xd8, 0x83, 0x21, 0x32, 0x42, 0x0a, 0x5c, 0xdc, 0x8e, 0xc9, 0xc9,
	0xa9, 0xc5, 0xc5, 0x21, 0xf9, 0xd9, 0xa9, 0x79, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8,
	0x42, 0x42, 0x4a, 0x5c, 0x3c, 0x41, 0xa9, 0x69, 0x45, 0xa9, 0xc5, 0x19, 0x10, 0x25, 0x4c, 0x60,
	0x25, 0x28, 0x62, 0x20, 0x53, 0x9c, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0x53, 0x42, 0x32, 0x73, 0x53,
	0x25, 0x98, 0x21, 0xa6, 0x20, 0x09, 0x09, 0xc9, 0x71, 0x71, 0xb9, 0x56, 0x14, 0x64, 0x16, 0xa5,
	0x82, 0x15, 0xb0, 0x80, 0x15, 0x20, 0x89, 0x24, 0x41, 0xdc, 0x66, 0x0c, 0x08, 0x00, 0x00, 0xff,
	0xff, 0xa0, 0x5f, 0x83, 0x77, 0xba, 0x00, 0x00, 0x00,
}
