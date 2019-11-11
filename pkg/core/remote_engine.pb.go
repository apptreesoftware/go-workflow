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
	return fileDescriptor_2d7480ac5bd0b3fb, []int{6}
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
	return fileDescriptor_2d7480ac5bd0b3fb, []int{7}
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
	return fileDescriptor_2d7480ac5bd0b3fb, []int{8}
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
	return fileDescriptor_2d7480ac5bd0b3fb, []int{9}
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
	return fileDescriptor_2d7480ac5bd0b3fb, []int{10}
}

func (m *UpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateRequest.Unmarshal(m, b)
}
func (m *UpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateRequest.Merge(m, src)
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
	return fileDescriptor_2d7480ac5bd0b3fb, []int{11}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
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
	proto.RegisterType((*UpdateRequest)(nil), "core.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "core.UpdateResponse")
}

func init() { proto.RegisterFile("remote_engine.proto", fileDescriptor_2d7480ac5bd0b3fb) }

var fileDescriptor_2d7480ac5bd0b3fb = []byte{
	// 596 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0x5f, 0x6f, 0xd3, 0x3e,
	0x14, 0xd5, 0xba, 0xfe, 0xf6, 0xe7, 0x36, 0xdb, 0x7e, 0x73, 0x3b, 0x29, 0x04, 0x54, 0x8d, 0x3c,
	0x21, 0xa4, 0xb5, 0x82, 0x09, 0x9e, 0xbb, 0xc1, 0x1e, 0x86, 0x00, 0x6d, 0x19, 0xf0, 0xc0, 0x0b,
	0x72, 0xdd, 0xdb, 0x34, 0x4a, 0x13, 0x07, 0xdb, 0x69, 0x35, 0xbe, 0x21, 0xdf, 0x0a, 0xc5, 0x76,
	0xb2, 0xa4, 0xed, 0xe0, 0xcd, 0xf7, 0x5c, 0xdf, 0xe3, 0x73, 0x8f, 0xaf, 0x0d, 0x5d, 0x81, 0x09,
	0x57, 0xf8, 0x03, 0xd3, 0x30, 0x4a, 0x71, 0x90, 0x09, 0xae, 0x38, 0x69, 0x33, 0x2e, 0xd0, 0x73,
	0x18, 0x4f, 0x12, 0x9e, 0x1a, 0xcc, 0xff, 0x02, 0xbd, 0x3b, 0x85, 0xd9, 0xc5, 0x82, 0x46, 0x73,
	0x3a, 0x9e, 0x63, 0x80, 0x3f, 0x73, 0x94, 0x8a, 0x78, 0xb0, 0x27, 0x15, 0x66, 0x9f, 0x69, 0x82,
	0xee, 0xd6, 0xe9, 0xd6, 0x8b, 0xfd, 0xa0, 0x8a, 0xc9, 0x29, 0x74, 0x8a, 0xf5, 0x37, 0x14, 0x32,
	0xe2, 0xa9, 0xdb, 0xd2, 0xe9, 0x3a, 0xe4, 0xbf, 0x81, 0x93, 0x15, 0x56, 0x99, 0xf1, 0x54, 0x22,
	0x79, 0x06, 0xfb, 0xb4, 0x04, 0x35, 0xef, 0x5e, 0xf0, 0x00, 0xf8, 0x07, 0xd0, 0xb9, 0x89, 0xd2,
	0xd0, 0x6a, 0xf0, 0xa7, 0xe0, 0x98, 0xd0, 0x16, 0xbb, 0xb0, 0x9b, 0xa0, 0x94, 0x34, 0x2c, 0x25,
	0x95, 0x61, 0x91, 0x59, 0x34, 0xd4, 0x94, 0x21, 0x39, 0x84, 0x16, 0x97, 0xee, 0xb6, 0x06, 0x5b,
	0x5c, 0x12, 0x02, 0x6d, 0x2a, 0xd8, 0xcc, 0x6d, 0x6b, 0x44, 0xaf, 0xfd, 0x11, 0x90, 0x8f, 0x9c,
	0x4e, 0x6e, 0x28, 0x8b, 0x69, 0x58, 0x39, 0xf0, 0x3f, 0x6c, 0x67, 0x71, 0x68, 0x4f, 0x2a, 0x96,
	0x8f, 0x9f, 0xe2, 0x5f, 0x43, 0xb7, 0xc1, 0xf0, 0x20, 0x58, 0xe6, 0x8c, 0xa1, 0x94, 0xb6, 0xd7,
	0x32, 0xac, 0xb7, 0xd2, 0x6a, 0xb4, 0xe2, 0xdf, 0xc2, 0x51, 0x90, 0xa7, 0x85, 0x7b, 0x15, 0x8d,
	0x07, 0x7b, 0x94, 0x31, 0xcc, 0x14, 0x4e, 0x2c, 0x4f, 0x15, 0x13, 0x1f, 0x1c, 0x14, 0x82, 0x8b,
	0x4f, 0x0d, 0xb6, 0x06, 0xe6, 0x0b, 0xe8, 0x5d, 0xe4, 0x6a, 0xc6, 0x45, 0xf4, 0x8b, 0xaa, 0x88,
	0xa7, 0x65, 0x87, 0x3d, 0xf8, 0x4f, 0xf1, 0x18, 0x53, 0x4d, 0xea, 0x04, 0x26, 0x20, 0x7d, 0x80,
	0x19, 0x97, 0xea, 0x4a, 0x4f, 0x8e, 0xe5, 0xab, 0x21, 0x85, 0x2f, 0x31, 0xde, 0x6b, 0x4b, 0x9d,
	0xa0, 0x58, 0x16, 0x9e, 0x32, 0x14, 0x4a, 0x7b, 0xea, 0x04, 0x7a, 0xed, 0xdf, 0xc2, 0xc9, 0xca,
	0x99, 0xb6, 0x99, 0x3e, 0x00, 0xb5, 0x89, 0xaa, 0x9d, 0x1a, 0xf2, 0x17, 0x67, 0x46, 0xe0, 0xbe,
	0x43, 0xa1, 0xa2, 0x69, 0xc4, 0xa8, 0xc2, 0xaf, 0xd9, 0x84, 0xaa, 0xfa, 0x65, 0x15, 0xa2, 0xb6,
	0xd6, 0x45, 0xb5, 0x6a, 0xa2, 0x9e, 0xc3, 0xc1, 0x5a, 0x59, 0x2e, 0xe6, 0xe5, 0x1d, 0xe7, 0x62,
	0xee, 0xbf, 0x84, 0xc3, 0x72, 0xcb, 0xbf, 0x2e, 0xf1, 0xf5, 0xef, 0x6d, 0xe8, 0x06, 0xfa, 0x9d,
	0x19, 0x6b, 0xee, 0x50, 0x2c, 0x22, 0x86, 0xe4, 0x0c, 0xda, 0xc5, 0xdc, 0x92, 0xe3, 0x41, 0xf1,
	0xe0, 0x06, 0xb5, 0x91, 0xf6, 0x48, 0x1d, 0xb2, 0x07, 0x9c, 0xc3, 0x8e, 0x39, 0x92, 0x74, 0x4d,
	0xb6, 0xa1, 0xd1, 0xeb, 0x35, 0x41, 0x5b, 0xf4, 0x1e, 0xf6, 0x4b, 0x7f, 0x91, 0x78, 0x66, 0xcb,
	0xa6, 0x4b, 0xf6, 0x9e, 0x6e, 0xcc, 0x59, 0x96, 0x4b, 0xe8, 0xd4, 0xe6, 0x96, 0xb8, 0x66, 0xef,
	0xfa, 0x63, 0xf0, 0x9e, 0x6c, 0xc8, 0x58, 0x8e, 0x0f, 0x70, 0x74, 0x2d, 0x1b, 0xaf, 0xbd, 0xd4,
	0xb3, 0xe9, 0x63, 0x29, 0xf5, 0x6c, 0xfe, 0x1e, 0xde, 0xc2, 0xae, 0x1d, 0x7e, 0x62, 0xdb, 0xae,
	0xde, 0x82, 0xa9, 0x3e, 0x59, 0x41, 0x6d, 0xdd, 0x08, 0x8e, 0x8d, 0x3f, 0xb5, 0x01, 0x21, 0x7d,
	0xb3, 0xf7, 0xb1, 0x99, 0xf1, 0x3a, 0x26, 0x7f, 0x95, 0x64, 0xea, 0xfe, 0xf2, 0xd5, 0xf7, 0x61,
	0x18, 0xa9, 0x59, 0x3e, 0x1e, 0x30, 0x9e, 0x0c, 0x69, 0x96, 0x29, 0x81, 0x28, 0xf9, 0x54, 0x2d,
	0xa9, 0xc0, 0x61, 0xc8, 0xcf, 0x96, 0x5c, 0xc4, 0xd3, 0x39, 0x5f, 0x0e, 0xb3, 0x38, 0x1c, 0x16,
	0x85, 0xe3, 0x1d, 0xfd, 0x83, 0x9e, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x81, 0xe9, 0x11, 0x86,
	0x6c, 0x05, 0x00, 0x00,
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
