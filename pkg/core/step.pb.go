// Code generated by protoc-gen-go. DO NOT EDIT.
// source: step.proto

package core // import "github.com/apptreesoftware/go-workflow/pkg/core"

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

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StepQueueWorkflowRequest struct {
	Environment          *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	Workflow             string       `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Input                []byte       `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte       `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32        `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *StepQueueWorkflowRequest) Reset()         { *m = StepQueueWorkflowRequest{} }
func (m *StepQueueWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*StepQueueWorkflowRequest) ProtoMessage()    {}
func (*StepQueueWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_bace0082f1fe2049, []int{0}
}
func (m *StepQueueWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepQueueWorkflowRequest.Unmarshal(m, b)
}
func (m *StepQueueWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepQueueWorkflowRequest.Marshal(b, m, deterministic)
}
func (dst *StepQueueWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepQueueWorkflowRequest.Merge(dst, src)
}
func (m *StepQueueWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_StepQueueWorkflowRequest.Size(m)
}
func (m *StepQueueWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StepQueueWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StepQueueWorkflowRequest proto.InternalMessageInfo

func (m *StepQueueWorkflowRequest) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *StepQueueWorkflowRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *StepQueueWorkflowRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

type StepQueueWorkflowResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *StepQueueWorkflowResponse) Reset()         { *m = StepQueueWorkflowResponse{} }
func (m *StepQueueWorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*StepQueueWorkflowResponse) ProtoMessage()    {}
func (*StepQueueWorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_bace0082f1fe2049, []int{1}
}
func (m *StepQueueWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepQueueWorkflowResponse.Unmarshal(m, b)
}
func (m *StepQueueWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepQueueWorkflowResponse.Marshal(b, m, deterministic)
}
func (dst *StepQueueWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepQueueWorkflowResponse.Merge(dst, src)
}
func (m *StepQueueWorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_StepQueueWorkflowResponse.Size(m)
}
func (m *StepQueueWorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StepQueueWorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StepQueueWorkflowResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StepQueueWorkflowRequest)(nil), "core.StepQueueWorkflowRequest")
	proto.RegisterType((*StepQueueWorkflowResponse)(nil), "core.StepQueueWorkflowResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EngineStepAPIClient is the client API for EngineStepAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EngineStepAPIClient interface {
	CachePush(ctx context.Context, in *CachePushRequest, opts ...grpc.CallOption) (*CachePushResponse, error)
	CachePull(ctx context.Context, in *CachePullRequest, opts ...grpc.CallOption) (*CachePullResponse, error)
	QueueWorkflow(ctx context.Context, in *StepQueueWorkflowRequest, opts ...grpc.CallOption) (*StepQueueWorkflowResponse, error)
}

type engineStepAPIClient struct {
	cc *grpc.ClientConn
}

func NewEngineStepAPIClient(cc *grpc.ClientConn) EngineStepAPIClient {
	return &engineStepAPIClient{cc}
}

