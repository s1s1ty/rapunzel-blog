// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/services.proto

package protos

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func init() { proto.RegisterFile("protos/services.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 530 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0x86, 0xad, 0xef, 0x53, 0x7f, 0x34, 0x21, 0x05, 0xdc, 0x1f, 0x5a, 0xf7, 0xa0, 0x92, 0xcf,
	0x5b, 0xa4, 0x44, 0x14, 0x24, 0x44, 0x91, 0x48, 0x20, 0x80, 0x72, 0x10, 0x12, 0x7a, 0xc2, 0xd9,
	0x92, 0x0e, 0x61, 0xa5, 0x36, 0x4e, 0x3d, 0xde, 0xa2, 0x5c, 0x06, 0xf7, 0xc0, 0x9d, 0x70, 0x63,
	0x68, 0x77, 0xbc, 0x59, 0x6f, 0xd7, 0xe9, 0x61, 0x9e, 0x77, 0xde, 0x99, 0xf1, 0xcc, 0x64, 0x61,
	0x7f, 0x61, 0x8d, 0x33, 0xf4, 0x9c, 0xd0, 0xde, 0xa5, 0x53, 0xa4, 0xb3, 0xe2, 0xb7, 0xdc, 0x64,
	0xac, 0x9e, 0x7a, 0x39, 0x23, 0xb4, 0x2c, 0xad, 0xd0, 0xc2, 0x90, 0xf3, 0x68, 0xcf, 0xa3, 0xa9,
	0xb9, 0xb9, 0xc1, 0x79, 0x49, 0x77, 0x03, 0x6a, 0xe6, 0x0c, 0x3b, 0x7f, 0x36, 0xa0, 0x75, 0x49,
	0x68, 0x27, 0x5c, 0x4f, 0x5e, 0xc0, 0xf6, 0x18, 0x67, 0x29, 0x39, 0xb4, 0xf2, 0x19, 0xc7, 0xd0,
	0xd9, 0x18, 0x6f, 0x19, 0xda, 0xc4, 0xa5, 0x66, 0xae, 0x02, 0x81, 0x42, 0x41, 0x0b, 0x79, 0x0a,
	0x1b, 0x43, 0x33, 0x4b, 0xe7, 0xf2, 0x49, 0x60, 0x2e, 0x88, 0x0a, 0x08, 0x15, 0x44, 0x0b, 0xd9,
	0x85, 0xad, 0x91, 0x35, 0x3f, 0xd2, 0x6b, 0x94, 0x32, 0x30, 0x78, 0xa6, 0x02, 0x46, 0x9e, 0x69,
	0x21, 0x5f, 0xc1, 0xe6, 0xe5, 0xe2, 0x2a, 0x71, 0x28, 0xf7, 0x03, 0x0f, 0xa3, 0xfc, 0x2d, 0x2a,
	0xc0, 0x54, 0x61, 0x2d, 0xe4, 0x47, 0xd8, 0xe9, 0xfd, 0x4c, 0xe6, 0x33, 0x1c, 0x25, 0x44, 0xbf,
	0x8c, 0xbd, 0x92, 0x47, 0x41, 0x86, 0xba, 0xa4, 0x02, 0x89, 0xea, 0x92, 0x16, 0xf2, 0x33, 0x3c,
	0x62, 0x36, 0x71, 0x89, 0xcb, 0x48, 0x1e, 0x37, 0xf2, 0x14, 0x53, 0x2d, 0x44, 0x75, 0xdc, 0xc8,
	0x54, 0x89, 0x5a, 0xc8, 0x3e, 0x00, 0xd3, 0xaf, 0xcb, 0x05, 0x46, 0x3a, 0xca, 0x83, 0x73, 0x29,
	0xd2, 0x51, 0x29, 0xf1, 0x54, 0x86, 0x66, 0x66, 0x32, 0x57, 0x9f, 0x0a, 0xa1, 0x65, 0x5c, 0x9f,
	0xca, 0x0a, 0x6b, 0x21, 0xbf, 0xc1, 0xde, 0x18, 0x09, 0x5d, 0xf9, 0xbc, 0x31, 0xde, 0x66, 0x48,
	0x4e, 0x9e, 0xd4, 0xf6, 0xdf, 0x0c, 0x50, 0x27, 0xb5, 0x3b, 0x68, 0x06, 0x68, 0x21, 0xdf, 0x43,
	0xbb, 0xa6, 0xc8, 0xc3, 0x75, 0x49, 0xd5, 0xe1, 0xba, 0x6c, 0x5a, 0x74, 0x7e, 0xff, 0x0f, 0xad,
	0x91, 0x21, 0x57, 0x9e, 0xe9, 0x39, 0x40, 0xcf, 0x62, 0xe2, 0x30, 0x87, 0xb5, 0x07, 0xe7, 0x80,
	0x25, 0xf5, 0x38, 0xbc, 0x1e, 0x53, 0xb4, 0x73, 0x0e, 0xc0, 0x07, 0x11, 0xf5, 0xb1, 0x14, 0xf3,
	0xbd, 0x04, 0xe8, 0xe3, 0x35, 0x7a, 0xdf, 0x2a, 0x60, 0x80, 0xee, 0xdd, 0xf2, 0x53, 0x5f, 0x1d,
	0xdc, 0x73, 0x4c, 0xb2, 0xe9, 0x14, 0x29, 0xdf, 0xed, 0xc5, 0x03, 0x77, 0x52, 0xb4, 0x1a, 0x88,
	0xb1, 0xc2, 0xa7, 0xb0, 0x35, 0x40, 0x17, 0xaf, 0x1a, 0x09, 0xef, 0x42, 0xfb, 0x43, 0x72, 0x67,
	0x32, 0x9b, 0xae, 0x6b, 0x35, 0x62, 0x7a, 0x01, 0xdb, 0xbe, 0x06, 0x55, 0xff, 0xc2, 0x22, 0xfe,
	0x4b, 0x86, 0x76, 0xa9, 0x76, 0xef, 0x59, 0x86, 0x69, 0x6e, 0xeb, 0xfc, 0xfd, 0x0f, 0x76, 0x7a,
	0xfc, 0x85, 0x29, 0xd7, 0xf2, 0x16, 0xda, 0x3c, 0x7b, 0xcf, 0x6b, 0xdb, 0xf6, 0xcc, 0x2f, 0x27,
	0xfc, 0x6b, 0x7b, 0x45, 0x8b, 0x3c, 0x01, 0x2f, 0xe1, 0x81, 0x04, 0x7e, 0x4b, 0xf1, 0x04, 0x6f,
	0xa0, 0xcd, 0x8b, 0x2a, 0x13, 0x34, 0x06, 0x70, 0xd4, 0xf4, 0x55, 0xeb, 0xea, 0x02, 0x0c, 0xd0,
	0xad, 0xf5, 0xc6, 0x6b, 0xbe, 0x86, 0x56, 0x65, 0x8a, 0x8f, 0xf0, 0xa0, 0x69, 0xe4, 0x29, 0x7e,
	0xe7, 0x2f, 0x7b, 0xf7, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xce, 0x95, 0x01, 0x3f, 0xf9, 0x05,
	0x00, 0x00,
}
