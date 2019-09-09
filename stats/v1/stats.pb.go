// Code generated by protoc-gen-go. DO NOT EDIT.
// source: stats.proto

package v1

import (
	v1 "code.storageos.net/storageos/service/common/v1"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type StatsStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsStatusRequest) Reset()         { *m = StatsStatusRequest{} }
func (m *StatsStatusRequest) String() string { return proto.CompactTextString(m) }
func (*StatsStatusRequest) ProtoMessage()    {}
func (*StatsStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{0}
}

func (m *StatsStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsStatusRequest.Unmarshal(m, b)
}
func (m *StatsStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsStatusRequest.Marshal(b, m, deterministic)
}
func (m *StatsStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsStatusRequest.Merge(m, src)
}
func (m *StatsStatusRequest) XXX_Size() int {
	return xxx_messageInfo_StatsStatusRequest.Size(m)
}
func (m *StatsStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_StatsStatusRequest proto.InternalMessageInfo

type StatsStatus struct {
	// The version control info string.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo,proto3" json:"version_info,omitempty"`
	// Generic daemon status.
	Status               *v1.DaemonStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *StatsStatus) Reset()         { *m = StatsStatus{} }
func (m *StatsStatus) String() string { return proto.CompactTextString(m) }
func (*StatsStatus) ProtoMessage()    {}
func (*StatsStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{1}
}

func (m *StatsStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsStatus.Unmarshal(m, b)
}
func (m *StatsStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsStatus.Marshal(b, m, deterministic)
}
func (m *StatsStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsStatus.Merge(m, src)
}
func (m *StatsStatus) XXX_Size() int {
	return xxx_messageInfo_StatsStatus.Size(m)
}
func (m *StatsStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsStatus.DiscardUnknown(m)
}

var xxx_messageInfo_StatsStatus proto.InternalMessageInfo

func (m *StatsStatus) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

func (m *StatsStatus) GetStatus() *v1.DaemonStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type StatsVolumeListQuery struct {
	// Optional list of volume IDs to query.
	VolumeIds            []uint32 `protobuf:"varint,1,rep,packed,name=volume_ids,json=volumeIds,proto3" json:"volume_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsVolumeListQuery) Reset()         { *m = StatsVolumeListQuery{} }
func (m *StatsVolumeListQuery) String() string { return proto.CompactTextString(m) }
func (*StatsVolumeListQuery) ProtoMessage()    {}
func (*StatsVolumeListQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{2}
}

func (m *StatsVolumeListQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsVolumeListQuery.Unmarshal(m, b)
}
func (m *StatsVolumeListQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsVolumeListQuery.Marshal(b, m, deterministic)
}
func (m *StatsVolumeListQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsVolumeListQuery.Merge(m, src)
}
func (m *StatsVolumeListQuery) XXX_Size() int {
	return xxx_messageInfo_StatsVolumeListQuery.Size(m)
}
func (m *StatsVolumeListQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsVolumeListQuery.DiscardUnknown(m)
}

var xxx_messageInfo_StatsVolumeListQuery proto.InternalMessageInfo

func (m *StatsVolumeListQuery) GetVolumeIds() []uint32 {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

type StatsVolumeLabel struct {
	// The label key.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// The label value.
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsVolumeLabel) Reset()         { *m = StatsVolumeLabel{} }
func (m *StatsVolumeLabel) String() string { return proto.CompactTextString(m) }
func (*StatsVolumeLabel) ProtoMessage()    {}
func (*StatsVolumeLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{3}
}

func (m *StatsVolumeLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsVolumeLabel.Unmarshal(m, b)
}
func (m *StatsVolumeLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsVolumeLabel.Marshal(b, m, deterministic)
}
func (m *StatsVolumeLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsVolumeLabel.Merge(m, src)
}
func (m *StatsVolumeLabel) XXX_Size() int {
	return xxx_messageInfo_StatsVolumeLabel.Size(m)
}
func (m *StatsVolumeLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsVolumeLabel.DiscardUnknown(m)
}

var xxx_messageInfo_StatsVolumeLabel proto.InternalMessageInfo

func (m *StatsVolumeLabel) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StatsVolumeLabel) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type StatsVolumeCredentials struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsVolumeCredentials) Reset()         { *m = StatsVolumeCredentials{} }
func (m *StatsVolumeCredentials) String() string { return proto.CompactTextString(m) }
func (*StatsVolumeCredentials) ProtoMessage()    {}
func (*StatsVolumeCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{4}
}

func (m *StatsVolumeCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsVolumeCredentials.Unmarshal(m, b)
}
func (m *StatsVolumeCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsVolumeCredentials.Marshal(b, m, deterministic)
}
func (m *StatsVolumeCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsVolumeCredentials.Merge(m, src)
}
func (m *StatsVolumeCredentials) XXX_Size() int {
	return xxx_messageInfo_StatsVolumeCredentials.Size(m)
}
func (m *StatsVolumeCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsVolumeCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_StatsVolumeCredentials proto.InternalMessageInfo

type StatsVolumeStatistics struct {
	// How much data has been written, e.g. volume will be full when matches provisioned size.
	CapacityUsedBytes uint64 `protobuf:"varint,1,opt,name=capacity_used_bytes,json=capacityUsedBytes,proto3" json:"capacity_used_bytes,omitempty"`
	// Since we compress, this should be less than what the client sees in 'df' on their mounted fs.
	ActualUsedBytes      uint64   `protobuf:"varint,2,opt,name=actual_used_bytes,json=actualUsedBytes,proto3" json:"actual_used_bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsVolumeStatistics) Reset()         { *m = StatsVolumeStatistics{} }
func (m *StatsVolumeStatistics) String() string { return proto.CompactTextString(m) }
func (*StatsVolumeStatistics) ProtoMessage()    {}
func (*StatsVolumeStatistics) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{5}
}

func (m *StatsVolumeStatistics) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsVolumeStatistics.Unmarshal(m, b)
}
func (m *StatsVolumeStatistics) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsVolumeStatistics.Marshal(b, m, deterministic)
}
func (m *StatsVolumeStatistics) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsVolumeStatistics.Merge(m, src)
}
func (m *StatsVolumeStatistics) XXX_Size() int {
	return xxx_messageInfo_StatsVolumeStatistics.Size(m)
}
func (m *StatsVolumeStatistics) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsVolumeStatistics.DiscardUnknown(m)
}

