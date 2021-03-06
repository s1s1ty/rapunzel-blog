// Code generated by protoc-gen-go. DO NOT EDIT.
// source: defs/error.proto

package defs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ErrorDetails struct {
	Field   string   `protobuf:"bytes,1,opt,name=Field" json:"Field,omitempty"`
	Details []string `protobuf:"bytes,2,rep,name=Details" json:"Details,omitempty"`
}

func (m *ErrorDetails) Reset()                    { *m = ErrorDetails{} }
func (m *ErrorDetails) String() string            { return proto.CompactTextString(m) }
func (*ErrorDetails) ProtoMessage()               {}
func (*ErrorDetails) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

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
func (*Error) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

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
	proto.RegisterType((*ErrorDetails)(nil), "defs.ErrorDetails")
	proto.RegisterType((*Error)(nil), "defs.Error")
}

func init() { proto.RegisterFile("defs/error.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0xfc, 0x01, 0xf5, 0x40, 0x08, 0x59, 0x0c, 0x1e, 0xa3, 0x4c, 0x59, 0x6a, 0x4b,
	0x50, 0x31, 0x32, 0x94, 0x80, 0xd4, 0xad, 0x8a, 0x98, 0xd8, 0x6c, 0x6c, 0xca, 0x51, 0x37, 0xae,
	0xec, 0x74, 0xe1, 0x6b, 0xf0, 0x85, 0x2b, 0x3b, 0x8d, 0xd4, 0x6e, 0xef, 0x9d, 0xdf, 0xfb, 0xf9,
	0x74, 0x70, 0xaf, 0xcd, 0x77, 0x10, 0xc6, 0x7b, 0xe7, 0xf9, 0xde, 0xbb, 0xc1, 0xd1, 0x22, 0x4e,
	0xea, 0x17, 0xb8, 0x7d, 0x8b, 0xc3, 0xd6, 0x0c, 0x12, 0x6d, 0xa0, 0x0f, 0x50, 0xbe, 0xa3, 0xb1,
	0x9a, 0x91, 0x8a, 0x34, 0xb3, 0x6e, 0x34, 0x94, 0xc1, 0xf5, 0x29, 0xc0, 0xb2, 0x2a, 0x6f, 0x66,
	0xdd, 0x64, 0xeb, 0x7f, 0x02, 0x65, 0x02, 0xd0, 0x3b, 0xc8, 0x56, 0xed, 0xa9, 0x96, 0xad, 0xda,
	0x48, 0xfa, 0xc0, 0xc1, 0x1a, 0x96, 0x8d, 0xa4, 0x64, 0x28, 0x85, 0xe2, 0xd5, 0x69, 0xc3, 0xf2,
	0x8a, 0x34, 0x65, 0x97, 0xf4, 0x39, 0xbd, 0x48, 0xd9, 0xc9, 0xd2, 0xe7, 0xcb, 0xed, 0x58, 0x59,
	0xe5, 0xcd, 0xcd, 0x23, 0xe5, 0x71, 0x75, 0x7e, 0xfe, 0xd2, 0x5d, 0xe4, 0x96, 0x12, 0xea, 0x1e,
	0xfb, 0x5f, 0xc9, 0x83, 0xdc, 0xa2, 0xe2, 0x5e, 0xee, 0x0f, 0xfd, 0x9f, 0xb1, 0x5c, 0xf6, 0xda,
	0x3b, 0xd4, 0xe3, 0x05, 0x96, 0x90, 0x3a, 0xeb, 0xa8, 0xd7, 0xe4, 0x73, 0xbe, 0xc1, 0xe1, 0xe7,
	0xa0, 0xf8, 0x97, 0xdb, 0x89, 0xb0, 0xd8, 0xa2, 0x0a, 0x8b, 0x1d, 0x8a, 0xa9, 0x3a, 0x57, 0xd6,
	0x6d, 0x44, 0xea, 0x89, 0xf8, 0xbb, 0xba, 0x4a, 0xfa, 0xe9, 0x18, 0x00, 0x00, 0xff, 0xff, 0x5d,
	0x92, 0xec, 0x5c, 0x59, 0x01, 0x00, 0x00,
}
