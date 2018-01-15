// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rdbplugin.proto

/*
Package rdbplugin_v1 is a generated protocol buffer package.

It is generated from these files:
	rdbplugin.proto

It has these top-level messages:
	RdbStatusRequest
	RdbStatus
	RdbVolumeListQuery
	RdbVolumeCredentials
	RdbVolumeStatistics
	RdbVolumeStatus
	RdbVolume
	RdbVolumeList
*/
package rdbplugin_v1

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

type RdbStatusRequest struct {
}

func (m *RdbStatusRequest) Reset()                    { *m = RdbStatusRequest{} }
func (m *RdbStatusRequest) String() string            { return proto.CompactTextString(m) }
func (*RdbStatusRequest) ProtoMessage()               {}
func (*RdbStatusRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type RdbStatus struct {
	// The version control info string.
	VersionInfo string `protobuf:"bytes,1,opt,name=version_info,json=versionInfo" json:"version_info,omitempty"`
}

func (m *RdbStatus) Reset()                    { *m = RdbStatus{} }
func (m *RdbStatus) String() string            { return proto.CompactTextString(m) }
func (*RdbStatus) ProtoMessage()               {}
func (*RdbStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *RdbStatus) GetVersionInfo() string {
	if m != nil {
		return m.VersionInfo
	}
	return ""
}

type RdbVolumeListQuery struct {
	// Optional list of volume IDs to query.
	VolumeIds []uint32 `protobuf:"varint,1,rep,packed,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
}

func (m *RdbVolumeListQuery) Reset()                    { *m = RdbVolumeListQuery{} }
func (m *RdbVolumeListQuery) String() string            { return proto.CompactTextString(m) }
func (*RdbVolumeListQuery) ProtoMessage()               {}
func (*RdbVolumeListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *RdbVolumeListQuery) GetVolumeIds() []uint32 {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

type RdbVolumeCredentials struct {
}

func (m *RdbVolumeCredentials) Reset()                    { *m = RdbVolumeCredentials{} }
func (m *RdbVolumeCredentials) String() string            { return proto.CompactTextString(m) }
func (*RdbVolumeCredentials) ProtoMessage()               {}
func (*RdbVolumeCredentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type RdbVolumeStatistics struct {
}

func (m *RdbVolumeStatistics) Reset()                    { *m = RdbVolumeStatistics{} }
func (m *RdbVolumeStatistics) String() string            { return proto.CompactTextString(m) }
func (*RdbVolumeStatistics) ProtoMessage()               {}
func (*RdbVolumeStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type RdbVolumeStatus struct {
}

func (m *RdbVolumeStatus) Reset()                    { *m = RdbVolumeStatus{} }
func (m *RdbVolumeStatus) String() string            { return proto.CompactTextString(m) }
func (*RdbVolumeStatus) ProtoMessage()               {}
func (*RdbVolumeStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

// *
// A volume used by the RDB plugin.
//
// Currently there's no clear merit to adding and removing a volume here,
// as the RDB plugin doesn't read ConfigFS and doesn't really do anything.
//
// However, the control plane should still configure a volume - it's likely
// to be useful. At the very least we can create blob resources ahead of time.
type RdbVolume struct {
	Cc *common_v1.DataplaneCommon `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The volume ID.
	VolumeId uint32 `protobuf:"varint,2,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
	// The volume size in bytes.
	VolumeSizeBytes uint64 `protobuf:"varint,3,opt,name=volume_size_bytes,json=volumeSizeBytes" json:"volume_size_bytes,omitempty"`
	// Volume credentials.
	Credentials *RdbVolumeCredentials `protobuf:"bytes,4,opt,name=credentials" json:"credentials,omitempty"`
	// Volume statistics.
	Stats *RdbVolumeStatistics `protobuf:"bytes,5,opt,name=stats" json:"stats,omitempty"`
	// Volume status.
	Status *RdbVolumeStatus `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
}

func (m *RdbVolume) Reset()                    { *m = RdbVolume{} }
func (m *RdbVolume) String() string            { return proto.CompactTextString(m) }
func (*RdbVolume) ProtoMessage()               {}
func (*RdbVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RdbVolume) GetCc() *common_v1.DataplaneCommon {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *RdbVolume) GetVolumeId() uint32 {
	if m != nil {
		return m.VolumeId
	}
	return 0
}

func (m *RdbVolume) GetVolumeSizeBytes() uint64 {
	if m != nil {
		return m.VolumeSizeBytes
	}
	return 0
}

func (m *RdbVolume) GetCredentials() *RdbVolumeCredentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *RdbVolume) GetStats() *RdbVolumeStatistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *RdbVolume) GetStatus() *RdbVolumeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type RdbVolumeList struct {
	Volumes []*RdbVolume `protobuf:"bytes,1,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *RdbVolumeList) Reset()                    { *m = RdbVolumeList{} }
func (m *RdbVolumeList) String() string            { return proto.CompactTextString(m) }
func (*RdbVolumeList) ProtoMessage()               {}
func (*RdbVolumeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RdbVolumeList) GetVolumes() []*RdbVolume {
	if m != nil {
		return m.Volumes
	}
	return nil
}

func init() {
	proto.RegisterType((*RdbStatusRequest)(nil), "rdbplugin.v1.RdbStatusRequest")
	proto.RegisterType((*RdbStatus)(nil), "rdbplugin.v1.RdbStatus")
	proto.RegisterType((*RdbVolumeListQuery)(nil), "rdbplugin.v1.RdbVolumeListQuery")
	proto.RegisterType((*RdbVolumeCredentials)(nil), "rdbplugin.v1.RdbVolumeCredentials")
	proto.RegisterType((*RdbVolumeStatistics)(nil), "rdbplugin.v1.RdbVolumeStatistics")
	proto.RegisterType((*RdbVolumeStatus)(nil), "rdbplugin.v1.RdbVolumeStatus")
	proto.RegisterType((*RdbVolume)(nil), "rdbplugin.v1.RdbVolume")
	proto.RegisterType((*RdbVolumeList)(nil), "rdbplugin.v1.RdbVolumeList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RdbPlugin service

type RdbPluginClient interface {
	// *
	// Get program status.
	Status(ctx context.Context, in *RdbStatusRequest, opts ...grpc.CallOption) (*RdbStatus, error)
	// *
	// Create a volume. Currently has no actual effect other than
	// to add an entry reporting that the volume is 'configured', and so will
	// have results returned by VolumeList.
	//
	// returns RpcResult
	VolumeCreate(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Create a volume. Currently has no actual effect other than
	// to update the entry reporting that the volume is 'configured', and so will
	// have results returned by VolumeList.
	//
	// returns RpcResult
	VolumeUpdate(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Delete the storage allocated for a volume.
	//
	// Deletion can take some time, this call starts the deletion process then
	// returns. Operation status must be obtained elsewhere.
	//
	// returns RpcResult
	VolumeDelete(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error)
	// *
	// Return a list of configured volumes. Optionally filter using the RdbVolumeListQuery message.
	//
	// returns An RdbVolumeList message containing RdbVolume messages,
	//         if any are found matching the filter.
	VolumeList(ctx context.Context, in *RdbVolumeListQuery, opts ...grpc.CallOption) (*RdbVolumeList, error)
}

type rdbPluginClient struct {
	cc *grpc.ClientConn
}

func NewRdbPluginClient(cc *grpc.ClientConn) RdbPluginClient {
	return &rdbPluginClient{cc}
}

func (c *rdbPluginClient) Status(ctx context.Context, in *RdbStatusRequest, opts ...grpc.CallOption) (*RdbStatus, error) {
	out := new(RdbStatus)
	err := grpc.Invoke(ctx, "/rdbplugin.v1.RdbPlugin/Status", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbPluginClient) VolumeCreate(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/rdbplugin.v1.RdbPlugin/VolumeCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbPluginClient) VolumeUpdate(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/rdbplugin.v1.RdbPlugin/VolumeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbPluginClient) VolumeDelete(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*common_v1.RpcResult, error) {
	out := new(common_v1.RpcResult)
	err := grpc.Invoke(ctx, "/rdbplugin.v1.RdbPlugin/VolumeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbPluginClient) VolumeList(ctx context.Context, in *RdbVolumeListQuery, opts ...grpc.CallOption) (*RdbVolumeList, error) {
	out := new(RdbVolumeList)
	err := grpc.Invoke(ctx, "/rdbplugin.v1.RdbPlugin/VolumeList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RdbPlugin service

type RdbPluginServer interface {
	// *
	// Get program status.
	Status(context.Context, *RdbStatusRequest) (*RdbStatus, error)
	// *
	// Create a volume. Currently has no actual effect other than
	// to add an entry reporting that the volume is 'configured', and so will
	// have results returned by VolumeList.
	//
	// returns RpcResult
	VolumeCreate(context.Context, *RdbVolume) (*common_v1.RpcResult, error)
	// *
	// Create a volume. Currently has no actual effect other than
	// to update the entry reporting that the volume is 'configured', and so will
	// have results returned by VolumeList.
	//
	// returns RpcResult
	VolumeUpdate(context.Context, *RdbVolume) (*common_v1.RpcResult, error)
	// *
	// Delete the storage allocated for a volume.
	//
	// Deletion can take some time, this call starts the deletion process then
	// returns. Operation status must be obtained elsewhere.
	//
	// returns RpcResult
	VolumeDelete(context.Context, *RdbVolume) (*common_v1.RpcResult, error)
	// *
	// Return a list of configured volumes. Optionally filter using the RdbVolumeListQuery message.
	//
	// returns An RdbVolumeList message containing RdbVolume messages,
	//         if any are found matching the filter.
	VolumeList(context.Context, *RdbVolumeListQuery) (*RdbVolumeList, error)
}

func RegisterRdbPluginServer(s *grpc.Server, srv RdbPluginServer) {
	s.RegisterService(&_RdbPlugin_serviceDesc, srv)
}

func _RdbPlugin_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdbplugin.v1.RdbPlugin/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginServer).Status(ctx, req.(*RdbStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RdbPlugin_VolumeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginServer).VolumeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdbplugin.v1.RdbPlugin/VolumeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginServer).VolumeCreate(ctx, req.(*RdbVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _RdbPlugin_VolumeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginServer).VolumeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdbplugin.v1.RdbPlugin/VolumeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginServer).VolumeUpdate(ctx, req.(*RdbVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _RdbPlugin_VolumeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginServer).VolumeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdbplugin.v1.RdbPlugin/VolumeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginServer).VolumeDelete(ctx, req.(*RdbVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _RdbPlugin_VolumeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbVolumeListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginServer).VolumeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rdbplugin.v1.RdbPlugin/VolumeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginServer).VolumeList(ctx, req.(*RdbVolumeListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _RdbPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rdbplugin.v1.RdbPlugin",
	HandlerType: (*RdbPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _RdbPlugin_Status_Handler,
		},
		{
			MethodName: "VolumeCreate",
			Handler:    _RdbPlugin_VolumeCreate_Handler,
		},
		{
			MethodName: "VolumeUpdate",
			Handler:    _RdbPlugin_VolumeUpdate_Handler,
		},
		{
			MethodName: "VolumeDelete",
			Handler:    _RdbPlugin_VolumeDelete_Handler,
		},
		{
			MethodName: "VolumeList",
			Handler:    _RdbPlugin_VolumeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rdbplugin.proto",
}

func init() { proto.RegisterFile("rdbplugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0xeb, 0xb4, 0x0d, 0x64, 0xec, 0x28, 0x74, 0x29, 0xd4, 0x72, 0x55, 0xe4, 0xee, 0xc9,
	0xea, 0xc1, 0x52, 0x52, 0x21, 0x8e, 0x48, 0x4d, 0x2e, 0x95, 0x40, 0x82, 0xad, 0xe0, 0x1a, 0xd9,
	0xeb, 0x2d, 0x5a, 0xc9, 0xf1, 0x1a, 0xcf, 0xae, 0xa5, 0xf6, 0xc1, 0x78, 0x20, 0x9e, 0x04, 0x75,
	0xed, 0x3a, 0x8e, 0xc0, 0x3d, 0xd0, 0xeb, 0x37, 0xf3, 0x8f, 0xff, 0xf9, 0xc7, 0x0b, 0xb3, 0x2a,
	0x4b, 0xcb, 0xdc, 0xfc, 0x90, 0x45, 0x5c, 0x56, 0x4a, 0x2b, 0xe2, 0x6d, 0x41, 0x3d, 0x0f, 0x3c,
	0xae, 0x36, 0x1b, 0xd5, 0xd6, 0x28, 0x81, 0x57, 0x2c, 0x4b, 0x6f, 0x74, 0xa2, 0x0d, 0x32, 0xf1,
	0xd3, 0x08, 0xd4, 0x34, 0x86, 0x49, 0xc7, 0xc8, 0x39, 0x78, 0xb5, 0xa8, 0x50, 0xaa, 0x62, 0x2d,
	0x8b, 0x5b, 0xe5, 0x3b, 0xa1, 0x13, 0x4d, 0x98, 0xdb, 0xb2, 0xeb, 0xe2, 0x56, 0xd1, 0x4b, 0x20,
	0x2c, 0x4b, 0xbf, 0xab, 0xdc, 0x6c, 0xc4, 0x27, 0x89, 0xfa, 0xab, 0x11, 0xd5, 0x1d, 0x39, 0x03,
	0xa8, 0x2d, 0x5a, 0xcb, 0x0c, 0x7d, 0x27, 0xdc, 0x8f, 0xa6, 0x6c, 0xd2, 0x90, 0xeb, 0x0c, 0xe9,
	0x5b, 0x38, 0xee, 0x44, 0xcb, 0x4a, 0x64, 0xa2, 0xd0, 0x32, 0xc9, 0x91, 0xbe, 0x81, 0xd7, 0x1d,
	0x7f, 0xb0, 0x20, 0x51, 0x4b, 0x8e, 0xf4, 0x08, 0x66, 0x3b, 0xd8, 0x20, 0xfd, 0x35, 0xb2, 0x3e,
	0x1b, 0x46, 0x2e, 0x60, 0xc4, 0xb9, 0x75, 0xe7, 0x2e, 0x82, 0xb8, 0xdd, 0xb1, 0x9e, 0xc7, 0xab,
	0x44, 0x27, 0x65, 0x9e, 0x14, 0x62, 0x69, 0x11, 0x1b, 0x71, 0x4e, 0x4e, 0x61, 0xd2, 0x59, 0xf3,
	0x47, 0xa1, 0x13, 0x4d, 0xd9, 0xcb, 0x47, 0x67, 0xe4, 0x02, 0x8e, 0xda, 0x22, 0xca, 0x7b, 0xb1,
	0x4e, 0xef, 0xb4, 0x40, 0x7f, 0x3f, 0x74, 0xa2, 0x03, 0x36, 0x6b, 0x0a, 0x37, 0xf2, 0x5e, 0x5c,
	0x3d, 0x60, 0xb2, 0x02, 0x97, 0x6f, 0xbd, 0xfb, 0x07, 0xf6, 0xeb, 0x34, 0xee, 0xe7, 0x1d, 0xff,
	0x6b, 0x4b, 0xd6, 0x97, 0x91, 0x0f, 0x70, 0x88, 0x3a, 0xd1, 0xe8, 0x1f, 0x5a, 0xfd, 0xf9, 0x80,
	0x7e, 0x9b, 0x06, 0x6b, 0xfa, 0xc9, 0x7b, 0x18, 0xa3, 0xcd, 0xc2, 0x1f, 0x5b, 0xe5, 0xd9, 0x13,
	0x4a, 0x83, 0xac, 0x6d, 0xa6, 0x57, 0x30, 0xdd, 0xb9, 0x17, 0x99, 0xc3, 0x8b, 0x66, 0xb3, 0xe6,
	0x4e, 0xee, 0xe2, 0x64, 0x60, 0x10, 0x7b, 0xec, 0x5b, 0xfc, 0x6e, 0xc2, 0xff, 0x62, 0x7b, 0xc8,
	0x12, 0xc6, 0xed, 0xef, 0xf2, 0xee, 0x2f, 0xe5, 0xce, 0xbf, 0x15, 0x9c, 0x0c, 0xd4, 0xe9, 0x1e,
	0xf9, 0x08, 0x5e, 0x17, 0x54, 0xa2, 0x05, 0x19, 0x32, 0x11, 0x1c, 0xf7, 0xce, 0xcb, 0x4a, 0xce,
	0x04, 0x9a, 0x5c, 0xf7, 0x07, 0x7c, 0x2b, 0xb3, 0xe7, 0x0d, 0x58, 0x89, 0x5c, 0xfc, 0xcf, 0x80,
	0xcf, 0x00, 0xbd, 0x58, 0xc3, 0x01, 0x79, 0xf7, 0x46, 0x82, 0xd3, 0x27, 0x3a, 0xe8, 0x5e, 0x3a,
	0xb6, 0x6f, 0xf4, 0xf2, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x95, 0xa5, 0xd6, 0xba, 0xd2, 0x03,
	0x00, 0x00,
}
