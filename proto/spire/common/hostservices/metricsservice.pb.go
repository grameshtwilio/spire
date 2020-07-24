// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/common/hostservices/metricsservice.proto

package hostservices

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

type Label struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Label) Reset()         { *m = Label{} }
func (m *Label) String() string { return proto.CompactTextString(m) }
func (*Label) ProtoMessage()    {}
func (*Label) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{0}
}

func (m *Label) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Label.Unmarshal(m, b)
}
func (m *Label) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Label.Marshal(b, m, deterministic)
}
func (m *Label) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Label.Merge(m, src)
}
func (m *Label) XXX_Size() int {
	return xxx_messageInfo_Label.Size(m)
}
func (m *Label) XXX_DiscardUnknown() {
	xxx_messageInfo_Label.DiscardUnknown(m)
}

var xxx_messageInfo_Label proto.InternalMessageInfo

func (m *Label) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Label) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type SetGaugeRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGaugeRequest) Reset()         { *m = SetGaugeRequest{} }
func (m *SetGaugeRequest) String() string { return proto.CompactTextString(m) }
func (*SetGaugeRequest) ProtoMessage()    {}
func (*SetGaugeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{1}
}

func (m *SetGaugeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGaugeRequest.Unmarshal(m, b)
}
func (m *SetGaugeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGaugeRequest.Marshal(b, m, deterministic)
}
func (m *SetGaugeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGaugeRequest.Merge(m, src)
}
func (m *SetGaugeRequest) XXX_Size() int {
	return xxx_messageInfo_SetGaugeRequest.Size(m)
}
func (m *SetGaugeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGaugeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGaugeRequest proto.InternalMessageInfo

func (m *SetGaugeRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SetGaugeRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *SetGaugeRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type SetGaugeResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGaugeResponse) Reset()         { *m = SetGaugeResponse{} }
func (m *SetGaugeResponse) String() string { return proto.CompactTextString(m) }
func (*SetGaugeResponse) ProtoMessage()    {}
func (*SetGaugeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{2}
}

func (m *SetGaugeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGaugeResponse.Unmarshal(m, b)
}
func (m *SetGaugeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGaugeResponse.Marshal(b, m, deterministic)
}
func (m *SetGaugeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGaugeResponse.Merge(m, src)
}
func (m *SetGaugeResponse) XXX_Size() int {
	return xxx_messageInfo_SetGaugeResponse.Size(m)
}
func (m *SetGaugeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGaugeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetGaugeResponse proto.InternalMessageInfo

type EmitKeyRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmitKeyRequest) Reset()         { *m = EmitKeyRequest{} }
func (m *EmitKeyRequest) String() string { return proto.CompactTextString(m) }
func (*EmitKeyRequest) ProtoMessage()    {}
func (*EmitKeyRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{3}
}

func (m *EmitKeyRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmitKeyRequest.Unmarshal(m, b)
}
func (m *EmitKeyRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmitKeyRequest.Marshal(b, m, deterministic)
}
func (m *EmitKeyRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmitKeyRequest.Merge(m, src)
}
func (m *EmitKeyRequest) XXX_Size() int {
	return xxx_messageInfo_EmitKeyRequest.Size(m)
}
func (m *EmitKeyRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EmitKeyRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EmitKeyRequest proto.InternalMessageInfo

func (m *EmitKeyRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *EmitKeyRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type EmitKeyResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EmitKeyResponse) Reset()         { *m = EmitKeyResponse{} }
func (m *EmitKeyResponse) String() string { return proto.CompactTextString(m) }
func (*EmitKeyResponse) ProtoMessage()    {}
func (*EmitKeyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{4}
}

func (m *EmitKeyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EmitKeyResponse.Unmarshal(m, b)
}
func (m *EmitKeyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EmitKeyResponse.Marshal(b, m, deterministic)
}
func (m *EmitKeyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EmitKeyResponse.Merge(m, src)
}
func (m *EmitKeyResponse) XXX_Size() int {
	return xxx_messageInfo_EmitKeyResponse.Size(m)
}
func (m *EmitKeyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EmitKeyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EmitKeyResponse proto.InternalMessageInfo

type IncrCounterRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrCounterRequest) Reset()         { *m = IncrCounterRequest{} }
func (m *IncrCounterRequest) String() string { return proto.CompactTextString(m) }
func (*IncrCounterRequest) ProtoMessage()    {}
func (*IncrCounterRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{5}
}

func (m *IncrCounterRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrCounterRequest.Unmarshal(m, b)
}
func (m *IncrCounterRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrCounterRequest.Marshal(b, m, deterministic)
}
func (m *IncrCounterRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrCounterRequest.Merge(m, src)
}
func (m *IncrCounterRequest) XXX_Size() int {
	return xxx_messageInfo_IncrCounterRequest.Size(m)
}
func (m *IncrCounterRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrCounterRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IncrCounterRequest proto.InternalMessageInfo

func (m *IncrCounterRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *IncrCounterRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *IncrCounterRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type IncrCounterResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IncrCounterResponse) Reset()         { *m = IncrCounterResponse{} }
func (m *IncrCounterResponse) String() string { return proto.CompactTextString(m) }
func (*IncrCounterResponse) ProtoMessage()    {}
func (*IncrCounterResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{6}
}

func (m *IncrCounterResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IncrCounterResponse.Unmarshal(m, b)
}
func (m *IncrCounterResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IncrCounterResponse.Marshal(b, m, deterministic)
}
func (m *IncrCounterResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IncrCounterResponse.Merge(m, src)
}
func (m *IncrCounterResponse) XXX_Size() int {
	return xxx_messageInfo_IncrCounterResponse.Size(m)
}
func (m *IncrCounterResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IncrCounterResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IncrCounterResponse proto.InternalMessageInfo

type AddSampleRequest struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	Val                  float32  `protobuf:"fixed32,2,opt,name=val,proto3" json:"val,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSampleRequest) Reset()         { *m = AddSampleRequest{} }
func (m *AddSampleRequest) String() string { return proto.CompactTextString(m) }
func (*AddSampleRequest) ProtoMessage()    {}
func (*AddSampleRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{7}
}

func (m *AddSampleRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSampleRequest.Unmarshal(m, b)
}
func (m *AddSampleRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSampleRequest.Marshal(b, m, deterministic)
}
func (m *AddSampleRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSampleRequest.Merge(m, src)
}
func (m *AddSampleRequest) XXX_Size() int {
	return xxx_messageInfo_AddSampleRequest.Size(m)
}
func (m *AddSampleRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSampleRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddSampleRequest proto.InternalMessageInfo

func (m *AddSampleRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AddSampleRequest) GetVal() float32 {
	if m != nil {
		return m.Val
	}
	return 0
}

func (m *AddSampleRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type AddSampleResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddSampleResponse) Reset()         { *m = AddSampleResponse{} }
func (m *AddSampleResponse) String() string { return proto.CompactTextString(m) }
func (*AddSampleResponse) ProtoMessage()    {}
func (*AddSampleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{8}
}

func (m *AddSampleResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddSampleResponse.Unmarshal(m, b)
}
func (m *AddSampleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddSampleResponse.Marshal(b, m, deterministic)
}
func (m *AddSampleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddSampleResponse.Merge(m, src)
}
func (m *AddSampleResponse) XXX_Size() int {
	return xxx_messageInfo_AddSampleResponse.Size(m)
}
func (m *AddSampleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AddSampleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AddSampleResponse proto.InternalMessageInfo

type MeasureSinceRequest struct {
	Key []string `protobuf:"bytes,1,rep,name=key,proto3" json:"key,omitempty"`
	// Unix time in nanoseconds
	Time                 int64    `protobuf:"varint,2,opt,name=time,proto3" json:"time,omitempty"`
	Labels               []*Label `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeasureSinceRequest) Reset()         { *m = MeasureSinceRequest{} }
func (m *MeasureSinceRequest) String() string { return proto.CompactTextString(m) }
func (*MeasureSinceRequest) ProtoMessage()    {}
func (*MeasureSinceRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{9}
}

func (m *MeasureSinceRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeasureSinceRequest.Unmarshal(m, b)
}
func (m *MeasureSinceRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeasureSinceRequest.Marshal(b, m, deterministic)
}
func (m *MeasureSinceRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasureSinceRequest.Merge(m, src)
}
func (m *MeasureSinceRequest) XXX_Size() int {
	return xxx_messageInfo_MeasureSinceRequest.Size(m)
}
func (m *MeasureSinceRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasureSinceRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MeasureSinceRequest proto.InternalMessageInfo

func (m *MeasureSinceRequest) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *MeasureSinceRequest) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *MeasureSinceRequest) GetLabels() []*Label {
	if m != nil {
		return m.Labels
	}
	return nil
}

type MeasureSinceResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MeasureSinceResponse) Reset()         { *m = MeasureSinceResponse{} }
func (m *MeasureSinceResponse) String() string { return proto.CompactTextString(m) }
func (*MeasureSinceResponse) ProtoMessage()    {}
func (*MeasureSinceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_64d53fc6d4ca8da9, []int{10}
}

func (m *MeasureSinceResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MeasureSinceResponse.Unmarshal(m, b)
}
func (m *MeasureSinceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MeasureSinceResponse.Marshal(b, m, deterministic)
}
func (m *MeasureSinceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MeasureSinceResponse.Merge(m, src)
}
func (m *MeasureSinceResponse) XXX_Size() int {
	return xxx_messageInfo_MeasureSinceResponse.Size(m)
}
func (m *MeasureSinceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MeasureSinceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MeasureSinceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Label)(nil), "spire.common.hostservices.Label")
	proto.RegisterType((*SetGaugeRequest)(nil), "spire.common.hostservices.SetGaugeRequest")
	proto.RegisterType((*SetGaugeResponse)(nil), "spire.common.hostservices.SetGaugeResponse")
	proto.RegisterType((*EmitKeyRequest)(nil), "spire.common.hostservices.EmitKeyRequest")
	proto.RegisterType((*EmitKeyResponse)(nil), "spire.common.hostservices.EmitKeyResponse")
	proto.RegisterType((*IncrCounterRequest)(nil), "spire.common.hostservices.IncrCounterRequest")
	proto.RegisterType((*IncrCounterResponse)(nil), "spire.common.hostservices.IncrCounterResponse")
	proto.RegisterType((*AddSampleRequest)(nil), "spire.common.hostservices.AddSampleRequest")
	proto.RegisterType((*AddSampleResponse)(nil), "spire.common.hostservices.AddSampleResponse")
	proto.RegisterType((*MeasureSinceRequest)(nil), "spire.common.hostservices.MeasureSinceRequest")
	proto.RegisterType((*MeasureSinceResponse)(nil), "spire.common.hostservices.MeasureSinceResponse")
}

