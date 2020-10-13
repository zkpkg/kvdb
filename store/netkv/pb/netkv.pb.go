// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netkv.proto

package pbnetkv

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ReadOptions struct {
	KeyOnly              bool     `protobuf:"varint,1,opt,name=key_only,json=keyOnly,proto3" json:"key_only,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReadOptions) Reset()         { *m = ReadOptions{} }
func (m *ReadOptions) String() string { return proto.CompactTextString(m) }
func (*ReadOptions) ProtoMessage()    {}
func (*ReadOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{0}
}

func (m *ReadOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReadOptions.Unmarshal(m, b)
}
func (m *ReadOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReadOptions.Marshal(b, m, deterministic)
}
func (m *ReadOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReadOptions.Merge(m, src)
}
func (m *ReadOptions) XXX_Size() int {
	return xxx_messageInfo_ReadOptions.Size(m)
}
func (m *ReadOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ReadOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ReadOptions proto.InternalMessageInfo

func (m *ReadOptions) GetKeyOnly() bool {
	if m != nil {
		return m.KeyOnly
	}
	return false
}

type KeyValue struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                []byte   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValue) Reset()         { *m = KeyValue{} }
func (m *KeyValue) String() string { return proto.CompactTextString(m) }
func (*KeyValue) ProtoMessage()    {}
func (*KeyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{1}
}

func (m *KeyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValue.Unmarshal(m, b)
}
func (m *KeyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValue.Marshal(b, m, deterministic)
}
func (m *KeyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValue.Merge(m, src)
}
func (m *KeyValue) XXX_Size() int {
	return xxx_messageInfo_KeyValue.Size(m)
}
func (m *KeyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValue.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValue proto.InternalMessageInfo

func (m *KeyValue) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *KeyValue) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type KeyValues struct {
	Kvs                  []*KeyValue `protobuf:"bytes,1,rep,name=kvs,proto3" json:"kvs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *KeyValues) Reset()         { *m = KeyValues{} }
func (m *KeyValues) String() string { return proto.CompactTextString(m) }
func (*KeyValues) ProtoMessage()    {}
func (*KeyValues) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{2}
}

func (m *KeyValues) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValues.Unmarshal(m, b)
}
func (m *KeyValues) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValues.Marshal(b, m, deterministic)
}
func (m *KeyValues) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValues.Merge(m, src)
}
func (m *KeyValues) XXX_Size() int {
	return xxx_messageInfo_KeyValues.Size(m)
}
func (m *KeyValues) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValues.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValues proto.InternalMessageInfo

func (m *KeyValues) GetKvs() []*KeyValue {
	if m != nil {
		return m.Kvs
	}
	return nil
}

type Keys struct {
	Keys                 [][]byte `protobuf:"bytes,1,rep,name=keys,proto3" json:"keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Keys) Reset()         { *m = Keys{} }
func (m *Keys) String() string { return proto.CompactTextString(m) }
func (*Keys) ProtoMessage()    {}
func (*Keys) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{3}
}

func (m *Keys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Keys.Unmarshal(m, b)
}
func (m *Keys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Keys.Marshal(b, m, deterministic)
}
func (m *Keys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Keys.Merge(m, src)
}
func (m *Keys) XXX_Size() int {
	return xxx_messageInfo_Keys.Size(m)
}
func (m *Keys) XXX_DiscardUnknown() {
	xxx_messageInfo_Keys.DiscardUnknown(m)
}

var xxx_messageInfo_Keys proto.InternalMessageInfo

func (m *Keys) GetKeys() [][]byte {
	if m != nil {
		return m.Keys
	}
	return nil
}

type Values struct {
	Values               [][]byte `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Values) Reset()         { *m = Values{} }
func (m *Values) String() string { return proto.CompactTextString(m) }
func (*Values) ProtoMessage()    {}
func (*Values) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{4}
}

func (m *Values) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Values.Unmarshal(m, b)
}
func (m *Values) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Values.Marshal(b, m, deterministic)
}
func (m *Values) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Values.Merge(m, src)
}
func (m *Values) XXX_Size() int {
	return xxx_messageInfo_Values.Size(m)
}
func (m *Values) XXX_DiscardUnknown() {
	xxx_messageInfo_Values.DiscardUnknown(m)
}

