// Code generated by protoc-gen-go. DO NOT EDIT.
// source: filesystem.proto

/*
Package filesystem_v1 is a generated protocol buffer package.

It is generated from these files:
	filesystem.proto

It has these top-level messages:
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
import common "."

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
func (FsVolume_VolumeDeviceType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3, 0} }

type FsVolumeListQuery struct {
	// A possibly-empty list of volume IDs to query.
	VolumeIds []uint32 `protobuf:"varint,1,rep,packed,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
}

func (m *FsVolumeListQuery) Reset()                    { *m = FsVolumeListQuery{} }
func (m *FsVolumeListQuery) String() string            { return proto.CompactTextString(m) }
func (*FsVolumeListQuery) ProtoMessage()               {}
func (*FsVolumeListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
func (*FsVolumeStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type FsVolumeStatus struct {
	// State of the volume device node (file/block device).
	NodeState FsVolumeState `protobuf:"varint,1,opt,name=node_state,json=nodeState,enum=filesystem.v1.FsVolumeState" json:"node_state,omitempty"`
}

func (m *FsVolumeStatus) Reset()                    { *m = FsVolumeStatus{} }
func (m *FsVolumeStatus) String() string            { return proto.CompactTextString(m) }
func (*FsVolumeStatus) ProtoMessage()               {}
func (*FsVolumeStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *FsVolumeStatus) GetNodeState() FsVolumeState {
	if m != nil {
		return m.NodeState
	}
	return FsVolumeState_NONE
}

// *
// A StorageOS volume to be presented via the FUSE filesystem.
type FsVolume struct {
	Cc *common.DataplaneCommon `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The volume ID to represent.
	VolumeId uint32 `protobuf:"varint,2,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
	// The type for this volume.
	NodeType FsVolume_VolumeDeviceType `protobuf:"varint,3,opt,name=node_type,json=nodeType,enum=filesystem.v1.FsVolume_VolumeDeviceType" json:"node_type,omitempty"`
	// An opaque value interpreted based on node_type.
	DeviceNumber uint32 `protobuf:"varint,4,opt,name=device_number,json=deviceNumber" json:"device_number,omitempty"`
	// The name of the volume to present in the filesystem.
	Filename string `protobuf:"bytes,5,opt,name=filename" json:"filename,omitempty"`
	// True if this volume should be backed by another volume.
	LinkedVolume bool `protobuf:"varint,6,opt,name=linked_volume,json=linkedVolume" json:"linked_volume,omitempty"`
	// If linked_volume is true, this is the volume to which we link.
	TargetVolumeId uint32 `protobuf:"varint,7,opt,name=target_volume_id,json=targetVolumeId" json:"target_volume_id,omitempty"`
	// The volume size in bytes.
	VolumeSizeBytes uint64 `protobuf:"varint,8,opt,name=volume_size_bytes,json=volumeSizeBytes" json:"volume_size_bytes,omitempty"`
	// Volume statistics.
	Stats *FsVolumeStatistics `protobuf:"bytes,9,opt,name=stats" json:"stats,omitempty"`
	// Volume status, e.g. readiness.
	Status *FsVolumeStatus `protobuf:"bytes,10,opt,name=status" json:"status,omitempty"`
}

func (m *FsVolume) Reset()                    { *m = FsVolume{} }
func (m *FsVolume) String() string            { return proto.CompactTextString(m) }
func (*FsVolume) ProtoMessage()               {}
func (*FsVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *FsVolume) GetCc() *common.DataplaneCommon {
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

func (m *FsVolume) GetLinkedVolume() bool {
	if m != nil {
		return m.LinkedVolume
	}
	return false
}

func (m *FsVolume) GetTargetVolumeId() uint32 {
	if m != nil {
		return m.TargetVolumeId
	}
	return 0
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
func (*FsVolumeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *FsVolumeList) GetVolumes() []*FsVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type FsPresentationListQuery struct {
	// A possibly-empty list of volume IDs to query.
	PresentationId []uint32 `protobuf:"varint,1,rep,packed,name=presentation_id,json=presentationId" json:"presentation_id,omitempty"`
}

func (m *FsPresentationListQuery) Reset()                    { *m = FsPresentationListQuery{} }
func (m *FsPresentationListQuery) String() string            { return proto.CompactTextString(m) }
func (*FsPresentationListQuery) ProtoMessage()               {}
func (*FsPresentationListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *FsPresentationListQuery) GetPresentationId() []uint32 {
	if m != nil {
		return m.PresentationId
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
	Cc *common.DataplaneCommon `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The inode the user mounts or opens.
	SourceId uint32 `protobuf:"varint,2,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	// The underlying inode of the StorageOS volume.
	TargetId uint32 `protobuf:"varint,3,opt,name=target_id,json=targetId" json:"target_id,omitempty"`
	// Presentation inode status, e.g. readiness.
	Status *FsVolumeStatus `protobuf:"bytes,5,opt,name=status" json:"status,omitempty"`
}

func (m *FsPresentation) Reset()                    { *m = FsPresentation{} }
func (m *FsPresentation) String() string            { return proto.CompactTextString(m) }
func (*FsPresentation) ProtoMessage()               {}
func (*FsPresentation) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *FsPresentation) GetCc() *common.DataplaneCommon {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *FsPresentation) GetSourceId() uint32 {
	if m != nil {
		return m.SourceId
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
func (*FsPresentationList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *FsPresentationList) GetPresentations() []*FsPresentation {
	if m != nil {
		return m.Presentations
	}
	return nil
}

func init() {
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
	// Create the specified FsVolume.
	//
	// returns RpcResult
	VolumeCreate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Update the specified FsVolume.
	//
	// returns RpcResult
	VolumeUpdate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Delete the specified FsVolume.
	//
	// returns RpcResult
	VolumeDelete(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
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
	PresentationCreate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Update configuration for a Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationUpdate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Remove configuration for the Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationDelete(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common.RpcResult, error)
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

func (c *fsClient) VolumeCreate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/VolumeCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) VolumeUpdate(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/VolumeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) VolumeDelete(ctx context.Context, in *FsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
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

func (c *fsClient) PresentationCreate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/PresentationCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) PresentationUpdate(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/filesystem.v1.Fs/PresentationUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClient) PresentationDelete(ctx context.Context, in *FsPresentation, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
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
	// Create the specified FsVolume.
	//
	// returns RpcResult
	VolumeCreate(context.Context, *FsVolume) (*common.RpcResult, error)
	// *
	// Update the specified FsVolume.
	//
	// returns RpcResult
	VolumeUpdate(context.Context, *FsVolume) (*common.RpcResult, error)
	// *
	// Delete the specified FsVolume.
	//
	// returns RpcResult
	VolumeDelete(context.Context, *FsVolume) (*common.RpcResult, error)
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
	PresentationCreate(context.Context, *FsPresentation) (*common.RpcResult, error)
	// *
	// Update configuration for a Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationUpdate(context.Context, *FsPresentation) (*common.RpcResult, error)
	// *
	// Remove configuration for the Presentation volume specified in the FsPresentation message.
	//
	// returns RpcResult
	PresentationDelete(context.Context, *FsPresentation) (*common.RpcResult, error)
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
	// 643 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xd1, 0x6a, 0xdb, 0x30,
	0x14, 0x8d, 0xd3, 0xa4, 0xb5, 0x6f, 0xe3, 0xd4, 0x15, 0x83, 0x9a, 0x74, 0x05, 0xd7, 0x1b, 0x9b,
	0xe9, 0x20, 0xd0, 0x8c, 0xb1, 0x87, 0x3d, 0x35, 0x49, 0x03, 0x61, 0x6d, 0xba, 0xa9, 0x5b, 0xa1,
	0x4f, 0xc6, 0xb5, 0xb5, 0x61, 0x96, 0xd8, 0xc6, 0x92, 0x0b, 0xe9, 0xe7, 0x6c, 0x5f, 0xb7, 0xbf,
	0x18, 0x92, 0x9c, 0x46, 0x49, 0x49, 0xd6, 0x6e, 0x4f, 0x71, 0x8e, 0xee, 0x39, 0xba, 0xf7, 0xdc,
	0x63, 0x0c, 0xd6, 0xb7, 0x78, 0x4c, 0xe8, 0x94, 0x32, 0x32, 0x69, 0x67, 0x79, 0xca, 0x52, 0x64,
	0x2a, 0xc8, 0xed, 0x71, 0xab, 0x11, 0xa6, 0x93, 0x49, 0x9a, 0xc8, 0x43, 0xb7, 0x03, 0xbb, 0x03,
	0x7a, 0x95, 0x8e, 0x8b, 0x09, 0x39, 0x8b, 0x29, 0xfb, 0x5c, 0x90, 0x7c, 0x8a, 0x0e, 0x00, 0x6e,
	0x05, 0xe4, 0xc7, 0x11, 0xb5, 0x35, 0x67, 0xc3, 0x33, 0xb1, 0x21, 0x91, 0x61, 0x44, 0xdd, 0x67,
	0x80, 0x66, 0x9c, 0x4b, 0x16, 0xb0, 0x98, 0xb2, 0x38, 0xa4, 0xee, 0x39, 0x34, 0x55, 0xb4, 0xa0,
	0xe8, 0x03, 0x40, 0x92, 0x46, 0xc4, 0xa7, 0x2c, 0x60, 0xc4, 0xd6, 0x1c, 0xcd, 0x6b, 0x76, 0x9e,
	0xb7, 0x17, 0xba, 0x69, 0xab, 0x14, 0x82, 0x0d, 0x5e, 0x2f, 0x1e, 0xdd, 0xdf, 0x1b, 0xa0, 0xcf,
	0x0e, 0x91, 0x03, 0xd5, 0x30, 0x14, 0x0a, 0xdb, 0x1d, 0xab, 0xdd, 0x0f, 0x58, 0x90, 0x8d, 0x83,
	0x84, 0xf4, 0xc4, 0x24, 0xb8, 0x1a, 0x86, 0x68, 0x1f, 0x8c, 0xfb, 0x96, 0xed, 0xaa, 0xa3, 0x79,
	0x26, 0xd6, 0x67, 0x1d, 0xa3, 0x53, 0x10, 0xc2, 0x3e, 0x9b, 0x66, 0xc4, 0xde, 0x10, 0x7d, 0x78,
	0x2b, 0xfa, 0x68, 0xcb, 0x9f, 0x3e, 0xb9, 0x8d, 0x43, 0xf2, 0x65, 0x9a, 0x11, 0xac, 0x73, 0x2a,
	0x7f, 0x42, 0x2f, 0xc0, 0x8c, 0x04, 0xee, 0x27, 0xc5, 0xe4, 0x86, 0xe4, 0x76, 0x4d, 0xdc, 0xd3,
	0x90, 0xe0, 0x48, 0x60, 0xa8, 0x05, 0x3a, 0x57, 0x4e, 0x82, 0x09, 0xb1, 0xeb, 0x8e, 0xe6, 0x19,
	0xf8, 0xfe, 0x3f, 0x17, 0x18, 0xc7, 0xc9, 0x0f, 0x12, 0xf9, 0xb2, 0x35, 0x7b, 0xd3, 0xd1, 0x3c,
	0x1d, 0x37, 0x24, 0x58, 0xce, 0xea, 0x81, 0xc5, 0x82, 0xfc, 0x3b, 0x61, 0xfe, 0x7c, 0xa0, 0x2d,
	0x71, 0x51, 0x53, 0xe2, 0x57, 0xb3, 0xb1, 0x8e, 0x60, 0xb7, 0x2c, 0xa1, 0xf1, 0x1d, 0xf1, 0x6f,
	0xa6, 0x8c, 0x50, 0x5b, 0x77, 0x34, 0xaf, 0x86, 0x77, 0xe4, 0xc1, 0x65, 0x7c, 0x47, 0xba, 0x1c,
	0x46, 0xef, 0xa1, 0xce, 0xd7, 0x40, 0x6d, 0x43, 0x98, 0x78, 0xb8, 0x66, 0x0d, 0x72, 0x9f, 0x58,
	0xd6, 0xa3, 0x77, 0xb0, 0x49, 0xc5, 0x3a, 0x6d, 0x10, 0xcc, 0x83, 0x35, 0xcc, 0x82, 0xe2, 0xb2,
	0xd8, 0x7d, 0x03, 0xd6, 0xb2, 0x93, 0x48, 0x87, 0xda, 0x60, 0x78, 0x76, 0x6a, 0x55, 0x90, 0x09,
	0xc6, 0xa8, 0xdb, 0xf7, 0xbb, 0x67, 0x17, 0xbd, 0x8f, 0x96, 0xe6, 0x9e, 0x40, 0x43, 0x0d, 0x21,
	0x3a, 0x86, 0x2d, 0xd9, 0xbf, 0x0c, 0xdf, 0x76, 0x67, 0x6f, 0xc5, 0xa5, 0x78, 0x56, 0xe7, 0x76,
	0x61, 0x6f, 0x40, 0x3f, 0xe5, 0x84, 0x92, 0x84, 0x8f, 0x90, 0x26, 0xf3, 0x34, 0xbf, 0x86, 0x9d,
	0x4c, 0x39, 0xe0, 0x7e, 0xca, 0x48, 0x37, 0x55, 0x78, 0x18, 0xb9, 0xbf, 0x34, 0x1e, 0x61, 0x55,
	0xe4, 0x71, 0xc1, 0xa3, 0x69, 0x91, 0x87, 0x6a, 0xf0, 0x24, 0x30, 0x8c, 0xf8, 0x61, 0xb9, 0xcb,
	0x38, 0x12, 0xc1, 0x33, 0xb1, 0x2e, 0x81, 0x61, 0xa4, 0x38, 0x5b, 0x7f, 0x8a, 0xb3, 0xd7, 0xfc,
	0xed, 0x5b, 0x9e, 0x14, 0xf5, 0xc0, 0x54, 0xa7, 0x99, 0x19, 0xf7, 0x50, 0x53, 0x65, 0xe2, 0x45,
	0xce, 0xd1, 0x4b, 0x30, 0x17, 0xde, 0x47, 0xbe, 0xb1, 0xd1, 0xc5, 0x88, 0x6f, 0xcc, 0x80, 0x3a,
	0x3e, 0x3d, 0xe9, 0x5f, 0x5b, 0x5a, 0xe7, 0x67, 0x0d, 0xaa, 0x03, 0x1e, 0x8c, 0x86, 0x2c, 0xed,
	0xe5, 0x84, 0xd7, 0xae, 0xda, 0x51, 0x0b, 0xda, 0x38, 0x0b, 0x31, 0xa1, 0xc5, 0x98, 0xb9, 0x95,
	0x39, 0xed, 0x6b, 0x16, 0xfd, 0x0b, 0xad, 0x4f, 0xc6, 0xe4, 0xf1, 0xb4, 0x73, 0x00, 0x25, 0x57,
	0xce, 0x0a, 0xd2, 0x7d, 0x56, 0x5a, 0xfb, 0x6b, 0x2a, 0xdc, 0x0a, 0x3a, 0x01, 0xa4, 0xfa, 0x57,
	0x4e, 0xbe, 0xde, 0xe4, 0xa5, 0x8e, 0x96, 0x24, 0x4a, 0x17, 0xfe, 0x47, 0xa2, 0x74, 0xe4, 0x49,
	0x12, 0x3e, 0x58, 0x0f, 0x22, 0xf4, 0x6a, 0xad, 0xc0, 0xdc, 0xa3, 0xc3, 0xbf, 0xd6, 0xb9, 0x95,
	0x9b, 0x4d, 0xf1, 0x79, 0x79, 0xfb, 0x27, 0x00, 0x00, 0xff, 0xff, 0x38, 0x59, 0xd0, 0xd4, 0x8f,
	0x06, 0x00, 0x00,
}
