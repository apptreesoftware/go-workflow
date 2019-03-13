// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote_engine.proto

package core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type StepAvailableRequest struct {
	StepName             string   `protobuf:"bytes,1,opt,name=stepName,proto3" json:"stepName,omitempty"`
	StepVersion          string   `protobuf:"bytes,2,opt,name=stepVersion,proto3" json:"stepVersion,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StepAvailableRequest) Reset()         { *m = StepAvailableRequest{} }
func (m *StepAvailableRequest) String() string { return proto.CompactTextString(m) }
func (*StepAvailableRequest) ProtoMessage()    {}
func (*StepAvailableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{0}
}

func (m *StepAvailableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepAvailableRequest.Unmarshal(m, b)
}
func (m *StepAvailableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepAvailableRequest.Marshal(b, m, deterministic)
}
func (m *StepAvailableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepAvailableRequest.Merge(m, src)
}
func (m *StepAvailableRequest) XXX_Size() int {
	return xxx_messageInfo_StepAvailableRequest.Size(m)
}
func (m *StepAvailableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StepAvailableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StepAvailableRequest proto.InternalMessageInfo

func (m *StepAvailableRequest) GetStepName() string {
	if m != nil {
		return m.StepName
	}
	return ""
}

func (m *StepAvailableRequest) GetStepVersion() string {
	if m != nil {
		return m.StepVersion
	}
	return ""
}

type StepAvailableResponse struct {
	Available            bool     `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StepAvailableResponse) Reset()         { *m = StepAvailableResponse{} }
func (m *StepAvailableResponse) String() string { return proto.CompactTextString(m) }
func (*StepAvailableResponse) ProtoMessage()    {}
func (*StepAvailableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{1}
}