func (c *engineStepAPIClient) CachePush(ctx context.Context, in *CachePushRequest, opts ...grpc.CallOption) (*CachePushResponse, error) {
	out := new(CachePushResponse)
	err := c.cc.Invoke(ctx, "/core.EngineStepAPI/CachePush", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineStepAPIClient) CachePull(ctx context.Context, in *CachePullRequest, opts ...grpc.CallOption) (*CachePullResponse, error) {
	out := new(CachePullResponse)
	err := c.cc.Invoke(ctx, "/core.EngineStepAPI/CachePull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *engineStepAPIClient) QueueWorkflow(ctx context.Context, in *StepQueueWorkflowRequest, opts ...grpc.CallOption) (*StepQueueWorkflowResponse, error) {
	out := new(StepQueueWorkflowResponse)
	err := c.cc.Invoke(ctx, "/core.EngineStepAPI/QueueWorkflow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EngineStepAPIServer is the server API for EngineStepAPI service.
type EngineStepAPIServer interface {
	CachePush(context.Context, *CachePushRequest) (*CachePushResponse, error)
	CachePull(context.Context, *CachePullRequest) (*CachePullResponse, error)
	QueueWorkflow(context.Context, *StepQueueWorkflowRequest) (*StepQueueWorkflowResponse, error)
}

func RegisterEngineStepAPIServer(s *grpc.Server, srv EngineStepAPIServer) {
	s.RegisterService(&_EngineStepAPI_serviceDesc, srv)
}

func _EngineStepAPI_CachePush_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CachePushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineStepAPIServer).CachePush(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.EngineStepAPI/CachePush",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineStepAPIServer).CachePush(ctx, req.(*CachePushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EngineStepAPI_CachePull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CachePullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineStepAPIServer).CachePull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.EngineStepAPI/CachePull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineStepAPIServer).CachePull(ctx, req.(*CachePullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EngineStepAPI_QueueWorkflow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepQueueWorkflowRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EngineStepAPIServer).QueueWorkflow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.EngineStepAPI/QueueWorkflow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EngineStepAPIServer).QueueWorkflow(ctx, req.(*StepQueueWorkflowRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EngineStepAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.EngineStepAPI",
	HandlerType: (*EngineStepAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CachePush",
			Handler:    _EngineStepAPI_CachePush_Handler,
		},
		{
			MethodName: "CachePull",
			Handler:    _EngineStepAPI_CachePull_Handler,
		},
		{
			MethodName: "QueueWorkflow",
			Handler:    _EngineStepAPI_QueueWorkflow_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "step.proto",
}

func init() { proto.RegisterFile("step.proto", fileDescriptor_step_bace0082f1fe2049) }

var fileDescriptor_step_bace0082f1fe2049 = []byte{
	// 277 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xbd, 0x4e, 0xc3, 0x40,
	0x10, 0x84, 0x75, 0xfc, 0x89, 0x6c, 0x92, 0x82, 0x13, 0x02, 0x63, 0x24, 0x88, 0x52, 0xb9, 0xc1,
	0x16, 0x49, 0x4b, 0x03, 0x28, 0x05, 0x9d, 0x31, 0x05, 0x12, 0x9d, 0x63, 0x6d, 0x6c, 0x2b, 0xf6,
	0xed, 0x71, 0x3f, 0xb8, 0xe6, 0x41, 0x79, 0x17, 0x64, 0xce, 0x84, 0x00, 0x49, 0xb9, 0xb3, 0x33,
	0x9f, 0x46, 0xbb, 0x00, 0xda, 0xa0, 0x0c, 0xa5, 0x22, 0x43, 0x7c, 0x2f, 0x23, 0x85, 0xfe, 0x20,
	0xa3, 0xba, 0x26, 0xe1, 0xb4, 0xf1, 0x3b, 0x03, 0xef, 0xc9, 0xa0, 0x7c, 0xb4, 0x68, 0xf1, 0x99,
	0xd4, 0x72, 0x51, 0x51, 0x93, 0xe0, 0xab, 0x45, 0x6d, 0xf8, 0x14, 0xfa, 0x28, 0xde, 0x4a, 0x45,
	0xa2, 0x46, 0x61, 0x3c, 0x36, 0x62, 0x41, 0x7f, 0x72, 0x14, 0xb6, 0x98, 0x70, 0xf6, 0xb3, 0x48,
	0xd6, 0x5d, 0xdc, 0x87, 0xc3, 0xa6, 0xe3, 0x78, 0x3b, 0x23, 0x16, 0xf4, 0x92, 0xd5, 0xcc, 0x8f,
	0x61, 0xbf, 0x14, 0xd2, 0x1a, 0x6f, 0x77, 0xc4, 0x82, 0x41, 0xe2, 0x86, 0xf1, 0x39, 0x9c, 0x6d,
	0xa8, 0xa0, 0x25, 0x09, 0x8d, 0x93, 0x0f, 0x06, 0xc3, 0x99, 0xc8, 0x4b, 0x81, 0xad, 0xe7, 0x36,
	0x7e, 0xe0, 0x37, 0xd0, 0xbb, 0x4f, 0xb3, 0x02, 0x63, 0xab, 0x0b, 0x7e, 0xe2, 0xda, 0xac, 0x84,
	0xae, 0xba, 0x7f, 0xfa, 0x4f, 0x77, 0xbc, 0xb5, 0x74, 0x55, 0xfd, 0x49, 0x57, 0xd5, 0xe6, 0x74,
	0xab, 0x77, 0xe9, 0x18, 0x86, 0xbf, 0x6a, 0xf2, 0x0b, 0xe7, 0xdc, 0x76, 0x42, 0xff, 0x72, 0xeb,
	0xde, 0x11, 0xef, 0xae, 0x5f, 0xa2, 0xbc, 0x34, 0x85, 0x9d, 0x87, 0x19, 0xd5, 0x51, 0x2a, 0xa5,
	0x51, 0x88, 0x9a, 0x16, 0xa6, 0x49, 0x15, 0x46, 0x39, 0x5d, 0x7d, 0x1f, 0x2f, 0x92, 0xcb, 0x3c,
	0x6a, 0x61, 0xf3, 0x83, 0xaf, 0xd7, 0x4d, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x63, 0xd9, 0x18,
	0x1e, 0xdc, 0x01, 0x00, 0x00,
}