var xxx_messageInfo_StatsVolumeStatistics proto.InternalMessageInfo

func (m *StatsVolumeStatistics) GetCapacityUsedBytes() uint64 {
	if m != nil {
		return m.CapacityUsedBytes
	}
	return 0
}

func (m *StatsVolumeStatistics) GetActualUsedBytes() uint64 {
	if m != nil {
		return m.ActualUsedBytes
	}
	return 0
}

type StatsVolumeStatus struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatsVolumeStatus) Reset()         { *m = StatsVolumeStatus{} }
func (m *StatsVolumeStatus) String() string { return proto.CompactTextString(m) }
func (*StatsVolumeStatus) ProtoMessage()    {}
func (*StatsVolumeStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{6}
}

func (m *StatsVolumeStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsVolumeStatus.Unmarshal(m, b)
}
func (m *StatsVolumeStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsVolumeStatus.Marshal(b, m, deterministic)
}
func (m *StatsVolumeStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsVolumeStatus.Merge(m, src)
}
func (m *StatsVolumeStatus) XXX_Size() int {
	return xxx_messageInfo_StatsVolumeStatus.Size(m)
}
func (m *StatsVolumeStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsVolumeStatus.DiscardUnknown(m)
}

var xxx_messageInfo_StatsVolumeStatus proto.InternalMessageInfo

