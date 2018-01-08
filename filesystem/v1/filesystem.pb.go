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
import common_v1 "code.storageos.net/scm/storageos/service/common/v1"

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

type FsVolume_VolumeControlStatus int32

const (
	FsVolume_NONE        FsVolume_VolumeControlStatus = 0
	FsVolume_ACTIVE      FsVolume_VolumeControlStatus = 1
	FsVolume_UNAVAILABLE FsVolume_VolumeControlStatus = 2
	FsVolume_FAILED      FsVolume_VolumeControlStatus = 3
	FsVolume_DELETING    FsVolume_VolumeControlStatus = 4
)

var FsVolume_VolumeControlStatus_name = map[int32]string{
	0: "NONE",
	1: "ACTIVE",
	2: "UNAVAILABLE",
	3: "FAILED",
	4: "DELETING",
}
var FsVolume_VolumeControlStatus_value = map[string]int32{
	"NONE":        0,
	"ACTIVE":      1,
	"UNAVAILABLE": 2,
	"FAILED":      3,
	"DELETING":    4,
}

func (x FsVolume_VolumeControlStatus) String() string {
	return proto.EnumName(FsVolume_VolumeControlStatus_name, int32(x))
}
func (FsVolume_VolumeControlStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{5, 1}
}

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
	// The control plane's desired state for this volume.
	ControlStatus FsVolume_VolumeControlStatus `protobuf:"varint,10,opt,name=control_status,json=controlStatus,enum=filesystem.v1.FsVolume_VolumeControlStatus" json:"control_status,omitempty"`
	// An opaque value interpreted based on node_type.
	DeviceNumber uint32 `protobuf:"varint,4,opt,name=device_number,json=deviceNumber" json:"device_number,omitempty"`
	// The filename of the underlying volume to present in the filesystem.
	Filename string `protobuf:"bytes,5,opt,name=filename" json:"filename,omitempty"`
	// The filename of the presentation node, to which users actually connect.
	PresentationFilename string `protobuf:"bytes,9,opt,name=presentation_filename,json=presentationFilename" json:"presentation_filename,omitempty"`
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

