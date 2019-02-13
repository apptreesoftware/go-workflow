// Code generated by protoc-gen-go. DO NOT EDIT.
// source: step.proto

package step // import "github.com/apptreesoftware/go-workflow/pkg/step"

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

type CachePushRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Metadata             []byte       `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Record               []byte       `protobuf:"bytes,3,opt,name=record,proto3" json:"record,omitempty"`
	CacheName            string       `protobuf:"bytes,4,opt,name=cacheName,proto3" json:"cacheName,omitempty"`
	Environment          *Environment `protobuf:"bytes,5,opt,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte       `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32        `json:"-" yaml:"-" xml:"-"`
}

func (m *CachePushRequest) Reset()         { *m = CachePushRequest{} }
func (m *CachePushRequest) String() string { return proto.CompactTextString(m) }
func (*CachePushRequest) ProtoMessage()    {}
func (*CachePushRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{0}
}
func (m *CachePushRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CachePushRequest.Unmarshal(m, b)
}
func (m *CachePushRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CachePushRequest.Marshal(b, m, deterministic)
}
func (dst *CachePushRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CachePushRequest.Merge(dst, src)
}
func (m *CachePushRequest) XXX_Size() int {
	return xxx_messageInfo_CachePushRequest.Size(m)
}
func (m *CachePushRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CachePushRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CachePushRequest proto.InternalMessageInfo

func (m *CachePushRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CachePushRequest) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *CachePushRequest) GetRecord() []byte {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *CachePushRequest) GetCacheName() string {
	if m != nil {
		return m.CacheName
	}
	return ""
}

func (m *CachePushRequest) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

type CachePushResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-"`
}

func (m *CachePushResponse) Reset()         { *m = CachePushResponse{} }
func (m *CachePushResponse) String() string { return proto.CompactTextString(m) }
func (*CachePushResponse) ProtoMessage()    {}
func (*CachePushResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{1}
}
func (m *CachePushResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CachePushResponse.Unmarshal(m, b)
}
func (m *CachePushResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CachePushResponse.Marshal(b, m, deterministic)
}
func (dst *CachePushResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CachePushResponse.Merge(dst, src)
}
func (m *CachePushResponse) XXX_Size() int {
	return xxx_messageInfo_CachePushResponse.Size(m)
}
func (m *CachePushResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CachePushResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CachePushResponse proto.InternalMessageInfo

type Environment struct {
	App                  string   `protobuf:"bytes,1,opt,name=app,proto3" json:"app,omitempty"`
	WorkflowId           string   `protobuf:"bytes,2,opt,name=workflowId,proto3" json:"workflowId,omitempty"`
	RunId                string   `protobuf:"bytes,3,opt,name=runId,proto3" json:"runId,omitempty"`
	StepName             string   `protobuf:"bytes,4,opt,name=stepName,proto3" json:"stepName,omitempty"`
	StepVersion          string   `protobuf:"bytes,5,opt,name=stepVersion,proto3" json:"stepVersion,omitempty"`
	InputFile            string   `protobuf:"bytes,6,opt,name=inputFile,proto3" json:"inputFile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-"`
}

func (m *Environment) Reset()         { *m = Environment{} }
func (m *Environment) String() string { return proto.CompactTextString(m) }
func (*Environment) ProtoMessage()    {}
func (*Environment) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{2}
}
func (m *Environment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Environment.Unmarshal(m, b)
}
func (m *Environment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Environment.Marshal(b, m, deterministic)
}
func (dst *Environment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Environment.Merge(dst, src)
}
func (m *Environment) XXX_Size() int {
	return xxx_messageInfo_Environment.Size(m)
}
func (m *Environment) XXX_DiscardUnknown() {
	xxx_messageInfo_Environment.DiscardUnknown(m)
}

var xxx_messageInfo_Environment proto.InternalMessageInfo

func (m *Environment) GetApp() string {
	if m != nil {
		return m.App
	}
	return ""
}

func (m *Environment) GetWorkflowId() string {
	if m != nil {
		return m.WorkflowId
	}
	return ""
}

func (m *Environment) GetRunId() string {
	if m != nil {
		return m.RunId
	}
	return ""
}

func (m *Environment) GetStepName() string {
	if m != nil {
		return m.StepName
	}
	return ""
}

func (m *Environment) GetStepVersion() string {
	if m != nil {
		return m.StepVersion
	}
	return ""
}