//*
// A volume used by the Stats module.
//
// Volumes contain the labels used to decorate the stats within Prometheus as well as the current stats.
// Stats returned in StatsVolumeStatistics are read-only and set in the data plane.
type StatsVolume struct {
	Cc *v1.DataplaneCommon `protobuf:"bytes,1,opt,name=cc,proto3" json:"cc,omitempty"`
	// The volume ID.
	VolumeId uint32 `protobuf:"varint,2,opt,name=volume_id,json=volumeId,proto3" json:"volume_id,omitempty"`
	// List of volume lables.
	Labels []*StatsVolumeLabel `protobuf:"bytes,3,rep,name=labels,proto3" json:"labels,omitempty"`
	// Volume credentials.
	Credentials *StatsVolumeCredentials `protobuf:"bytes,4,opt,name=credentials,proto3" json:"credentials,omitempty"`
	// Volume statistics.  Statistics are read-only and will be ignored in write operations.
	Stats *StatsVolumeStatistics `protobuf:"bytes,5,opt,name=stats,proto3" json:"stats,omitempty"`
	// Volume status.
	Status               *StatsVolumeStatus `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *StatsVolume) Reset()         { *m = StatsVolume{} }
func (m *StatsVolume) String() string { return proto.CompactTextString(m) }
func (*StatsVolume) ProtoMessage()    {}
func (*StatsVolume) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{7}
}

func (m *StatsVolume) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsVolume.Unmarshal(m, b)
}
func (m *StatsVolume) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsVolume.Marshal(b, m, deterministic)
}
func (m *StatsVolume) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsVolume.Merge(m, src)
}
func (m *StatsVolume) XXX_Size() int {
	return xxx_messageInfo_StatsVolume.Size(m)
}
func (m *StatsVolume) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsVolume.DiscardUnknown(m)
}

var xxx_messageInfo_StatsVolume proto.InternalMessageInfo

func (m *StatsVolume) GetCc() *v1.DataplaneCommon {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *StatsVolume) GetVolumeId() uint32 {
	if m != nil {
		return m.VolumeId
	}
	return 0
}

func (m *StatsVolume) GetLabels() []*StatsVolumeLabel {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *StatsVolume) GetCredentials() *StatsVolumeCredentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *StatsVolume) GetStats() *StatsVolumeStatistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *StatsVolume) GetStatus() *StatsVolumeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

//* A list of volumes used by the Stats module.
//
type StatsVolumeList struct {
	Volumes              []*StatsVolume `protobuf:"bytes,1,rep,name=volumes,proto3" json:"volumes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *StatsVolumeList) Reset()         { *m = StatsVolumeList{} }
func (m *StatsVolumeList) String() string { return proto.CompactTextString(m) }
func (*StatsVolumeList) ProtoMessage()    {}
func (*StatsVolumeList) Descriptor() ([]byte, []int) {
	return fileDescriptor_b4756a0aec8b9d44, []int{8}
}

func (m *StatsVolumeList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatsVolumeList.Unmarshal(m, b)
}
func (m *StatsVolumeList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatsVolumeList.Marshal(b, m, deterministic)
}
func (m *StatsVolumeList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatsVolumeList.Merge(m, src)
}
func (m *StatsVolumeList) XXX_Size() int {
	return xxx_messageInfo_StatsVolumeList.Size(m)
}
func (m *StatsVolumeList) XXX_DiscardUnknown() {
	xxx_messageInfo_StatsVolumeList.DiscardUnknown(m)
}

var xxx_messageInfo_StatsVolumeList proto.InternalMessageInfo

func (m *StatsVolumeList) GetVolumes() []*StatsVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func init() {
	proto.RegisterType((*StatsStatusRequest)(nil), "stats.v1.StatsStatusRequest")
	proto.RegisterType((*StatsStatus)(nil), "stats.v1.StatsStatus")
	proto.RegisterType((*StatsVolumeListQuery)(nil), "stats.v1.StatsVolumeListQuery")
	proto.RegisterType((*StatsVolumeLabel)(nil), "stats.v1.StatsVolumeLabel")
	proto.RegisterType((*StatsVolumeCredentials)(nil), "stats.v1.StatsVolumeCredentials")
	proto.RegisterType((*StatsVolumeStatistics)(nil), "stats.v1.StatsVolumeStatistics")
	proto.RegisterType((*StatsVolumeStatus)(nil), "stats.v1.StatsVolumeStatus")
	proto.RegisterType((*StatsVolume)(nil), "stats.v1.StatsVolume")
	proto.RegisterType((*StatsVolumeList)(nil), "stats.v1.StatsVolumeList")
}

func init() { proto.RegisterFile("stats.proto", fileDescriptor_b4756a0aec8b9d44) }

var fileDescriptor_b4756a0aec8b9d44 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x4e, 0xdb, 0x4a,
	0x10, 0x25, 0x09, 0xc9, 0x25, 0x63, 0xb8, 0xc0, 0x92, 0x5c, 0x7c, 0x0d, 0x5c, 0x72, 0xfd, 0x84,
	0x90, 0x6a, 0x8b, 0x20, 0x5e, 0x2a, 0x55, 0x95, 0x02, 0x6a, 0x45, 0x5b, 0x1e, 0xba, 0x88, 0x3e,
	0xf4, 0x25, 0x5a, 0xec, 0x01, 0x59, 0x75, 0xbc, 0xa9, 0x77, 0x6d, 0x29, 0x7f, 0xd2, 0x6f, 0xe9,
	0xd7, 0x55, 0xde, 0xb5, 0xc3, 0x92, 0x38, 0xad, 0xaa, 0xbe, 0x20, 0xfb, 0xcc, 0x39, 0x73, 0x66,
	0x0e, 0x63, 0x05, 0x2c, 0x21, 0x99, 0x14, 0xde, 0x34, 0xe5, 0x92, 0x93, 0x0d, 0xfd, 0x92, 0x9f,
	0x39, 0x9b, 0x01, 0x9f, 0x4c, 0x78, 0xa2, 0x71, 0xb7, 0x07, 0xe4, 0xb6, 0xa8, 0x14, 0x7f, 0x32,
	0x41, 0xf1, 0x6b, 0x86, 0x42, 0xba, 0x0c, 0x2c, 0x03, 0x25, 0xff, 0xc3, 0x66, 0x8e, 0xa9, 0x88,
	0x78, 0x32, 0x8e, 0x92, 0x07, 0x6e, 0x37, 0x06, 0x8d, 0x93, 0x2e, 0xb5, 0x4a, 0xec, 0x3a, 0x79,
	0xe0, 0xc4, 0x87, 0x8e, 0x50, 0x64, 0xbb, 0x39, 0x68, 0x9c, 0x58, 0xc3, 0x7d, 0xaf, 0xb4, 0xc9,
	0xcf, 0xbc, 0x2b, 0x86, 0x13, 0x9e, 0x94, 0x0e, 0x25, 0xcd, 0xbd, 0x80, 0x9e, 0xb2, 0xf8, 0xc4,
	0xe3, 0x6c, 0x82, 0x1f, 0x22, 0x21, 0x3f, 0x66, 0x98, 0xce, 0xc8, 0x11, 0x40, 0xae, 0xa0, 0x71,
	0x14, 0x0a, 0xbb, 0x31, 0x68, 0x9d, 0x6c, 0xd1, 0xae, 0x46, 0xae, 0x43, 0xe1, 0xbe, 0x84, 0x1d,
	0x53, 0xc6, 0xee, 0x31, 0x26, 0x3b, 0xd0, 0xfa, 0x82, 0xb3, 0x72, 0xaa, 0xe2, 0x91, 0xf4, 0xa0,
	0x9d, 0xb3, 0x38, 0x43, 0x35, 0x4c, 0x97, 0xea, 0x17, 0xd7, 0x86, 0x7f, 0x0c, 0xed, 0x65, 0x8a,
	0x21, 0x26, 0x32, 0x62, 0xb1, 0x70, 0x05, 0xf4, 0x8d, 0x4a, 0xf1, 0x18, 0x09, 0x19, 0x05, 0x82,
	0x78, 0xb0, 0x17, 0xb0, 0x29, 0x0b, 0x22, 0x39, 0x1b, 0x67, 0x02, 0xc3, 0xf1, 0xfd, 0x4c, 0xa2,
	0x50, 0x56, 0xeb, 0x74, 0xb7, 0x2a, 0xdd, 0x09, 0x0c, 0x47, 0x45, 0x81, 0x9c, 0xc2, 0x2e, 0x0b,
	0x64, 0xc6, 0x62, 0x93, 0xdd, 0x54, 0xec, 0x6d, 0x5d, 0x98, 0x73, 0xdd, 0x3d, 0xd8, 0x5d, 0x30,
	0xcd, 0x84, 0xfb, 0xbd, 0x59, 0x46, 0xaf, 0x51, 0x72, 0x0a, 0xcd, 0x20, 0x50, 0x7e, 0xd6, 0xd0,
	0x79, 0x96, 0xa9, 0x64, 0xd3, 0x98, 0x25, 0x78, 0xa9, 0x20, 0xda, 0x0c, 0x02, 0x72, 0x00, 0xdd,
	0x79, 0x74, 0xca, 0x74, 0x8b, 0x6e, 0x54, 0xc9, 0x91, 0x21, 0x74, 0xe2, 0x22, 0x2d, 0x61, 0xb7,
	0x06, 0x2d, 0xd5, 0xac, 0xba, 0x08, 0x6f, 0x31, 0x50, 0x5a, 0x32, 0xc9, 0x08, 0xac, 0xe0, 0x29,
	0x25, 0x7b, 0x5d, 0x4d, 0x31, 0xa8, 0x15, 0x1a, 0x69, 0x52, 0x53, 0x44, 0x2e, 0xa0, 0xad, 0xf8,
	0x76, 0x5b, 0xa9, 0x8f, 0x6b, 0xd5, 0x4f, 0x89, 0x53, 0xcd, 0x26, 0xe7, 0xf3, 0x7b, 0xea, 0x28,
	0xdd, 0xc1, 0x4a, 0x9d, 0x71, 0x53, 0x23, 0xd8, 0x5e, 0xb8, 0x29, 0xe2, 0xc3, 0x5f, 0x3a, 0x02,
	0x7d, 0x4b, 0xd6, 0xb0, 0x5f, 0xdb, 0x88, 0x56, 0xac, 0xe1, 0xb7, 0x0e, 0xb4, 0x55, 0x81, 0xbc,
	0x86, 0x4e, 0x79, 0xff, 0x87, 0x0b, 0x9a, 0x67, 0x1f, 0x8b, 0xd3, 0xaf, 0xad, 0xba, 0x6b, 0xe4,
	0x15, 0x6c, 0xce, 0xc3, 0x61, 0x12, 0x49, 0xbd, 0xb5, 0xd3, 0x33, 0xfe, 0xad, 0x74, 0x1a, 0x50,
	0x14, 0x59, 0x2c, 0x4d, 0xf9, 0xdd, 0x34, 0xfc, 0x13, 0xf9, 0x15, 0xc6, 0xf8, 0xfb, 0xf2, 0x6b,
	0x00, 0x23, 0xc6, 0xff, 0xea, 0xaf, 0xa5, 0xfa, 0x6a, 0x9d, 0x7f, 0x57, 0xd6, 0xdd, 0x35, 0xf2,
	0x06, 0xb6, 0x2e, 0x79, 0xf2, 0x10, 0x3d, 0xbe, 0x45, 0x39, 0xe2, 0x3c, 0x26, 0xa6, 0xa7, 0xae,
	0xbc, 0xc7, 0x99, 0x73, 0xb4, 0x84, 0x96, 0x7c, 0x8a, 0xd3, 0x78, 0xa6, 0x46, 0xda, 0xd1, 0xb8,
	0x0e, 0x44, 0xb5, 0xea, 0x2f, 0x89, 0x0a, 0xd8, 0x39, 0x5c, 0x82, 0xb5, 0xe6, 0xa9, 0xd5, 0xdf,
	0x1a, 0x2e, 0x46, 0x54, 0x8d, 0x9c, 0x25, 0x85, 0xb9, 0x5d, 0x9d, 0x49, 0xb9, 0xdd, 0x3b, 0xd8,
	0x9e, 0x4f, 0x7b, 0x2b, 0xd3, 0x28, 0x79, 0x5c, 0xb1, 0xdf, 0x71, 0xdd, 0x7e, 0x5a, 0x51, 0x8d,
	0x75, 0x03, 0xc4, 0x9c, 0xb6, 0x6c, 0xb7, 0xbf, 0x24, 0xd4, 0x85, 0x5f, 0x6e, 0x79, 0x53, 0x05,
	0x56, 0x8c, 0x5a, 0x36, 0xfb, 0xd9, 0x9e, 0x07, 0x2b, 0x8c, 0xf4, 0xa6, 0x23, 0xff, 0xf3, 0x8b,
	0x80, 0x87, 0xe8, 0x09, 0xc9, 0x53, 0xf6, 0x88, 0x5c, 0x78, 0x09, 0x4a, 0x7f, 0xfe, 0xe6, 0x0b,
	0x4c, 0xf3, 0x28, 0x40, 0x5f, 0x9d, 0x82, 0x9f, 0x9f, 0xdd, 0x77, 0xd4, 0x6f, 0xcc, 0xf9, 0x8f,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc5, 0x13, 0xb8, 0x92, 0x8a, 0x06, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type StatsClient interface {
	//*
	// Get program status.
	Status(ctx context.Context, in *StatsStatusRequest, opts ...grpc.CallOption) (*StatsStatus, error)
	//*
	// Add configuration for the given StatsVolume message.
	//
	// returns RpcResult
	VolumeCreate(ctx context.Context, in *StatsVolume, opts ...grpc.CallOption) (*v1.RpcResult, error)
	//*
	// Update configuration for the given StatsVolume message.
	//
	// returns RpcResult
	VolumeUpdate(ctx context.Context, in *StatsVolume, opts ...grpc.CallOption) (*v1.RpcResult, error)
	//*
	// Remove configuration for the given StatsVolume message.
	//
	// returns RpcResult
	VolumeDelete(ctx context.Context, in *StatsVolume, opts ...grpc.CallOption) (*v1.RpcResult, error)
	//*
	// Return a list of configured volumes. Optionally filter using the StatsVolumeListQuery message.
	//
	// returns An StatsVolumeList message containing StatsVolume messages,
	//         if any are found matching the filter.
	VolumeList(ctx context.Context, in *StatsVolumeListQuery, opts ...grpc.CallOption) (*StatsVolumeList, error)
	// Config services, from common.v1.
	ConfigGetBool(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetBoolReply, error)
	ConfigUpdateBool(ctx context.Context, in *v1.ConfigBool, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error)
	ConfigListBool(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigBoolList, error)
	ConfigGetString(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetStringReply, error)
	ConfigUpdateString(ctx context.Context, in *v1.ConfigString, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error)
	ConfigListString(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigStringList, error)
}

type statsClient struct {
	cc *grpc.ClientConn
}

func NewStatsClient(cc *grpc.ClientConn) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) Status(ctx context.Context, in *StatsStatusRequest, opts ...grpc.CallOption) (*StatsStatus, error) {
	out := new(StatsStatus)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) VolumeCreate(ctx context.Context, in *StatsVolume, opts ...grpc.CallOption) (*v1.RpcResult, error) {
	out := new(v1.RpcResult)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/VolumeCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) VolumeUpdate(ctx context.Context, in *StatsVolume, opts ...grpc.CallOption) (*v1.RpcResult, error) {
	out := new(v1.RpcResult)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/VolumeUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) VolumeDelete(ctx context.Context, in *StatsVolume, opts ...grpc.CallOption) (*v1.RpcResult, error) {
	out := new(v1.RpcResult)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/VolumeDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) VolumeList(ctx context.Context, in *StatsVolumeListQuery, opts ...grpc.CallOption) (*StatsVolumeList, error) {
	out := new(StatsVolumeList)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/VolumeList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) ConfigGetBool(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetBoolReply, error) {
	out := new(v1.ConfigGetBoolReply)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/ConfigGetBool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) ConfigUpdateBool(ctx context.Context, in *v1.ConfigBool, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
	out := new(v1.ConfigUpdateReply)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/ConfigUpdateBool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) ConfigListBool(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigBoolList, error) {
	out := new(v1.ConfigBoolList)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/ConfigListBool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) ConfigGetString(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetStringReply, error) {
	out := new(v1.ConfigGetStringReply)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/ConfigGetString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) ConfigUpdateString(ctx context.Context, in *v1.ConfigString, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
	out := new(v1.ConfigUpdateReply)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/ConfigUpdateString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) ConfigListString(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigStringList, error) {
	out := new(v1.ConfigStringList)
	err := c.cc.Invoke(ctx, "/stats.v1.Stats/ConfigListString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServer is the server API for Stats service.
type StatsServer interface {
	//*
	// Get program status.
	Status(context.Context, *StatsStatusRequest) (*StatsStatus, error)
	//*
	// Add configuration for the given StatsVolume message.
	//
	// returns RpcResult
	VolumeCreate(context.Context, *StatsVolume) (*v1.RpcResult, error)
	//*
	// Update configuration for the given StatsVolume message.
	//
	// returns RpcResult
	VolumeUpdate(context.Context, *StatsVolume) (*v1.RpcResult, error)
	//*
	// Remove configuration for the given StatsVolume message.
	//
	// returns RpcResult
	VolumeDelete(context.Context, *StatsVolume) (*v1.RpcResult, error)
	//*
	// Return a list of configured volumes. Optionally filter using the StatsVolumeListQuery message.
	//
	// returns An StatsVolumeList message containing StatsVolume messages,
	//         if any are found matching the filter.
	VolumeList(context.Context, *StatsVolumeListQuery) (*StatsVolumeList, error)
	// Config services, from common.v1.
	ConfigGetBool(context.Context, *v1.ConfigKey) (*v1.ConfigGetBoolReply, error)
	ConfigUpdateBool(context.Context, *v1.ConfigBool) (*v1.ConfigUpdateReply, error)
	ConfigListBool(context.Context, *v1.ConfigListQuery) (*v1.ConfigBoolList, error)
	ConfigGetString(context.Context, *v1.ConfigKey) (*v1.ConfigGetStringReply, error)
	ConfigUpdateString(context.Context, *v1.ConfigString) (*v1.ConfigUpdateReply, error)
	ConfigListString(context.Context, *v1.ConfigListQuery) (*v1.ConfigStringList, error)
}

// UnimplementedStatsServer can be embedded to have forward compatible implementations.
type UnimplementedStatsServer struct {
}

func (*UnimplementedStatsServer) Status(ctx context.Context, req *StatsStatusRequest) (*StatsStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedStatsServer) VolumeCreate(ctx context.Context, req *StatsVolume) (*v1.RpcResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumeCreate not implemented")
}
func (*UnimplementedStatsServer) VolumeUpdate(ctx context.Context, req *StatsVolume) (*v1.RpcResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumeUpdate not implemented")
}
func (*UnimplementedStatsServer) VolumeDelete(ctx context.Context, req *StatsVolume) (*v1.RpcResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumeDelete not implemented")
}
func (*UnimplementedStatsServer) VolumeList(ctx context.Context, req *StatsVolumeListQuery) (*StatsVolumeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VolumeList not implemented")
}
func (*UnimplementedStatsServer) ConfigGetBool(ctx context.Context, req *v1.ConfigKey) (*v1.ConfigGetBoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigGetBool not implemented")
}
func (*UnimplementedStatsServer) ConfigUpdateBool(ctx context.Context, req *v1.ConfigBool) (*v1.ConfigUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigUpdateBool not implemented")
}
func (*UnimplementedStatsServer) ConfigListBool(ctx context.Context, req *v1.ConfigListQuery) (*v1.ConfigBoolList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigListBool not implemented")
}
func (*UnimplementedStatsServer) ConfigGetString(ctx context.Context, req *v1.ConfigKey) (*v1.ConfigGetStringReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigGetString not implemented")
}
func (*UnimplementedStatsServer) ConfigUpdateString(ctx context.Context, req *v1.ConfigString) (*v1.ConfigUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigUpdateString not implemented")
}
func (*UnimplementedStatsServer) ConfigListString(ctx context.Context, req *v1.ConfigListQuery) (*v1.ConfigStringList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigListString not implemented")
}

func RegisterStatsServer(s *grpc.Server, srv StatsServer) {
	s.RegisterService(&_Stats_serviceDesc, srv)
}

func _Stats_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Status(ctx, req.(*StatsStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_VolumeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).VolumeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/VolumeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).VolumeCreate(ctx, req.(*StatsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_VolumeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).VolumeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/VolumeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).VolumeUpdate(ctx, req.(*StatsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_VolumeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).VolumeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/VolumeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).VolumeDelete(ctx, req.(*StatsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_VolumeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatsVolumeListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).VolumeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/VolumeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).VolumeList(ctx, req.(*StatsVolumeListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_ConfigGetBool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).ConfigGetBool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/ConfigGetBool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).ConfigGetBool(ctx, req.(*v1.ConfigKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_ConfigUpdateBool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigBool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).ConfigUpdateBool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/ConfigUpdateBool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).ConfigUpdateBool(ctx, req.(*v1.ConfigBool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_ConfigListBool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).ConfigListBool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/ConfigListBool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).ConfigListBool(ctx, req.(*v1.ConfigListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_ConfigGetString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).ConfigGetString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/ConfigGetString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).ConfigGetString(ctx, req.(*v1.ConfigKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_ConfigUpdateString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).ConfigUpdateString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/ConfigUpdateString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).ConfigUpdateString(ctx, req.(*v1.ConfigString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_ConfigListString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).ConfigListString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/stats.v1.Stats/ConfigListString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).ConfigListString(ctx, req.(*v1.ConfigListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Stats_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stats.v1.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Stats_Status_Handler,
		},
		{
			MethodName: "VolumeCreate",
			Handler:    _Stats_VolumeCreate_Handler,
		},
		{
			MethodName: "VolumeUpdate",
			Handler:    _Stats_VolumeUpdate_Handler,
		},
		{
			MethodName: "VolumeDelete",
			Handler:    _Stats_VolumeDelete_Handler,
		},
		{
			MethodName: "VolumeList",
			Handler:    _Stats_VolumeList_Handler,
		},
		{
			MethodName: "ConfigGetBool",
			Handler:    _Stats_ConfigGetBool_Handler,
		},
		{
			MethodName: "ConfigUpdateBool",
			Handler:    _Stats_ConfigUpdateBool_Handler,
		},
		{
			MethodName: "ConfigListBool",
			Handler:    _Stats_ConfigListBool_Handler,
		},
		{
			MethodName: "ConfigGetString",
			Handler:    _Stats_ConfigGetString_Handler,
		},
		{
			MethodName: "ConfigUpdateString",
			Handler:    _Stats_ConfigUpdateString_Handler,
		},
		{
			MethodName: "ConfigListString",
			Handler:    _Stats_ConfigListString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stats.proto",
}
