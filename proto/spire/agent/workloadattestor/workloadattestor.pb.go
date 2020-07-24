// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/agent/workloadattestor/workloadattestor.proto

package workloadattestor

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	common "github.com/spiffe/spire/proto/spire/common"
	plugin "github.com/spiffe/spire/proto/spire/common/plugin"
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

//* Represents the workload PID.
type AttestRequest struct {
	//* Workload PID
	Pid                  int32    `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestRequest) Reset()         { *m = AttestRequest{} }
func (m *AttestRequest) String() string { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()    {}
func (*AttestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_410d8a5728772cda, []int{0}
}

func (m *AttestRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestRequest.Unmarshal(m, b)
}
func (m *AttestRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestRequest.Marshal(b, m, deterministic)
}
func (m *AttestRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestRequest.Merge(m, src)
}
func (m *AttestRequest) XXX_Size() int {
	return xxx_messageInfo_AttestRequest.Size(m)
}
func (m *AttestRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AttestRequest proto.InternalMessageInfo

func (m *AttestRequest) GetPid() int32 {
	if m != nil {
		return m.Pid
	}
	return 0
}

//* Represents a list of selectors resolved for a given PID.
type AttestResponse struct {
	//* List of selectors
	Selectors            []*common.Selector `protobuf:"bytes,1,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AttestResponse) Reset()         { *m = AttestResponse{} }
func (m *AttestResponse) String() string { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()    {}
func (*AttestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_410d8a5728772cda, []int{1}
}

func (m *AttestResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttestResponse.Unmarshal(m, b)
}
func (m *AttestResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttestResponse.Marshal(b, m, deterministic)
}
func (m *AttestResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttestResponse.Merge(m, src)
}
func (m *AttestResponse) XXX_Size() int {
	return xxx_messageInfo_AttestResponse.Size(m)
}
func (m *AttestResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AttestResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AttestResponse proto.InternalMessageInfo

func (m *AttestResponse) GetSelectors() []*common.Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "spire.agent.workloadattestor.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "spire.agent.workloadattestor.AttestResponse")
}

func init() {
	proto.RegisterFile("spire/agent/workloadattestor/workloadattestor.proto", fileDescriptor_410d8a5728772cda)
}

var fileDescriptor_410d8a5728772cda = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xe9, 0x77, 0x7c, 0x07, 0x8b, 0x4c, 0x46, 0x0e, 0x32, 0x8b, 0x87, 0x3a, 0x50, 0xe6,
	0x0f, 0x52, 0xd8, 0x3c, 0x89, 0x97, 0x29, 0x28, 0xde, 0xa4, 0x1e, 0x84, 0xdd, 0xba, 0xee, 0x4d,
	0x0d, 0x76, 0x79, 0x63, 0x92, 0xe2, 0x1f, 0xe7, 0x3f, 0x27, 0x4b, 0x52, 0xa5, 0x2a, 0x73, 0xa7,
	0x84, 0x3e, 0x9f, 0xe7, 0x7d, 0x9e, 0xf6, 0x2d, 0x99, 0x1a, 0x25, 0x34, 0xa4, 0x79, 0x09, 0xd2,
	0xa6, 0x6f, 0xa8, 0x5f, 0x2a, 0xcc, 0x97, 0xb9, 0xb5, 0x60, 0x2c, 0xea, 0x1f, 0x0f, 0x98, 0xd2,
	0x68, 0x91, 0x1e, 0x38, 0x13, 0x73, 0x26, 0xf6, 0x9d, 0x89, 0xf7, 0xfd, 0xc8, 0x02, 0x57, 0x2b,
	0x94, 0xe1, 0xf0, 0xc6, 0x38, 0x69, 0x49, 0xaa, 0xaa, 0x4b, 0xd1, 0x1c, 0x9e, 0x18, 0x1d, 0x92,
	0xfe, 0xcc, 0x0d, 0xca, 0xe0, 0xb5, 0x06, 0x63, 0xe9, 0x80, 0x74, 0x94, 0x58, 0x0e, 0xa3, 0x24,
	0x1a, 0xff, 0xcf, 0xd6, 0xd7, 0xd1, 0x2d, 0xd9, 0x6d, 0x10, 0xa3, 0x50, 0x1a, 0xa0, 0x17, 0xa4,
	0x67, 0xa0, 0x82, 0xc2, 0xa2, 0x36, 0xc3, 0x28, 0xe9, 0x8c, 0x77, 0x26, 0x7b, 0xcc, 0x77, 0x0c,
	0xf1, 0x8f, 0x41, 0xce, 0xbe, 0xc0, 0xc9, 0xfb, 0x3f, 0x32, 0x78, 0x0a, 0xe5, 0x67, 0xa1, 0x3c,
	0x2d, 0x48, 0xd7, 0xdf, 0xe9, 0x19, 0xdb, 0xf4, 0x96, 0xac, 0xd5, 0x32, 0x3e, 0xdf, 0x0e, 0x0e,
	0x7d, 0xe7, 0xa4, 0x77, 0x83, 0x92, 0x8b, 0xb2, 0xd6, 0x40, 0x8f, 0xda, 0x4d, 0xc3, 0xd7, 0xf8,
	0xd4, 0x9b, 0x84, 0xe3, 0xbf, 0xb0, 0x30, 0x9b, 0x93, 0xfe, 0x1d, 0xd8, 0x07, 0x27, 0xdf, 0x4b,
	0x8e, 0xf4, 0xe4, 0x57, 0x63, 0x8b, 0x69, 0x32, 0x4e, 0xb7, 0x41, 0x7d, 0xce, 0xf5, 0xd5, 0xfc,
	0xb2, 0x14, 0xf6, 0xb9, 0x5e, 0xac, 0xe9, 0xd4, 0x28, 0xc1, 0x39, 0xa4, 0x7e, 0xbd, 0x6e, 0x93,
	0xe9, 0xa6, 0x1f, 0x6b, 0xd1, 0x75, 0xcc, 0xf4, 0x23, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x03, 0x35,
	0x95, 0x7f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// WorkloadAttestorClient is the client API for WorkloadAttestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkloadAttestorClient interface {
	//* Returns a list of selectors resolved for a given PID
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error)
	//* Applies the plugin configuration and returns configuration errors
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	//* Returns the version and related metadata of the plugin
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type workloadAttestorClient struct {
	cc grpc.ClientConnInterface
}

func NewWorkloadAttestorClient(cc grpc.ClientConnInterface) WorkloadAttestorClient {
	return &workloadAttestorClient{cc}
}

func (c *workloadAttestorClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestResponse, error) {
	out := new(AttestResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.workloadattestor.WorkloadAttestor/Attest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.workloadattestor.WorkloadAttestor/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *workloadAttestorClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.workloadattestor.WorkloadAttestor/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkloadAttestorServer is the server API for WorkloadAttestor service.
type WorkloadAttestorServer interface {
	//* Returns a list of selectors resolved for a given PID
	Attest(context.Context, *AttestRequest) (*AttestResponse, error)
	//* Applies the plugin configuration and returns configuration errors
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	//* Returns the version and related metadata of the plugin
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

// UnimplementedWorkloadAttestorServer can be embedded to have forward compatible implementations.
type UnimplementedWorkloadAttestorServer struct {
}

func (*UnimplementedWorkloadAttestorServer) Attest(ctx context.Context, req *AttestRequest) (*AttestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attest not implemented")
}
func (*UnimplementedWorkloadAttestorServer) Configure(ctx context.Context, req *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedWorkloadAttestorServer) GetPluginInfo(ctx context.Context, req *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}

func RegisterWorkloadAttestorServer(s *grpc.Server, srv WorkloadAttestorServer) {
	s.RegisterService(&_WorkloadAttestor_serviceDesc, srv)
}

func _WorkloadAttestor_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.workloadattestor.WorkloadAttestor/Attest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.workloadattestor.WorkloadAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _WorkloadAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.workloadattestor.WorkloadAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkloadAttestorServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkloadAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.workloadattestor.WorkloadAttestor",
	HandlerType: (*WorkloadAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Attest",
			Handler:    _WorkloadAttestor_Attest_Handler,
		},
		{
			MethodName: "Configure",
			Handler:    _WorkloadAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _WorkloadAttestor_GetPluginInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spire/agent/workloadattestor/workloadattestor.proto",
}