var xxx_messageInfo_Values proto.InternalMessageInfo

func (m *Values) GetValues() [][]byte {
	if m != nil {
		return m.Values
	}
	return nil
}

type ScanRequest struct {
	Start                []byte       `protobuf:"bytes,1,opt,name=start,proto3" json:"start,omitempty"`
	ExclusiveEnd         []byte       `protobuf:"bytes,2,opt,name=exclusive_end,json=exclusiveEnd,proto3" json:"exclusive_end,omitempty"`
	Limit                uint64       `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Options              *ReadOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *ScanRequest) Reset()         { *m = ScanRequest{} }
func (m *ScanRequest) String() string { return proto.CompactTextString(m) }
func (*ScanRequest) ProtoMessage()    {}
func (*ScanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{5}
}

func (m *ScanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScanRequest.Unmarshal(m, b)
}
func (m *ScanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScanRequest.Marshal(b, m, deterministic)
}
func (m *ScanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScanRequest.Merge(m, src)
}
func (m *ScanRequest) XXX_Size() int {
	return xxx_messageInfo_ScanRequest.Size(m)
}
func (m *ScanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ScanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ScanRequest proto.InternalMessageInfo

func (m *ScanRequest) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *ScanRequest) GetExclusiveEnd() []byte {
	if m != nil {
		return m.ExclusiveEnd
	}
	return nil
}

func (m *ScanRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ScanRequest) GetOptions() *ReadOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type BatchPrefixRequest struct {
	Prefixes             [][]byte     `protobuf:"bytes,1,rep,name=prefixes,proto3" json:"prefixes,omitempty"`
	LimitPerPrefix       uint64       `protobuf:"varint,2,opt,name=limit_per_prefix,json=limitPerPrefix,proto3" json:"limit_per_prefix,omitempty"`
	Options              *ReadOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *BatchPrefixRequest) Reset()         { *m = BatchPrefixRequest{} }
func (m *BatchPrefixRequest) String() string { return proto.CompactTextString(m) }
func (*BatchPrefixRequest) ProtoMessage()    {}
func (*BatchPrefixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{6}
}

func (m *BatchPrefixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchPrefixRequest.Unmarshal(m, b)
}
func (m *BatchPrefixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchPrefixRequest.Marshal(b, m, deterministic)
}
func (m *BatchPrefixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchPrefixRequest.Merge(m, src)
}
func (m *BatchPrefixRequest) XXX_Size() int {
	return xxx_messageInfo_BatchPrefixRequest.Size(m)
}
func (m *BatchPrefixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchPrefixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchPrefixRequest proto.InternalMessageInfo

func (m *BatchPrefixRequest) GetPrefixes() [][]byte {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

func (m *BatchPrefixRequest) GetLimitPerPrefix() uint64 {
	if m != nil {
		return m.LimitPerPrefix
	}
	return 0
}

func (m *BatchPrefixRequest) GetOptions() *ReadOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type BatchScanRequest struct {
	Start                [][]byte `protobuf:"bytes,1,rep,name=start,proto3" json:"start,omitempty"`
	ExclusiveEnd         [][]byte `protobuf:"bytes,2,rep,name=exclusive_end,json=exclusiveEnd,proto3" json:"exclusive_end,omitempty"`
	LimitPerScan         uint64   `protobuf:"varint,3,opt,name=limit_per_scan,json=limitPerScan,proto3" json:"limit_per_scan,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BatchScanRequest) Reset()         { *m = BatchScanRequest{} }
func (m *BatchScanRequest) String() string { return proto.CompactTextString(m) }
func (*BatchScanRequest) ProtoMessage()    {}
func (*BatchScanRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{7}
}

func (m *BatchScanRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BatchScanRequest.Unmarshal(m, b)
}
func (m *BatchScanRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BatchScanRequest.Marshal(b, m, deterministic)
}
func (m *BatchScanRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BatchScanRequest.Merge(m, src)
}
func (m *BatchScanRequest) XXX_Size() int {
	return xxx_messageInfo_BatchScanRequest.Size(m)
}
func (m *BatchScanRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BatchScanRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BatchScanRequest proto.InternalMessageInfo

func (m *BatchScanRequest) GetStart() [][]byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *BatchScanRequest) GetExclusiveEnd() [][]byte {
	if m != nil {
		return m.ExclusiveEnd
	}
	return nil
}

func (m *BatchScanRequest) GetLimitPerScan() uint64 {
	if m != nil {
		return m.LimitPerScan
	}
	return 0
}

type PrefixRequest struct {
	Prefix               []byte       `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	Limit                uint64       `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Options              *ReadOptions `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *PrefixRequest) Reset()         { *m = PrefixRequest{} }
func (m *PrefixRequest) String() string { return proto.CompactTextString(m) }
func (*PrefixRequest) ProtoMessage()    {}
func (*PrefixRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{8}
}

func (m *PrefixRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrefixRequest.Unmarshal(m, b)
}
func (m *PrefixRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrefixRequest.Marshal(b, m, deterministic)
}
func (m *PrefixRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrefixRequest.Merge(m, src)
}
func (m *PrefixRequest) XXX_Size() int {
	return xxx_messageInfo_PrefixRequest.Size(m)
}
func (m *PrefixRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrefixRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrefixRequest proto.InternalMessageInfo

func (m *PrefixRequest) GetPrefix() []byte {
	if m != nil {
		return m.Prefix
	}
	return nil
}

func (m *PrefixRequest) GetLimit() uint64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *PrefixRequest) GetOptions() *ReadOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

type EmptyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmptyResponse) Reset()         { *m = EmptyResponse{} }
func (m *EmptyResponse) String() string { return proto.CompactTextString(m) }
func (*EmptyResponse) ProtoMessage()    {}
func (*EmptyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_25aabd6fb5784ada, []int{9}
}

func (m *EmptyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmptyResponse.Unmarshal(m, b)
}
func (m *EmptyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmptyResponse.Marshal(b, m, deterministic)
}
func (m *EmptyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmptyResponse.Merge(m, src)
}
func (m *EmptyResponse) XXX_Size() int {
	return xxx_messageInfo_EmptyResponse.Size(m)
}
func (m *EmptyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmptyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmptyResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReadOptions)(nil), "dfuse.netkv.v1.ReadOptions")
	proto.RegisterType((*KeyValue)(nil), "dfuse.netkv.v1.KeyValue")
	proto.RegisterType((*KeyValues)(nil), "dfuse.netkv.v1.KeyValues")
	proto.RegisterType((*Keys)(nil), "dfuse.netkv.v1.Keys")
	proto.RegisterType((*Values)(nil), "dfuse.netkv.v1.Values")
	proto.RegisterType((*ScanRequest)(nil), "dfuse.netkv.v1.ScanRequest")
	proto.RegisterType((*BatchPrefixRequest)(nil), "dfuse.netkv.v1.BatchPrefixRequest")
	proto.RegisterType((*BatchScanRequest)(nil), "dfuse.netkv.v1.BatchScanRequest")
	proto.RegisterType((*PrefixRequest)(nil), "dfuse.netkv.v1.PrefixRequest")
	proto.RegisterType((*EmptyResponse)(nil), "dfuse.netkv.v1.EmptyResponse")
}

func init() { proto.RegisterFile("netkv.proto", fileDescriptor_25aabd6fb5784ada) }

var fileDescriptor_25aabd6fb5784ada = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x4d, 0x6b, 0xdb, 0x40,
	0x10, 0x45, 0x91, 0x22, 0xcb, 0x23, 0xdb, 0x35, 0x4b, 0x08, 0x8a, 0x4b, 0x41, 0xa8, 0x3d, 0x88,
	0x1e, 0x4c, 0xeb, 0x52, 0x7a, 0x29, 0x14, 0xdc, 0x84, 0x52, 0x4c, 0x9b, 0xb0, 0x85, 0x1c, 0x7a,
	0x31, 0x8a, 0x3d, 0xa1, 0x42, 0x8a, 0xa4, 0x6a, 0x57, 0x22, 0xfa, 0x19, 0x3d, 0xf4, 0x77, 0xf6,
	0x2f, 0x14, 0xed, 0xae, 0x1c, 0x7f, 0x44, 0x6e, 0x7d, 0xd3, 0xcc, 0xbc, 0x7d, 0xfb, 0xde, 0x9b,
	0x45, 0x60, 0x27, 0xc8, 0xa3, 0x72, 0x9c, 0xe5, 0x29, 0x4f, 0xc9, 0x60, 0x79, 0x5b, 0x30, 0x1c,
	0xcb, 0x56, 0xf9, 0xda, 0xf3, 0xc1, 0xa6, 0x18, 0x2c, 0x2f, 0x33, 0x1e, 0xa6, 0x09, 0x23, 0x67,
	0x60, 0x45, 0x58, 0xcd, 0xd3, 0x24, 0xae, 0x1c, 0xcd, 0xd5, 0x7c, 0x8b, 0x76, 0x22, 0xac, 0x2e,
	0x93, 0xb8, 0xf2, 0x26, 0x60, 0xcd, 0xb0, 0xba, 0x0e, 0xe2, 0x02, 0xc9, 0x10, 0xf4, 0x08, 0x25,
	0xa2, 0x47, 0xeb, 0x4f, 0x72, 0x02, 0xc7, 0x65, 0x3d, 0x72, 0x8e, 0x44, 0x4f, 0x16, 0xde, 0x3b,
	0xe8, 0x36, 0x67, 0x18, 0x79, 0x09, 0x7a, 0x54, 0x32, 0x47, 0x73, 0x75, 0xdf, 0x9e, 0x38, 0xe3,
	0x4d, 0x21, 0xe3, 0x06, 0x47, 0x6b, 0x90, 0x37, 0x02, 0x63, 0x86, 0x15, 0x23, 0x04, 0x8c, 0x08,
	0x2b, 0x79, 0xa8, 0x47, 0xc5, 0xb7, 0xe7, 0x82, 0xa9, 0x18, 0x4f, 0xc1, 0x14, 0xf7, 0x34, 0x73,
	0x55, 0x79, 0xbf, 0x35, 0xb0, 0xbf, 0x2d, 0x82, 0x84, 0xe2, 0xcf, 0x02, 0x19, 0xaf, 0xc5, 0x31,
	0x1e, 0xe4, 0x5c, 0x09, 0x96, 0x05, 0x79, 0x0e, 0x7d, 0xbc, 0x5f, 0xc4, 0x05, 0x0b, 0x4b, 0x9c,
	0x63, 0xb2, 0x54, 0xd2, 0x7b, 0xab, 0xe6, 0x45, 0xb2, 0xac, 0x8f, 0xc6, 0xe1, 0x5d, 0xc8, 0x1d,
	0xdd, 0xd5, 0x7c, 0x83, 0xca, 0x82, 0xbc, 0x85, 0x4e, 0x2a, 0x13, 0x73, 0x0c, 0x57, 0xf3, 0xed,
	0xc9, 0xd3, 0x6d, 0x3b, 0x6b, 0xa1, 0xd2, 0x06, 0xeb, 0xfd, 0xd2, 0x80, 0x4c, 0x03, 0xbe, 0xf8,
	0x71, 0x95, 0xe3, 0x6d, 0x78, 0xdf, 0xc8, 0x1b, 0x81, 0x95, 0x89, 0xc6, 0xca, 0xc8, 0xaa, 0x26,
	0x3e, 0x0c, 0xc5, 0x95, 0xf3, 0x0c, 0xf3, 0xb9, 0xec, 0x0a, 0x9d, 0x06, 0x1d, 0x88, 0xfe, 0x15,
	0xe6, 0x92, 0x6c, 0x5d, 0x93, 0x7e, 0x80, 0x26, 0x06, 0x43, 0x21, 0xa9, 0x25, 0x2f, 0x7d, 0x6f,
	0x5e, 0xfa, 0x4e, 0x5e, 0x2f, 0x60, 0xf0, 0xa0, 0x97, 0x2d, 0x82, 0x44, 0x05, 0xd7, 0x6b, 0xd4,
	0xd6, 0xf7, 0x78, 0x1c, 0xfa, 0x9b, 0x11, 0x9c, 0x82, 0xa9, 0xcc, 0xc9, 0x15, 0xa9, 0xea, 0x21,
	0xfe, 0xa3, 0x96, 0xf8, 0x0f, 0xb1, 0xfa, 0x04, 0xfa, 0x17, 0x77, 0x19, 0xaf, 0x28, 0xb2, 0x2c,
	0x4d, 0x18, 0x4e, 0xfe, 0xe8, 0x70, 0xfc, 0x15, 0xf9, 0xec, 0x9a, 0x9c, 0x83, 0x25, 0x17, 0x53,
	0x70, 0x72, 0xd6, 0xf6, 0x34, 0xd9, 0xe8, 0xd9, 0xf6, 0x68, 0x83, 0x8f, 0xbc, 0x57, 0x2c, 0x9f,
	0x90, 0x93, 0x93, 0x47, 0x58, 0xd8, 0xa8, 0xf5, 0xd9, 0xbf, 0xd2, 0xc8, 0x07, 0x30, 0xea, 0x70,
	0xc8, 0x8e, 0x99, 0xb5, 0xd5, 0xec, 0x25, 0xf8, 0x0c, 0xdd, 0xd5, 0x2a, 0x89, 0xbb, 0x0d, 0xdc,
	0xde, 0xf2, 0x5e, 0xaa, 0x29, 0xd8, 0x02, 0x7f, 0x8e, 0x31, 0x72, 0x6c, 0x31, 0xf3, 0x8f, 0x34,
	0x3e, 0x82, 0xa9, 0x9e, 0xe6, 0x0e, 0x70, 0x63, 0xf9, 0x7b, 0x85, 0x7c, 0x51, 0x42, 0x14, 0x93,
	0xf7, 0xa8, 0xab, 0xff, 0xa6, 0x9b, 0x76, 0xbf, 0x77, 0xb2, 0x1b, 0x31, 0xb8, 0x31, 0xc5, 0x0f,
	0xf1, 0xcd, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20, 0x4f, 0x5f, 0x52, 0x1f, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// NetKVClient is the client API for NetKV service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NetKVClient interface {
	BatchPut(ctx context.Context, in *KeyValues, opts ...grpc.CallOption) (*EmptyResponse, error)
	// TODO: we need to be able to get individual responses
	// regarding Not-Foundness of each key, otherwise, we'll have
	// a hard time knowing which key was not found, etc..
	// This will happen more with sparse trxdb.
	BatchGet(ctx context.Context, in *Keys, opts ...grpc.CallOption) (NetKV_BatchGetClient, error)
	Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (NetKV_ScanClient, error)
	BatchScan(ctx context.Context, in *BatchScanRequest, opts ...grpc.CallOption) (NetKV_BatchScanClient, error)
	BatchDelete(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*EmptyResponse, error)
	Prefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (NetKV_PrefixClient, error)
	BatchPrefix(ctx context.Context, in *BatchPrefixRequest, opts ...grpc.CallOption) (NetKV_BatchPrefixClient, error)
}

type netKVClient struct {
	cc *grpc.ClientConn
}

func NewNetKVClient(cc *grpc.ClientConn) NetKVClient {
	return &netKVClient{cc}
}

func (c *netKVClient) BatchPut(ctx context.Context, in *KeyValues, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/dfuse.netkv.v1.NetKV/BatchPut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netKVClient) BatchGet(ctx context.Context, in *Keys, opts ...grpc.CallOption) (NetKV_BatchGetClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetKV_serviceDesc.Streams[0], "/dfuse.netkv.v1.NetKV/BatchGet", opts...)
	if err != nil {
		return nil, err
	}
	x := &netKVBatchGetClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetKV_BatchGetClient interface {
	Recv() (*KeyValue, error)
	grpc.ClientStream
}

type netKVBatchGetClient struct {
	grpc.ClientStream
}

func (x *netKVBatchGetClient) Recv() (*KeyValue, error) {
	m := new(KeyValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *netKVClient) Scan(ctx context.Context, in *ScanRequest, opts ...grpc.CallOption) (NetKV_ScanClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetKV_serviceDesc.Streams[1], "/dfuse.netkv.v1.NetKV/Scan", opts...)
	if err != nil {
		return nil, err
	}
	x := &netKVScanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetKV_ScanClient interface {
	Recv() (*KeyValue, error)
	grpc.ClientStream
}

type netKVScanClient struct {
	grpc.ClientStream
}

func (x *netKVScanClient) Recv() (*KeyValue, error) {
	m := new(KeyValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *netKVClient) BatchScan(ctx context.Context, in *BatchScanRequest, opts ...grpc.CallOption) (NetKV_BatchScanClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetKV_serviceDesc.Streams[2], "/dfuse.netkv.v1.NetKV/BatchScan", opts...)
	if err != nil {
		return nil, err
	}
	x := &netKVBatchScanClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetKV_BatchScanClient interface {
	Recv() (*KeyValue, error)
	grpc.ClientStream
}

type netKVBatchScanClient struct {
	grpc.ClientStream
}

func (x *netKVBatchScanClient) Recv() (*KeyValue, error) {
	m := new(KeyValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *netKVClient) BatchDelete(ctx context.Context, in *Keys, opts ...grpc.CallOption) (*EmptyResponse, error) {
	out := new(EmptyResponse)
	err := c.cc.Invoke(ctx, "/dfuse.netkv.v1.NetKV/BatchDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *netKVClient) Prefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (NetKV_PrefixClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetKV_serviceDesc.Streams[3], "/dfuse.netkv.v1.NetKV/Prefix", opts...)
	if err != nil {
		return nil, err
	}
	x := &netKVPrefixClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetKV_PrefixClient interface {
	Recv() (*KeyValue, error)
	grpc.ClientStream
}

type netKVPrefixClient struct {
	grpc.ClientStream
}

func (x *netKVPrefixClient) Recv() (*KeyValue, error) {
	m := new(KeyValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *netKVClient) BatchPrefix(ctx context.Context, in *BatchPrefixRequest, opts ...grpc.CallOption) (NetKV_BatchPrefixClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NetKV_serviceDesc.Streams[4], "/dfuse.netkv.v1.NetKV/BatchPrefix", opts...)
	if err != nil {
		return nil, err
	}
	x := &netKVBatchPrefixClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type NetKV_BatchPrefixClient interface {
	Recv() (*KeyValue, error)
	grpc.ClientStream
}

type netKVBatchPrefixClient struct {
	grpc.ClientStream
}

func (x *netKVBatchPrefixClient) Recv() (*KeyValue, error) {
	m := new(KeyValue)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// NetKVServer is the server API for NetKV service.
type NetKVServer interface {
	BatchPut(context.Context, *KeyValues) (*EmptyResponse, error)
	// TODO: we need to be able to get individual responses
	// regarding Not-Foundness of each key, otherwise, we'll have
	// a hard time knowing which key was not found, etc..
	// This will happen more with sparse trxdb.
	BatchGet(*Keys, NetKV_BatchGetServer) error
	Scan(*ScanRequest, NetKV_ScanServer) error
	BatchScan(*BatchScanRequest, NetKV_BatchScanServer) error
	BatchDelete(context.Context, *Keys) (*EmptyResponse, error)
	Prefix(*PrefixRequest, NetKV_PrefixServer) error
	BatchPrefix(*BatchPrefixRequest, NetKV_BatchPrefixServer) error
}

// UnimplementedNetKVServer can be embedded to have forward compatible implementations.
type UnimplementedNetKVServer struct {
}

func (*UnimplementedNetKVServer) BatchPut(ctx context.Context, req *KeyValues) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchPut not implemented")
}
func (*UnimplementedNetKVServer) BatchGet(req *Keys, srv NetKV_BatchGetServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchGet not implemented")
}
func (*UnimplementedNetKVServer) Scan(req *ScanRequest, srv NetKV_ScanServer) error {
	return status.Errorf(codes.Unimplemented, "method Scan not implemented")
}
func (*UnimplementedNetKVServer) BatchScan(req *BatchScanRequest, srv NetKV_BatchScanServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchScan not implemented")
}
func (*UnimplementedNetKVServer) BatchDelete(ctx context.Context, req *Keys) (*EmptyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchDelete not implemented")
}
func (*UnimplementedNetKVServer) Prefix(req *PrefixRequest, srv NetKV_PrefixServer) error {
	return status.Errorf(codes.Unimplemented, "method Prefix not implemented")
}
func (*UnimplementedNetKVServer) BatchPrefix(req *BatchPrefixRequest, srv NetKV_BatchPrefixServer) error {
	return status.Errorf(codes.Unimplemented, "method BatchPrefix not implemented")
}

func RegisterNetKVServer(s *grpc.Server, srv NetKVServer) {
	s.RegisterService(&_NetKV_serviceDesc, srv)
}

func _NetKV_BatchPut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyValues)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetKVServer).BatchPut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.netkv.v1.NetKV/BatchPut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetKVServer).BatchPut(ctx, req.(*KeyValues))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetKV_BatchGet_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Keys)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetKVServer).BatchGet(m, &netKVBatchGetServer{stream})
}

type NetKV_BatchGetServer interface {
	Send(*KeyValue) error
	grpc.ServerStream
}

type netKVBatchGetServer struct {
	grpc.ServerStream
}

func (x *netKVBatchGetServer) Send(m *KeyValue) error {
	return x.ServerStream.SendMsg(m)
}

func _NetKV_Scan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ScanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetKVServer).Scan(m, &netKVScanServer{stream})
}

type NetKV_ScanServer interface {
	Send(*KeyValue) error
	grpc.ServerStream
}

type netKVScanServer struct {
	grpc.ServerStream
}

func (x *netKVScanServer) Send(m *KeyValue) error {
	return x.ServerStream.SendMsg(m)
}

func _NetKV_BatchScan_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BatchScanRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetKVServer).BatchScan(m, &netKVBatchScanServer{stream})
}

type NetKV_BatchScanServer interface {
	Send(*KeyValue) error
	grpc.ServerStream
}

type netKVBatchScanServer struct {
	grpc.ServerStream
}

func (x *netKVBatchScanServer) Send(m *KeyValue) error {
	return x.ServerStream.SendMsg(m)
}

func _NetKV_BatchDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Keys)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetKVServer).BatchDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dfuse.netkv.v1.NetKV/BatchDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetKVServer).BatchDelete(ctx, req.(*Keys))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetKV_Prefix_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrefixRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetKVServer).Prefix(m, &netKVPrefixServer{stream})
}

type NetKV_PrefixServer interface {
	Send(*KeyValue) error
	grpc.ServerStream
}

type netKVPrefixServer struct {
	grpc.ServerStream
}

func (x *netKVPrefixServer) Send(m *KeyValue) error {
	return x.ServerStream.SendMsg(m)
}

func _NetKV_BatchPrefix_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BatchPrefixRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(NetKVServer).BatchPrefix(m, &netKVBatchPrefixServer{stream})
}

type NetKV_BatchPrefixServer interface {
	Send(*KeyValue) error
	grpc.ServerStream
}

type netKVBatchPrefixServer struct {
	grpc.ServerStream
}

func (x *netKVBatchPrefixServer) Send(m *KeyValue) error {
	return x.ServerStream.SendMsg(m)
}

var _NetKV_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dfuse.netkv.v1.NetKV",
	HandlerType: (*NetKVServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BatchPut",
			Handler:    _NetKV_BatchPut_Handler,
		},
		{
			MethodName: "BatchDelete",
			Handler:    _NetKV_BatchDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BatchGet",
			Handler:       _NetKV_BatchGet_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Scan",
			Handler:       _NetKV_Scan_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BatchScan",
			Handler:       _NetKV_BatchScan_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Prefix",
			Handler:       _NetKV_Prefix_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "BatchPrefix",
			Handler:       _NetKV_BatchPrefix_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "netkv.proto",
}