func (m *Environment) GetInputFile() string {
	if m != nil {
		return m.InputFile
	}
	return ""
}

type CachePullRequest struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CacheName            string       `protobuf:"bytes,2,opt,name=cacheName,proto3" json:"cacheName,omitempty"`
	Environment          *Environment `protobuf:"bytes,3,opt,name=environment,proto3" json:"environment,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte       `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32        `json:"-" yaml:"-" xml:"-"`
}

func (m *CachePullRequest) Reset()         { *m = CachePullRequest{} }
func (m *CachePullRequest) String() string { return proto.CompactTextString(m) }
func (*CachePullRequest) ProtoMessage()    {}
func (*CachePullRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{3}
}
func (m *CachePullRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CachePullRequest.Unmarshal(m, b)
}
func (m *CachePullRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CachePullRequest.Marshal(b, m, deterministic)
}
func (dst *CachePullRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CachePullRequest.Merge(dst, src)
}
func (m *CachePullRequest) XXX_Size() int {
	return xxx_messageInfo_CachePullRequest.Size(m)
}
func (m *CachePullRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CachePullRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CachePullRequest proto.InternalMessageInfo

func (m *CachePullRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CachePullRequest) GetCacheName() string {
	if m != nil {
		return m.CacheName
	}
	return ""
}

func (m *CachePullRequest) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

type CachePullResponse struct {
	Record               []byte   `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	Metadata             []byte   `protobuf:"bytes,2,opt,name=metadata,proto3" json:"metadata,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-"`
}

func (m *CachePullResponse) Reset()         { *m = CachePullResponse{} }
func (m *CachePullResponse) String() string { return proto.CompactTextString(m) }
func (*CachePullResponse) ProtoMessage()    {}
func (*CachePullResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{4}
}
func (m *CachePullResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CachePullResponse.Unmarshal(m, b)
}
func (m *CachePullResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CachePullResponse.Marshal(b, m, deterministic)
}
func (dst *CachePullResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CachePullResponse.Merge(dst, src)
}
func (m *CachePullResponse) XXX_Size() int {
	return xxx_messageInfo_CachePullResponse.Size(m)
}
func (m *CachePullResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CachePullResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CachePullResponse proto.InternalMessageInfo

func (m *CachePullResponse) GetRecord() []byte {
	if m != nil {
		return m.Record
	}
	return nil
}

func (m *CachePullResponse) GetMetadata() []byte {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type Package struct {
	Name                 string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Lang                 string                  `protobuf:"bytes,2,opt,name=lang,proto3" json:"lang,omitempty"`
	Exec                 string                  `protobuf:"bytes,3,opt,name=exec,proto3" json:"exec,omitempty"`
	Steps                map[string]*PackageStep `protobuf:"bytes,4,rep,name=steps,proto3" json:"steps,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte                  `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32                   `json:"-" yaml:"-" xml:"-"`
}

func (m *Package) Reset()         { *m = Package{} }
func (m *Package) String() string { return proto.CompactTextString(m) }
func (*Package) ProtoMessage()    {}
func (*Package) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{5}
}
func (m *Package) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Package.Unmarshal(m, b)
}
func (m *Package) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Package.Marshal(b, m, deterministic)
}
func (dst *Package) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Package.Merge(dst, src)
}
func (m *Package) XXX_Size() int {
	return xxx_messageInfo_Package.Size(m)
}
func (m *Package) XXX_DiscardUnknown() {
	xxx_messageInfo_Package.DiscardUnknown(m)
}

var xxx_messageInfo_Package proto.InternalMessageInfo

func (m *Package) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Package) GetLang() string {
	if m != nil {
		return m.Lang
	}
	return ""
}

func (m *Package) GetExec() string {
	if m != nil {
		return m.Exec
	}
	return ""
}

func (m *Package) GetSteps() map[string]*PackageStep {
	if m != nil {
		return m.Steps
	}
	return nil
}

