// Code generated by protoc-gen-go. DO NOT EDIT.
// source: defs/common.proto

package defs

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Pager struct {
	CurrentPage int32 `protobuf:"varint,1,opt,name=current_page,json=currentPage" json:"current_page,omitempty"`
	TotalPage   int32 `protobuf:"varint,2,opt,name=total_page,json=totalPage" json:"total_page,omitempty"`
	Total       int32 `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	Current     int32 `protobuf:"varint,4,opt,name=current" json:"current,omitempty"`
}

func (m *Pager) Reset()                    { *m = Pager{} }
func (m *Pager) String() string            { return proto.CompactTextString(m) }
func (*Pager) ProtoMessage()               {}
func (*Pager) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *Pager) GetCurrentPage() int32 {
	if m != nil {
		return m.CurrentPage
	}
	return 0
}

func (m *Pager) GetTotalPage() int32 {
	if m != nil {
		return m.TotalPage
	}
	return 0
}

func (m *Pager) GetTotal() int32 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *Pager) GetCurrent() int32 {
	if m != nil {
		return m.Current
	}
	return 0
}

type Query struct {
	Field string `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *Query) GetField() string {
	if m != nil {
		return m.Field
	}
	return ""
}

func (m *Query) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type GetByQuery struct {
	Query []*Query `protobuf:"bytes,1,rep,name=query" json:"query,omitempty"`
}

func (m *GetByQuery) Reset()                    { *m = GetByQuery{} }
func (m *GetByQuery) String() string            { return proto.CompactTextString(m) }
func (*GetByQuery) ProtoMessage()               {}
func (*GetByQuery) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetByQuery) GetQuery() []*Query {
	if m != nil {
		return m.Query
	}
	return nil
}

type GetByID struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetByID) Reset()                    { *m = GetByID{} }
func (m *GetByID) String() string            { return proto.CompactTextString(m) }
func (*GetByID) ProtoMessage()               {}
func (*GetByID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetByID) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func init() {
	proto.RegisterType((*Pager)(nil), "defs.Pager")
	proto.RegisterType((*Query)(nil), "defs.Query")
	proto.RegisterType((*GetByQuery)(nil), "defs.GetByQuery")
	proto.RegisterType((*GetByID)(nil), "defs.GetByID")
}

func init() { proto.RegisterFile("defs/common.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x90, 0xc1, 0x6e, 0xb3, 0x30,
	0x10, 0x84, 0x05, 0x09, 0x7f, 0xc4, 0xf2, 0xab, 0x52, 0x51, 0x0f, 0xf4, 0x50, 0x29, 0xe1, 0x94,
	0x4b, 0x6c, 0xa9, 0xc9, 0x13, 0xd0, 0x4a, 0x55, 0x6f, 0x94, 0x63, 0x2f, 0x95, 0x0d, 0x0e, 0x75,
	0x03, 0x36, 0x35, 0xa6, 0x12, 0x7d, 0xfa, 0xca, 0xeb, 0x70, 0xdb, 0x6f, 0x66, 0x87, 0xc1, 0x0b,
	0xb7, 0x8d, 0x38, 0x8f, 0xb4, 0xd6, 0x7d, 0xaf, 0x15, 0x19, 0x8c, 0xb6, 0x3a, 0x5d, 0x3b, 0x29,
	0x9f, 0x21, 0x2a, 0x59, 0x2b, 0x4c, 0xba, 0x83, 0xff, 0xf5, 0x64, 0x8c, 0x50, 0xf6, 0x63, 0x60,
	0xad, 0xc8, 0x82, 0x6d, 0xb0, 0x8f, 0xaa, 0xe4, 0xaa, 0xb9, 0x9d, 0xf4, 0x01, 0xc0, 0x6a, 0xcb,
	0x3a, 0xbf, 0x10, 0xe2, 0x42, 0x8c, 0x0a, 0xda, 0x77, 0x10, 0x21, 0x64, 0x2b, 0x74, 0x3c, 0xa4,
	0x19, 0x6c, 0xae, 0xdf, 0xc8, 0xd6, 0xa8, 0x2f, 0x98, 0x1f, 0x21, 0x7a, 0x9b, 0x84, 0x99, 0x5d,
	0xf0, 0x2c, 0x45, 0xd7, 0x60, 0x67, 0x5c, 0x79, 0x70, 0xea, 0x0f, 0xeb, 0x26, 0x5f, 0x14, 0x57,
	0x1e, 0x72, 0x0a, 0xf0, 0x22, 0x6c, 0x31, 0xfb, 0xe4, 0x0e, 0xa2, 0x6f, 0x37, 0x64, 0xc1, 0x76,
	0xb5, 0x4f, 0x1e, 0x13, 0xe2, 0xde, 0x44, 0xd0, 0xab, 0xbc, 0x93, 0xdf, 0xc3, 0x06, 0x03, 0xaf,
	0xcf, 0xe9, 0x0d, 0x84, 0x72, 0x29, 0x09, 0x65, 0x53, 0x70, 0xc8, 0x95, 0x54, 0x5f, 0x8c, 0x8c,
	0xec, 0x22, 0x39, 0x31, 0x6c, 0x98, 0xd4, 0xaf, 0xe8, 0x08, 0x53, 0x8d, 0xd1, 0xb2, 0xf1, 0x77,
	0x2a, 0x92, 0x27, 0xbc, 0x5a, 0xe9, 0xa0, 0x0c, 0xde, 0x0f, 0xad, 0xb4, 0x9f, 0x13, 0x27, 0xb5,
	0xee, 0xe9, 0x78, 0xba, 0x48, 0x3e, 0x9e, 0x7a, 0x49, 0x97, 0xec, 0x81, 0x77, 0xba, 0xa5, 0x18,
	0xa4, 0xee, 0x5f, 0xf8, 0x3f, 0x9c, 0x8f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x99, 0xd5, 0x16,
	0xf1, 0x81, 0x01, 0x00, 0x00,
}
