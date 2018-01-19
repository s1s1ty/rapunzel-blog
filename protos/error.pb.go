// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/error.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	protos/error.proto
	protos/services.proto
	protos/session.proto
	protos/user.proto

It has these top-level messages:
	ErrorDetails
	Error
	Session
	User
	ReqRegistration
	ResRegistration
	ReqLogin
	ResLogin
	ReqUpdateUser
	ResUpdateUser
	ReqChangePassword
	ResChangePassword
	ReqChangeUserStatus
	ResChangeUserStatus
	ReqChangeUserType
	ResChangeUserType
	ReqUserLogout
	ResUserLogout
	ReqResetPasswordRequest
	ResResetPasswordRequest
	ReqResetPassword
	ResResetPassword
*/
package protos

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

type ErrorDetails struct {
	Field   string   `protobuf:"bytes,1,opt,name=Field" json:"Field,omitempty"`
	Details []string `protobuf:"bytes,2,rep,name=Details" json:"Details,omitempty"`
}

func (m *ErrorDetails) Reset()                    { *m = ErrorDetails{} }
func (m *ErrorDetails) String() string            { return proto.CompactTextString(m) }
func (*ErrorDetails) ProtoMessage()               {}
func (*ErrorDetails) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ErrorDetails) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *ErrorDetails) GetDetails() []string {
	if m != nil {
		return m.Details
	}
	return nil
}

type Error struct {
	ID           string          `protobuf:"bytes,1,opt,name=ID" json:"ID,omitempty"`
	Title        string          `protobuf:"bytes,2,opt,name=Title" json:"Title,omitempty"`
	Code         int32           `protobuf:"varint,3,opt,name=Code" json:"Code,omitempty"`
	Details      string          `protobuf:"bytes,4,opt,name=Details" json:"Details,omitempty"`
	ErrorDetails []*ErrorDetails `protobuf:"bytes,5,rep,name=ErrorDetails" json:"ErrorDetails,omitempty"`
}

func (m *Error) Reset()                    { *m = Error{} }
func (m *Error) String() string            { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()               {}
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Error) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Error) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Error) GetCode() int32 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

func (m *Error) GetErrorDetails() []*ErrorDetails {
	if m != nil {
		return m.ErrorDetails
	}
	return nil
}

func init() {
	proto.RegisterType((*ErrorDetails)(nil), "protos.ErrorDetails")
	proto.RegisterType((*Error)(nil), "protos.Error")
}

func init() { proto.RegisterFile("protos/error.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x28, 0xca, 0x2f,
	0xc9, 0x2f, 0xd6, 0x4f, 0x2d, 0x2a, 0xca, 0x2f, 0xd2, 0x03, 0x73, 0x84, 0xd8, 0x20, 0x62, 0x4a,
	0x76, 0x5c, 0x3c, 0xae, 0x20, 0x61, 0x97, 0xd4, 0x92, 0xc4, 0xcc, 0x9c, 0x62, 0x21, 0x11, 0x2e,
	0x56, 0xb7, 0xcc, 0xd4, 0x9c, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48,
	0x82, 0x8b, 0x1d, 0xaa, 0x40, 0x82, 0x49, 0x81, 0x59, 0x83, 0x33, 0x08, 0xc6, 0x55, 0x9a, 0xca,
	0xc8, 0xc5, 0x0a, 0x36, 0x40, 0x88, 0x8f, 0x8b, 0xc9, 0xd3, 0x05, 0xaa, 0x8d, 0xc9, 0xd3, 0x05,
	0x64, 0x52, 0x48, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0x13, 0xc4, 0x24, 0x30, 0x47, 0x48, 0x88, 0x8b,
	0xc5, 0x39, 0x3f, 0x25, 0x55, 0x82, 0x59, 0x81, 0x51, 0x83, 0x35, 0x08, 0xcc, 0x46, 0x36, 0x9d,
	0x05, 0xac, 0x16, 0xc6, 0x15, 0xb2, 0x40, 0x75, 0x9d, 0x04, 0xab, 0x02, 0xb3, 0x06, 0xb7, 0x91,
	0x08, 0xc4, 0x0f, 0xc5, 0x7a, 0xc8, 0x72, 0x41, 0x28, 0x2a, 0x93, 0x20, 0xfe, 0x33, 0x06, 0x04,
	0x00, 0x00, 0xff, 0xff, 0xe7, 0xef, 0xa1, 0xb3, 0xfc, 0x00, 0x00, 0x00,
}