type PackageStep struct {
	Description          string                 `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	Inputs               map[string]*InputInfo  `protobuf:"bytes,2,rep,name=inputs,proto3" json:"inputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Outputs              map[string]*OutputInfo `protobuf:"bytes,3,rep,name=outputs,proto3" json:"outputs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte                 `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32                  `json:"-" yaml:"-" xml:"-"`
}

func (m *PackageStep) Reset()         { *m = PackageStep{} }
func (m *PackageStep) String() string { return proto.CompactTextString(m) }
func (*PackageStep) ProtoMessage()    {}
func (*PackageStep) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{6}
}
func (m *PackageStep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PackageStep.Unmarshal(m, b)
}
func (m *PackageStep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PackageStep.Marshal(b, m, deterministic)
}
func (dst *PackageStep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PackageStep.Merge(dst, src)
}
func (m *PackageStep) XXX_Size() int {
	return xxx_messageInfo_PackageStep.Size(m)
}
func (m *PackageStep) XXX_DiscardUnknown() {
	xxx_messageInfo_PackageStep.DiscardUnknown(m)
}

var xxx_messageInfo_PackageStep proto.InternalMessageInfo

func (m *PackageStep) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PackageStep) GetInputs() map[string]*InputInfo {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *PackageStep) GetOutputs() map[string]*OutputInfo {
	if m != nil {
		return m.Outputs
	}
	return nil
}

type InputInfo struct {
	Required             bool     `protobuf:"varint,1,opt,name=required,proto3" json:"required,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-"`
}

func (m *InputInfo) Reset()         { *m = InputInfo{} }
func (m *InputInfo) String() string { return proto.CompactTextString(m) }
func (*InputInfo) ProtoMessage()    {}
func (*InputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{7}
}
func (m *InputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputInfo.Unmarshal(m, b)
}
func (m *InputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputInfo.Marshal(b, m, deterministic)
}
func (dst *InputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputInfo.Merge(dst, src)
}
func (m *InputInfo) XXX_Size() int {
	return xxx_messageInfo_InputInfo.Size(m)
}
func (m *InputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_InputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_InputInfo proto.InternalMessageInfo

func (m *InputInfo) GetRequired() bool {
	if m != nil {
		return m.Required
	}
	return false
}

func (m *InputInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

type OutputInfo struct {
	Description          string   `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-"`
}

func (m *OutputInfo) Reset()         { *m = OutputInfo{} }
func (m *OutputInfo) String() string { return proto.CompactTextString(m) }
func (*OutputInfo) ProtoMessage()    {}
func (*OutputInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_step_4f21a37eff3ade7d, []int{8}
}
func (m *OutputInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputInfo.Unmarshal(m, b)
}
func (m *OutputInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputInfo.Marshal(b, m, deterministic)
}
func (dst *OutputInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputInfo.Merge(dst, src)
}
func (m *OutputInfo) XXX_Size() int {
	return xxx_messageInfo_OutputInfo.Size(m)
}
func (m *OutputInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputInfo.DiscardUnknown(m)
}

var xxx_messageInfo_OutputInfo proto.InternalMessageInfo

