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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

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

func init() {
	proto.RegisterType((*SpawnWorkflowRequest)(nil), "core.SpawnWorkflowRequest")
	proto.RegisterType((*SpawnWorkflowResponse)(nil), "core.SpawnWorkflowResponse")
	proto.RegisterType((*StepCompleteRequest)(nil), "core.StepCompleteRequest")
	proto.RegisterType((*StepCompleteResponse)(nil), "core.StepCompleteResponse")
}

func init() { proto.RegisterFile("engine_api.proto", fileDescriptor_4a508ff9519b0c05) }

var fileDescriptor_4a508ff9519b0c05 = []byte{
	// 423 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x95, 0x9b, 0x52, 0xc8, 0xf4, 0x43, 0xb0, 0x84, 0x60, 0x96, 0x4b, 0x94, 0x53, 0x2e, 0xc4,
	0x6a, 0x7b, 0xe1, 0xc0, 0xa5, 0x44, 0x01, 0x7a, 0x40, 0x8a, 0xdc, 0x03, 0x12, 0x17, 0xe4, 0x3a,
	0x53, 0x67, 0x15, 0x7b, 0x67, 0xbb, 0xde, 0xc5, 0x80, 0xc4, 0xbf, 0xe0, 0xc4, 0xaf, 0x45, 0x6b,
	0x6f, 0x82, 0x9b, 0xfa, 0xd6, 0xe3, 0x7b, 0x33, 0xf3, 0xe6, 0xcd, 0xec, 0x2c, 0x3c, 0x45, 0x99,
	0x09, 0x89, 0xdf, 0x12, 0x25, 0xa6, 0x4a, 0x93, 0x21, 0xb6, 0x9f, 0x92, 0x46, 0x7e, 0x94, 0x52,
	0x51, 0x90, 0x6c, 0xb8, 0xf1, 0x6f, 0x18, 0x5c, 0xa9, 0xa4, 0x92, 0x5f, 0x48, 0xaf, 0x6f, 0x72,
	0xaa, 0x62, 0xbc, 0xb5, 0x58, 0x1a, 0x76, 0x0e, 0x87, 0x28, 0xbf, 0x0b, 0x4d, 0xb2, 0x40, 0x69,
	0xc2, 0x60, 0x14, 0x4c, 0x0e, 0xcf, 0x9e, 0x4d, 0x9d, 0xc2, 0x74, 0xfe, 0x3f, 0x10, 0xb7, 0xb3,
	0x18, 0x87, 0x27, 0x95, 0xd7, 0x09, 0xf7, 0x46, 0xc1, 0xa4, 0x1f, 0x6f, 0x31, 0x1b, 0xc0, 0x23,
	0x21, 0x95, 0x35, 0x61, 0x6f, 0x14, 0x4c, 0x8e, 0xe2, 0x06, 0x8c, 0x4f, 0xe1, 0xc5, 0x4e, 0xfb,
	0x52, 0x91, 0x2c, 0x91, 0x85, 0xf0, 0xf8, 0xd6, 0xa2, 0xc5, 0xcb, 0x65, 0xdd, 0xbb, 0x1f, 0x6f,
	0xe0, 0xf8, 0x4f, 0x00, 0xcf, 0xaf, 0x0c, 0xaa, 0x19, 0x15, 0x2a, 0x47, 0x83, 0x0f, 0x72, 0x3c,
	0x84, 0x03, 0xb2, 0xc6, 0xd9, 0xda, 0xab, 0x6d, 0x79, 0xe4, 0x26, 0xc1, 0x1f, 0xc2, 0xcc, 0x68,
	0x89, 0xb5, 0xe1, 0x5e, 0xbc, 0xc5, 0x6e, 0x12, 0xd4, 0x9a, 0x74, 0xb8, 0x5f, 0x1b, 0x6b, 0xc0,
	0x78, 0x08, 0x83, 0xbb, 0xae, 0x9a, 0x41, 0xce, 0xfe, 0xf6, 0xa0, 0x3f, 0xaf, 0x5f, 0xe2, 0x62,
	0x71, 0xc9, 0xde, 0xc2, 0xc9, 0x85, 0x35, 0x2b, 0xd2, 0xe2, 0x17, 0x2e, 0x17, 0x42, 0x66, 0x8c,
	0x79, 0x87, 0x85, 0x32, 0x3f, 0x3f, 0x63, 0x59, 0x26, 0x19, 0xf2, 0x0e, 0x8e, 0xcd, 0xe1, 0xe4,
	0x23, 0x1a, 0xd7, 0x62, 0x91, 0xa4, 0x6b, 0xc7, 0x84, 0x4d, 0x56, 0x8b, 0xf2, 0xab, 0xe0, 0xaf,
	0x3a, 0x22, 0x7e, 0xaf, 0x9f, 0xe0, 0xf8, 0xce, 0xc2, 0x19, 0xf7, 0xb9, 0x1d, 0x47, 0xc0, 0x5f,
	0x77, 0xc6, 0xbc, 0xd2, 0x3b, 0xe8, 0xcf, 0x92, 0x74, 0x85, 0x0b, 0x5b, 0xae, 0xd8, 0xb0, 0xc9,
	0xdc, 0x12, 0x1b, 0x85, 0x97, 0xf7, 0xf8, 0x7b, 0xd5, 0x79, 0xbe, 0x53, 0x9d, 0xe7, 0xdd, 0xd5,
	0x8e, 0xf7, 0xd5, 0x1f, 0xe0, 0xb8, 0xbd, 0xec, 0x25, 0x6b, 0x4d, 0xbc, 0x73, 0x17, 0x9c, 0x77,
	0x85, 0x1a, 0x9d, 0xf7, 0xa7, 0x5f, 0xa3, 0x4c, 0x98, 0x95, 0xbd, 0x9e, 0xa6, 0x54, 0x44, 0x89,
	0x52, 0x46, 0x23, 0x96, 0x74, 0x63, 0xaa, 0x44, 0x63, 0x94, 0xd1, 0x9b, 0xcd, 0xf9, 0x46, 0x6a,
	0x9d, 0x45, 0x4e, 0xe7, 0xfa, 0xa0, 0xfe, 0x37, 0xe7, 0xff, 0x02, 0x00, 0x00, 0xff, 0xff, 0x68,
	0x9f, 0xda, 0xf2, 0x5f, 0x03, 0x00, 0x00,
}
