// Code generated by protoc-gen-go. DO NOT EDIT.
// source: platform_api.proto

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

type EnvironmentStatus int32

const (
	EnvironmentStatus_Creating  EnvironmentStatus = 0
	EnvironmentStatus_Ready     EnvironmentStatus = 1
	EnvironmentStatus_Failed    EnvironmentStatus = 2
	EnvironmentStatus_Upgrading EnvironmentStatus = 3
)

var EnvironmentStatus_name = map[int32]string{
	0: "Creating",
	1: "Ready",
	2: "Failed",
	3: "Upgrading",
}
var EnvironmentStatus_value = map[string]int32{
	"Creating":  0,
	"Ready":     1,
	"Failed":    2,
	"Upgrading": 3,
}

func (x EnvironmentStatus) String() string {
	return proto.EnumName(EnvironmentStatus_name, int32(x))
}
func (EnvironmentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{0}
}

type EnvironmentConfiguration struct {
	EnvironmentName      string            `protobuf:"bytes,1,opt,name=environmentName,proto3" json:"environmentName,omitempty"`
	ResourcePlan         string            `protobuf:"bytes,2,opt,name=resourcePlan,proto3" json:"resourcePlan,omitempty"`
	Region               string            `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	Url                  string            `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Status               EnvironmentStatus `protobuf:"varint,5,opt,name=status,proto3,enum=core.EnvironmentStatus" json:"status,omitempty"`
	StatusMessage        string            `protobuf:"bytes,6,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *EnvironmentConfiguration) Reset()         { *m = EnvironmentConfiguration{} }
func (m *EnvironmentConfiguration) String() string { return proto.CompactTextString(m) }
func (*EnvironmentConfiguration) ProtoMessage()    {}
func (*EnvironmentConfiguration) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{0}
}
func (m *EnvironmentConfiguration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentConfiguration.Unmarshal(m, b)
}
func (m *EnvironmentConfiguration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentConfiguration.Marshal(b, m, deterministic)
}
func (dst *EnvironmentConfiguration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentConfiguration.Merge(dst, src)
}
func (m *EnvironmentConfiguration) XXX_Size() int {
	return xxx_messageInfo_EnvironmentConfiguration.Size(m)
}
func (m *EnvironmentConfiguration) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentConfiguration.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentConfiguration proto.InternalMessageInfo

func (m *EnvironmentConfiguration) GetEnvironmentName() string {
	if m != nil {
		return m.EnvironmentName
	}
	return ""
}

func (m *EnvironmentConfiguration) GetResourcePlan() string {
	if m != nil {
		return m.ResourcePlan
	}
	return ""
}

func (m *EnvironmentConfiguration) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func (m *EnvironmentConfiguration) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *EnvironmentConfiguration) GetStatus() EnvironmentStatus {
	if m != nil {
		return m.Status
	}
	return EnvironmentStatus_Creating
}

func (m *EnvironmentConfiguration) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

type UpdateEnvironmentRequest struct {
	EnvironmentName      string            `protobuf:"bytes,1,opt,name=environmentName,proto3" json:"environmentName,omitempty"`
	ResourcePlan         string            `protobuf:"bytes,2,opt,name=resourcePlan,proto3" json:"resourcePlan,omitempty"`
	Status               EnvironmentStatus `protobuf:"varint,3,opt,name=status,proto3,enum=core.EnvironmentStatus" json:"status,omitempty"`
	Url                  string            `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	StatusMessage        string            `protobuf:"bytes,5,opt,name=statusMessage,proto3" json:"statusMessage,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte            `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32             `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *UpdateEnvironmentRequest) Reset()         { *m = UpdateEnvironmentRequest{} }
func (m *UpdateEnvironmentRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEnvironmentRequest) ProtoMessage()    {}
func (*UpdateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{1}
}
func (m *UpdateEnvironmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEnvironmentRequest.Unmarshal(m, b)
}
func (m *UpdateEnvironmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEnvironmentRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateEnvironmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEnvironmentRequest.Merge(dst, src)
}
func (m *UpdateEnvironmentRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEnvironmentRequest.Size(m)
}
func (m *UpdateEnvironmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEnvironmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEnvironmentRequest proto.InternalMessageInfo

func (m *UpdateEnvironmentRequest) GetEnvironmentName() string {
	if m != nil {
		return m.EnvironmentName
	}
	return ""
}

func (m *UpdateEnvironmentRequest) GetResourcePlan() string {
	if m != nil {
		return m.ResourcePlan
	}
	return ""
}

func (m *UpdateEnvironmentRequest) GetStatus() EnvironmentStatus {
	if m != nil {
		return m.Status
	}
	return EnvironmentStatus_Creating
}

func (m *UpdateEnvironmentRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *UpdateEnvironmentRequest) GetStatusMessage() string {
	if m != nil {
		return m.StatusMessage
	}
	return ""
}

type EnvironmentRequest struct {
	EnvironmentName      string   `protobuf:"bytes,1,opt,name=environmentName,proto3" json:"environmentName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *EnvironmentRequest) Reset()         { *m = EnvironmentRequest{} }
func (m *EnvironmentRequest) String() string { return proto.CompactTextString(m) }
func (*EnvironmentRequest) ProtoMessage()    {}
func (*EnvironmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{2}
}
func (m *EnvironmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentRequest.Unmarshal(m, b)
}
func (m *EnvironmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentRequest.Marshal(b, m, deterministic)
}
func (dst *EnvironmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentRequest.Merge(dst, src)
}
func (m *EnvironmentRequest) XXX_Size() int {
	return xxx_messageInfo_EnvironmentRequest.Size(m)
}
func (m *EnvironmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentRequest proto.InternalMessageInfo

func (m *EnvironmentRequest) GetEnvironmentName() string {
	if m != nil {
		return m.EnvironmentName
	}
	return ""
}

type CreateEnvironmentRequest struct {
	EnvironmentName      string   `protobuf:"bytes,1,opt,name=environmentName,proto3" json:"environmentName,omitempty"`
	ResourcePlan         string   `protobuf:"bytes,2,opt,name=resourcePlan,proto3" json:"resourcePlan,omitempty"`
	Region               string   `protobuf:"bytes,3,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *CreateEnvironmentRequest) Reset()         { *m = CreateEnvironmentRequest{} }
func (m *CreateEnvironmentRequest) String() string { return proto.CompactTextString(m) }
func (*CreateEnvironmentRequest) ProtoMessage()    {}
func (*CreateEnvironmentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{3}
}
func (m *CreateEnvironmentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateEnvironmentRequest.Unmarshal(m, b)
}
func (m *CreateEnvironmentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateEnvironmentRequest.Marshal(b, m, deterministic)
}
func (dst *CreateEnvironmentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateEnvironmentRequest.Merge(dst, src)
}
func (m *CreateEnvironmentRequest) XXX_Size() int {
	return xxx_messageInfo_CreateEnvironmentRequest.Size(m)
}
func (m *CreateEnvironmentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateEnvironmentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateEnvironmentRequest proto.InternalMessageInfo

func (m *CreateEnvironmentRequest) GetEnvironmentName() string {
	if m != nil {
		return m.EnvironmentName
	}
	return ""
}

func (m *CreateEnvironmentRequest) GetResourcePlan() string {
	if m != nil {
		return m.ResourcePlan
	}
	return ""
}

func (m *CreateEnvironmentRequest) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

type DeleteEnvironmentResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *DeleteEnvironmentResponse) Reset()         { *m = DeleteEnvironmentResponse{} }
func (m *DeleteEnvironmentResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteEnvironmentResponse) ProtoMessage()    {}
func (*DeleteEnvironmentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{4}
}
func (m *DeleteEnvironmentResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteEnvironmentResponse.Unmarshal(m, b)
}
func (m *DeleteEnvironmentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteEnvironmentResponse.Marshal(b, m, deterministic)
}
func (dst *DeleteEnvironmentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteEnvironmentResponse.Merge(dst, src)
}
func (m *DeleteEnvironmentResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteEnvironmentResponse.Size(m)
}
func (m *DeleteEnvironmentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteEnvironmentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteEnvironmentResponse proto.InternalMessageInfo

type UpdateAvailableRequest struct {
	Os                   string   `protobuf:"bytes,1,opt,name=os,proto3" json:"os,omitempty"`
	Arch                 string   `protobuf:"bytes,2,opt,name=arch,proto3" json:"arch,omitempty"`
	Version              string   `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *UpdateAvailableRequest) Reset()         { *m = UpdateAvailableRequest{} }
func (m *UpdateAvailableRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateAvailableRequest) ProtoMessage()    {}
func (*UpdateAvailableRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{5}
}
func (m *UpdateAvailableRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvailableRequest.Unmarshal(m, b)
}
func (m *UpdateAvailableRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvailableRequest.Marshal(b, m, deterministic)
}
func (dst *UpdateAvailableRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvailableRequest.Merge(dst, src)
}
func (m *UpdateAvailableRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateAvailableRequest.Size(m)
}
func (m *UpdateAvailableRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvailableRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvailableRequest proto.InternalMessageInfo

func (m *UpdateAvailableRequest) GetOs() string {
	if m != nil {
		return m.Os
	}
	return ""
}

func (m *UpdateAvailableRequest) GetArch() string {
	if m != nil {
		return m.Arch
	}
	return ""
}

func (m *UpdateAvailableRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type UpdateAvailableResponse struct {
	Available            bool     `protobuf:"varint,1,opt,name=available,proto3" json:"available,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Url                  string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_unrecognized     []byte   `json:"-" yaml:"-" xml:"-" bson:"-"`
	XXX_sizecache        int32    `json:"-" yaml:"-" xml:"-" bson:"-"`
}

func (m *UpdateAvailableResponse) Reset()         { *m = UpdateAvailableResponse{} }
func (m *UpdateAvailableResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateAvailableResponse) ProtoMessage()    {}
func (*UpdateAvailableResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_platform_api_d79e73c2f3492dde, []int{6}
}
func (m *UpdateAvailableResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateAvailableResponse.Unmarshal(m, b)
}
func (m *UpdateAvailableResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateAvailableResponse.Marshal(b, m, deterministic)
}
func (dst *UpdateAvailableResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateAvailableResponse.Merge(dst, src)
}
func (m *UpdateAvailableResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateAvailableResponse.Size(m)
}
func (m *UpdateAvailableResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateAvailableResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateAvailableResponse proto.InternalMessageInfo

func (m *UpdateAvailableResponse) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *UpdateAvailableResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *UpdateAvailableResponse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*EnvironmentConfiguration)(nil), "core.EnvironmentConfiguration")
	proto.RegisterType((*UpdateEnvironmentRequest)(nil), "core.UpdateEnvironmentRequest")
	proto.RegisterType((*EnvironmentRequest)(nil), "core.EnvironmentRequest")
	proto.RegisterType((*CreateEnvironmentRequest)(nil), "core.CreateEnvironmentRequest")
	proto.RegisterType((*DeleteEnvironmentResponse)(nil), "core.DeleteEnvironmentResponse")
	proto.RegisterType((*UpdateAvailableRequest)(nil), "core.UpdateAvailableRequest")
	proto.RegisterType((*UpdateAvailableResponse)(nil), "core.UpdateAvailableResponse")
	proto.RegisterEnum("core.EnvironmentStatus", EnvironmentStatus_name, EnvironmentStatus_value)
}

func init() { proto.RegisterFile("platform_api.proto", fileDescriptor_platform_api_d79e73c2f3492dde) }

var fileDescriptor_platform_api_d79e73c2f3492dde = []byte{
	// 561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0xfd, 0x39, 0x4e, 0xf3, 0x6b, 0x46, 0x6d, 0x9a, 0x0c, 0xa2, 0x35, 0xa1, 0x94, 0xca, 0xe2,
	0x10, 0x21, 0x11, 0x8b, 0x72, 0x47, 0xa2, 0x05, 0x2a, 0x21, 0xb5, 0xaa, 0x52, 0x95, 0x03, 0x97,
	0x6a, 0xe3, 0x4c, 0xdc, 0x55, 0xed, 0x5d, 0xb3, 0xbb, 0x4e, 0x94, 0x1b, 0x5f, 0x91, 0x2f, 0xc0,
	0x9d, 0x6f, 0x81, 0xfc, 0xaf, 0xf9, 0x63, 0xa2, 0x56, 0x02, 0x71, 0x1b, 0xbf, 0x19, 0xbf, 0x7d,
	0x6f, 0xdf, 0xd8, 0x80, 0x71, 0xc8, 0xcc, 0x58, 0xaa, 0xe8, 0x9a, 0xc5, 0xbc, 0x1f, 0x2b, 0x69,
	0x24, 0xd6, 0x7d, 0xa9, 0xa8, 0xbb, 0xe5, 0xcb, 0x28, 0x92, 0x22, 0xc7, 0xba, 0xa8, 0x0d, 0xc5,
	0xd7, 0x21, 0x1f, 0x2a, 0xa6, 0x66, 0x39, 0xe6, 0xfe, 0xb4, 0xc0, 0xf9, 0x20, 0x26, 0x5c, 0x49,
	0x11, 0x91, 0x30, 0x27, 0x52, 0x8c, 0x79, 0x90, 0x28, 0x66, 0xb8, 0x14, 0xd8, 0x83, 0x1d, 0x9a,
	0xf7, 0xce, 0x59, 0x44, 0x8e, 0x75, 0x68, 0xf5, 0x9a, 0x83, 0x55, 0x18, 0x5d, 0xd8, 0x52, 0xa4,
	0x65, 0xa2, 0x7c, 0xba, 0x08, 0x99, 0x70, 0x6a, 0xd9, 0xd8, 0x12, 0x86, 0xbb, 0xd0, 0x50, 0x14,
	0x70, 0x29, 0x1c, 0x3b, 0xeb, 0x16, 0x4f, 0xd8, 0x06, 0x3b, 0x51, 0xa1, 0x53, 0xcf, 0xc0, 0xb4,
	0x44, 0x0f, 0x1a, 0xda, 0x30, 0x93, 0x68, 0x67, 0xe3, 0xd0, 0xea, 0xb5, 0x8e, 0xf6, 0xfa, 0xa9,
	0x9b, 0xfe, 0x82, 0xce, 0xcb, 0xac, 0x3d, 0x28, 0xc6, 0xf0, 0x05, 0x6c, 0xe7, 0xd5, 0x19, 0x69,
	0xcd, 0x02, 0x72, 0x1a, 0x19, 0xd9, 0x32, 0xe8, 0x7e, 0xb7, 0xc0, 0xb9, 0x8a, 0x47, 0xcc, 0xd0,
	0x02, 0xd3, 0x80, 0xbe, 0x26, 0xa4, 0xcd, 0x5f, 0xf6, 0x3a, 0x77, 0x60, 0x3f, 0xcc, 0x41, 0xf5,
	0x12, 0x2a, 0x9e, 0x36, 0x7e, 0xe7, 0xe9, 0x2d, 0xe0, 0x9f, 0x98, 0x71, 0xbf, 0x59, 0xe0, 0x9c,
	0x28, 0xfa, 0x17, 0x77, 0xb2, 0x26, 0x7f, 0xf7, 0x29, 0x3c, 0x79, 0x4f, 0x21, 0xad, 0x28, 0xd0,
	0xb1, 0x14, 0x9a, 0xdc, 0xcf, 0xb0, 0x9b, 0x47, 0xf6, 0x6e, 0xc2, 0x78, 0xc8, 0x86, 0x21, 0x95,
	0xe2, 0x5a, 0x50, 0x93, 0xba, 0xd0, 0x53, 0x93, 0x1a, 0x11, 0xea, 0x4c, 0xf9, 0x37, 0xc5, 0xd1,
	0x59, 0x8d, 0x0e, 0xfc, 0x3f, 0x21, 0xa5, 0xe7, 0x67, 0x96, 0x8f, 0xae, 0x0f, 0x7b, 0x15, 0xde,
	0xfc, 0x48, 0xdc, 0x87, 0x26, 0x2b, 0xc1, 0x8c, 0x7f, 0x73, 0x30, 0x07, 0x16, 0x29, 0x6b, 0x4b,
	0x94, 0x65, 0x84, 0xf6, 0x5d, 0x84, 0x2f, 0x4f, 0xa1, 0x53, 0x49, 0x1c, 0xb7, 0x60, 0x33, 0xbb,
	0x70, 0x2e, 0x82, 0xf6, 0x7f, 0xd8, 0x84, 0x8d, 0x01, 0xb1, 0xd1, 0xac, 0x6d, 0x21, 0x40, 0xe3,
	0x23, 0xe3, 0x21, 0x8d, 0xda, 0x35, 0xdc, 0x86, 0xe6, 0x55, 0x1c, 0x28, 0x36, 0x4a, 0xa7, 0xec,
	0xa3, 0x1f, 0x36, 0x3c, 0x5e, 0x60, 0x3a, 0x63, 0x82, 0x05, 0x94, 0x56, 0xe8, 0x41, 0xfd, 0x82,
	0x8b, 0x00, 0xb1, 0x58, 0xb0, 0x28, 0x36, 0xb3, 0x62, 0x37, 0xba, 0x8f, 0x72, 0xec, 0x98, 0x69,
	0xee, 0xdf, 0xb9, 0xbb, 0x84, 0x4e, 0x25, 0x6f, 0x3c, 0xc8, 0x27, 0xd7, 0x2d, 0x42, 0xf7, 0xa0,
	0xb2, 0xbe, 0xcb, 0x3f, 0x8a, 0x73, 0xe8, 0x54, 0x22, 0x44, 0xa7, 0xf2, 0x52, 0x49, 0xf7, 0x3c,
	0xef, 0xac, 0x4d, 0x3d, 0x15, 0x59, 0xf9, 0x50, 0x4b, 0x91, 0xeb, 0xbe, 0xe0, 0x7b, 0x45, 0x7e,
	0x82, 0xd6, 0x29, 0x99, 0x87, 0x29, 0xbc, 0xdf, 0xf0, 0xce, 0xca, 0xfa, 0xe0, 0xfe, 0xa2, 0xbc,
	0xd5, 0x6d, 0xed, 0x3e, 0x5b, 0xd3, 0xcd, 0x0d, 0x1f, 0xbf, 0xfe, 0xe2, 0x05, 0xdc, 0xdc, 0x24,
	0xc3, 0xbe, 0x2f, 0x23, 0x8f, 0xc5, 0xb1, 0x51, 0x44, 0x5a, 0x8e, 0xcd, 0x94, 0x29, 0xf2, 0x02,
	0xf9, 0x6a, 0x2a, 0xd5, 0xed, 0x38, 0x94, 0x53, 0x2f, 0xbe, 0x0d, 0xbc, 0x94, 0x6a, 0xd8, 0xc8,
	0x7e, 0xe0, 0x6f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x45, 0x96, 0x55, 0x4e, 0xfe, 0x05, 0x00,
	0x00,
}