func (m *FsVolume) GetControlStatus() FsVolume_VolumeControlStatus {
	if m != nil {
		return m.ControlStatus
	}
	return FsVolume_NONE
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

func (m *FsVolume) GetPresentationFilename() string {
	if m != nil {
		return m.PresentationFilename
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
//
// In order to discover information about the type of volume to present, the target
// volume has to be fetched regardless. As a result, there's no need to duplicate fields
// here, we just need enough information to get from the presentation node to the target,
// and to know when the presentation is actually available to the user.
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
	proto.RegisterEnum("filesystem.v1.FsVolume_VolumeControlStatus", FsVolume_VolumeControlStatus_name, FsVolume_VolumeControlStatus_value)
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
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x5d, 0x6f, 0xe2, 0x46,
	0x14, 0xc5, 0x7c, 0x05, 0x2e, 0x18, 0x9c, 0x69, 0xaa, 0x58, 0xa4, 0xa9, 0x88, 0x5b, 0xb5, 0x34,
	0x51, 0x91, 0x42, 0x54, 0xf5, 0xa1, 0x0f, 0x15, 0x60, 0x53, 0x59, 0x21, 0xa4, 0x9d, 0x24, 0x48,
	0x79, 0xb2, 0x88, 0x99, 0x54, 0x96, 0xc0, 0x66, 0x3d, 0x63, 0x24, 0xf2, 0xbe, 0xbf, 0x64, 0x7f,
	0xc5, 0xfe, 0xbb, 0xd5, 0x8c, 0x0d, 0x19, 0xc8, 0xc2, 0x2a, 0xd9, 0x7d, 0xc2, 0x3e, 0x73, 0xef,
	0xb9, 0x77, 0xce, 0xb9, 0xd7, 0x80, 0xf6, 0xe8, 0x4d, 0x08, 0x5d, 0x50, 0x46, 0xa6, 0xcd, 0x59,
	0x18, 0xb0, 0x00, 0xa9, 0x12, 0x32, 0x3f, 0xaf, 0x95, 0xdd, 0x60, 0x3a, 0x0d, 0xfc, 0xf8, 0xd0,
	0xd8, 0x87, 0x6a, 0x8f, 0xde, 0xb0, 0x11, 0x8b, 0x28, 0x26, 0xef, 0x22, 0x42, 0x99, 0xf1, 0x3b,
	0x14, 0x96, 0x10, 0x3a, 0x81, 0xf2, 0x9c, 0x84, 0xd4, 0x0b, 0x7c, 0xc7, 0xf3, 0x1f, 0x03, 0x5d,
	0xa9, 0x2b, 0x8d, 0x22, 0x2e, 0x25, 0x98, 0xed, 0x3f, 0x06, 0x46, 0x0b, 0xf6, 0x7b, 0x74, 0x18,
	0x4c, 0xa2, 0x29, 0xe9, 0x7b, 0x94, 0xfd, 0x17, 0x91, 0x70, 0x81, 0x8e, 0x01, 0xe6, 0x02, 0x72,
	0xbc, 0x31, 0xd5, 0x95, 0x7a, 0xa6, 0xa1, 0xe2, 0x62, 0x8c, 0xd8, 0x63, 0x6a, 0x1c, 0x00, 0x5a,
	0xe6, 0xf0, 0x42, 0x1e, 0x65, 0x9e, 0x4b, 0x8d, 0x2b, 0xa8, 0xc8, 0x68, 0x44, 0xd1, 0x5f, 0x00,
	0x7e, 0x30, 0x26, 0x0e, 0x65, 0x23, 0x46, 0x44, 0xf1, 0x4a, 0xeb, 0x87, 0xe6, 0xda, 0x7d, 0x9a,
	0x72, 0x0a, 0xc1, 0x45, 0x1e, 0x2f, 0x1e, 0x8d, 0xf7, 0x39, 0x7e, 0x91, 0xf8, 0x10, 0x9d, 0x42,
	0xda, 0x75, 0x05, 0x43, 0xa9, 0x55, 0x6b, 0x26, 0x12, 0xcc, 0xcf, 0x9b, 0xe6, 0x88, 0x8d, 0x66,
	0x93, 0x91, 0x4f, 0xba, 0x02, 0xc2, 0x69, 0xd7, 0x45, 0x47, 0x50, 0x5c, 0x35, 0xaf, 0xa7, 0xeb,
	0x4a, 0x43, 0xc5, 0x85, 0x65, 0xef, 0xc8, 0x02, 0x51, 0xc2, 0x61, 0x8b, 0x19, 0xd1, 0x33, 0xa2,
	0xa3, 0xc6, 0x96, 0x8e, 0x9a, 0xf1, 0x8f, 0x49, 0xe6, 0x9e, 0x4b, 0x6e, 0x17, 0x33, 0x82, 0x0b,
	0x3c, 0x95, 0x3f, 0x21, 0x0c, 0x15, 0x37, 0xf0, 0x59, 0x18, 0x4c, 0xc4, 0xe5, 0x22, 0xaa, 0x83,
	0xe0, 0x3a, 0xdb, 0xcd, 0xd5, 0x8d, 0x73, 0x12, 0xc3, 0x54, 0x57, 0x7e, 0x45, 0x3f, 0x81, 0x3a,
	0x16, 0xb5, 0x1c, 0x3f, 0x9a, 0x3e, 0x90, 0x50, 0xcf, 0x8a, 0xde, 0xcb, 0x31, 0x38, 0x10, 0x18,
	0xaa, 0x41, 0x81, 0x57, 0xf0, 0x47, 0x53, 0xa2, 0xe7, 0x84, 0x9b, 0xab, 0x77, 0x74, 0x01, 0xdf,
	0xcf, 0x42, 0x42, 0x89, 0xcf, 0x3d, 0x09, 0x7c, 0x67, 0x15, 0x58, 0x14, 0x81, 0x07, 0xf2, 0x61,
	0x6f, 0x99, 0x74, 0x0a, 0xfb, 0x89, 0x5a, 0xd4, 0x7b, 0x22, 0xce, 0xc3, 0x82, 0x11, 0xaa, 0xe7,
	0xeb, 0x4a, 0x23, 0x8b, 0xab, 0xf1, 0xc1, 0x8d, 0xf7, 0x44, 0x3a, 0x1c, 0x46, 0x7f, 0x42, 0x8e,
	0xdf, 0x96, 0xea, 0x7b, 0xc2, 0x88, 0x93, 0x1d, 0x56, 0xc6, 0x33, 0x81, 0xe3, 0x78, 0xf4, 0x07,
	0xe4, 0x13, 0x99, 0x0a, 0x22, 0xf3, 0x78, 0x47, 0x66, 0x44, 0x71, 0x12, 0x6c, 0x9c, 0x81, 0xb6,
	0xe9, 0x01, 0x2a, 0x40, 0xb6, 0x67, 0xf7, 0x2d, 0x2d, 0x85, 0x54, 0x28, 0x0e, 0x3a, 0xa6, 0xd3,
	0xe9, 0x5f, 0x77, 0x2f, 0x35, 0xc5, 0x18, 0xc2, 0x77, 0x9f, 0x11, 0x99, 0xc7, 0x0f, 0xae, 0x07,
	0x3c, 0x1e, 0x20, 0xdf, 0xee, 0xde, 0xda, 0x43, 0x4b, 0x53, 0x50, 0x15, 0x4a, 0x77, 0x83, 0xf6,
	0xb0, 0x6d, 0xf7, 0xdb, 0x9d, 0xbe, 0xa5, 0xa5, 0xf9, 0x61, 0xaf, 0x6d, 0xf7, 0x2d, 0x53, 0xcb,
	0xa0, 0x32, 0x14, 0x4c, 0xab, 0x6f, 0xdd, 0xda, 0x83, 0x7f, 0xb4, 0xac, 0xd1, 0x86, 0xb2, 0xbc,
	0x20, 0xe8, 0x1c, 0xf6, 0x62, 0x5d, 0xe2, 0xc5, 0x28, 0xb5, 0x0e, 0xb7, 0x5c, 0x06, 0x2f, 0xe3,
	0x0c, 0x13, 0x0e, 0x7b, 0xf4, 0x5f, 0x49, 0xfd, 0xe7, 0x4d, 0xfb, 0x0d, 0xb4, 0x35, 0xcf, 0x9e,
	0xf7, 0xad, 0x2a, 0xe3, 0x7c, 0xeb, 0x3e, 0x2a, 0x7c, 0xc1, 0x64, 0x9a, 0x57, 0xad, 0xc5, 0xaf,
	0x50, 0xdd, 0xa8, 0x94, 0x2c, 0x47, 0x65, 0xbd, 0x10, 0xdf, 0x1f, 0x36, 0x0a, 0xff, 0x27, 0x8c,
	0x87, 0x64, 0xe2, 0xfd, 0x89, 0x01, 0x7b, 0x2c, 0x39, 0x99, 0x7b, 0x8d, 0x93, 0xf7, 0xfc, 0x8b,
	0xb1, 0xa9, 0x00, 0xea, 0x82, 0x2a, 0xd7, 0x5e, 0x0a, 0xfa, 0x92, 0x53, 0xce, 0xc4, 0xeb, 0x39,
	0xa7, 0x3f, 0x83, 0xba, 0xf6, 0x0d, 0x91, 0x1c, 0x2f, 0x42, 0x0e, 0x5b, 0x6d, 0xf3, 0x5e, 0x53,
	0x5a, 0x1f, 0x72, 0x90, 0xee, 0x51, 0xd4, 0x85, 0x7c, 0x32, 0x17, 0x3f, 0xbe, 0x28, 0xb2, 0xf6,
	0x19, 0xad, 0x1d, 0x6e, 0x39, 0x37, 0x52, 0xe8, 0x6f, 0x28, 0x27, 0x93, 0x16, 0x12, 0x5e, 0x70,
	0xdb, 0x00, 0xd4, 0x0e, 0x24, 0x4b, 0xf0, 0xcc, 0xc5, 0x84, 0x46, 0x13, 0x26, 0x13, 0xdc, 0xcd,
	0xc6, 0x5f, 0x47, 0x60, 0x92, 0x09, 0x79, 0x0b, 0xc1, 0x15, 0x80, 0x34, 0xd2, 0xf5, 0x2d, 0xe9,
	0xab, 0x31, 0xad, 0x1d, 0xed, 0x88, 0x30, 0x52, 0xe8, 0x12, 0x90, 0x6c, 0x51, 0xa2, 0xcb, 0x6e,
	0x1f, 0xb7, 0xf6, 0xb6, 0x41, 0x96, 0x68, 0xf4, 0x6d, 0xc8, 0x12, 0xbd, 0xde, 0x48, 0xe6, 0x80,
	0xf6, 0x62, 0x86, 0x7f, 0xd9, 0x49, 0xf5, 0xac, 0xe0, 0xc9, 0x17, 0xe3, 0x8c, 0xd4, 0x43, 0x5e,
	0xfc, 0xab, 0x5f, 0x7c, 0x0a, 0x00, 0x00, 0xff, 0xff, 0x8e, 0x27, 0xc3, 0x2c, 0x06, 0x08, 0x00,
	0x00,
}