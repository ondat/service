// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filesystem.proto

/*
Package filesystem_v1 is a generated protocol buffer package.

It is generated from these files:
	filesystem.proto

It has these top-level messages:
	FsStatusRequest
	FsStatus
	FsVolumeListQuery
	FsVolumeStatistics
	FsVolumeStatus
	FsVolume
	FsVolumeList
	FsPresentationListQuery
	FsPresentation
	FsPresentationList
*/
package filesystem_v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common_v1 "."

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

type FsVolumeState int32

const (
	FsVolumeState_NONE  FsVolumeState = 0
	FsVolumeState_READY FsVolumeState = 1
)

var FsVolumeState_name = map[int32]string{
	0: "NONE",
	1: "READY",
}
var FsVolumeState_value = map[string]int32{
	"NONE":  0,
	"READY": 1,
}

func (x FsVolumeState) String() string {
	return proto.EnumName(FsVolumeState_name, int32(x))
}
func (FsVolumeState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FsVolume_VolumeDeviceType int32

const (
	FsVolume_FILE      FsVolume_VolumeDeviceType = 0
	FsVolume_NBD_BLOCK FsVolume_VolumeDeviceType = 1
)

var FsVolume_VolumeDeviceType_name = map[int32]string{
	0: "FILE",
	1: "NBD_BLOCK",
}
var FsVolume_VolumeDeviceType_value = map[string]int32{
	"FILE":      0,
	"NBD_BLOCK": 1,
}

func (x FsVolume_VolumeDeviceType) String() string {
	return proto.EnumName(FsVolume_VolumeDeviceType_name, int32(x))
}
func (FsVolume_VolumeDeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{5, 0} }

type FsStatusRequest struct {
}

func (m *FsStatusRequest) Reset()                    { *m = FsStatusRequest{} }
func (m *FsStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*FsStatusRequest) ProtoMessage()               {}
func (*FsStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type FsStatus struct {
	// The version control info string.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
}

func (m *FsStatus) Reset()                    { *m = FsStatus{} }
func (m *FsStatus) String() string            { return proto.CompactTextString(m) }
func (*FsStatus) ProtoMessage()               {}
func (*FsStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FsStatus) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

type FsVolumeListQuery struct {
	// A possibly-empty list of volume IDs to query.
	VolumeIds []uint32 `protobuf:"varint,1,rep,packed,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
}

func (m *FsVolumeListQuery) Reset()                    { *m = FsVolumeListQuery{} }
func (m *FsVolumeListQuery) String() string            { return proto.CompactTextString(m) }
func (*FsVolumeListQuery) ProtoMessage()               {}
func (*FsVolumeListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FsVolumeListQuery) GetVolumeIds() []uint32 {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

type FsVolumeStatistics struct {
}

func (m *FsVolumeStatistics) Reset()                    { *m = FsVolumeStatistics{} }
func (m *FsVolumeStatistics) String() string            { return proto.CompactTextString(m) }
func (*FsVolumeStatistics) ProtoMessage()               {}
func (*FsVolumeStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type FsVolumeStatus struct {
	// State of the volume device node (file/block device).
	NodeState FsVolumeState `protobuf:"varint,1,opt,name=node_state,json=nodeState,enum=filesystem.v1.FsVolumeState" json:"node_state,omitempty"`
}

func (m *FsVolumeStatus) Reset()                    { *m = FsVolumeStatus{} }
func (m *FsVolumeStatus) String() string            { return proto.CompactTextString(m) }
func (*FsVolumeStatus) ProtoMessage()               {}
func (*FsVolumeStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FsVolumeStatus) GetNodeState() FsVolumeState {
	if m != nil {
		return m.NodeState
	}
	return FsVolumeState_NONE
}

// *
// A StorageOS volume to be presented via the FUSE filesystem.
type FsVolume struct {
	Cc *common_v1.DataplaneCommon `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The volume ID to represent.
	VolumeId uint32 `protobuf:"varint,2,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
	// The type for this volume.
	NodeType FsVolume_VolumeDeviceType `protobuf:"varint,3,opt,name=node_type,json=nodeType,enum=filesystem.v1.FsVolume_VolumeDeviceType" json:"node_type,omitempty"`
	// An opaque value interpreted based on node_type.
	DeviceNumber uint32 `protobuf:"varint,4,opt,name=device_number,json=deviceNumber" json:"device_number,omitempty"`
	// The name of the volume to present in the filesystem.
	Filename string `protobuf:"bytes,5,opt,name=filename" json:"filename,omitempty"`
	// The volume size in bytes.
	VolumeSizeBytes uint64 `protobuf:"varint,6,opt,name=volume_size_bytes,json=volumeSizeBytes" json:"volume_size_bytes,omitempty"`
	// Volume statistics.
	Stats *FsVolumeStatistics `protobuf:"bytes,7,opt,name=stats" json:"stats,omitempty"`
	// Volume status, e.g. readiness.
	Status *FsVolumeStatus `protobuf:"bytes,8,opt,name=status" json:"status,omitempty"`
}

func (m *FsVolume) Reset()                    { *m = FsVolume{} }
func (m *FsVolume) String() string            { return proto.CompactTextString(m) }
func (*FsVolume) ProtoMessage()               {}
func (*FsVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FsVolume) GetCc() *common_v1.DataplaneCommon {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *FsVolume) GetVolumeId() uint32 {
	if m != nil {
		return m.VolumeId
	}
	return 0
}

func (m *FsVolume) GetNodeType() FsVolume_VolumeDeviceType {
	if m != nil {
		return m.NodeType
	}
	return FsVolume_FILE
}

func (m *FsVolume) GetDeviceNumber() uint32 {
	if m != nil {
		return m.DeviceNumber
	}
	return 0
}

func (m *FsVolume) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *FsVolume) GetVolumeSizeBytes() uint64 {
	if m != nil {
		return m.VolumeSizeBytes
	}
	return 0
}

func (m *FsVolume) GetStats() *FsVolumeStatistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *FsVolume) GetStatus() *FsVolumeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type FsVolumeList struct {
	Volumes []*FsVolume `protobuf:"bytes,1,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *FsVolumeList) Reset()                    { *m = FsVolumeList{} }
func (m *FsVolumeList) String() string            { return proto.CompactTextString(m) }
func (*FsVolumeList) ProtoMessage()               {}
func (*FsVolumeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FsVolumeList) GetVolumes() []*FsVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type FsPresentationListQuery struct {
	// A possibly-empty list of volume IDs to query.
	PresentationIds []uint32 `protobuf:"varint,1,rep,packed,name=presentation_ids,json=presentationIds" json:"presentation_ids,omitempty"`
}

func (m *FsPresentationListQuery) Reset()                    { *m = FsPresentationListQuery{} }
func (m *FsPresentationListQuery) String() string            { return proto.CompactTextString(m) }
func (*FsPresentationListQuery) ProtoMessage()               {}
func (*FsPresentationListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FsPresentationListQuery) GetPresentationIds() []uint32 {
	if m != nil {
		return m.PresentationIds
	}
	return nil
}

// *
// Presentation volume message for Fs RPCs.
//
// The minimum amount of information required to specify the 'presentation' or source volume,
// the volume that is presented to the user and (usually) mounted. All actual work
// is done on the target volume, which has actual storage associated with it.
type FsPresentation struct {
	Cc *common_v1.DataplaneCommon `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The inode the user mounts or opens.
	PresentationId uint32 `protobuf:"varint,2,opt,name=presentation_id,json=presentationId" json:"presentation_id,omitempty"`
	// The underlying inode of the StorageOS volume.
	TargetId uint32 `protobuf:"varint,3,opt,name=target_id,json=targetId" json:"target_id,omitempty"`
	// Presentation inode status, e.g. readiness.
	Status *FsVolumeStatus `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *FsPresentation) Reset()                    { *m = FsPresentation{} }
func (m *FsPresentation) String() string            { return proto.CompactTextString(m) }
func (*FsPresentation) ProtoMessage()               {}
func (*FsPresentation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *FsPresentation) GetCc() *common_v1.DataplaneCommon {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *FsPresentation) GetPresentationId() uint32 {
	if m != nil {
		return m.PresentationId
	}
	return 0
}

func (m *FsPresentation) GetTargetId() uint32 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func (m *FsPresentation) GetStatus() *FsVolumeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type FsPresentationList struct {
	Presentations []*FsPresentation `protobuf:"bytes,1,rep,name=presentations" json:"presentations,omitempty"`
}

func (m *FsPresentationList) Reset()                    { *m = FsPresentationList{} }
func (m *FsPresentationList) String() string            { return proto.CompactTextString(m) }
func (*FsPresentationList) ProtoMessage()               {}
func (*FsPresentationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *FsPresentationList) GetPresentations() []*FsPresentation {
	if m != nil {
		return m.Presentations
	}
	return nil
}

func init() {
	proto.RegisterType((*FsStatusRequest)(nil), "filesystem.v1.FsStatusRequest")
	proto.RegisterType((*FsStatus)(nil), "filesystem.v1.FsStatus")
	proto.RegisterType((*FsVolumeListQuery)(nil), "filesystem.v1.FsVolumeListQuery")
	proto.RegisterType((*FsVolumeStatistics)(nil), "filesystem.v1.FsVolumeStatistics")
	proto.RegisterType((*FsVolumeStatus)(nil), "filesystem.v1.FsVolumeStatus")
	proto.RegisterType((*FsVolume)(nil), "filesystem.v1.FsVolume")
	proto.RegisterType((*FsVolumeList)(nil), "filesystem.v1.FsVolumeList")
	proto.RegisterType((*FsPresentationListQuery)(nil), "filesystem.v1.FsPresentationListQuery")
	proto.RegisterType((*FsPresentation)(nil), "filesystem.v1.FsPresentation")
	proto.RegisterType((*FsPresentationList)(nil), "filesystem.v1.FsPresentationList")
	proto.RegisterEnum("filesystem.v1.FsVolumeState", FsVolumeState_name, FsVolumeState_value)
	proto.RegisterEnum("filesystem.v1.FsVolume_VolumeDeviceType", FsVolume_VolumeDeviceType_name, FsVolume_VolumeDeviceType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Fs service

type FsClient interface {
	// *
	// Get program status.
	Status(ctx context.Context, in *FsStatusRequest, opts ...grpc.CallOption) (*FsStatus, error)
	// *
	// Create the specified FsVolume.
	//
	// returns RpcResult
	VolumeCreate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Update the specified FsVolume.
	//
	// returns RpcResult
	VolumeUpdate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Delete the specified FsVolume.
	//
	// returns RpcResult
	VolumeDelete(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Return a list of FsVolume messages, optionally filtered using the supplied
	// FsVolumeListQuery message.
	//
	// returns A FsVolumeList message containing FsVolume objects,
	//         if any are found that match the filter.
	VolumeList(ctx context.Context, in *FsVolumeListQuery, opts ...grpc.CallOption) (*FsVolumeList, error)
	// *
	// Add configuration for a Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationCreate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Update configuration for a Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationUpdate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Remove configuration for the Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationDelete(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// List configured Presentation volumes, optionally filtered using a FsPresentationListQuery
	// message.
	//
	// returns A FsPresentationList message containing FsPresentation mesages,
	//         if any are found matching the filter.
	PresentationList(ctx context.Context, in *FsPresentationListQuery, opts ...grpc.CallOption) (*FsPresentationList, error)
}

type fsClient struct {
	cc *grpc.ClientConn
}

func NewFsClient(cc *grpc.ClientConn) FsClient {
	return &fsClient{cc}
}

func (c *fsClient) Status(ctx context.Context, in *FsStatusRequest, opts ...grpc.CallOption) (*FsStatus, error) {
	out := new(FsStatus)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) VolumeCreate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/VolumeCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) VolumeUpdate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/VolumeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) VolumeDelete(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/VolumeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) VolumeList(ctx context.Context, in *FsVolumeListQuery, opts ...grpc.CallOption) (*FsVolumeList, error) {
	out := new(FsVolumeList)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/VolumeList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) PresentationCreate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/PresentationCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) PresentationUpdate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/PresentationUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) PresentationDelete(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/PresentationDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) PresentationList(ctx context.Context, in *FsPresentationListQuery, opts ...grpc.CallOption) (*FsPresentationList, error) {
	out := new(FsPresentationList)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/PresentationList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Fs service

type FsServer interface {
	// *
	// Get program status.
	Status(context.Context, *FsStatusRequest) (*FsStatus, error)
	// *
	// Create the specified FsVolume.
	//
	// returns RpcResult
	VolumeCreate(context.Context, *FsVolume) (*common_v1.RpcResult, error)
	// *
	// Update the specified FsVolume.
	//
	// returns RpcResult
	VolumeUpdate(context.Context, *FsVolume) (*common_v1.RpcResult, error)
	// *
	// Delete the specified FsVolume.
	//
	// returns RpcResult
	VolumeDelete(context.Context, *FsVolume) (*common_v1.RpcResult, error)
	// *
	// Return a list of FsVolume messages, optionally filtered using the supplied
	// FsVolumeListQuery message.
	//
	// returns A FsVolumeList message containing FsVolume objects,
	//         if any are found that match the filter.
	VolumeList(context.Context, *FsVolumeListQuery) (*FsVolumeList, error)
	// *
	// Add configuration for a Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationCreate(context.Context, *FsPresentation) (*common_v1.RpcResult, error)
	// *
	// Update configuration for a Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationUpdate(context.Context, *FsPresentation) (*common_v1.RpcResult, error)
	// *
	// Remove configuration for the Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationDelete(context.Context, *FsPresentation) (*common_v1.RpcResult, error)
	// *
	// List configured Presentation volumes, optionally filtered using a FsPresentationListQuery
	// message.
	//
	// returns A FsPresentationList message containing FsPresentation mesages,
	//         if any are found matching the filter.
	PresentationList(context.Context, *FsPresentationListQuery) (*FsPresentationList, error)
}

func RegisterFsServer(s *grpc.Server, srv FsServer) {
	s.RegisterService(&_Fs_serviceDesc, srv)
}

func _Fs_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).Status(ctx, req.(*FsStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_VolumeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).VolumeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/VolumeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).VolumeCreate(ctx, req.(*FsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_VolumeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).VolumeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/VolumeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).VolumeUpdate(ctx, req.(*FsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_VolumeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).VolumeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/VolumeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).VolumeDelete(ctx, req.(*FsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_VolumeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsVolumeListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).VolumeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/VolumeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).VolumeList(ctx, req.(*FsVolumeListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_PresentationCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsPresentation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).PresentationCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/PresentationCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).PresentationCreate(ctx, req.(*FsPresentation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_PresentationUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsPresentation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).PresentationUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/PresentationUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).PresentationUpdate(ctx, req.(*FsPresentation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_PresentationDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsPresentation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).PresentationDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/PresentationDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).PresentationDelete(ctx, req.(*FsPresentation))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fs_PresentationList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FsPresentationListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServer).PresentationList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/filesystem.v1.Fs/PresentationList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServer).PresentationList(ctx, req.(*FsPresentationListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Fs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "filesystem.v1.Fs",
	HandlerType: (*FsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Fs_Status_Handler,
		},
		{
			MethodName: "VolumeCreate",
			Handler:    _Fs_VolumeCreate_Handler,
		},
		{
			MethodName: "VolumeUpdate",
			Handler:    _Fs_VolumeUpdate_Handler,
		},
		{
			MethodName: "VolumeDelete",
			Handler:    _Fs_VolumeDelete_Handler,
		},
		{
			MethodName: "VolumeList",
			Handler:    _Fs_VolumeList_Handler,
		},
		{
			MethodName: "PresentationCreate",
			Handler:    _Fs_PresentationCreate_Handler,
		},
		{
			MethodName: "PresentationUpdate",
			Handler:    _Fs_PresentationUpdate_Handler,
		},
		{
			MethodName: "PresentationDelete",
			Handler:    _Fs_PresentationDelete_Handler,
		},
		{
			MethodName: "PresentationList",
			Handler:    _Fs_PresentationList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "filesystem.proto",
}

func init() { proto.RegisterFile("filesystem.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0xd1, 0x6e, 0xda, 0x4a,
	0x10, 0xc5, 0x10, 0x08, 0x4c, 0x70, 0x70, 0x56, 0x91, 0x62, 0x91, 0x9b, 0x2b, 0xe2, 0x7b, 0xd5,
	0xd2, 0x54, 0x45, 0x0a, 0x55, 0xd5, 0x87, 0x3e, 0x54, 0x09, 0x04, 0x09, 0x25, 0x21, 0xed, 0xa6,
	0xad, 0x94, 0x27, 0xcb, 0x31, 0x93, 0xca, 0x12, 0xd8, 0xae, 0x77, 0x8d, 0x44, 0x7e, 0xa7, 0xfd,
	0x89, 0xfe, 0x5d, 0xb5, 0xbb, 0x26, 0x59, 0x88, 0xa0, 0x4a, 0xda, 0x27, 0xcc, 0xd9, 0x99, 0x73,
	0x66, 0xcf, 0x1c, 0x03, 0x58, 0x37, 0xc1, 0x08, 0xd9, 0x94, 0x71, 0x1c, 0xb7, 0xe2, 0x24, 0xe2,
	0x11, 0x31, 0x35, 0x64, 0x72, 0x58, 0xaf, 0xfa, 0xd1, 0x78, 0x1c, 0x85, 0xea, 0xd0, 0xd9, 0x82,
	0x5a, 0x8f, 0x5d, 0x72, 0x8f, 0xa7, 0x8c, 0xe2, 0xb7, 0x14, 0x19, 0x77, 0x5e, 0x41, 0x79, 0x06,
	0x91, 0x7d, 0xa8, 0x4e, 0x30, 0x61, 0x41, 0x14, 0xba, 0x41, 0x78, 0x13, 0xd9, 0x46, 0xc3, 0x68,
	0x56, 0xe8, 0x46, 0x86, 0xf5, 0xc3, 0x9b, 0xc8, 0x69, 0xc3, 0x56, 0x8f, 0x7d, 0x89, 0x46, 0xe9,
	0x18, 0xcf, 0x02, 0xc6, 0x3f, 0xa6, 0x98, 0x4c, 0xc9, 0x1e, 0xc0, 0x44, 0x42, 0x6e, 0x30, 0x64,
	0xb6, 0xd1, 0x28, 0x34, 0x4d, 0x5a, 0x51, 0x48, 0x7f, 0xc8, 0x9c, 0x6d, 0x20, 0xb3, 0x1e, 0x21,
	0x14, 0x30, 0x1e, 0xf8, 0xcc, 0x39, 0x87, 0x4d, 0x1d, 0x4d, 0x19, 0x79, 0x07, 0x10, 0x46, 0x43,
	0x74, 0x19, 0xf7, 0x38, 0x4a, 0xf1, 0xcd, 0xf6, 0x3f, 0xad, 0xb9, 0xfb, 0xb4, 0xf4, 0x16, 0xa4,
	0x15, 0x51, 0x2f, 0x1f, 0x9d, 0x1f, 0x05, 0x71, 0x11, 0x75, 0x48, 0x0e, 0x20, 0xef, 0xfb, 0x92,
	0x61, 0xa3, 0x5d, 0x6f, 0x65, 0x16, 0x4c, 0x0e, 0x5b, 0x5d, 0x8f, 0x7b, 0xf1, 0xc8, 0x0b, 0xb1,
	0x23, 0x21, 0x9a, 0xf7, 0x7d, 0xb2, 0x0b, 0x95, 0xbb, 0xe1, 0xed, 0x7c, 0xc3, 0x68, 0x9a, 0xb4,
	0x3c, 0x9b, 0x9d, 0x9c, 0x80, 0x94, 0x70, 0xf9, 0x34, 0x46, 0xbb, 0x20, 0x27, 0x6a, 0x2e, 0x99,
	0xa8, 0xa5, 0x3e, 0xba, 0x38, 0x09, 0x7c, 0xfc, 0x34, 0x8d, 0x91, 0x96, 0x45, 0xab, 0x78, 0x22,
	0xff, 0x81, 0x39, 0x94, 0xb8, 0x1b, 0xa6, 0xe3, 0x6b, 0x4c, 0xec, 0x35, 0xa9, 0x53, 0x55, 0xe0,
	0x40, 0x62, 0xa4, 0x0e, 0x65, 0xc1, 0x1c, 0x7a, 0x63, 0xb4, 0x8b, 0xd2, 0xf9, 0xbb, 0xef, 0xe4,
	0x00, 0xb6, 0xb2, 0x21, 0x59, 0x70, 0x8b, 0xee, 0xf5, 0x94, 0x23, 0xb3, 0x4b, 0x0d, 0xa3, 0xb9,
	0x46, 0x6b, 0xea, 0xe0, 0x32, 0xb8, 0xc5, 0x63, 0x01, 0x93, 0xb7, 0x50, 0x14, 0x0e, 0x32, 0x7b,
	0x5d, 0xde, 0x7f, 0x7f, 0x85, 0x83, 0x6a, 0x15, 0x54, 0xd5, 0x93, 0x37, 0x50, 0x62, 0x72, 0x13,
	0x76, 0x59, 0x76, 0xee, 0xad, 0xe8, 0x4c, 0x19, 0xcd, 0x8a, 0x9d, 0x97, 0x60, 0x2d, 0x5e, 0x9d,
	0x94, 0x61, 0xad, 0xd7, 0x3f, 0x3b, 0xb1, 0x72, 0xc4, 0x84, 0xca, 0xe0, 0xb8, 0xeb, 0x1e, 0x9f,
	0x5d, 0x74, 0x4e, 0x2d, 0xc3, 0x39, 0x82, 0xaa, 0x9e, 0x1f, 0x72, 0x08, 0xeb, 0x6a, 0x7e, 0x95,
	0x9b, 0x8d, 0xf6, 0xce, 0x12, 0x51, 0x3a, 0xab, 0x73, 0xba, 0xb0, 0xd3, 0x63, 0x1f, 0x12, 0x64,
	0x18, 0x8a, 0x2b, 0x44, 0xe1, 0x7d, 0x10, 0x5f, 0x80, 0x15, 0x6b, 0x07, 0x5a, 0x1c, 0x6b, 0x3a,
	0x2e, 0x42, 0xf9, 0xd3, 0x10, 0xf9, 0xd3, 0x69, 0x1e, 0x95, 0x9a, 0xe7, 0x50, 0x5b, 0x50, 0xca,
	0xb2, 0xb3, 0x39, 0x2f, 0x24, 0xe2, 0xc5, 0xbd, 0xe4, 0x2b, 0x72, 0x51, 0x52, 0x50, 0xf1, 0x52,
	0x40, 0x7f, 0xa8, 0x39, 0x5e, 0x7c, 0x8c, 0xe3, 0x57, 0xe2, 0x85, 0x5a, 0x74, 0x80, 0x74, 0xc0,
	0xd4, 0xb5, 0x67, 0x86, 0x3e, 0xe4, 0xd4, 0x3b, 0xe9, 0x7c, 0xcf, 0xc1, 0xff, 0x60, 0xce, 0xbd,
	0x62, 0x62, 0x93, 0x83, 0x8b, 0x81, 0xd8, 0x64, 0x05, 0x8a, 0xf4, 0xe4, 0xa8, 0x7b, 0x65, 0x19,
	0xed, 0xef, 0x45, 0xc8, 0xf7, 0x18, 0xe9, 0x40, 0x29, 0x7b, 0x75, 0xff, 0x7d, 0x20, 0x32, 0xf7,
	0x2b, 0x53, 0xdf, 0x59, 0x72, 0xee, 0xe4, 0xc8, 0x7b, 0xa8, 0x2a, 0xbd, 0x4e, 0x82, 0x42, 0x70,
	0x59, 0x00, 0xea, 0xdb, 0xda, 0x4a, 0x68, 0xec, 0x53, 0x64, 0xe9, 0x88, 0xeb, 0x04, 0x9f, 0xe3,
	0xe1, 0x9f, 0x11, 0x74, 0x71, 0x84, 0x4f, 0x21, 0x38, 0x07, 0xd0, 0x22, 0xdd, 0x58, 0xd2, 0x7e,
	0x17, 0xd3, 0xfa, 0xee, 0x8a, 0x0a, 0x27, 0x47, 0x4e, 0x81, 0xe8, 0x2b, 0xca, 0x7c, 0x59, 0xbd,
	0xc7, 0xa5, 0xb3, 0x2d, 0x90, 0x65, 0x1e, 0xfd, 0x1d, 0xb2, 0xcc, 0xaf, 0x27, 0x92, 0xb9, 0x60,
	0x3d, 0xc8, 0xf0, 0xb3, 0x95, 0x54, 0xf7, 0x0e, 0xee, 0xff, 0xb6, 0xce, 0xc9, 0x5d, 0x97, 0xe4,
	0x9f, 0xde, 0xeb, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x14, 0x82, 0xab, 0x25, 0x07, 0x00,
	0x00,
}
