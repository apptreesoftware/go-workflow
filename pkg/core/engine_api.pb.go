// Code generated by protoc-gen-go. DO NOT EDIT.
// source: engine_api.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
	return fileDescriptor_4a508ff9519b0c05, []int{0}
}

func (m *SpawnWorkflowRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnWorkflowRequest.Unmarshal(m, b)
}
func (m *SpawnWorkflowRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnWorkflowRequest.Marshal(b, m, deterministic)
}
func (m *SpawnWorkflowRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnWorkflowRequest.Merge(m, src)
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
	return fileDescriptor_4a508ff9519b0c05, []int{1}
}

func (m *SpawnWorkflowResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnWorkflowResponse.Unmarshal(m, b)
}
func (m *SpawnWorkflowResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnWorkflowResponse.Marshal(b, m, deterministic)
}
func (m *SpawnWorkflowResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnWorkflowResponse.Merge(m, src)
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
	return fileDescriptor_4a508ff9519b0c05, []int{2}
}

func (m *StepCompleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepCompleteRequest.Unmarshal(m, b)
}
func (m *StepCompleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepCompleteRequest.Marshal(b, m, deterministic)
}
func (m *StepCompleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepCompleteRequest.Merge(m, src)
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
	return fileDescriptor_4a508ff9519b0c05, []int{3}
}

func (m *StepCompleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepCompleteResponse.Unmarshal(m, b)
}
func (m *StepCompleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepCompleteResponse.Marshal(b, m, deterministic)
}
func (m *StepCompleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepCompleteResponse.Merge(m, src)
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
	return fileDescriptor_4a508ff9519b0c05, []int{4}
}

func (m *LogRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogRequest.Unmarshal(m, b)
}
func (m *LogRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogRequest.Marshal(b, m, deterministic)
}
func (m *LogRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogRequest.Merge(m, src)
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
	return fileDescriptor_4a508ff9519b0c05, []int{5}
}

func (m *LogResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogResponse.Unmarshal(m, b)
}
func (m *LogResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogResponse.Marshal(b, m, deterministic)
}
func (m *LogResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogResponse.Merge(m, src)
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
	proto.RegisterType((*LogResponse)(nil), "core.LogResponse")
}

func init() { proto.RegisterFile("engine_api.proto", fileDescriptor_4a508ff9519b0c05) }

