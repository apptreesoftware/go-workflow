// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine_api.proto

package core // import "github.com/apptreesoftware/go-workflow/pkg/core"

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

type SpawnWorkflowRequest struct {
	Environment          *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	Workflow             string       `protobuf:"bytes,2,opt,name=workflow,proto3" json:"workflow,omitempty"`
	Input                []byte       `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte       `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32        `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *SpawnWorkflowRequest) Reset()         { *m = SpawnWorkflowRequest{} }
func (m *SpawnWorkflowRequest) String() string { return proto.CompactTextString(m) }
func (*SpawnWorkflowRequest) ProtoMessage()    {}
func (*SpawnWorkflowRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_api_1d670826c665dc28, []int{0}
}
func (m *SpawnWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnWorkflowRequest.Unmarshal(m, b)
}
func (m *SpawnWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnWorkflowRequest.Marshal(b, m, deterministic)
}
func (dst *SpawnWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnWorkflowRequest.Merge(dst, src)
}
func (m *SpawnWorkflowRequest) XXX_Size() int {
	return xxx_messageInfo_SpawnWorkflowRequest.Size(m)
}
func (m *SpawnWorkflowRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnWorkflowRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnWorkflowRequest proto.InternalMessageInfo

func (m *SpawnWorkflowRequest) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *SpawnWorkflowRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *SpawnWorkflowRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

type SpawnWorkflowResponse struct {
	QueueId              string   `protobuf:"bytes,1,opt,name=queueId,proto3" json:"queueId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *SpawnWorkflowResponse) Reset()         { *m = SpawnWorkflowResponse{} }
func (m *SpawnWorkflowResponse) String() string { return proto.CompactTextString(m) }
func (*SpawnWorkflowResponse) ProtoMessage()    {}
func (*SpawnWorkflowResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_api_1d670826c665dc28, []int{1}
}
func (m *SpawnWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnWorkflowResponse.Unmarshal(m, b)
}
func (m *SpawnWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnWorkflowResponse.Marshal(b, m, deterministic)
}
func (dst *SpawnWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnWorkflowResponse.Merge(dst, src)
}
func (m *SpawnWorkflowResponse) XXX_Size() int {
	return xxx_messageInfo_SpawnWorkflowResponse.Size(m)
}
func (m *SpawnWorkflowResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnWorkflowResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnWorkflowResponse proto.InternalMessageInfo

func (m *SpawnWorkflowResponse) GetQueueId() string {
	if m != nil {
		return m.QueueId
	}
	return ""
}

type StepCompleteRequest struct {
	Environment          *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	Output               []byte       `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
	ExitCode             int64        `protobuf:"varint,3,opt,name=exitCode,proto3" json:"exitCode,omitempty"`
	Error                string       `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte       `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32        `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *StepCompleteRequest) Reset()         { *m = StepCompleteRequest{} }
func (m *StepCompleteRequest) String() string { return proto.CompactTextString(m) }
func (*StepCompleteRequest) ProtoMessage()    {}
func (*StepCompleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_api_1d670826c665dc28, []int{2}
}
func (m *StepCompleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepCompleteRequest.Unmarshal(m, b)
}
func (m *StepCompleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepCompleteRequest.Marshal(b, m, deterministic)
}
func (dst *StepCompleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepCompleteRequest.Merge(dst, src)
}
func (m *StepCompleteRequest) XXX_Size() int {
	return xxx_messageInfo_StepCompleteRequest.Size(m)
}
func (m *StepCompleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StepCompleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StepCompleteRequest proto.InternalMessageInfo

func (m *StepCompleteRequest) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *StepCompleteRequest) GetOutput() []byte {
	if m != nil {
		return m.Output
	}
	return nil
}

func (m *StepCompleteRequest) GetExitCode() int64 {
	if m != nil {
		return m.ExitCode
	}
	return 0
}

func (m *StepCompleteRequest) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StepCompleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *StepCompleteResponse) Reset()         { *m = StepCompleteResponse{} }
func (m *StepCompleteResponse) String() string { return proto.CompactTextString(m) }
func (*StepCompleteResponse) ProtoMessage()    {}
func (*StepCompleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_api_1d670826c665dc28, []int{3}
}
func (m *StepCompleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepCompleteResponse.Unmarshal(m, b)
}
func (m *StepCompleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepCompleteResponse.Marshal(b, m, deterministic)
}
func (dst *StepCompleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepCompleteResponse.Merge(dst, src)
}
func (m *StepCompleteResponse) XXX_Size() int {
	return xxx_messageInfo_StepCompleteResponse.Size(m)
}
func (m *StepCompleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StepCompleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StepCompleteResponse proto.InternalMessageInfo

type LogRequest struct {
	JobId                string   `protobuf:"bytes,1,opt,name=jobId,proto3" json:"jobId,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	LogLevel             string   `protobuf:"bytes,3,opt,name=logLevel,proto3" json:"logLevel,omitempty"`
	Project              string   `protobuf:"bytes,4,opt,name=project,proto3" json:"project,omitempty"`
	Workflow             string   `protobuf:"bytes,5,opt,name=workflow,proto3" json:"workflow,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *LogRequest) Reset()         { *m = LogRequest{} }
func (m *LogRequest) String() string { return proto.CompactTextString(m) }
func (*LogRequest) ProtoMessage()    {}
func (*LogRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_api_1d670826c665dc28, []int{4}
}
func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (dst *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(dst, src)
}
func (m *LogRequest) XXX_Size() int {
	return xxx_messageInfo_LogRequest.Size(m)
}
func (m *LogRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogRequest proto.InternalMessageInfo

func (m *LogRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *LogRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *LogRequest) GetLogLevel() string {
	if m != nil {
		return m.LogLevel
	}
	return ""
}

func (m *LogRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *LogRequest) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

type LogErrorRequest struct {
	Error                string   `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	StackTrace           string   `protobuf:"bytes,2,opt,name=stackTrace,proto3" json:"stackTrace,omitempty"`
	Message              string   `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *LogErrorRequest) Reset()         { *m = LogErrorRequest{} }
func (m *LogErrorRequest) String() string { return proto.CompactTextString(m) }
func (*LogErrorRequest) ProtoMessage()    {}
func (*LogErrorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_api_1d670826c665dc28, []int{5}
}
func (m *LogErrorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogErrorRequest.Unmarshal(m, b)
}
func (m *LogErrorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogErrorRequest.Marshal(b, m, deterministic)
}
func (dst *LogErrorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogErrorRequest.Merge(dst, src)
}
func (m *LogErrorRequest) XXX_Size() int {
	return xxx_messageInfo_LogErrorRequest.Size(m)
}
func (m *LogErrorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LogErrorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LogErrorRequest proto.InternalMessageInfo

func (m *LogErrorRequest) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *LogErrorRequest) GetStackTrace() string {
	if m != nil {
		return m.StackTrace
	}
	return ""
}

func (m *LogErrorRequest) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LogResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *LogResponse) Reset()         { *m = LogResponse{} }
func (m *LogResponse) String() string { return proto.CompactTextString(m) }
func (*LogResponse) ProtoMessage()    {}
func (*LogResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_engine_api_1d670826c665dc28, []int{6}
}
func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (dst *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(dst, src)
}
func (m *LogResponse) XXX_Size() int {
	return xxx_messageInfo_LogResponse.Size(m)
}
func (m *LogResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LogResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LogResponse proto.InternalMessageInfo

func (m *LogResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LogResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*SpawnWorkflowRequest)(nil), "core.SpawnWorkflowRequest")
	proto.RegisterType((*SpawnWorkflowResponse)(nil), "core.SpawnWorkflowResponse")
	proto.RegisterType((*StepCompleteRequest)(nil), "core.StepCompleteRequest")
	proto.RegisterType((*StepCompleteResponse)(nil), "core.StepCompleteResponse")
	proto.RegisterType((*LogRequest)(nil), "core.LogRequest")
	proto.RegisterType((*LogErrorRequest)(nil), "core.LogErrorRequest")
	proto.RegisterType((*LogResponse)(nil), "core.LogResponse")
}

func init() { proto.RegisterFile("engine_api.proto", fileDescriptor_engine_api_1d670826c665dc28) }

var fileDescriptor_engine_api_1d670826c665dc28 = []byte{
	// 580 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0xa4, 0x24, 0x93, 0xb6, 0xb4, 0x4b, 0x1a, 0x8c, 0x91, 0x50, 0xe4, 0x53, 0x2e,
	0x24, 0x4a, 0xcb, 0x81, 0x03, 0x97, 0x34, 0x0a, 0x50, 0x29, 0x48, 0x91, 0x83, 0x84, 0xc4, 0x05,
	0x39, 0xce, 0xd4, 0x71, 0x63, 0x7b, 0xb7, 0xeb, 0x75, 0x03, 0x48, 0xfc, 0x05, 0x4e, 0xfc, 0x27,
	0xfe, 0x16, 0x5a, 0x7b, 0xed, 0xd8, 0xa9, 0x39, 0xf5, 0xf8, 0xde, 0xec, 0x8c, 0xdf, 0x9b, 0x0f,
	0xc3, 0x29, 0x86, 0xae, 0x17, 0xe2, 0x37, 0x9b, 0x79, 0x03, 0xc6, 0xa9, 0xa0, 0xa4, 0xee, 0x50,
	0x8e, 0xc6, 0x91, 0x43, 0x83, 0x80, 0x86, 0x29, 0x67, 0xfe, 0x82, 0xce, 0x82, 0xd9, 0xdb, 0xf0,
	0x0b, 0xe5, 0x9b, 0x1b, 0x9f, 0x6e, 0x2d, 0xbc, 0x8b, 0x31, 0x12, 0xe4, 0x12, 0xda, 0x18, 0xde,
	0x7b, 0x9c, 0x86, 0x01, 0x86, 0x42, 0xd7, 0x7a, 0x5a, 0xbf, 0x7d, 0x71, 0x36, 0x90, 0x15, 0x06,
	0xd3, 0x5d, 0xc0, 0x2a, 0xbe, 0x22, 0x06, 0x34, 0xb7, 0xaa, 0x8e, 0x7e, 0xd0, 0xd3, 0xfa, 0x2d,
	0x2b, 0xc7, 0xa4, 0x03, 0x0d, 0x2f, 0x64, 0xb1, 0xd0, 0x6b, 0x3d, 0xad, 0x7f, 0x64, 0xa5, 0xc0,
	0x1c, 0xc1, 0xf9, 0xde, 0xe7, 0x23, 0x46, 0xc3, 0x08, 0x89, 0x0e, 0x4f, 0xee, 0x62, 0x8c, 0xf1,
	0x7a, 0x95, 0x7c, 0xbb, 0x65, 0x65, 0xd0, 0xfc, 0xa3, 0xc1, 0xb3, 0x85, 0x40, 0x36, 0xa1, 0x01,
	0xf3, 0x51, 0xe0, 0xa3, 0x14, 0x77, 0xe1, 0x90, 0xc6, 0x42, 0xca, 0x3a, 0x48, 0x64, 0x29, 0x24,
	0x9d, 0xe0, 0x77, 0x4f, 0x4c, 0xe8, 0x0a, 0x13, 0xc1, 0x35, 0x2b, 0xc7, 0xd2, 0x09, 0x72, 0x4e,
	0xb9, 0x5e, 0x4f, 0x84, 0xa5, 0xc0, 0xec, 0x42, 0xa7, 0xac, 0x2a, 0x35, 0x62, 0xfe, 0xd6, 0x00,
	0x66, 0xd4, 0xcd, 0x54, 0x76, 0xa0, 0x71, 0x4b, 0x97, 0xb9, 0xab, 0x14, 0x48, 0xb7, 0x01, 0x46,
	0x91, 0xed, 0xa2, 0xea, 0x5b, 0x06, 0xa5, 0x10, 0x9f, 0xba, 0x33, 0xbc, 0x47, 0x3f, 0x11, 0xd2,
	0xb2, 0x72, 0x2c, 0xb3, 0x18, 0xa7, 0xb7, 0xe8, 0x08, 0x25, 0x25, 0x83, 0xa5, 0x41, 0x34, 0xca,
	0x83, 0x30, 0x6d, 0x78, 0x3a, 0xa3, 0xee, 0x54, 0x8a, 0x2e, 0x88, 0x4a, 0x1d, 0x69, 0x05, 0x47,
	0xe4, 0x15, 0x40, 0x24, 0x6c, 0x67, 0xf3, 0x99, 0xdb, 0x4e, 0xa6, 0xab, 0xc0, 0x14, 0x45, 0xd7,
	0x4a, 0xa2, 0xcd, 0x31, 0xb4, 0x13, 0xcb, 0xbb, 0x59, 0x46, 0xb1, 0xe3, 0x60, 0x14, 0x25, 0x1f,
	0x68, 0x5a, 0x19, 0xfc, 0xbf, 0xef, 0x8b, 0xbf, 0x75, 0x68, 0x4d, 0x93, 0x05, 0x1e, 0xcf, 0xaf,
	0xc9, 0x5b, 0x38, 0x19, 0xc7, 0x62, 0x4d, 0xb9, 0xf7, 0x13, 0x57, 0x73, 0x2f, 0x74, 0x09, 0x51,
	0x83, 0x0d, 0x98, 0xf8, 0xf1, 0x29, 0xcd, 0x31, 0x2a, 0x38, 0x32, 0x85, 0x93, 0x0f, 0x28, 0xe4,
	0x64, 0xe6, 0xb6, 0xb3, 0x91, 0x8c, 0x9e, 0xbe, 0x2a, 0x50, 0xaa, 0x0d, 0xc6, 0x8b, 0x8a, 0x88,
	0xb2, 0xf0, 0x11, 0x8e, 0x4b, 0x7b, 0x4a, 0x0c, 0xf5, 0xb6, 0xe2, 0x76, 0x8c, 0x97, 0x95, 0x31,
	0x55, 0xe9, 0x1d, 0xb4, 0x26, 0xb6, 0xb3, 0xc6, 0x79, 0x1c, 0xad, 0x49, 0x37, 0x7d, 0x99, 0x13,
	0x59, 0x85, 0xe7, 0x0f, 0xf8, 0x07, 0xd9, 0xbe, 0xbf, 0x97, 0xed, 0xfb, 0xd5, 0xd9, 0x92, 0x57,
	0xd9, 0x57, 0xd0, 0x4e, 0xc8, 0x05, 0xda, 0xdc, 0x59, 0x67, 0x9d, 0x28, 0x50, 0x7b, 0x9d, 0x28,
	0x45, 0x54, 0x8d, 0xf7, 0x70, 0x5c, 0xdc, 0xf3, 0x15, 0x29, 0x74, 0x6d, 0xef, 0x24, 0x0d, 0xa3,
	0x2a, 0xa4, 0xea, 0x8c, 0x92, 0xb3, 0xc8, 0xc6, 0x74, 0x9a, 0xbe, 0xdc, 0x1d, 0x8a, 0x71, 0x56,
	0x60, 0x54, 0xca, 0x1b, 0x68, 0x66, 0x9b, 0x4b, 0xce, 0xf3, 0x70, 0x71, 0x93, 0x2b, 0xb2, 0xae,
	0x46, 0x5f, 0x87, 0xae, 0x27, 0xd6, 0xf1, 0x72, 0xe0, 0xd0, 0x60, 0x68, 0x33, 0x26, 0x38, 0x62,
	0x44, 0x6f, 0xc4, 0xd6, 0xe6, 0x38, 0x74, 0xe9, 0xeb, 0xec, 0x32, 0x86, 0x6c, 0xe3, 0x0e, 0x65,
	0xfa, 0xf2, 0x30, 0xf9, 0x37, 0x5e, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x97, 0xad, 0xa3, 0xde,
	0x43, 0x05, 0x00, 0x00,
}
