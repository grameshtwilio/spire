// Code generated by protoc-gen-go. DO NOT EDIT.
// source: spire/server/nodeattestor/nodeattestor.proto

package nodeattestor

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

//* Represents a request to attest a node.
type AttestRequest struct {
	//* A type which contains attestation data for specific platform.
	AttestationData *common.AttestationData `protobuf:"bytes,1,opt,name=attestation_data,json=attestationData,proto3" json:"attestation_data,omitempty"`
	//* Challenge response
	Response             []byte   `protobuf:"bytes,3,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AttestRequest) Reset()         { *m = AttestRequest{} }
func (m *AttestRequest) String() string { return proto.CompactTextString(m) }
func (*AttestRequest) ProtoMessage()    {}
func (*AttestRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8296ad380b689384, []int{0}
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

func (m *AttestRequest) GetAttestationData() *common.AttestationData {
	if m != nil {
		return m.AttestationData
	}
	return nil
}

func (m *AttestRequest) GetResponse() []byte {
	if m != nil {
		return m.Response
	}
	return nil
}

//* Represents a response when attesting a node.
type AttestResponse struct {
	//* SPIFFE ID of the attested node
	AgentId string `protobuf:"bytes,2,opt,name=agent_id,json=agentId,proto3" json:"agent_id,omitempty"`
	//* Challenge required for attestation
	Challenge []byte `protobuf:"bytes,3,opt,name=challenge,proto3" json:"challenge,omitempty"`
	//* Optional list of selectors
	Selectors            []*common.Selector `protobuf:"bytes,4,rep,name=selectors,proto3" json:"selectors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AttestResponse) Reset()         { *m = AttestResponse{} }
func (m *AttestResponse) String() string { return proto.CompactTextString(m) }
func (*AttestResponse) ProtoMessage()    {}
func (*AttestResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8296ad380b689384, []int{1}
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

func (m *AttestResponse) GetAgentId() string {
	if m != nil {
		return m.AgentId
	}
	return ""
}

func (m *AttestResponse) GetChallenge() []byte {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *AttestResponse) GetSelectors() []*common.Selector {
	if m != nil {
		return m.Selectors
	}
	return nil
}

func init() {
	proto.RegisterType((*AttestRequest)(nil), "spire.agent.nodeattestor.AttestRequest")
	proto.RegisterType((*AttestResponse)(nil), "spire.agent.nodeattestor.AttestResponse")
}

func init() {
	proto.RegisterFile("spire/server/nodeattestor/nodeattestor.proto", fileDescriptor_8296ad380b689384)
}

var fileDescriptor_8296ad380b689384 = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4f, 0x4b, 0xfb, 0x30,
	0x18, 0xc7, 0x49, 0x37, 0xf6, 0x6b, 0xb3, 0xed, 0xe7, 0xc8, 0x41, 0xba, 0xa2, 0x50, 0x06, 0x6a,
	0x15, 0x69, 0x65, 0x0a, 0xe2, 0x71, 0x2a, 0xe8, 0x76, 0x10, 0xa9, 0xb7, 0x5d, 0x46, 0xb6, 0x3e,
	0xed, 0x0a, 0x5d, 0x52, 0x9b, 0xd4, 0x8b, 0x67, 0x5f, 0x91, 0x6f, 0x50, 0x6c, 0xda, 0xcd, 0x82,
	0x63, 0x9e, 0xd2, 0x27, 0xdf, 0xcf, 0xf3, 0x27, 0xdf, 0x3e, 0xf8, 0x5c, 0xa4, 0x71, 0x06, 0x9e,
	0x80, 0xec, 0x0d, 0x32, 0x8f, 0xf1, 0x00, 0xa8, 0x94, 0x20, 0x24, 0xaf, 0x07, 0x6e, 0x9a, 0x71,
	0xc9, 0x89, 0x59, 0xd0, 0x2e, 0x8d, 0x80, 0x49, 0xf7, 0xa7, 0x6e, 0xd9, 0xaa, 0xce, 0x82, 0xaf,
	0x56, 0x9c, 0x79, 0x69, 0x92, 0x47, 0x71, 0x75, 0xa8, 0x5c, 0xab, 0x5f, 0x23, 0xd4, 0xa1, 0xa4,
	0xc1, 0x3b, 0xee, 0x8e, 0x8a, 0x42, 0x3e, 0xbc, 0xe6, 0x20, 0x24, 0x79, 0xc4, 0x3d, 0x55, 0x99,
	0xca, 0x98, 0xb3, 0x59, 0x40, 0x25, 0x35, 0x91, 0x8d, 0x9c, 0xf6, 0xf0, 0xd0, 0x55, 0x23, 0x94,
	0xf9, 0xa3, 0x0d, 0x75, 0x4f, 0x25, 0xf5, 0xf7, 0x68, 0xfd, 0x82, 0x58, 0x58, 0xcf, 0x40, 0xa4,
	0x9c, 0x09, 0x30, 0x1b, 0x36, 0x72, 0x3a, 0xfe, 0x3a, 0x9e, 0x34, 0x75, 0xad, 0xd7, 0x18, 0x7c,
	0x20, 0xfc, 0xbf, 0xea, 0xae, 0x04, 0xd2, 0xc7, 0x7a, 0xf1, 0xc4, 0x59, 0x1c, 0x98, 0x9a, 0x8d,
	0x1c, 0xc3, 0xff, 0x57, 0xc4, 0xe3, 0x80, 0x1c, 0x60, 0x63, 0xb1, 0xa4, 0x49, 0x02, 0x2c, 0xaa,
	0x0a, 0x6e, 0x2e, 0xc8, 0x15, 0x36, 0x04, 0x24, 0xb0, 0x90, 0x3c, 0x13, 0x66, 0xd3, 0x6e, 0x38,
	0xed, 0xe1, 0x7e, 0x7d, 0xe0, 0x97, 0x52, 0xf6, 0x37, 0xe0, 0xa4, 0xa9, 0xa3, 0x9e, 0x36, 0xfc,
	0xd4, 0x70, 0xe7, 0x89, 0x07, 0x30, 0x2a, 0x2d, 0x25, 0x33, 0xdc, 0x52, 0xdf, 0xe4, 0xc4, 0xdd,
	0xe6, 0xbb, 0x5b, 0xf3, 0xcd, 0x72, 0x76, 0x83, 0xea, 0x89, 0x0e, 0xba, 0x40, 0x64, 0x8a, 0x8d,
	0x3b, 0xce, 0xc2, 0x38, 0xca, 0x33, 0x20, 0x47, 0xf5, 0x39, 0xcb, 0x5f, 0xb7, 0xd6, 0xab, 0x0e,
	0xc7, 0xbb, 0xb0, 0xd2, 0xc2, 0x10, 0x77, 0x1f, 0x40, 0x3e, 0x17, 0xf2, 0x98, 0x85, 0x9c, 0x9c,
	0xfe, 0x9a, 0x58, 0x63, 0xaa, 0x1e, 0x67, 0x7f, 0x41, 0x55, 0x9f, 0xdb, 0x9b, 0xe9, 0x75, 0x14,
	0xcb, 0x65, 0x3e, 0xff, 0xa6, 0x3d, 0x91, 0xc6, 0x61, 0x08, 0x9e, 0xda, 0xb4, 0x62, 0xb7, 0xbc,
	0xad, 0xfb, 0x3d, 0x6f, 0x15, 0xc0, 0xe5, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3e, 0x21, 0x99,
	0x1c, 0x03, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// NodeAttestorClient is the client API for NodeAttestor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type NodeAttestorClient interface {
	//* Attesta a node.
	Attest(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_AttestClient, error)
	//* Responsible for configuration of the plugin.
	Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error)
	//* Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error)
}

