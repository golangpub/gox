// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base.proto

package base // import "github.com/gopub/gox/protobuf/base"

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

type VoidValue struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoidValue) Reset()         { *m = VoidValue{} }
func (m *VoidValue) String() string { return proto.CompactTextString(m) }
func (*VoidValue) ProtoMessage()    {}
func (*VoidValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{0}
}
func (m *VoidValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoidValue.Unmarshal(m, b)
}
func (m *VoidValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoidValue.Marshal(b, m, deterministic)
}
func (dst *VoidValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoidValue.Merge(dst, src)
}
func (m *VoidValue) XXX_Size() int {
	return xxx_messageInfo_VoidValue.Size(m)
}
func (m *VoidValue) XXX_DiscardUnknown() {
	xxx_messageInfo_VoidValue.DiscardUnknown(m)
}

var xxx_messageInfo_VoidValue proto.InternalMessageInfo

type BoolValue struct {
	Value                bool     `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BoolValue) Reset()         { *m = BoolValue{} }
func (m *BoolValue) String() string { return proto.CompactTextString(m) }
func (*BoolValue) ProtoMessage()    {}
func (*BoolValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{1}
}
func (m *BoolValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BoolValue.Unmarshal(m, b)
}
func (m *BoolValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BoolValue.Marshal(b, m, deterministic)
}
func (dst *BoolValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BoolValue.Merge(dst, src)
}
func (m *BoolValue) XXX_Size() int {
	return xxx_messageInfo_BoolValue.Size(m)
}
func (m *BoolValue) XXX_DiscardUnknown() {
	xxx_messageInfo_BoolValue.DiscardUnknown(m)
}

var xxx_messageInfo_BoolValue proto.InternalMessageInfo

func (m *BoolValue) GetValue() bool {
	if m != nil {
		return m.Value
	}
	return false
}

type Int64Value struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64Value) Reset()         { *m = Int64Value{} }
func (m *Int64Value) String() string { return proto.CompactTextString(m) }
func (*Int64Value) ProtoMessage()    {}
func (*Int64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{2}
}
func (m *Int64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64Value.Unmarshal(m, b)
}
func (m *Int64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64Value.Marshal(b, m, deterministic)
}
func (dst *Int64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64Value.Merge(dst, src)
}
func (m *Int64Value) XXX_Size() int {
	return xxx_messageInfo_Int64Value.Size(m)
}
func (m *Int64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64Value.DiscardUnknown(m)
}

var xxx_messageInfo_Int64Value proto.InternalMessageInfo

func (m *Int64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Float64Value struct {
	Value                int64    `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Float64Value) Reset()         { *m = Float64Value{} }
func (m *Float64Value) String() string { return proto.CompactTextString(m) }
func (*Float64Value) ProtoMessage()    {}
func (*Float64Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{3}
}
func (m *Float64Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float64Value.Unmarshal(m, b)
}
func (m *Float64Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float64Value.Marshal(b, m, deterministic)
}
func (dst *Float64Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float64Value.Merge(dst, src)
}
func (m *Float64Value) XXX_Size() int {
	return xxx_messageInfo_Float64Value.Size(m)
}
func (m *Float64Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Float64Value.DiscardUnknown(m)
}

var xxx_messageInfo_Float64Value proto.InternalMessageInfo