func init() {
	proto.RegisterFile("spire/common/hostservices/metricsservice.proto", fileDescriptor_64d53fc6d4ca8da9)
}

var fileDescriptor_64d53fc6d4ca8da9 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x94, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0xd5, 0xa5, 0x1b, 0xf4, 0x0d, 0x6d, 0xdd, 0xeb, 0x40, 0x21, 0xa7, 0x28, 0xa7, 0x32,
	0x20, 0x11, 0x03, 0x09, 0x38, 0x02, 0x42, 0x08, 0xc1, 0x2e, 0xc9, 0x8d, 0x13, 0x69, 0xf6, 0xda,
	0x5a, 0x24, 0x71, 0xb0, 0x9d, 0x4a, 0xfd, 0x67, 0xf8, 0x5b, 0x51, 0x6c, 0x53, 0xd2, 0x42, 0xa3,
	0x20, 0xa4, 0xde, 0x9e, 0xad, 0xef, 0xcb, 0xef, 0x8b, 0xfd, 0xc9, 0x10, 0xca, 0x8a, 0x09, 0x8a,
	0x32, 0x5e, 0x14, 0xbc, 0x8c, 0x96, 0x5c, 0x2a, 0x49, 0x62, 0xc5, 0x32, 0x92, 0x51, 0x41, 0x4a,
	0xb0, 0x4c, 0xda, 0x75, 0x58, 0x09, 0xae, 0x38, 0x3e, 0xd4, 0xfa, 0xd0, 0xe8, 0xc3, 0xb6, 0x3e,
	0x78, 0x06, 0xc7, 0x9f, 0xd3, 0x19, 0xe5, 0x88, 0x30, 0x2c, 0xd3, 0x82, 0xdc, 0x81, 0x3f, 0x98,
	0x8e, 0x62, 0x3d, 0xe3, 0x25, 0x1c, 0xaf, 0xd2, 0xbc, 0x26, 0xf7, 0x48, 0x6f, 0x9a, 0x45, 0xc0,
	0xe1, 0x3c, 0x21, 0xf5, 0x21, 0xad, 0x17, 0x14, 0xd3, 0xf7, 0x9a, 0xa4, 0xc2, 0x31, 0x38, 0xdf,
	0x68, 0xed, 0x0e, 0x7c, 0x67, 0x3a, 0x8a, 0x9b, 0xb1, 0xd9, 0x59, 0xa5, 0xb9, 0x36, 0x1e, 0xc5,
	0xcd, 0x88, 0xaf, 0xe0, 0x24, 0x6f, 0x48, 0xd2, 0x75, 0x7c, 0x67, 0x7a, 0x7a, 0xed, 0x87, 0x7b,
	0x53, 0x85, 0x3a, 0x52, 0x6c, 0xf5, 0x01, 0xc2, 0xf8, 0x37, 0x50, 0x56, 0xbc, 0x94, 0x14, 0xbc,
	0x80, 0xb3, 0xf7, 0x05, 0x53, 0x9f, 0x68, 0xfd, 0x0f, 0x19, 0x82, 0x0b, 0x38, 0xdf, 0xb8, 0xec,
	0x87, 0x04, 0xe0, 0xc7, 0x32, 0x13, 0xef, 0x78, 0x5d, 0x2a, 0x12, 0x87, 0xf9, 0xa1, 0xfb, 0x30,
	0xd9, 0x62, 0xda, 0x28, 0x15, 0x8c, 0xdf, 0xdc, 0xde, 0x26, 0x69, 0x51, 0xe5, 0x07, 0x3a, 0xd9,
	0x09, 0x5c, 0xb4, 0x88, 0x36, 0x46, 0x0d, 0x93, 0x1b, 0x4a, 0x65, 0x2d, 0x28, 0x61, 0x65, 0xd6,
	0x91, 0x04, 0x61, 0xa8, 0x58, 0x61, 0xda, 0xe1, 0xc4, 0x7a, 0xfe, 0x8f, 0x2c, 0x0f, 0xe0, 0x72,
	0x1b, 0x6b, 0xe2, 0x5c, 0xff, 0x18, 0xc2, 0xd9, 0x8d, 0x69, 0x75, 0x62, 0x9c, 0x98, 0xc1, 0xdd,
	0x5f, 0x85, 0xc0, 0xab, 0x0e, 0xc0, 0x4e, 0x4d, 0xbd, 0xc7, 0xbd, 0xb4, 0x86, 0x8b, 0x5f, 0xe1,
	0x8e, 0xed, 0x0a, 0x3e, 0xea, 0xf0, 0x6d, 0xb7, 0xd0, 0xbb, 0xea, 0x23, 0xb5, 0x84, 0x1c, 0x4e,
	0x5b, 0x35, 0xc0, 0xa7, 0x1d, 0xd6, 0x3f, 0x2b, 0xea, 0x85, 0x7d, 0xe5, 0x96, 0x36, 0x87, 0xd1,
	0xe6, 0xae, 0xb1, 0xeb, 0x24, 0x76, 0x3b, 0xe8, 0x3d, 0xe9, 0x27, 0xb6, 0x1c, 0x0e, 0xf7, 0xda,
	0xf7, 0x88, 0x5d, 0x39, 0xff, 0xd2, 0x33, 0x2f, 0xea, 0xad, 0x37, 0xc0, 0xb7, 0xaf, 0xbf, 0xbc,
	0x5c, 0x30, 0xb5, 0xac, 0x67, 0x8d, 0x25, 0x92, 0x15, 0x9b, 0xcf, 0x29, 0x32, 0x2f, 0xa4, 0x7e,
	0xfe, 0xa2, 0xbd, 0xaf, 0xe5, 0xec, 0x44, 0x0b, 0x9e, 0xff, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x88,
	0x31, 0x8f, 0xf0, 0x51, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MetricsServiceClient interface {
	SetGauge(ctx context.Context, in *SetGaugeRequest, opts ...grpc.CallOption) (*SetGaugeResponse, error)
	EmitKey(ctx context.Context, in *EmitKeyRequest, opts ...grpc.CallOption) (*EmitKeyResponse, error)
	IncrCounter(ctx context.Context, in *IncrCounterRequest, opts ...grpc.CallOption) (*IncrCounterResponse, error)
	AddSample(ctx context.Context, in *AddSampleRequest, opts ...grpc.CallOption) (*AddSampleResponse, error)
	MeasureSince(ctx context.Context, in *MeasureSinceRequest, opts ...grpc.CallOption) (*MeasureSinceResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) SetGauge(ctx context.Context, in *SetGaugeRequest, opts ...grpc.CallOption) (*SetGaugeResponse, error) {
	out := new(SetGaugeResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/SetGauge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) EmitKey(ctx context.Context, in *EmitKeyRequest, opts ...grpc.CallOption) (*EmitKeyResponse, error) {
	out := new(EmitKeyResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/EmitKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) IncrCounter(ctx context.Context, in *IncrCounterRequest, opts ...grpc.CallOption) (*IncrCounterResponse, error) {
	out := new(IncrCounterResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/IncrCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) AddSample(ctx context.Context, in *AddSampleRequest, opts ...grpc.CallOption) (*AddSampleResponse, error) {
	out := new(AddSampleResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/AddSample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricsServiceClient) MeasureSince(ctx context.Context, in *MeasureSinceRequest, opts ...grpc.CallOption) (*MeasureSinceResponse, error) {
	out := new(MeasureSinceResponse)
	err := c.cc.Invoke(ctx, "/spire.common.hostservices.MetricsService/MeasureSince", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
type MetricsServiceServer interface {
	SetGauge(context.Context, *SetGaugeRequest) (*SetGaugeResponse, error)
	EmitKey(context.Context, *EmitKeyRequest) (*EmitKeyResponse, error)
	IncrCounter(context.Context, *IncrCounterRequest) (*IncrCounterResponse, error)
	AddSample(context.Context, *AddSampleRequest) (*AddSampleResponse, error)
	MeasureSince(context.Context, *MeasureSinceRequest) (*MeasureSinceResponse, error)
}

// UnimplementedMetricsServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (*UnimplementedMetricsServiceServer) SetGauge(ctx context.Context, req *SetGaugeRequest) (*SetGaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGauge not implemented")
}
func (*UnimplementedMetricsServiceServer) EmitKey(ctx context.Context, req *EmitKeyRequest) (*EmitKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmitKey not implemented")
}
func (*UnimplementedMetricsServiceServer) IncrCounter(ctx context.Context, req *IncrCounterRequest) (*IncrCounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrCounter not implemented")
}
func (*UnimplementedMetricsServiceServer) AddSample(ctx context.Context, req *AddSampleRequest) (*AddSampleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSample not implemented")
}
func (*UnimplementedMetricsServiceServer) MeasureSince(ctx context.Context, req *MeasureSinceRequest) (*MeasureSinceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MeasureSince not implemented")
}

func RegisterMetricsServiceServer(s *grpc.Server, srv MetricsServiceServer) {
	s.RegisterService(&_MetricsService_serviceDesc, srv)
}

func _MetricsService_SetGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).SetGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/SetGauge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).SetGauge(ctx, req.(*SetGaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_EmitKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmitKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).EmitKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/EmitKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).EmitKey(ctx, req.(*EmitKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_IncrCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrCounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).IncrCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/IncrCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).IncrCounter(ctx, req.(*IncrCounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_AddSample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSampleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).AddSample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/AddSample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).AddSample(ctx, req.(*AddSampleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MetricsService_MeasureSince_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MeasureSinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).MeasureSince(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.common.hostservices.MetricsService/MeasureSince",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).MeasureSince(ctx, req.(*MeasureSinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _MetricsService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.common.hostservices.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetGauge",
			Handler:    _MetricsService_SetGauge_Handler,
		},
		{
			MethodName: "EmitKey",
			Handler:    _MetricsService_EmitKey_Handler,
		},
		{
			MethodName: "IncrCounter",
			Handler:    _MetricsService_IncrCounter_Handler,
		},
		{
			MethodName: "AddSample",
			Handler:    _MetricsService_AddSample_Handler,
		},
		{
			MethodName: "MeasureSince",
			Handler:    _MetricsService_MeasureSince_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/common/hostservices/metricsservice.proto",
}