type nodeAttestorClient struct {
	cc grpc.ClientConnInterface
}

func NewNodeAttestorClient(cc grpc.ClientConnInterface) NodeAttestorClient {
	return &nodeAttestorClient{cc}
}

func (c *nodeAttestorClient) Attest(ctx context.Context, opts ...grpc.CallOption) (NodeAttestor_AttestClient, error) {
	stream, err := c.cc.NewStream(ctx, &_NodeAttestor_serviceDesc.Streams[0], "/spire.agent.nodeattestor.NodeAttestor/Attest", opts...)
	if err != nil {
		return nil, err
	}
	x := &nodeAttestorAttestClient{stream}
	return x, nil
}

type NodeAttestor_AttestClient interface {
	Send(*AttestRequest) error
	Recv() (*AttestResponse, error)
	grpc.ClientStream
}

type nodeAttestorAttestClient struct {
	grpc.ClientStream
}

func (x *nodeAttestorAttestClient) Send(m *AttestRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *nodeAttestorAttestClient) Recv() (*AttestResponse, error) {
	m := new(AttestResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *nodeAttestorClient) Configure(ctx context.Context, in *plugin.ConfigureRequest, opts ...grpc.CallOption) (*plugin.ConfigureResponse, error) {
	out := new(plugin.ConfigureResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/Configure", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeAttestorClient) GetPluginInfo(ctx context.Context, in *plugin.GetPluginInfoRequest, opts ...grpc.CallOption) (*plugin.GetPluginInfoResponse, error) {
	out := new(plugin.GetPluginInfoResponse)
	err := c.cc.Invoke(ctx, "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NodeAttestorServer is the server API for NodeAttestor service.
type NodeAttestorServer interface {
	//* Attesta a node.
	Attest(NodeAttestor_AttestServer) error
	//* Responsible for configuration of the plugin.
	Configure(context.Context, *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error)
	//* Returns the  version and related metadata of the installed plugin.
	GetPluginInfo(context.Context, *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error)
}

// UnimplementedNodeAttestorServer can be embedded to have forward compatible implementations.
type UnimplementedNodeAttestorServer struct {
}

func (*UnimplementedNodeAttestorServer) Attest(srv NodeAttestor_AttestServer) error {
	return status.Errorf(codes.Unimplemented, "method Attest not implemented")
}
func (*UnimplementedNodeAttestorServer) Configure(ctx context.Context, req *plugin.ConfigureRequest) (*plugin.ConfigureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Configure not implemented")
}
func (*UnimplementedNodeAttestorServer) GetPluginInfo(ctx context.Context, req *plugin.GetPluginInfoRequest) (*plugin.GetPluginInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginInfo not implemented")
}

func RegisterNodeAttestorServer(s *grpc.Server, srv NodeAttestorServer) {
	s.RegisterService(&_NodeAttestor_serviceDesc, srv)
}

func _NodeAttestor_Attest_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(NodeAttestorServer).Attest(&nodeAttestorAttestServer{stream})
}

type NodeAttestor_AttestServer interface {
	Send(*AttestResponse) error
	Recv() (*AttestRequest, error)
	grpc.ServerStream
}

type nodeAttestorAttestServer struct {
	grpc.ServerStream
}

func (x *nodeAttestorAttestServer) Send(m *AttestResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *nodeAttestorAttestServer) Recv() (*AttestRequest, error) {
	m := new(AttestRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _NodeAttestor_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.ConfigureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).Configure(ctx, req.(*plugin.ConfigureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NodeAttestor_GetPluginInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(plugin.GetPluginInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spire.agent.nodeattestor.NodeAttestor/GetPluginInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeAttestorServer).GetPluginInfo(ctx, req.(*plugin.GetPluginInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NodeAttestor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "spire.agent.nodeattestor.NodeAttestor",
	HandlerType: (*NodeAttestorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _NodeAttestor_Configure_Handler,
		},
		{
			MethodName: "GetPluginInfo",
			Handler:    _NodeAttestor_GetPluginInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Attest",
			Handler:       _NodeAttestor_Attest_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "spire/server/nodeattestor/nodeattestor.proto",
}