var fileDescriptor_4a508ff9519b0c05 = []byte{
	// 515 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x41, 0x6f, 0xd3, 0x4c,
	0x10, 0x95, 0x9b, 0xa6, 0x5f, 0x33, 0x69, 0xab, 0x76, 0xbf, 0x10, 0x8c, 0xb9, 0x44, 0x3e, 0xe5,
	0x42, 0xa2, 0xb4, 0x17, 0x0e, 0x5c, 0x42, 0x14, 0xa0, 0x52, 0x91, 0x22, 0xf7, 0x80, 0xc4, 0x05,
	0x39, 0xce, 0xd4, 0x71, 0x63, 0x7b, 0xb6, 0xeb, 0x75, 0x03, 0x48, 0xfc, 0x05, 0x4e, 0xfc, 0x2d,
	0xfe, 0x13, 0x5a, 0xef, 0x3a, 0x71, 0x52, 0x73, 0xe2, 0xf8, 0xde, 0xcc, 0x1b, 0xbf, 0x99, 0x7d,
	0x32, 0x9c, 0x63, 0x1a, 0x46, 0x29, 0x7e, 0xf1, 0x79, 0x34, 0xe0, 0x82, 0x24, 0xb1, 0xc3, 0x80,
	0x04, 0x3a, 0x27, 0x01, 0x25, 0x09, 0xa5, 0x9a, 0x73, 0x7f, 0x40, 0xe7, 0x96, 0xfb, 0xeb, 0xf4,
	0x13, 0x89, 0xd5, 0x5d, 0x4c, 0x6b, 0x0f, 0x1f, 0x72, 0xcc, 0x24, 0xbb, 0x82, 0x36, 0xa6, 0x8f,
	0x91, 0xa0, 0x34, 0xc1, 0x54, 0xda, 0x56, 0xcf, 0xea, 0xb7, 0x2f, 0x2f, 0x06, 0x6a, 0xc2, 0x60,
	0xba, 0x2d, 0x78, 0xd5, 0x2e, 0xe6, 0xc0, 0xf1, 0xda, 0xcc, 0xb1, 0x0f, 0x7a, 0x56, 0xbf, 0xe5,
	0x6d, 0x30, 0xeb, 0x40, 0x33, 0x4a, 0x79, 0x2e, 0xed, 0x46, 0xcf, 0xea, 0x9f, 0x78, 0x1a, 0xb8,
	0x23, 0x78, 0xb6, 0xf7, 0xf9, 0x8c, 0x53, 0x9a, 0x21, 0xb3, 0xe1, 0xbf, 0x87, 0x1c, 0x73, 0xbc,
	0x5e, 0x14, 0xdf, 0x6e, 0x79, 0x25, 0x74, 0x7f, 0x59, 0xf0, 0xff, 0xad, 0x44, 0x3e, 0xa1, 0x84,
	0xc7, 0x28, 0xf1, 0x9f, 0x1c, 0x77, 0xe1, 0x88, 0x72, 0xa9, 0x6c, 0x1d, 0x14, 0xb6, 0x0c, 0x52,
	0x9b, 0xe0, 0xd7, 0x48, 0x4e, 0x68, 0x81, 0x85, 0xe1, 0x86, 0xb7, 0xc1, 0x6a, 0x13, 0x14, 0x82,
	0x84, 0x7d, 0x58, 0x18, 0xd3, 0xc0, 0xed, 0x42, 0x67, 0xd7, 0x95, 0x5e, 0xc4, 0xfd, 0x69, 0x01,
	0xdc, 0x50, 0x58, 0xba, 0xec, 0x40, 0xf3, 0x9e, 0xe6, 0x9b, 0xad, 0x34, 0x50, 0xdb, 0x26, 0x98,
	0x65, 0x7e, 0x88, 0xe6, 0x6e, 0x25, 0x54, 0x46, 0x62, 0x0a, 0x6f, 0xf0, 0x11, 0xe3, 0xc2, 0x48,
	0xcb, 0xdb, 0x60, 0xa5, 0xe2, 0x82, 0xee, 0x31, 0x90, 0xc6, 0x4a, 0x09, 0x77, 0x1e, 0xa2, 0xb9,
	0xfb, 0x10, 0xee, 0x18, 0xda, 0x85, 0x9f, 0xed, 0xa1, 0xb3, 0x3c, 0x08, 0x30, 0xcb, 0x0a, 0x4b,
	0xc7, 0x5e, 0x09, 0xff, 0x6e, 0xea, 0xf2, 0x77, 0x03, 0x5a, 0xd3, 0x22, 0x5d, 0xe3, 0xd9, 0x35,
	0x7b, 0x0d, 0x67, 0xe3, 0x5c, 0x2e, 0x49, 0x44, 0xdf, 0x71, 0x31, 0x8b, 0xd2, 0x90, 0x31, 0x73,
	0xf5, 0x84, 0xcb, 0x6f, 0x1f, 0xb5, 0xc6, 0xa9, 0xe1, 0xd8, 0x14, 0xce, 0xde, 0xa3, 0x54, 0x67,
	0x9b, 0xf9, 0xc1, 0x4a, 0x31, 0xb6, 0xee, 0xaa, 0x50, 0xe6, 0x70, 0xce, 0x8b, 0x9a, 0x8a, 0x59,
	0xe1, 0x03, 0x9c, 0xee, 0x84, 0x88, 0x39, 0xa6, 0xb7, 0x26, 0xd8, 0xce, 0xcb, 0xda, 0x9a, 0x99,
	0xf4, 0x06, 0x5a, 0x13, 0x3f, 0x58, 0xe2, 0x2c, 0xcf, 0x96, 0xac, 0xab, 0x3b, 0x37, 0x44, 0x39,
	0xe1, 0xf9, 0x13, 0xfe, 0x89, 0x3a, 0x8e, 0xf7, 0xd4, 0x71, 0x5c, 0xaf, 0x56, 0xbc, 0x51, 0xbf,
	0x83, 0xd3, 0x6a, 0x80, 0x16, 0xac, 0xb2, 0xf1, 0x5e, 0xd6, 0x1d, 0xa7, 0xae, 0x64, 0xe6, 0x8c,
	0x8a, 0xbc, 0x95, 0x27, 0x3e, 0xd7, 0x9d, 0xdb, 0x04, 0x3a, 0x17, 0x15, 0x46, 0x4b, 0xde, 0x8e,
	0x3e, 0x0f, 0xc3, 0x48, 0x2e, 0xf3, 0xf9, 0x20, 0xa0, 0x64, 0xe8, 0x73, 0x2e, 0x05, 0x62, 0x46,
	0x77, 0x72, 0xed, 0x0b, 0x1c, 0x86, 0xf4, 0xaa, 0x0c, 0xcf, 0x90, 0xaf, 0xc2, 0xa1, 0x92, 0xcf,
	0x8f, 0x8a, 0xdf, 0xc7, 0xd5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x60, 0xc3, 0xf0, 0x66,
	0x04, 0x00, 0x00,
}
