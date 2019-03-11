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
	return fileDescriptor_engine_api_e1b180795c825aaa, []int{0}
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
	return fileDescriptor_engine_api_e1b180795c825aaa, []int{1}
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

func init() {
	proto.RegisterType((*SpawnWorkflowRequest)(nil), "core.SpawnWorkflowRequest")
	proto.RegisterType((*SpawnWorkflowResponse)(nil), "core.SpawnWorkflowResponse")
}

func init() { proto.RegisterFile("engine_api.proto", fileDescriptor_engine_api_e1b180795c825aaa) }

var fileDescriptor_engine_api_e1b180795c825aaa = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4d, 0x4f, 0xfa, 0x40,
	0x10, 0xc6, 0x53, 0xfe, 0x7f, 0x5f, 0xba, 0x20, 0xd1, 0x0d, 0x6a, 0xad, 0x17, 0xc2, 0x89, 0x8b,
	0x6d, 0x80, 0x8b, 0x07, 0x2f, 0x68, 0x88, 0x72, 0x30, 0x69, 0xca, 0xc1, 0xc4, 0x8b, 0x29, 0x65,
	0x68, 0x37, 0xb4, 0xbb, 0xcb, 0xbe, 0xd8, 0x68, 0xe2, 0x97, 0xf4, 0x13, 0x99, 0xd2, 0x82, 0x08,
	0x3d, 0xce, 0x6f, 0x9e, 0xe7, 0xd9, 0x99, 0x69, 0xd1, 0x29, 0xd0, 0x88, 0x50, 0x78, 0x0b, 0x38,
	0x71, 0xb8, 0x60, 0x8a, 0xe1, 0xff, 0x21, 0x13, 0x60, 0x37, 0x42, 0x96, 0xa6, 0x8c, 0x16, 0xac,
	0xf3, 0x85, 0x5a, 0x13, 0x1e, 0x64, 0xf4, 0x85, 0x89, 0xc5, 0x3c, 0x61, 0x99, 0x0f, 0x4b, 0x0d,
	0x52, 0xe1, 0x01, 0xaa, 0x03, 0x7d, 0x27, 0x82, 0xd1, 0x14, 0xa8, 0xb2, 0x8c, 0xb6, 0xd1, 0xad,
	0xf7, 0xcf, 0x9c, 0x3c, 0xc1, 0x19, 0xfd, 0x36, 0xfc, 0x6d, 0x15, 0xb6, 0xd1, 0x71, 0x56, 0xe6,
	0x58, 0xb5, 0xb6, 0xd1, 0x35, 0xfd, 0x4d, 0x8d, 0x5b, 0xe8, 0x80, 0x50, 0xae, 0x95, 0xf5, 0xaf,
	0x6d, 0x74, 0x1b, 0x7e, 0x51, 0x74, 0x7a, 0xe8, 0x7c, 0xe7, 0x79, 0xc9, 0x19, 0x95, 0x80, 0x2d,
	0x74, 0xb4, 0xd4, 0xa0, 0x61, 0x3c, 0x5b, 0xbd, 0x6d, 0xfa, 0xeb, 0xb2, 0xff, 0x5d, 0x43, 0xe6,
	0x68, 0xb5, 0xda, 0xd0, 0x1b, 0xe3, 0x5b, 0xd4, 0x1c, 0x6a, 0x15, 0x33, 0x41, 0x3e, 0x61, 0xe6,
	0x11, 0x1a, 0x61, 0x5c, 0x0e, 0x99, 0x72, 0xf5, 0xf1, 0x0c, 0x52, 0x06, 0x11, 0xd8, 0x15, 0x0c,
	0x8f, 0x50, 0xf3, 0x11, 0xd4, 0x44, 0x01, 0xf7, 0x82, 0x70, 0x91, 0x13, 0xab, 0x50, 0x6d, 0xa1,
	0xf2, 0x1a, 0xf6, 0x55, 0x45, 0xa7, 0x1c, 0xf4, 0x09, 0x9d, 0xfc, 0xd9, 0x00, 0xdb, 0xa5, 0xb6,
	0xe2, 0xaa, 0xf6, 0x75, 0x65, 0xaf, 0x4c, 0xba, 0x43, 0xe6, 0x43, 0x10, 0xc6, 0xe0, 0x69, 0x19,
	0xe3, 0x8b, 0x42, 0xb9, 0x01, 0xeb, 0x84, 0xcb, 0x3d, 0xbe, 0xe7, 0x4e, 0x92, 0x1d, 0x77, 0x92,
	0x54, 0xbb, 0x73, 0x5e, 0xb8, 0xef, 0x7b, 0xaf, 0x6e, 0x44, 0x54, 0xac, 0xa7, 0x4e, 0xc8, 0x52,
	0x37, 0xe0, 0x5c, 0x09, 0x00, 0xc9, 0xe6, 0x2a, 0x0b, 0x04, 0xb8, 0x11, 0xbb, 0x59, 0x7f, 0x47,
	0x97, 0x2f, 0x22, 0x37, 0x0f, 0x99, 0x1e, 0xae, 0x7e, 0xa0, 0xc1, 0x4f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xc1, 0xcb, 0xeb, 0x8d, 0x68, 0x02, 0x00, 0x00,
}