func (m *Float64Value) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type Int64List struct {
	Value                []int64  `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Int64List) Reset()         { *m = Int64List{} }
func (m *Int64List) String() string { return proto.CompactTextString(m) }
func (*Int64List) ProtoMessage()    {}
func (*Int64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{4}
}
func (m *Int64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Int64List.Unmarshal(m, b)
}
func (m *Int64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Int64List.Marshal(b, m, deterministic)
}
func (dst *Int64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Int64List.Merge(dst, src)
}
func (m *Int64List) XXX_Size() int {
	return xxx_messageInfo_Int64List.Size(m)
}
func (m *Int64List) XXX_DiscardUnknown() {
	xxx_messageInfo_Int64List.DiscardUnknown(m)
}

var xxx_messageInfo_Int64List proto.InternalMessageInfo

func (m *Int64List) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type Float64List struct {
	Value                []int64  `protobuf:"varint,1,rep,packed,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Float64List) Reset()         { *m = Float64List{} }
func (m *Float64List) String() string { return proto.CompactTextString(m) }
func (*Float64List) ProtoMessage()    {}
func (*Float64List) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{5}
}
func (m *Float64List) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Float64List.Unmarshal(m, b)
}
func (m *Float64List) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Float64List.Marshal(b, m, deterministic)
}
func (dst *Float64List) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Float64List.Merge(dst, src)
}
func (m *Float64List) XXX_Size() int {
	return xxx_messageInfo_Float64List.Size(m)
}
func (m *Float64List) XXX_DiscardUnknown() {
	xxx_messageInfo_Float64List.DiscardUnknown(m)
}

var xxx_messageInfo_Float64List proto.InternalMessageInfo

func (m *Float64List) GetValue() []int64 {
	if m != nil {
		return m.Value
	}
	return nil
}

type StringValue struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringValue) Reset()         { *m = StringValue{} }
func (m *StringValue) String() string { return proto.CompactTextString(m) }
func (*StringValue) ProtoMessage()    {}
func (*StringValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{6}
}
func (m *StringValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringValue.Unmarshal(m, b)
}
func (m *StringValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringValue.Marshal(b, m, deterministic)
}
func (dst *StringValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringValue.Merge(dst, src)
}
func (m *StringValue) XXX_Size() int {
	return xxx_messageInfo_StringValue.Size(m)
}
func (m *StringValue) XXX_DiscardUnknown() {
	xxx_messageInfo_StringValue.DiscardUnknown(m)
}

var xxx_messageInfo_StringValue proto.InternalMessageInfo

func (m *StringValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StringList struct {
	Value                []string `protobuf:"bytes,1,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StringList) Reset()         { *m = StringList{} }
func (m *StringList) String() string { return proto.CompactTextString(m) }
func (*StringList) ProtoMessage()    {}
func (*StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{7}
}
func (m *StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StringList.Unmarshal(m, b)
}
func (m *StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StringList.Marshal(b, m, deterministic)
}
func (dst *StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StringList.Merge(dst, src)
}
func (m *StringList) XXX_Size() int {
	return xxx_messageInfo_StringList.Size(m)
}
func (m *StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_StringList proto.InternalMessageInfo

func (m *StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type ListQuery struct {
	SinceId              int64    `protobuf:"varint,1,opt,name=since_id,json=sinceId,proto3" json:"since_id,omitempty"`
	Count                int32    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListQuery) Reset()         { *m = ListQuery{} }
func (m *ListQuery) String() string { return proto.CompactTextString(m) }
func (*ListQuery) ProtoMessage()    {}
func (*ListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{8}
}
func (m *ListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListQuery.Unmarshal(m, b)
}
func (m *ListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListQuery.Marshal(b, m, deterministic)
}
func (dst *ListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListQuery.Merge(dst, src)
}
func (m *ListQuery) XXX_Size() int {
	return xxx_messageInfo_ListQuery.Size(m)
}
func (m *ListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_ListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_ListQuery proto.InternalMessageInfo

func (m *ListQuery) GetSinceId() int64 {
	if m != nil {
		return m.SinceId
	}
	return 0
}

func (m *ListQuery) GetCount() int32 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Area struct {
	MinLat               float64  `protobuf:"fixed64,1,opt,name=min_lat,json=minLat,proto3" json:"min_lat,omitempty"`
	MaxLat               float64  `protobuf:"fixed64,2,opt,name=max_lat,json=maxLat,proto3" json:"max_lat,omitempty"`
	MinLng               float64  `protobuf:"fixed64,3,opt,name=min_lng,json=minLng,proto3" json:"min_lng,omitempty"`
	MaxLng               float64  `protobuf:"fixed64,4,opt,name=max_lng,json=maxLng,proto3" json:"max_lng,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Area) Reset()         { *m = Area{} }
func (m *Area) String() string { return proto.CompactTextString(m) }
func (*Area) ProtoMessage()    {}
func (*Area) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{9}
}
func (m *Area) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Area.Unmarshal(m, b)
}
func (m *Area) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Area.Marshal(b, m, deterministic)
}
func (dst *Area) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Area.Merge(dst, src)
}
func (m *Area) XXX_Size() int {
	return xxx_messageInfo_Area.Size(m)
}
func (m *Area) XXX_DiscardUnknown() {
	xxx_messageInfo_Area.DiscardUnknown(m)
}