func (m *OutputInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*CachePushRequest)(nil), "CachePushRequest")
	proto.RegisterType((*CachePushResponse)(nil), "CachePushResponse")
	proto.RegisterType((*Environment)(nil), "Environment")
	proto.RegisterType((*CachePullRequest)(nil), "CachePullRequest")
	proto.RegisterType((*CachePullResponse)(nil), "CachePullResponse")
	proto.RegisterType((*Package)(nil), "Package")
	proto.RegisterMapType((map[string]*PackageStep)(nil), "Package.StepsEntry")
	proto.RegisterType((*PackageStep)(nil), "PackageStep")
	proto.RegisterMapType((map[string]*InputInfo)(nil), "PackageStep.InputsEntry")
	proto.RegisterMapType((map[string]*OutputInfo)(nil), "PackageStep.OutputsEntry")
	proto.RegisterType((*InputInfo)(nil), "InputInfo")
	proto.RegisterType((*OutputInfo)(nil), "OutputInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheClient is the client API for Cache service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheClient interface {
	Push(ctx context.Context, in *CachePushRequest, opts ...grpc.CallOption) (*CachePushResponse, error)
	Pull(ctx context.Context, in *CachePullRequest, opts ...grpc.CallOption) (*CachePullResponse, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Push(ctx context.Context, in *CachePushRequest, opts ...grpc.CallOption) (*CachePushResponse, error) {
	out := new(CachePushResponse)
	err := c.cc.Invoke(ctx, "/Cache/Push", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Pull(ctx context.Context, in *CachePullRequest, opts ...grpc.CallOption) (*CachePullResponse, error) {
	out := new(CachePullResponse)
	err := c.cc.Invoke(ctx, "/Cache/Pull", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServer is the server API for Cache service.
type CacheServer interface {
	Push(context.Context, *CachePushRequest) (*CachePushResponse, error)
	Pull(context.Context, *CachePullRequest) (*CachePullResponse, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Push_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CachePushRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Push(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cache/Push",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Push(ctx, req.(*CachePushRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Pull_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CachePullRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Pull(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Cache/Pull",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Pull(ctx, req.(*CachePullRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Push",
			Handler:    _Cache_Push_Handler,
		},
		{
			MethodName: "Pull",
			Handler:    _Cache_Pull_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "step.proto",
}

func init() { proto.RegisterFile("step.proto", fileDescriptor_step_4f21a37eff3ade7d) }

var fileDescriptor_step_4f21a37eff3ade7d = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0x9d, 0x3f, 0xad, 0xc7, 0xd1, 0x4f, 0xed, 0xf6, 0x27, 0x64, 0x22, 0x84, 0x8c, 0x4f,
	0xe1, 0x50, 0x07, 0xd2, 0x0b, 0xe2, 0x08, 0x4a, 0xab, 0x5c, 0xa0, 0x32, 0x12, 0x07, 0x6e, 0x5b,
	0x7b, 0x92, 0x58, 0x71, 0xbc, 0xee, 0x7a, 0xdd, 0x90, 0xef, 0x83, 0xb8, 0x73, 0xe3, 0xe3, 0xa1,
	0xdd, 0x75, 0xec, 0x4d, 0x5a, 0xc8, 0x6d, 0xf7, 0xed, 0xcc, 0xdb, 0x37, 0x6f, 0x66, 0x17, 0xa0,
	0x14, 0x58, 0x84, 0x05, 0x67, 0x82, 0x05, 0x3f, 0x2c, 0x38, 0xfb, 0x48, 0xe3, 0x25, 0xde, 0x56,
	0xe5, 0x32, 0xc2, 0xfb, 0x0a, 0x4b, 0x41, 0xfe, 0x03, 0x3b, 0x4d, 0x3c, 0xcb, 0xb7, 0x46, 0x4e,
	0x64, 0xa7, 0x09, 0x19, 0xc2, 0xe9, 0x1a, 0x05, 0x4d, 0xa8, 0xa0, 0x9e, 0xed, 0x5b, 0xa3, 0x41,
	0xd4, 0xec, 0xc9, 0x33, 0xe8, 0x73, 0x8c, 0x19, 0x4f, 0xbc, 0x8e, 0x3a, 0xa9, 0x77, 0xe4, 0x05,
	0x38, 0xb1, 0xe4, 0xfd, 0x44, 0xd7, 0xe8, 0x75, 0x15, 0x55, 0x0b, 0x90, 0x10, 0x5c, 0xcc, 0x1f,
	0x52, 0xce, 0xf2, 0x35, 0xe6, 0xc2, 0xeb, 0xf9, 0xd6, 0xc8, 0x9d, 0x0c, 0xc2, 0x69, 0x8b, 0x45,
	0x66, 0x40, 0x70, 0x01, 0xe7, 0x86, 0xca, 0xb2, 0x60, 0x79, 0x89, 0xc1, 0x2f, 0x0b, 0x5c, 0x23,
	0x83, 0x9c, 0x41, 0x87, 0x16, 0x45, 0xad, 0x5b, 0x2e, 0xc9, 0x4b, 0x80, 0x0d, 0xe3, 0xab, 0x79,
	0xc6, 0x36, 0xb3, 0x44, 0x49, 0x77, 0x22, 0x03, 0x21, 0xff, 0x43, 0x8f, 0x57, 0xf9, 0x4c, 0x6b,
	0x77, 0x22, 0xbd, 0x91, 0xe5, 0x4a, 0x87, 0x0c, 0xe5, 0xcd, 0x9e, 0xf8, 0xe0, 0xca, 0xf5, 0x57,
	0xe4, 0x65, 0xca, 0x72, 0x25, 0xdc, 0x89, 0x4c, 0x48, 0x16, 0x9e, 0xe6, 0x45, 0x25, 0xae, 0xd3,
	0x0c, 0xbd, 0xbe, 0x2e, 0xbc, 0x01, 0x82, 0xa2, 0xb1, 0x3b, 0xcb, 0xfe, 0x66, 0xf7, 0x9e, 0x75,
	0xf6, 0x11, 0xeb, 0x3a, 0xc7, 0xac, 0xbb, 0x69, 0xac, 0x93, 0x37, 0x6a, 0xeb, 0x8c, 0xae, 0x59,
	0x7b, 0x5d, 0xfb, 0x47, 0xa7, 0x83, 0xdf, 0x16, 0x9c, 0xdc, 0xd2, 0x78, 0x45, 0x17, 0x48, 0x08,
	0x74, 0x73, 0xa9, 0x4e, 0x8b, 0x56, 0x6b, 0x89, 0x65, 0x34, 0x5f, 0xd4, 0x8a, 0xd5, 0x5a, 0x62,
	0xf8, 0x1d, 0xe3, 0xda, 0x5f, 0xb5, 0x26, 0xaf, 0xa1, 0x27, 0xfd, 0x2a, 0xbd, 0xae, 0xdf, 0x19,
	0xb9, 0x93, 0x8b, 0xb0, 0x26, 0x0d, 0xbf, 0x48, 0x74, 0x9a, 0x0b, 0xbe, 0x8d, 0x74, 0xc4, 0xf0,
	0x1a, 0xa0, 0x05, 0x65, 0x7f, 0x57, 0xb8, 0xdd, 0xf5, 0x77, 0x85, 0x5b, 0x12, 0x40, 0xef, 0x81,
	0x66, 0x95, 0x76, 0x49, 0xba, 0x50, 0x53, 0xc9, 0xa4, 0x48, 0x1f, 0xbd, 0xb7, 0xdf, 0x59, 0xc1,
	0x4f, 0x1b, 0x5c, 0xe3, 0x48, 0x76, 0x31, 0xc1, 0x32, 0xe6, 0x69, 0x21, 0x64, 0x17, 0x35, 0xa3,
	0x09, 0x91, 0x37, 0xd0, 0x57, 0x4d, 0x2b, 0x3d, 0x5b, 0xa9, 0xf4, 0x4c, 0xea, 0x70, 0xa6, 0x8e,
	0xb4, 0xd4, 0x3a, 0x8e, 0x5c, 0xc1, 0x09, 0xab, 0x84, 0x4a, 0xe9, 0xa8, 0x94, 0xe7, 0x7b, 0x29,
	0x9f, 0xf5, 0x99, 0xce, 0xd9, 0x45, 0x0e, 0xa7, 0xe0, 0x1a, 0x5c, 0x4f, 0x54, 0xe8, 0xef, 0x57,
	0x08, 0xfa, 0xea, 0x59, 0x3e, 0x67, 0x46, 0x7d, 0xc3, 0x1b, 0x18, 0x98, 0xfc, 0x4f, 0xf0, 0xbc,
	0xda, 0xe7, 0x71, 0x6b, 0x3d, 0x07, 0x44, 0xc1, 0x0c, 0x9c, 0xe6, 0x02, 0x39, 0x0c, 0x1c, 0xef,
	0xab, 0x94, 0xa3, 0x1e, 0x93, 0xd3, 0xa8, 0xd9, 0x1f, 0x3a, 0x68, 0x3f, 0x72, 0x30, 0x08, 0x01,
	0xda, 0x3b, 0x8e, 0x3b, 0x3e, 0x41, 0xe8, 0xa9, 0x39, 0x25, 0x97, 0xd0, 0x95, 0xcf, 0x9c, 0x9c,
	0x87, 0x87, 0x1f, 0xd3, 0x90, 0x84, 0x8f, 0x7e, 0x01, 0x1d, 0x9e, 0x65, 0x6d, 0x78, 0xf3, 0xb0,
	0xda, 0xf0, 0x76, 0xf2, 0x3f, 0xbc, 0xfd, 0x36, 0x5e, 0xa4, 0x62, 0x59, 0xdd, 0x85, 0x31, 0x5b,
	0x8f, 0x69, 0x51, 0x08, 0x8e, 0x58, 0xb2, 0xb9, 0xd8, 0x50, 0x8e, 0xe3, 0x05, 0xbb, 0xdc, 0x7d,
	0x0f, 0xe3, 0x62, 0xb5, 0x18, 0xcb, 0x31, 0xbc, 0xeb, 0xab, 0xaf, 0xf2, 0xea, 0x4f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x20, 0xb2, 0xcf, 0x73, 0x38, 0x05, 0x00, 0x00,
}
