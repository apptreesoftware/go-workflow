// Code generated by protoc-gen-go. DO NOT EDIT.
// source: remote_engine.proto

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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{0}
}
func (m *StepAvailableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepAvailableRequest.Unmarshal(m, b)
}
func (m *StepAvailableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepAvailableRequest.Marshal(b, m, deterministic)
}
func (dst *StepAvailableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepAvailableRequest.Merge(dst, src)
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{1}
}
func (m *StepAvailableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StepAvailableResponse.Unmarshal(m, b)
}
func (m *StepAvailableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StepAvailableResponse.Marshal(b, m, deterministic)
}
func (dst *StepAvailableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StepAvailableResponse.Merge(dst, src)
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{2}
}
func (m *PingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingRequest.Unmarshal(m, b)
}
func (m *PingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingRequest.Marshal(b, m, deterministic)
}
func (dst *PingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingRequest.Merge(dst, src)
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
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Os                   string   `protobuf:"bytes,3,opt,name=os,proto3" json:"os,omitempty"`
	Arch                 string   `protobuf:"bytes,4,opt,name=arch,proto3" json:"arch,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingResponse) Reset()         { *m = PingResponse{} }
func (m *PingResponse) String() string { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()    {}
func (*PingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{3}
}
func (m *PingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingResponse.Unmarshal(m, b)
}
func (m *PingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingResponse.Marshal(b, m, deterministic)
}
func (dst *PingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingResponse.Merge(dst, src)
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

func (m *PingResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *PingResponse) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *PingResponse) GetArch() string {
	if m != nil {
		return m.Arch
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{4}
}
func (m *LoadPackageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPackageRequest.Unmarshal(m, b)
}
func (m *LoadPackageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPackageRequest.Marshal(b, m, deterministic)
}
func (dst *LoadPackageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPackageRequest.Merge(dst, src)
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{5}
}
func (m *LoadPackageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadPackageResponse.Unmarshal(m, b)
}
func (m *LoadPackageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadPackageResponse.Marshal(b, m, deterministic)
}
func (dst *LoadPackageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadPackageResponse.Merge(dst, src)
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{6}
}
func (m *RunStepResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RunStepResponse.Unmarshal(m, b)
}
func (m *RunStepResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RunStepResponse.Marshal(b, m, deterministic)
}
func (dst *RunStepResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RunStepResponse.Merge(dst, src)
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
	Key                  []byte   `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
	Cert                 []byte   `protobuf:"bytes,4,opt,name=cert,proto3" json:"cert,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthorizationRequest) Reset()         { *m = AuthorizationRequest{} }
func (m *AuthorizationRequest) String() string { return proto.CompactTextString(m) }
func (*AuthorizationRequest) ProtoMessage()    {}
func (*AuthorizationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{7}
}
func (m *AuthorizationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationRequest.Unmarshal(m, b)
}
func (m *AuthorizationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationRequest.Marshal(b, m, deterministic)
}
func (dst *AuthorizationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationRequest.Merge(dst, src)
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

func (m *AuthorizationRequest) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *AuthorizationRequest) GetCert() []byte {
	if m != nil {
		return m.Cert
	}
	return nil
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{8}
}
func (m *AuthorizationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthorizationResponse.Unmarshal(m, b)
}
func (m *AuthorizationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthorizationResponse.Marshal(b, m, deterministic)
}
func (dst *AuthorizationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthorizationResponse.Merge(dst, src)
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{9}
}
func (m *CertificateUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateUpdateRequest.Unmarshal(m, b)
}
func (m *CertificateUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateUpdateRequest.Marshal(b, m, deterministic)
}
func (dst *CertificateUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateUpdateRequest.Merge(dst, src)
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
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{10}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type UpdateRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateRequest) Reset()         { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()    {}
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{11}
}
func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(dst, src)
}
func (m *UpdateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateRequest.Size(m)
}
func (m *UpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateRequest proto.InternalMessageInfo

func (m *UpdateRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type UpdateResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_remote_engine_0a63d5404e431d47, []int{12}
}
func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(dst, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

func (m *UpdateResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func init() {
	proto.RegisterType((*StepAvailableRequest)(nil), "core.StepAvailableRequest")
	proto.RegisterType((*StepAvailableResponse)(nil), "core.StepAvailableResponse")
	proto.RegisterType((*PingRequest)(nil), "core.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "core.PingResponse")
	proto.RegisterType((*LoadPackageRequest)(nil), "core.LoadPackageRequest")
	proto.RegisterType((*LoadPackageResponse)(nil), "core.LoadPackageResponse")
	proto.RegisterType((*RunStepResponse)(nil), "core.RunStepResponse")
	proto.RegisterType((*AuthorizationRequest)(nil), "core.AuthorizationRequest")
	proto.RegisterType((*AuthorizationResponse)(nil), "core.AuthorizationResponse")
	proto.RegisterType((*CertificateUpdateRequest)(nil), "core.CertificateUpdateRequest")
	proto.RegisterType((*Empty)(nil), "core.Empty")
	proto.RegisterType((*UpdateRequest)(nil), "core.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "core.UpdateResponse")
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
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
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

func (c *remoteEngineServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/core.RemoteEngineService/Update", in, out, opts...)
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
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
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

func _RemoteEngineService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemoteEngineServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/core.RemoteEngineService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemoteEngineServiceServer).Update(ctx, req.(*UpdateRequest))
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
			MethodName: "Update",
			Handler:    _RemoteEngineService_Update_Handler,
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

func init() { proto.RegisterFile("remote_engine.proto", fileDescriptor_remote_engine_0a63d5404e431d47) }

var fileDescriptor_remote_engine_0a63d5404e431d47 = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xdf, 0x6f, 0xd3, 0x30,
	0x10, 0xd6, 0xba, 0xee, 0xd7, 0x35, 0xdb, 0x98, 0xdb, 0x49, 0x21, 0xa0, 0x69, 0xf8, 0x09, 0x21,
	0xad, 0x15, 0x4c, 0xf0, 0xbc, 0x0d, 0xf6, 0x30, 0x04, 0x68, 0xcb, 0x80, 0x07, 0x5e, 0x90, 0xe7,
	0x5d, 0xd3, 0x28, 0x4d, 0x1c, 0x6c, 0x67, 0xd5, 0xf8, 0x0f, 0xf9, 0xaf, 0x50, 0x6c, 0x27, 0x4b,
	0xda, 0x0e, 0xde, 0x7c, 0xdf, 0xf9, 0xee, 0xbe, 0xfb, 0x7c, 0x67, 0xe8, 0x4b, 0x4c, 0x85, 0xc6,
	0x9f, 0x98, 0x45, 0x71, 0x86, 0xc3, 0x5c, 0x0a, 0x2d, 0x48, 0x97, 0x0b, 0x89, 0x81, 0xc7, 0x45,
	0x9a, 0x8a, 0xcc, 0x62, 0xf4, 0x2b, 0x0c, 0xae, 0x35, 0xe6, 0xa7, 0x77, 0x2c, 0x9e, 0xb2, 0x9b,
	0x29, 0x86, 0xf8, 0xab, 0x40, 0xa5, 0x49, 0x00, 0x9b, 0x4a, 0x63, 0xfe, 0x85, 0xa5, 0xe8, 0xaf,
	0x1c, 0xae, 0xbc, 0xdc, 0x0a, 0x6b, 0x9b, 0x1c, 0x42, 0xaf, 0x3c, 0x7f, 0x47, 0xa9, 0x62, 0x91,
	0xf9, 0x1d, 0xe3, 0x6e, 0x42, 0xf4, 0x2d, 0xec, 0xcf, 0x65, 0x55, 0xb9, 0xc8, 0x14, 0x92, 0xe7,
	0xb0, 0xc5, 0x2a, 0xd0, 0xe4, 0xdd, 0x0c, 0x1f, 0x00, 0xba, 0x0d, 0xbd, 0xcb, 0x38, 0x8b, 0x1c,
	0x07, 0x3a, 0x06, 0xcf, 0x9a, 0x2e, 0xd8, 0x87, 0x8d, 0x14, 0x95, 0x62, 0x51, 0x45, 0xa9, 0x32,
	0x4b, 0xcf, 0x5d, 0x8b, 0x4d, 0x65, 0x92, 0x1d, 0xe8, 0x08, 0xe5, 0xaf, 0x1a, 0xb0, 0x23, 0x14,
	0x21, 0xd0, 0x65, 0x92, 0x4f, 0xfc, 0xae, 0x41, 0xcc, 0x99, 0x9e, 0x00, 0xf9, 0x24, 0xd8, 0xed,
	0x25, 0xe3, 0x09, 0x8b, 0x6a, 0x05, 0x9e, 0xc0, 0x6a, 0x9e, 0x44, 0xae, 0x52, 0x79, 0x7c, 0xbc,
	0x0a, 0xbd, 0x80, 0x7e, 0x2b, 0xc3, 0x03, 0x61, 0x55, 0x70, 0x8e, 0x4a, 0xb9, 0x5e, 0x2b, 0xb3,
	0xd9, 0x4a, 0xa7, 0xd5, 0x0a, 0xbd, 0x82, 0xdd, 0xb0, 0xc8, 0x4a, 0xf5, 0xea, 0x34, 0x01, 0x6c,
	0x32, 0xce, 0x31, 0xd7, 0x78, 0xeb, 0xf2, 0xd4, 0x36, 0xa1, 0xe0, 0xa1, 0x94, 0x42, 0x7e, 0x6e,
	0x65, 0x6b, 0x61, 0x54, 0xc2, 0xe0, 0xb4, 0xd0, 0x13, 0x21, 0xe3, 0xdf, 0x4c, 0xc7, 0x22, 0xab,
	0x3a, 0x1c, 0xc0, 0x9a, 0x16, 0x09, 0x66, 0x26, 0xa9, 0x17, 0x5a, 0x83, 0x1c, 0x00, 0x4c, 0x84,
	0xd2, 0xe7, 0x66, 0x72, 0x5c, 0xbe, 0x06, 0x52, 0xea, 0x92, 0xe0, 0xbd, 0x91, 0xd4, 0x0b, 0xcb,
	0x63, 0xa9, 0x29, 0x47, 0xa9, 0x8d, 0xa6, 0x5e, 0x68, 0xce, 0xf4, 0x0a, 0xf6, 0xe7, 0x6a, 0xba,
	0x66, 0x0e, 0x00, 0x98, 0x73, 0xd4, 0xed, 0x34, 0x90, 0x7f, 0x28, 0x73, 0x02, 0xfe, 0x7b, 0x94,
	0x3a, 0x1e, 0xc7, 0x9c, 0x69, 0xfc, 0x96, 0xdf, 0x32, 0xdd, 0x7c, 0xac, 0x92, 0xd4, 0xca, 0x22,
	0xa9, 0x4e, 0x83, 0xd4, 0x06, 0xac, 0x9d, 0xa7, 0xb9, 0xbe, 0xa7, 0x2f, 0x60, 0x7b, 0x21, 0xbe,
	0x90, 0xd3, 0xea, 0xb1, 0x0b, 0x39, 0xa5, 0xaf, 0x60, 0xa7, 0xba, 0xf2, 0xbf, 0xd7, 0x7c, 0xf3,
	0x67, 0x15, 0xfa, 0xa1, 0x59, 0x38, 0xab, 0xd1, 0x35, 0xca, 0xbb, 0x98, 0x23, 0x39, 0x82, 0x6e,
	0x39, 0xc0, 0x64, 0x6f, 0x58, 0x6e, 0xde, 0xb0, 0x31, 0xdb, 0x01, 0x69, 0x42, 0xae, 0xc0, 0x31,
	0xac, 0xdb, 0x92, 0xa4, 0x6f, 0xbd, 0x2d, 0x8e, 0xc1, 0xa0, 0x0d, 0xba, 0xa0, 0x0f, 0xb0, 0x55,
	0x09, 0x8d, 0x24, 0xb0, 0x57, 0x96, 0xbd, 0x76, 0xf0, 0x6c, 0xa9, 0xcf, 0x65, 0x39, 0x83, 0x5e,
	0x63, 0x80, 0x89, 0x6f, 0xef, 0x2e, 0x6e, 0x45, 0xf0, 0x74, 0x89, 0xc7, 0xe5, 0xf8, 0x08, 0xbb,
	0x17, 0xaa, 0xb5, 0xf6, 0x15, 0x9f, 0x65, 0x3f, 0x4c, 0xc5, 0x67, 0xf9, 0x3f, 0xf1, 0x0e, 0x36,
	0xdc, 0x16, 0x10, 0xd7, 0x76, 0xbd, 0x14, 0x36, 0x7a, 0x7f, 0x0e, 0x75, 0x71, 0x27, 0xb0, 0x67,
	0xf5, 0x69, 0x4c, 0x0a, 0x39, 0xb0, 0x77, 0x1f, 0x1b, 0x9e, 0xa0, 0x67, 0xfd, 0x66, 0x34, 0xce,
	0x5e, 0xff, 0x18, 0x45, 0xb1, 0x9e, 0x14, 0x37, 0x43, 0x2e, 0xd2, 0x11, 0xcb, 0x73, 0x2d, 0x11,
	0x95, 0x18, 0xeb, 0x19, 0x93, 0x38, 0x8a, 0xc4, 0xd1, 0x4c, 0xc8, 0x64, 0x3c, 0x15, 0xb3, 0x51,
	0x9e, 0x44, 0xa3, 0x32, 0xf0, 0x66, 0xdd, 0x7c, 0xa5, 0xc7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x96, 0x14, 0xb8, 0xd5, 0x75, 0x05, 0x00, 0x00,
}