var xxx_messageInfo_Area proto.InternalMessageInfo

func (m *Area) GetMinLat() float64 {
	if m != nil {
		return m.MinLat
	}
	return 0
}

func (m *Area) GetMaxLat() float64 {
	if m != nil {
		return m.MaxLat
	}
	return 0
}

func (m *Area) GetMinLng() float64 {
	if m != nil {
		return m.MinLng
	}
	return 0
}

func (m *Area) GetMaxLng() float64 {
	if m != nil {
		return m.MaxLng
	}
	return 0
}

type Coordinate struct {
	Longitude            float64  `protobuf:"fixed64,1,opt,name=longitude,proto3" json:"longitude,omitempty"`
	Latitude             float64  `protobuf:"fixed64,2,opt,name=latitude,proto3" json:"latitude,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Coordinate) Reset()         { *m = Coordinate{} }
func (m *Coordinate) String() string { return proto.CompactTextString(m) }
func (*Coordinate) ProtoMessage()    {}
func (*Coordinate) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{10}
}
func (m *Coordinate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Coordinate.Unmarshal(m, b)
}
func (m *Coordinate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Coordinate.Marshal(b, m, deterministic)
}
func (dst *Coordinate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinate.Merge(dst, src)
}
func (m *Coordinate) XXX_Size() int {
	return xxx_messageInfo_Coordinate.Size(m)
}
func (m *Coordinate) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinate.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinate proto.InternalMessageInfo

func (m *Coordinate) GetLongitude() float64 {
	if m != nil {
		return m.Longitude
	}
	return 0
}

func (m *Coordinate) GetLatitude() float64 {
	if m != nil {
		return m.Latitude
	}
	return 0
}

type PhoneNumber struct {
	CountryCode          int32    `protobuf:"varint,1,opt,name=country_code,json=countryCode,proto3" json:"country_code,omitempty"`
	NationalNumber       int64    `protobuf:"varint,2,opt,name=national_number,json=nationalNumber,proto3" json:"national_number,omitempty"`
	Extension            string   `protobuf:"bytes,3,opt,name=extension,proto3" json:"extension,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PhoneNumber) Reset()         { *m = PhoneNumber{} }
func (m *PhoneNumber) String() string { return proto.CompactTextString(m) }
func (*PhoneNumber) ProtoMessage()    {}
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{11}
}
func (m *PhoneNumber) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PhoneNumber.Unmarshal(m, b)
}
func (m *PhoneNumber) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PhoneNumber.Marshal(b, m, deterministic)
}
func (dst *PhoneNumber) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PhoneNumber.Merge(dst, src)
}
func (m *PhoneNumber) XXX_Size() int {
	return xxx_messageInfo_PhoneNumber.Size(m)
}
func (m *PhoneNumber) XXX_DiscardUnknown() {
	xxx_messageInfo_PhoneNumber.DiscardUnknown(m)
}

var xxx_messageInfo_PhoneNumber proto.InternalMessageInfo

func (m *PhoneNumber) GetCountryCode() int32 {
	if m != nil {
		return m.CountryCode
	}
	return 0
}

func (m *PhoneNumber) GetNationalNumber() int64 {
	if m != nil {
		return m.NationalNumber
	}
	return 0
}

func (m *PhoneNumber) GetExtension() string {
	if m != nil {
		return m.Extension
	}
	return ""
}