func (m *StepAvailableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepAvailableResponse.Unmarshal(m, b)
}
func (m *StepAvailableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepAvailableResponse.Marshal(b, m, deterministic)
}
func (m *StepAvailableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepAvailableResponse.Merge(m, src)
}
func (m *StepAvailableResponse) XXX_Size() int {
	return xxx_messageInfo_StepAvailableResponse.Size(m)
}
func (m *StepAvailableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StepAvailableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StepAvailableResponse proto.InternalMessageInfo

func (m *StepAvailableResponse) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

type PingRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingRequest) Reset()         { *m = PingRequest{} }
func (m *PingRequest) String() string { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()    {}
func (*PingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{2}
}

func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (m *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(m, src)
}
func (m *PingRequest) XXX_Size() int {
	return xxx_messageInfo_PingRequest.Size(m)
}
func (m *PingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PingRequest proto.InternalMessageInfo

type PingResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{3}
}

func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (m *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(m, src)
}
func (m *PingResponse) XXX_Size() int {
	return xxx_messageInfo_PingResponse.Size(m)
}
func (m *PingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PingResponse proto.InternalMessageInfo

func (m *PingResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type LoadPackageRequest struct {
	Pkg                  string   `protobuf:"bytes,1,opt,name=pkg,proto3" json:"pkg,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadPackageRequest) Reset()         { *m = LoadPackageRequest{} }
func (m *LoadPackageRequest) String() string { return proto.CompactTextString(m) }
func (*LoadPackageRequest) ProtoMessage()    {}
func (*LoadPackageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{4}
}

func (m *LoadPackageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPackageRequest.Unmarshal(m, b)
}
func (m *LoadPackageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPackageRequest.Marshal(b, m, deterministic)
}
func (m *LoadPackageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPackageRequest.Merge(m, src)
}
func (m *LoadPackageRequest) XXX_Size() int {
	return xxx_messageInfo_LoadPackageRequest.Size(m)
}
func (m *LoadPackageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadPackageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadPackageRequest proto.InternalMessageInfo

func (m *LoadPackageRequest) GetPkg() string {
	if m != nil {
		return m.Pkg
	}
	return ""
}

func (m *LoadPackageRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type LoadPackageResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoadPackageResponse) Reset()         { *m = LoadPackageResponse{} }
func (m *LoadPackageResponse) String() string { return proto.CompactTextString(m) }
func (*LoadPackageResponse) ProtoMessage()    {}
func (*LoadPackageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{5}
}

func (m *LoadPackageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPackageResponse.Unmarshal(m, b)
}
func (m *LoadPackageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPackageResponse.Marshal(b, m, deterministic)
}
func (m *LoadPackageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPackageResponse.Merge(m, src)
}
func (m *LoadPackageResponse) XXX_Size() int {
	return xxx_messageInfo_LoadPackageResponse.Size(m)
}
func (m *LoadPackageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadPackageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadPackageResponse proto.InternalMessageInfo

func (m *LoadPackageResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *LoadPackageResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type RunStepRequest struct {
	Environment          *Environment `protobuf:"bytes,1,opt,name=environment,proto3" json:"environment,omitempty"`
	Input                []byte       `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	StepConfig           []byte       `protobuf:"bytes,3,opt,name=stepConfig,proto3" json:"stepConfig,omitempty"`
	StepInstanceId       string       `protobuf:"bytes,4,opt,name=stepInstanceId,proto3" json:"stepInstanceId,omitempty"`
	JobId                string       `protobuf:"bytes,5,opt,name=jobId,proto3" json:"jobId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *RunStepRequest) Reset()         { *m = RunStepRequest{} }
func (m *RunStepRequest) String() string { return proto.CompactTextString(m) }
func (*RunStepRequest) ProtoMessage()    {}
func (*RunStepRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{6}
}

func (m *RunStepRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunStepRequest.Unmarshal(m, b)
}
func (m *RunStepRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunStepRequest.Marshal(b, m, deterministic)
}
func (m *RunStepRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunStepRequest.Merge(m, src)
}
func (m *RunStepRequest) XXX_Size() int {
	return xxx_messageInfo_RunStepRequest.Size(m)
}
func (m *RunStepRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RunStepRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RunStepRequest proto.InternalMessageInfo

func (m *RunStepRequest) GetEnvironment() *Environment {
	if m != nil {
		return m.Environment
	}
	return nil
}

func (m *RunStepRequest) GetInput() []byte {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *RunStepRequest) GetStepConfig() []byte {
	if m != nil {
		return m.StepConfig
	}
	return nil
}

func (m *RunStepRequest) GetStepInstanceId() string {
	if m != nil {
		return m.StepInstanceId
	}
	return ""
}

func (m *RunStepRequest) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

type RunStepResponse struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	ErrorMessage         string   `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RunStepResponse) Reset()         { *m = RunStepResponse{} }
func (m *RunStepResponse) String() string { return proto.CompactTextString(m) }
func (*RunStepResponse) ProtoMessage()    {}
func (*RunStepResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{7}
}

func (m *RunStepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunStepResponse.Unmarshal(m, b)
}
func (m *RunStepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunStepResponse.Marshal(b, m, deterministic)
}
func (m *RunStepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunStepResponse.Merge(m, src)
}
func (m *RunStepResponse) XXX_Size() int {
	return xxx_messageInfo_RunStepResponse.Size(m)
}
func (m *RunStepResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RunStepResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RunStepResponse proto.InternalMessageInfo

func (m *RunStepResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

func (m *RunStepResponse) GetErrorMessage() string {
	if m != nil {
		return m.ErrorMessage
	}
	return ""
}

type AuthorizationRequest struct {
	Token                []byte   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	HostEngine           string   `protobuf:"bytes,2,opt,name=hostEngine,proto3" json:"hostEngine,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{8}
}

func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (m *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(m, src)
}
func (m *AuthorizationRequest) XXX_Size() int {
	return xxx_messageInfo_AuthorizationRequest.Size(m)
}
func (m *AuthorizationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationRequest proto.InternalMessageInfo

func (m *AuthorizationRequest) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *AuthorizationRequest) GetHostEngine() string {
	if m != nil {
		return m.HostEngine
	}
	return ""
}

type AuthorizationResponse struct {
	Authorized           bool     `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizationResponse) Reset()         { *m = AuthorizationResponse{} }
func (m *AuthorizationResponse) String() string { return proto.CompactTextString(m) }
func (*AuthorizationResponse) ProtoMessage()    {}
func (*AuthorizationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{9}
}

func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (m *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(m, src)
}
func (m *AuthorizationResponse) XXX_Size() int {
	return xxx_messageInfo_AuthorizationResponse.Size(m)
}
func (m *AuthorizationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthorizationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AuthorizationResponse proto.InternalMessageInfo

func (m *AuthorizationResponse) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func (m *AuthorizationResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type CertificateUpdateRequest struct {
	Key                  []byte   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Cert                 []byte   `protobuf:"bytes,2,opt,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CertificateUpdateRequest) Reset()         { *m = CertificateUpdateRequest{} }
func (m *CertificateUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*CertificateUpdateRequest) ProtoMessage()    {}
func (*CertificateUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{10}
}

func (m *CertificateUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateUpdateRequest.Unmarshal(m, b)
}
func (m *CertificateUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateUpdateRequest.Marshal(b, m, deterministic)
}
func (m *CertificateUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateUpdateRequest.Merge(m, src)
}
func (m *CertificateUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_CertificateUpdateRequest.Size(m)
}
func (m *CertificateUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateUpdateRequest proto.InternalMessageInfo

func (m *CertificateUpdateRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *CertificateUpdateRequest) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7480ac5bd0b3fb, []int{11}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*StepAvailableRequest)(nil), "core.StepAvailableRequest")
	proto.RegisterType((*StepAvailableResponse)(nil), "core.StepAvailableResponse")
	proto.RegisterType((*PingRequest)(nil), "core.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "core.PingResponse")
	proto.RegisterType((*LoadPackageRequest)(nil), "core.LoadPackageRequest")
	proto.RegisterType((*LoadPackageResponse)(nil), "core.LoadPackageResponse")
	proto.RegisterType((*RunStepRequest)(nil), "core.RunStepRequest")
	proto.RegisterType((*RunStepResponse)(nil), "core.RunStepResponse")
	proto.RegisterType((*AuthorizationRequest)(nil), "core.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "core.AuthorizationResponse")
	proto.RegisterType((*CertificateUpdateRequest)(nil), "core.CertificateUpdateRequest")
	proto.RegisterType((*Empty)(nil), "core.Empty")
}

func init() { proto.RegisterFile("remote_engine.proto", fileDescriptor_2d7480ac5bd0b3fb) }

var fileDescriptor_2d7480ac5bd0b3fb = []byte{
	// 612 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x53, 0x13, 0x31,
	0x14, 0x1d, 0xa0, 0x15, 0xb8, 0xad, 0x20, 0xa1, 0xcc, 0xac, 0xab, 0xc3, 0x30, 0xfb, 0xe0, 0xf0,
	0x42, 0x3b, 0xc2, 0xe8, 0x33, 0x1f, 0xf2, 0x50, 0x07, 0x1d, 0x58, 0xd4, 0x07, 0x5f, 0x9c, 0x34,
	0xbd, 0x5d, 0x62, 0xd9, 0x24, 0x26, 0x69, 0x19, 0xfc, 0x1d, 0xfe, 0x19, 0xff, 0x9d, 0x93, 0x6c,
	0xb6, 0xdd, 0xad, 0xd5, 0xb7, 0x9c, 0x73, 0x6f, 0xce, 0x9e, 0xdc, 0x8f, 0x85, 0x5d, 0x8d, 0xb9,
	0xb4, 0xf8, 0x0d, 0x45, 0xc6, 0x05, 0x76, 0x95, 0x96, 0x56, 0x92, 0x06, 0x93, 0x1a, 0xe3, 0x36,
	0x93, 0x79, 0x2e, 0x45, 0xc1, 0x25, 0x9f, 0xa0, 0x73, 0x6b, 0x51, 0x9d, 0x4d, 0x29, 0xbf, 0xa7,
	0x83, 0x7b, 0x4c, 0xf1, 0xc7, 0x04, 0x8d, 0x25, 0x31, 0x6c, 0x18, 0x8b, 0xea, 0x23, 0xcd, 0x31,
	0x5a, 0x39, 0x58, 0x39, 0xdc, 0x4c, 0x67, 0x98, 0x1c, 0x40, 0xcb, 0x9d, 0xbf, 0xa0, 0x36, 0x5c,
	0x8a, 0x68, 0xd5, 0x87, 0xab, 0x54, 0xf2, 0x06, 0xf6, 0x16, 0x54, 0x8d, 0x92, 0xc2, 0x20, 0x79,
	0x09, 0x9b, 0xb4, 0x24, 0xbd, 0xee, 0x46, 0x3a, 0x27, 0x92, 0xa7, 0xd0, 0xba, 0xe6, 0x22, 0x0b,
	0x1e, 0x92, 0x43, 0x68, 0x17, 0x30, 0x5c, 0x8e, 0x60, 0x3d, 0x47, 0x63, 0x68, 0x56, 0x5a, 0x2a,
	0x61, 0x72, 0x0a, 0xe4, 0x4a, 0xd2, 0xe1, 0x35, 0x65, 0x63, 0x9a, 0xcd, 0xde, 0xf0, 0x0c, 0xd6,
	0xd4, 0x38, 0x0b, 0xb9, 0xee, 0xe8, 0x14, 0xa6, 0x35, 0xd7, 0x25, 0x4c, 0xfa, 0xb0, 0x5b, 0x53,
	0x98, 0x7f, 0xd2, 0x4c, 0x18, 0x43, 0x63, 0x82, 0xdb, 0x12, 0x56, 0xcd, 0xac, 0xd6, 0xcd, 0xfc,
	0x5e, 0x81, 0xad, 0x74, 0x22, 0x5c, 0x01, 0x4a, 0x27, 0x27, 0xd0, 0x42, 0x31, 0xe5, 0x5a, 0x8a,
	0x1c, 0x85, 0xf5, 0x52, 0xad, 0xe3, 0x9d, 0xae, 0xeb, 0x47, 0xf7, 0x72, 0x1e, 0x48, 0xab, 0x59,
	0xa4, 0x03, 0x4d, 0x2e, 0xd4, 0xc4, 0x7a, 0xfd, 0x76, 0x5a, 0x00, 0xb2, 0x0f, 0xe0, 0x2a, 0x7d,
	0x21, 0xc5, 0x88, 0x67, 0xd1, 0x9a, 0x0f, 0x55, 0x18, 0xf2, 0x0a, 0xb6, 0x1c, 0xea, 0x0b, 0x63,
	0xa9, 0x60, 0xd8, 0x1f, 0x46, 0x0d, 0x6f, 0x6f, 0x81, 0x75, 0xea, 0xdf, 0xe5, 0xa0, 0x3f, 0x8c,
	0x9a, 0x3e, 0x5c, 0x80, 0xe4, 0x06, 0xb6, 0x67, 0xd6, 0x43, 0x09, 0x62, 0xd8, 0xa0, 0x8c, 0xa1,
	0xb2, 0x38, 0x0c, 0x35, 0x98, 0x61, 0x92, 0x40, 0x1b, 0xb5, 0x96, 0xfa, 0x43, 0xad, 0x12, 0x35,
	0x2e, 0xb9, 0x82, 0xce, 0xd9, 0xc4, 0xde, 0x49, 0xcd, 0x7f, 0x52, 0xcb, 0xa5, 0x28, 0x6b, 0xd2,
	0x81, 0xa6, 0x95, 0x63, 0x14, 0x5e, 0xb4, 0x9d, 0x16, 0xc0, 0x3d, 0xef, 0x4e, 0x1a, 0x7b, 0xe9,
	0xe7, 0x36, 0xe8, 0x55, 0x98, 0xe4, 0x06, 0xf6, 0x16, 0xd4, 0x82, 0xcd, 0x7d, 0x00, 0x1a, 0x02,
	0x33, 0xa3, 0x15, 0xe6, 0x3f, 0xfd, 0x3a, 0x85, 0xe8, 0x02, 0xb5, 0xe5, 0x23, 0xce, 0xa8, 0xc5,
	0xcf, 0x6a, 0x48, 0x6d, 0x75, 0x84, 0xc6, 0xf8, 0x18, 0x2c, 0xba, 0x23, 0x21, 0xd0, 0x60, 0xa8,
	0xcb, 0xa6, 0xf8, 0x73, 0xb2, 0x0e, 0xcd, 0xcb, 0x5c, 0xd9, 0xc7, 0xe3, 0x5f, 0x6b, 0xb0, 0x9b,
	0xfa, 0xcd, 0x2b, 0xec, 0xde, 0xa2, 0x9e, 0x72, 0x86, 0xe4, 0x08, 0x1a, 0x6e, 0x92, 0x49, 0x68,
	0x79, 0x65, 0xc8, 0x63, 0x52, 0xa5, 0xc2, 0x5b, 0xde, 0xc1, 0x66, 0xf9, 0x48, 0x24, 0x71, 0x91,
	0xb0, 0xac, 0x86, 0xf1, 0x8b, 0xa5, 0xb1, 0xa0, 0x72, 0x0e, 0xad, 0xca, 0x48, 0x93, 0xa8, 0xc8,
	0xfd, 0x7b, 0x4f, 0xe2, 0xe7, 0x4b, 0x22, 0x41, 0xe3, 0x3d, 0x6c, 0xf7, 0x4d, 0x6d, 0x95, 0x4b,
	0x3f, 0xcb, 0xfe, 0x1a, 0xa5, 0x9f, 0xe5, 0xbb, 0xff, 0x16, 0xd6, 0xc3, 0x6c, 0x91, 0x4e, 0x91,
	0x57, 0xdf, 0x92, 0x78, 0x6f, 0x81, 0x0d, 0xf7, 0x4e, 0x61, 0xa7, 0x68, 0x4a, 0xa5, 0x4b, 0x64,
	0xbf, 0xc8, 0xfd, 0x57, 0xe3, 0xe2, 0x56, 0x58, 0x2e, 0xd7, 0x96, 0xf3, 0xd7, 0x5f, 0x7b, 0x19,
	0xb7, 0x77, 0x93, 0x41, 0x97, 0xc9, 0xbc, 0x47, 0x95, 0xb2, 0x1a, 0xd1, 0xc8, 0x91, 0x7d, 0xa0,
	0x1a, 0x7b, 0x99, 0x3c, 0x7a, 0x90, 0x7a, 0x3c, 0xba, 0x97, 0x0f, 0x3d, 0x35, 0xce, 0x7a, 0xee,
	0xe2, 0xe0, 0x89, 0xff, 0x3d, 0x9e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x61, 0x50, 0x76, 0xc2,
	0x49, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RemoteEngineServiceClient is the client API for RemoteEngineService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RemoteEngineServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	Authorize(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error)
	LoadPackage(ctx context.Context, in *LoadPackageRequest, opts ...grpc.CallOption) (*LoadPackageResponse, error)
	IsStepAvailable(ctx context.Context, in *StepAvailableRequest, opts ...grpc.CallOption) (*StepAvailableResponse, error)
	RunStep(ctx context.Context, in *RunStepRequest, opts ...grpc.CallOption) (*RunStepResponse, error)
	UpdateCertificate(ctx context.Context, in *CertificateUpdateRequest, opts ...grpc.CallOption) (*Empty, error)
}

type remoteEngineServiceClient struct {
	cc *grpc.ClientConn
}

func NewRemoteEngineServiceClient(cc *grpc.ClientConn) RemoteEngineServiceClient {
	return &remoteEngineServiceClient{cc}
}

func (c *remoteEngineServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, "/core.RemoteEngineService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteEngineServiceClient) Authorize(ctx context.Context, in *AuthorizationRequest, opts ...grpc.CallOption) (*AuthorizationResponse, error) {
	out := new(AuthorizationResponse)
	err := c.cc.Invoke(ctx, "/core.RemoteEngineService/Authorize", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteEngineServiceClient) LoadPackage(ctx context.Context, in *LoadPackageRequest, opts ...grpc.CallOption) (*LoadPackageResponse, error) {
	out := new(LoadPackageResponse)
	err := c.cc.Invoke(ctx, "/core.RemoteEngineService/LoadPackage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteEngineServiceClient) IsStepAvailable(ctx context.Context, in *StepAvailableRequest, opts ...grpc.CallOption) (*StepAvailableResponse, error) {
	out := new(StepAvailableResponse)
	err := c.cc.Invoke(ctx, "/core.RemoteEngineService/IsStepAvailable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteEngineServiceClient) RunStep(ctx context.Context, in *RunStepRequest, opts ...grpc.CallOption) (*RunStepResponse, error) {
	out := new(RunStepResponse)
	err := c.cc.Invoke(ctx, "/core.RemoteEngineService/RunStep", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *remoteEngineServiceClient) UpdateCertificate(ctx context.Context, in *CertificateUpdateRequest, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/core.RemoteEngineService/UpdateCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemoteEngineServiceServer is the server API for RemoteEngineService service.
type RemoteEngineServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	Authorize(context.Context, *AuthorizationRequest) (*AuthorizationResponse, error)
	LoadPackage(context.Context, *LoadPackageRequest) (*LoadPackageResponse, error)
	IsStepAvailable(context.Context, *StepAvailableRequest) (*StepAvailableResponse, error)
	RunStep(context.Context, *RunStepRequest) (*RunStepResponse, error)
	UpdateCertificate(context.Context, *CertificateUpdateRequest) (*Empty, error)
}

func RegisterRemoteEngineServiceServer(s *grpc.Server, srv RemoteEngineServiceServer) {
	s.RegisterService(&_RemoteEngineService_serviceDesc, srv)
}

func _RemoteEngineService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteEngineServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RemoteEngineService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteEngineServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteEngineService_Authorize_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteEngineServiceServer).Authorize(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RemoteEngineService/Authorize",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteEngineServiceServer).Authorize(ctx, req.(*AuthorizationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteEngineService_LoadPackage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadPackageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteEngineServiceServer).LoadPackage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RemoteEngineService/LoadPackage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteEngineServiceServer).LoadPackage(ctx, req.(*LoadPackageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteEngineService_IsStepAvailable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StepAvailableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteEngineServiceServer).IsStepAvailable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RemoteEngineService/IsStepAvailable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteEngineServiceServer).IsStepAvailable(ctx, req.(*StepAvailableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteEngineService_RunStep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunStepRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteEngineServiceServer).RunStep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RemoteEngineService/RunStep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteEngineServiceServer).RunStep(ctx, req.(*RunStepRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RemoteEngineService_UpdateCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CertificateUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteEngineServiceServer).UpdateCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RemoteEngineService/UpdateCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteEngineServiceServer).UpdateCertificate(ctx, req.(*CertificateUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RemoteEngineService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "core.RemoteEngineService",
	HandlerType: (*RemoteEngineServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _RemoteEngineService_Ping_Handler,
		},
		{
			MethodName: "Authorize",
			Handler:    _RemoteEngineService_Authorize_Handler,
		},
		{
			MethodName: "LoadPackage",
			Handler:    _RemoteEngineService_LoadPackage_Handler,
		},
		{
			MethodName: "IsStepAvailable",
			Handler:    _RemoteEngineService_IsStepAvailable_Handler,
		},
		{
			MethodName: "RunStep",
			Handler:    _RemoteEngineService_RunStep_Handler,
		},
		{
			MethodName: "UpdateCertificate",
			Handler:    _RemoteEngineService_UpdateCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "remote_engine.proto",
}
