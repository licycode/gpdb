// Code generated by protoc-gen-go. DO NOT EDIT.
// source: hub_to_agent.proto

package idl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PingAgentsRequest struct {
}

func (m *PingAgentsRequest) Reset()                    { *m = PingAgentsRequest{} }
func (m *PingAgentsRequest) String() string            { return proto.CompactTextString(m) }
func (*PingAgentsRequest) ProtoMessage()               {}
func (*PingAgentsRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type PingAgentsReply struct {
}

func (m *PingAgentsReply) Reset()                    { *m = PingAgentsReply{} }
func (m *PingAgentsReply) String() string            { return proto.CompactTextString(m) }
func (*PingAgentsReply) ProtoMessage()               {}
func (*PingAgentsReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type CheckUpgradeStatusRequest struct {
}

func (m *CheckUpgradeStatusRequest) Reset()                    { *m = CheckUpgradeStatusRequest{} }
func (m *CheckUpgradeStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusRequest) ProtoMessage()               {}
func (*CheckUpgradeStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

type CheckUpgradeStatusReply struct {
	ProcessList string `protobuf:"bytes,1,opt,name=process_list,json=processList" json:"process_list,omitempty"`
}

func (m *CheckUpgradeStatusReply) Reset()                    { *m = CheckUpgradeStatusReply{} }
func (m *CheckUpgradeStatusReply) String() string            { return proto.CompactTextString(m) }
func (*CheckUpgradeStatusReply) ProtoMessage()               {}
func (*CheckUpgradeStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *CheckUpgradeStatusReply) GetProcessList() string {
	if m != nil {
		return m.ProcessList
	}
	return ""
}

type CheckConversionStatusRequest struct {
}

func (m *CheckConversionStatusRequest) Reset()                    { *m = CheckConversionStatusRequest{} }
func (m *CheckConversionStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckConversionStatusRequest) ProtoMessage()               {}
func (*CheckConversionStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

type CheckConversionStatusReply struct {
	Status []string `protobuf:"bytes,1,rep,name=status" json:"status,omitempty"`
}

func (m *CheckConversionStatusReply) Reset()                    { *m = CheckConversionStatusReply{} }
func (m *CheckConversionStatusReply) String() string            { return proto.CompactTextString(m) }
func (*CheckConversionStatusReply) ProtoMessage()               {}
func (*CheckConversionStatusReply) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

func (m *CheckConversionStatusReply) GetStatus() []string {
	if m != nil {
		return m.Status
	}
	return nil
}

type FileSysUsage struct {
	Filesystem string  `protobuf:"bytes,1,opt,name=filesystem" json:"filesystem,omitempty"`
	Usage      float64 `protobuf:"fixed64,2,opt,name=usage" json:"usage,omitempty"`
}

func (m *FileSysUsage) Reset()                    { *m = FileSysUsage{} }
func (m *FileSysUsage) String() string            { return proto.CompactTextString(m) }
func (*FileSysUsage) ProtoMessage()               {}
func (*FileSysUsage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *FileSysUsage) GetFilesystem() string {
	if m != nil {
		return m.Filesystem
	}
	return ""
}

func (m *FileSysUsage) GetUsage() float64 {
	if m != nil {
		return m.Usage
	}
	return 0
}

type CheckDiskUsageRequestToAgent struct {
}

func (m *CheckDiskUsageRequestToAgent) Reset()                    { *m = CheckDiskUsageRequestToAgent{} }
func (m *CheckDiskUsageRequestToAgent) String() string            { return proto.CompactTextString(m) }
func (*CheckDiskUsageRequestToAgent) ProtoMessage()               {}
func (*CheckDiskUsageRequestToAgent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type CheckDiskUsageReplyFromAgent struct {
	ListOfFileSysUsage []*FileSysUsage `protobuf:"bytes,1,rep,name=list_of_file_sys_usage,json=listOfFileSysUsage" json:"list_of_file_sys_usage,omitempty"`
}

func (m *CheckDiskUsageReplyFromAgent) Reset()                    { *m = CheckDiskUsageReplyFromAgent{} }
func (m *CheckDiskUsageReplyFromAgent) String() string            { return proto.CompactTextString(m) }
func (*CheckDiskUsageReplyFromAgent) ProtoMessage()               {}
func (*CheckDiskUsageReplyFromAgent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *CheckDiskUsageReplyFromAgent) GetListOfFileSysUsage() []*FileSysUsage {
	if m != nil {
		return m.ListOfFileSysUsage
	}
	return nil
}

func init() {
	proto.RegisterType((*PingAgentsRequest)(nil), "idl.PingAgentsRequest")
	proto.RegisterType((*PingAgentsReply)(nil), "idl.PingAgentsReply")
	proto.RegisterType((*CheckUpgradeStatusRequest)(nil), "idl.CheckUpgradeStatusRequest")
	proto.RegisterType((*CheckUpgradeStatusReply)(nil), "idl.CheckUpgradeStatusReply")
	proto.RegisterType((*CheckConversionStatusRequest)(nil), "idl.CheckConversionStatusRequest")
	proto.RegisterType((*CheckConversionStatusReply)(nil), "idl.CheckConversionStatusReply")
	proto.RegisterType((*FileSysUsage)(nil), "idl.FileSysUsage")
	proto.RegisterType((*CheckDiskUsageRequestToAgent)(nil), "idl.CheckDiskUsageRequestToAgent")
	proto.RegisterType((*CheckDiskUsageReplyFromAgent)(nil), "idl.CheckDiskUsageReplyFromAgent")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Agent service

type AgentClient interface {
	CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error)
	CheckConversionStatus(ctx context.Context, in *CheckConversionStatusRequest, opts ...grpc.CallOption) (*CheckConversionStatusReply, error)
	CheckDiskUsageOnAgents(ctx context.Context, in *CheckDiskUsageRequestToAgent, opts ...grpc.CallOption) (*CheckDiskUsageReplyFromAgent, error)
	PingAgents(ctx context.Context, in *PingAgentsRequest, opts ...grpc.CallOption) (*PingAgentsReply, error)
}

type agentClient struct {
	cc *grpc.ClientConn
}

func NewAgentClient(cc *grpc.ClientConn) AgentClient {
	return &agentClient{cc}
}

func (c *agentClient) CheckUpgradeStatus(ctx context.Context, in *CheckUpgradeStatusRequest, opts ...grpc.CallOption) (*CheckUpgradeStatusReply, error) {
	out := new(CheckUpgradeStatusReply)
	err := grpc.Invoke(ctx, "/idl.Agent/CheckUpgradeStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CheckConversionStatus(ctx context.Context, in *CheckConversionStatusRequest, opts ...grpc.CallOption) (*CheckConversionStatusReply, error) {
	out := new(CheckConversionStatusReply)
	err := grpc.Invoke(ctx, "/idl.Agent/CheckConversionStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) CheckDiskUsageOnAgents(ctx context.Context, in *CheckDiskUsageRequestToAgent, opts ...grpc.CallOption) (*CheckDiskUsageReplyFromAgent, error) {
	out := new(CheckDiskUsageReplyFromAgent)
	err := grpc.Invoke(ctx, "/idl.Agent/CheckDiskUsageOnAgents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *agentClient) PingAgents(ctx context.Context, in *PingAgentsRequest, opts ...grpc.CallOption) (*PingAgentsReply, error) {
	out := new(PingAgentsReply)
	err := grpc.Invoke(ctx, "/idl.Agent/PingAgents", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Agent service

type AgentServer interface {
	CheckUpgradeStatus(context.Context, *CheckUpgradeStatusRequest) (*CheckUpgradeStatusReply, error)
	CheckConversionStatus(context.Context, *CheckConversionStatusRequest) (*CheckConversionStatusReply, error)
	CheckDiskUsageOnAgents(context.Context, *CheckDiskUsageRequestToAgent) (*CheckDiskUsageReplyFromAgent, error)
	PingAgents(context.Context, *PingAgentsRequest) (*PingAgentsReply, error)
}

func RegisterAgentServer(s *grpc.Server, srv AgentServer) {
	s.RegisterService(&_Agent_serviceDesc, srv)
}

func _Agent_CheckUpgradeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckUpgradeStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CheckUpgradeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CheckUpgradeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CheckUpgradeStatus(ctx, req.(*CheckUpgradeStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CheckConversionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckConversionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CheckConversionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CheckConversionStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CheckConversionStatus(ctx, req.(*CheckConversionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_CheckDiskUsageOnAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDiskUsageRequestToAgent)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).CheckDiskUsageOnAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/CheckDiskUsageOnAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).CheckDiskUsageOnAgents(ctx, req.(*CheckDiskUsageRequestToAgent))
	}
	return interceptor(ctx, in, info, handler)
}

func _Agent_PingAgents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingAgentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AgentServer).PingAgents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.Agent/PingAgents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AgentServer).PingAgents(ctx, req.(*PingAgentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Agent_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.Agent",
	HandlerType: (*AgentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckUpgradeStatus",
			Handler:    _Agent_CheckUpgradeStatus_Handler,
		},
		{
			MethodName: "CheckConversionStatus",
			Handler:    _Agent_CheckConversionStatus_Handler,
		},
		{
			MethodName: "CheckDiskUsageOnAgents",
			Handler:    _Agent_CheckDiskUsageOnAgents_Handler,
		},
		{
			MethodName: "PingAgents",
			Handler:    _Agent_PingAgents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "hub_to_agent.proto",
}

func init() { proto.RegisterFile("hub_to_agent.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 368 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x8e, 0xda, 0x30,
	0x10, 0x25, 0x20, 0x90, 0x18, 0x90, 0x2a, 0x5c, 0x9a, 0xd2, 0x14, 0x51, 0xc8, 0x89, 0x13, 0x07,
	0xda, 0x23, 0x97, 0x0a, 0xca, 0xa9, 0x12, 0xab, 0x00, 0xc7, 0x55, 0x36, 0x80, 0x09, 0x16, 0x26,
	0xce, 0x66, 0x9c, 0x95, 0xf2, 0xaf, 0xfb, 0x31, 0x2b, 0x3b, 0x11, 0x04, 0x85, 0xec, 0xd1, 0xef,
	0xcd, 0xbc, 0x79, 0x79, 0x33, 0x01, 0x72, 0x8a, 0x77, 0xae, 0x14, 0xae, 0xe7, 0xd3, 0x40, 0x4e,
	0xc2, 0x48, 0x48, 0x41, 0x6a, 0xec, 0xc0, 0xed, 0xaf, 0xd0, 0x79, 0x62, 0x81, 0xff, 0x57, 0xe1,
	0xe8, 0xd0, 0xd7, 0x98, 0xa2, 0xb4, 0x3b, 0xf0, 0x25, 0x0f, 0x86, 0x3c, 0xb1, 0x7f, 0xc2, 0x8f,
	0xf9, 0x89, 0xee, 0xcf, 0xdb, 0xd0, 0x8f, 0xbc, 0x03, 0x5d, 0x4b, 0x4f, 0xc6, 0xd7, 0xfa, 0x19,
	0x7c, 0x7f, 0x44, 0x86, 0x3c, 0x21, 0x23, 0x68, 0x87, 0x91, 0xd8, 0x53, 0x44, 0x97, 0x33, 0x94,
	0x3d, 0x63, 0x68, 0x8c, 0x9b, 0x4e, 0x2b, 0xc3, 0xfe, 0x33, 0x94, 0xf6, 0x00, 0xfa, 0xba, 0x7b,
	0x2e, 0x82, 0x37, 0x1a, 0x21, 0x13, 0xc1, 0xbd, 0xfa, 0x1f, 0xb0, 0x4a, 0x78, 0x35, 0xc0, 0x84,
	0x06, 0xea, 0x67, 0xcf, 0x18, 0xd6, 0xc6, 0x4d, 0x27, 0x7b, 0xd9, 0x0b, 0x68, 0x2f, 0x19, 0xa7,
	0xeb, 0x04, 0xb7, 0xe8, 0xf9, 0x94, 0x0c, 0x00, 0x8e, 0x8c, 0x53, 0x4c, 0x50, 0xd2, 0x4b, 0x66,
	0x23, 0x87, 0x90, 0x2e, 0xd4, 0x63, 0x55, 0xd8, 0xab, 0x0e, 0x8d, 0xb1, 0xe1, 0xa4, 0x8f, 0xab,
	0xb7, 0x05, 0xc3, 0xb3, 0xd6, 0xc9, 0x4c, 0x6d, 0x84, 0x0e, 0xc7, 0xa6, 0x45, 0x3e, 0xe4, 0xc9,
	0x32, 0x12, 0x17, 0xcd, 0x93, 0x7f, 0x60, 0xaa, 0xcf, 0x76, 0xc5, 0xd1, 0x55, 0xb3, 0x5c, 0x4c,
	0xd0, 0x4d, 0xc7, 0x28, 0xb7, 0xad, 0x69, 0x67, 0xc2, 0x0e, 0x7c, 0x92, 0x37, 0xea, 0x10, 0xd5,
	0xb0, 0x3a, 0xe6, 0xb1, 0xe9, 0x7b, 0x15, 0xea, 0xa9, 0xe0, 0x06, 0x48, 0x31, 0x6a, 0x32, 0xd0,
	0x32, 0xa5, 0x0b, 0xb2, 0xfa, 0xa5, 0xbc, 0xda, 0x6d, 0x85, 0x3c, 0xc3, 0xb7, 0x87, 0x11, 0x93,
	0xd1, 0xad, 0xb1, 0x64, 0x3d, 0xd6, 0xaf, 0xcf, 0x4a, 0x52, 0xf9, 0x17, 0x30, 0xef, 0x53, 0x5a,
	0x05, 0xe9, 0x6d, 0xe5, 0xf5, 0x4b, 0x22, 0xb6, 0x1e, 0x97, 0xe4, 0x53, 0xb6, 0x2b, 0x64, 0x06,
	0x70, 0xbb, 0x58, 0x62, 0xea, 0x96, 0xc2, 0x5d, 0x5b, 0xdd, 0x02, 0xae, 0xfd, 0xed, 0x1a, 0xfa,
	0x87, 0xf8, 0xfd, 0x11, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x39, 0x40, 0x5e, 0x26, 0x03, 0x00, 0x00,
}