type FullName struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	MiddleName           string   `protobuf:"bytes,2,opt,name=middle_name,json=middleName,proto3" json:"middle_name,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FullName) Reset()         { *m = FullName{} }
func (m *FullName) String() string { return proto.CompactTextString(m) }
func (*FullName) ProtoMessage()    {}
func (*FullName) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_ce139bde246bf41c, []int{12}
}
func (m *FullName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FullName.Unmarshal(m, b)
}
func (m *FullName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FullName.Marshal(b, m, deterministic)
}
func (dst *FullName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FullName.Merge(dst, src)
}
func (m *FullName) XXX_Size() int {
	return xxx_messageInfo_FullName.Size(m)
}
func (m *FullName) XXX_DiscardUnknown() {
	xxx_messageInfo_FullName.DiscardUnknown(m)
}

var xxx_messageInfo_FullName proto.InternalMessageInfo

func (m *FullName) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *FullName) GetMiddleName() string {
	if m != nil {
		return m.MiddleName
	}
	return ""
}

func (m *FullName) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func init() {
	proto.RegisterType((*VoidValue)(nil), "gopub.protobuf.VoidValue")
	proto.RegisterType((*BoolValue)(nil), "gopub.protobuf.BoolValue")
	proto.RegisterType((*Int64Value)(nil), "gopub.protobuf.Int64Value")
	proto.RegisterType((*Float64Value)(nil), "gopub.protobuf.Float64Value")
	proto.RegisterType((*Int64List)(nil), "gopub.protobuf.Int64List")
	proto.RegisterType((*Float64List)(nil), "gopub.protobuf.Float64List")
	proto.RegisterType((*StringValue)(nil), "gopub.protobuf.StringValue")
	proto.RegisterType((*StringList)(nil), "gopub.protobuf.StringList")
	proto.RegisterType((*ListQuery)(nil), "gopub.protobuf.ListQuery")
	proto.RegisterType((*Area)(nil), "gopub.protobuf.Area")
	proto.RegisterType((*Coordinate)(nil), "gopub.protobuf.Coordinate")
	proto.RegisterType((*PhoneNumber)(nil), "gopub.protobuf.PhoneNumber")
	proto.RegisterType((*FullName)(nil), "gopub.protobuf.FullName")
}

func init() { proto.RegisterFile("base.proto", fileDescriptor_base_ce139bde246bf41c) }

var fileDescriptor_base_ce139bde246bf41c = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x4f, 0x6b, 0xd4, 0x40,
	0x14, 0x67, 0x77, 0xbb, 0x6d, 0xe6, 0xa5, 0x54, 0x18, 0x04, 0x57, 0xad, 0xd8, 0x1d, 0x0b, 0xf6,
	0xd4, 0x3d, 0x28, 0x9e, 0xbc, 0xd8, 0xc2, 0x42, 0xa1, 0x14, 0x8d, 0xd0, 0x83, 0x97, 0x30, 0xd9,
	0x4c, 0xa7, 0x23, 0x93, 0xf7, 0x4a, 0x32, 0xa3, 0xe9, 0xb7, 0x97, 0xcc, 0x24, 0x9b, 0x82, 0x8b,
	0xde, 0xf2, 0xfb, 0xfb, 0x66, 0xc8, 0x1b, 0x80, 0x42, 0x36, 0xea, 0xfc, 0xa1, 0x26, 0x47, 0xfc,
	0x48, 0xd3, 0x83, 0x2f, 0x22, 0x28, 0xfc, 0x9d, 0x48, 0x81, 0xdd, 0x92, 0x29, 0x6f, 0xa5, 0xf5,
	0x4a, 0x2c, 0x81, 0x5d, 0x10, 0xd9, 0x00, 0xf8, 0x73, 0x98, 0xff, 0xea, 0x3e, 0x16, 0x93, 0x93,
	0xc9, 0x59, 0x92, 0x45, 0x20, 0x04, 0xc0, 0x15, 0xba, 0x4f, 0x1f, 0x77, 0x78, 0x66, 0x83, 0xe7,
	0x14, 0x0e, 0xd7, 0x96, 0xe4, 0x7f, 0x5c, 0x4b, 0x60, 0xa1, 0xe9, 0xda, 0x34, 0xee, 0xa9, 0x65,
	0x36, 0x5a, 0xde, 0x41, 0xda, 0x17, 0xfd, 0xdb, 0xf4, 0xdd, 0xd5, 0x06, 0xf5, 0x8e, 0x61, 0xec,
	0xc9, 0xb1, 0xa3, 0xe9, 0xef, 0xa2, 0xad, 0xe7, 0x33, 0xb0, 0x4e, 0xfd, 0xe6, 0x55, 0xfd, 0xc8,
	0x5f, 0x42, 0xd2, 0x18, 0xdc, 0xa8, 0xdc, 0x94, 0xfd, 0xb1, 0x0f, 0x02, 0xbe, 0x2a, 0xbb, 0xf4,
	0x86, 0x3c, 0xba, 0xc5, 0xf4, 0x64, 0x72, 0x36, 0xcf, 0x22, 0x10, 0x3f, 0x61, 0xef, 0x4b, 0xad,
	0x24, 0x7f, 0x01, 0x07, 0x95, 0xc1, 0xdc, 0x4a, 0x17, 0x72, 0x93, 0x6c, 0xbf, 0x32, 0x78, 0x2d,
	0x5d, 0x10, 0x64, 0x1b, 0x84, 0x69, 0x2f, 0xc8, 0x76, 0x10, 0xba, 0x04, 0xea, 0xc5, 0x6c, 0x4c,
	0xa0, 0xde, 0x26, 0x50, 0x2f, 0xf6, 0xc6, 0x04, 0x6a, 0xb1, 0x06, 0xb8, 0x24, 0xaa, 0x4b, 0x83,
	0xd2, 0x29, 0x7e, 0x0c, 0xcc, 0x12, 0x6a, 0xe3, 0x7c, 0xa9, 0xfa, 0x99, 0x23, 0xc1, 0x5f, 0x41,
	0x62, 0xa5, 0x8b, 0x62, 0x9c, 0xbb, 0xc5, 0xe2, 0x37, 0xa4, 0x5f, 0xef, 0x09, 0xd5, 0x8d, 0xaf,
	0x0a, 0x55, 0xf3, 0x25, 0x1c, 0x86, 0xbb, 0xd4, 0x8f, 0xf9, 0x86, 0xfa, 0xae, 0x79, 0x96, 0xf6,
	0xdc, 0x25, 0x95, 0x8a, 0xbf, 0x87, 0x67, 0x28, 0x9d, 0x21, 0x94, 0x36, 0xc7, 0x90, 0x0a, 0xa5,
	0xb3, 0xec, 0x68, 0xa0, 0xfb, 0xae, 0x63, 0x60, 0xaa, 0x75, 0x0a, 0x1b, 0x43, 0x18, 0xae, 0xc5,
	0xb2, 0x91, 0x10, 0x1a, 0x92, 0xb5, 0xb7, 0xf6, 0x46, 0x56, 0x8a, 0xbf, 0x01, 0xb8, 0x33, 0x75,
	0xe3, 0x72, 0x94, 0xd5, 0xf0, 0xd7, 0x58, 0x60, 0x82, 0xfc, 0x16, 0xd2, 0xca, 0x94, 0xa5, 0x55,
	0x51, 0x9f, 0x06, 0x1d, 0x22, 0x15, 0x0c, 0xaf, 0x81, 0x59, 0x39, 0xc4, 0xe3, 0xa4, 0xa4, 0x23,
	0x3a, 0xf1, 0xe2, 0xf4, 0x87, 0xd0, 0xc6, 0xdd, 0xfb, 0xe2, 0x7c, 0x43, 0xd5, 0x2a, 0xec, 0xfe,
	0x4a, 0x53, 0xbb, 0x1a, 0xf6, 0x7f, 0xd5, 0x3d, 0x8d, 0x62, 0x3f, 0xc0, 0x0f, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb6, 0x52, 0x6c, 0xf9, 0x29, 0x03, 0x00, 0x00,
}
