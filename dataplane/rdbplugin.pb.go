// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rdbplugin.proto

/*
Package storageos_rpc is a generated protocol buffer package.

It is generated from these files:
	rdbplugin.proto

It has these top-level messages:
	RdbVolumeListQuery
	RdbVolumeCredentials
	RdbVolumeStats
	RdbVolume
*/
package storageos_rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import storageos_rpc1 "."

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

type RdbVolumeListQuery struct {
	// Optional list of volume IDs to query.
	VolumeIds []*RdbVolume `protobuf:"bytes,1,rep,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
}

func (m *RdbVolumeListQuery) Reset()                    { *m = RdbVolumeListQuery{} }
func (m *RdbVolumeListQuery) String() string            { return proto.CompactTextString(m) }
func (*RdbVolumeListQuery) ProtoMessage()               {}
func (*RdbVolumeListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *RdbVolumeListQuery) GetVolumeIds() []*RdbVolume {
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
func (*RdbVolumeCredentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type RdbVolumeStats struct {
}

func (m *RdbVolumeStats) Reset()                    { *m = RdbVolumeStats{} }
func (m *RdbVolumeStats) String() string            { return proto.CompactTextString(m) }
func (*RdbVolumeStats) ProtoMessage()               {}
func (*RdbVolumeStats) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// *
// A volume used by the RDB plugin.
//
// Currently there's no clear merit to adding and removing a volume here,
// as the RDB plugin doesn't read ConfigFS and doesn't really do anything.
//
// However, the control plane should still configure a volume - it's likely
// to be useful. At the very least we can create blob resources ahead of time.
type RdbVolume struct {
	Cc *storageos_rpc1.DataplaneCommonConfig `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The volume ID.
	VolumeId uint32 `protobuf:"varint,2,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
	// The volume size in bytes.
	VolumeSizeBytes uint64 `protobuf:"varint,3,opt,name=volume_size_bytes,json=volumeSizeBytes" json:"volume_size_bytes,omitempty"`
	// Volume credentials.
	Credentials *RdbVolumeCredentials `protobuf:"bytes,4,opt,name=credentials" json:"credentials,omitempty"`
	// Volume statistics.
	Stats *RdbVolumeStats `protobuf:"bytes,5,opt,name=stats" json:"stats,omitempty"`
}

func (m *RdbVolume) Reset()                    { *m = RdbVolume{} }
func (m *RdbVolume) String() string            { return proto.CompactTextString(m) }
func (*RdbVolume) ProtoMessage()               {}
func (*RdbVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *RdbVolume) GetCc() *storageos_rpc1.DataplaneCommonConfig {
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

func (m *RdbVolume) GetStats() *RdbVolumeStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*RdbVolumeListQuery)(nil), "storageos_rpc.RdbVolumeListQuery")
	proto.RegisterType((*RdbVolumeCredentials)(nil), "storageos_rpc.RdbVolumeCredentials")
	proto.RegisterType((*RdbVolumeStats)(nil), "storageos_rpc.RdbVolumeStats")
	proto.RegisterType((*RdbVolume)(nil), "storageos_rpc.RdbVolume")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for RdbPluginConfig service

type RdbPluginConfigClient interface {
	// *
	// Configure a volume. Currently has no actual effect other authentication
	// to add an entry reporting that the volume is 'configured', and so will
	// have results returned by VolumeList.
	//
	// returns RpcResult
	VolumeConfigure(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*storageos_rpc1.RpcResult, error)
	// *
	// Unconfigure a volume. As for VolumeConfigure, this has no direct on-disk
	// effect at the moment, and only removes the volume from the list of
	// 'configured' items returned by VolumeList.
	//
	// returns RpcResult
	VolumeUnconfigure(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*storageos_rpc1.RpcResult, error)
	VolumeDelete(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*storageos_rpc1.RpcResult, error)
	VolumeList(ctx context.Context, in *RdbVolumeListQuery, opts ...grpc.CallOption) (RdbPluginConfig_VolumeListClient, error)
}

type rdbPluginConfigClient struct {
	cc *grpc.ClientConn
}

func NewRdbPluginConfigClient(cc *grpc.ClientConn) RdbPluginConfigClient {
	return &rdbPluginConfigClient{cc}
}

func (c *rdbPluginConfigClient) VolumeConfigure(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*storageos_rpc1.RpcResult, error) {
	out := new(storageos_rpc1.RpcResult)
	err := grpc.Invoke(ctx, "/storageos_rpc.RdbPluginConfig/VolumeConfigure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbPluginConfigClient) VolumeUnconfigure(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*storageos_rpc1.RpcResult, error) {
	out := new(storageos_rpc1.RpcResult)
	err := grpc.Invoke(ctx, "/storageos_rpc.RdbPluginConfig/VolumeUnconfigure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbPluginConfigClient) VolumeDelete(ctx context.Context, in *RdbVolume, opts ...grpc.CallOption) (*storageos_rpc1.RpcResult, error) {
	out := new(storageos_rpc1.RpcResult)
	err := grpc.Invoke(ctx, "/storageos_rpc.RdbPluginConfig/VolumeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rdbPluginConfigClient) VolumeList(ctx context.Context, in *RdbVolumeListQuery, opts ...grpc.CallOption) (RdbPluginConfig_VolumeListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_RdbPluginConfig_serviceDesc.Streams[0], c.cc, "/storageos_rpc.RdbPluginConfig/VolumeList", opts...)
	if err != nil {
		return nil, err
	}
	x := &rdbPluginConfigVolumeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RdbPluginConfig_VolumeListClient interface {
	Recv() (*RdbVolume, error)
	grpc.ClientStream
}

type rdbPluginConfigVolumeListClient struct {
	grpc.ClientStream
}

func (x *rdbPluginConfigVolumeListClient) Recv() (*RdbVolume, error) {
	m := new(RdbVolume)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RdbPluginConfig service

type RdbPluginConfigServer interface {
	// *
	// Configure a volume. Currently has no actual effect other authentication
	// to add an entry reporting that the volume is 'configured', and so will
	// have results returned by VolumeList.
	//
	// returns RpcResult
	VolumeConfigure(context.Context, *RdbVolume) (*storageos_rpc1.RpcResult, error)
	// *
	// Unconfigure a volume. As for VolumeConfigure, this has no direct on-disk
	// effect at the moment, and only removes the volume from the list of
	// 'configured' items returned by VolumeList.
	//
	// returns RpcResult
	VolumeUnconfigure(context.Context, *RdbVolume) (*storageos_rpc1.RpcResult, error)
	VolumeDelete(context.Context, *RdbVolume) (*storageos_rpc1.RpcResult, error)
	VolumeList(*RdbVolumeListQuery, RdbPluginConfig_VolumeListServer) error
}

func RegisterRdbPluginConfigServer(s *grpc.Server, srv RdbPluginConfigServer) {
	s.RegisterService(&_RdbPluginConfig_serviceDesc, srv)
}

func _RdbPluginConfig_VolumeConfigure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginConfigServer).VolumeConfigure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageos_rpc.RdbPluginConfig/VolumeConfigure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginConfigServer).VolumeConfigure(ctx, req.(*RdbVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _RdbPluginConfig_VolumeUnconfigure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginConfigServer).VolumeUnconfigure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageos_rpc.RdbPluginConfig/VolumeUnconfigure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginConfigServer).VolumeUnconfigure(ctx, req.(*RdbVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _RdbPluginConfig_VolumeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RdbVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RdbPluginConfigServer).VolumeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/storageos_rpc.RdbPluginConfig/VolumeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RdbPluginConfigServer).VolumeDelete(ctx, req.(*RdbVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _RdbPluginConfig_VolumeList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RdbVolumeListQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RdbPluginConfigServer).VolumeList(m, &rdbPluginConfigVolumeListServer{stream})
}

type RdbPluginConfig_VolumeListServer interface {
	Send(*RdbVolume) error
	grpc.ServerStream
}

type rdbPluginConfigVolumeListServer struct {
	grpc.ServerStream
}

func (x *rdbPluginConfigVolumeListServer) Send(m *RdbVolume) error {
	return x.ServerStream.SendMsg(m)
}

var _RdbPluginConfig_serviceDesc = grpc.ServiceDesc{
	ServiceName: "storageos_rpc.RdbPluginConfig",
	HandlerType: (*RdbPluginConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VolumeConfigure",
			Handler:    _RdbPluginConfig_VolumeConfigure_Handler,
		},
		{
			MethodName: "VolumeUnconfigure",
			Handler:    _RdbPluginConfig_VolumeUnconfigure_Handler,
		},
		{
			MethodName: "VolumeDelete",
			Handler:    _RdbPluginConfig_VolumeDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "VolumeList",
			Handler:       _RdbPluginConfig_VolumeList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rdbplugin.proto",
}

func init() { proto.RegisterFile("rdbplugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 355 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x5d, 0xbb, 0x4d, 0xdc, 0xdd, 0x66, 0x5d, 0x10, 0x09, 0x13, 0xa1, 0x56, 0x1f, 0x8a, 0x0f,
	0x43, 0x36, 0xc1, 0x77, 0x37, 0x91, 0x81, 0x03, 0xcd, 0xd0, 0xd7, 0xd1, 0xa6, 0x71, 0x04, 0xba,
	0xa6, 0x24, 0xa9, 0xb0, 0x7d, 0x80, 0x1f, 0xe5, 0xd7, 0xc9, 0xd2, 0xd9, 0x4d, 0xa5, 0x2f, 0xfa,
	0xd8, 0x73, 0x4e, 0xce, 0x3d, 0xe7, 0xf6, 0x82, 0x23, 0xa3, 0x30, 0x8d, 0xb3, 0x39, 0x4f, 0x7a,
	0xa9, 0x14, 0x5a, 0xa0, 0xb6, 0xd2, 0x42, 0x06, 0x73, 0x26, 0xd4, 0x4c, 0xa6, 0xb4, 0xdb, 0xa2,
	0x62, 0xb1, 0x10, 0x1b, 0xd2, 0x9b, 0x00, 0x22, 0x51, 0xf8, 0x22, 0xe2, 0x6c, 0xc1, 0x1e, 0xb8,
	0xd2, 0x4f, 0x19, 0x93, 0x4b, 0x74, 0x03, 0xf0, 0x66, 0xa0, 0x19, 0x8f, 0x14, 0xb6, 0xdc, 0xaa,
	0xdf, 0xec, 0xe3, 0xde, 0x37, 0x9f, 0x5e, 0xf1, 0x8c, 0x34, 0x72, 0xed, 0x38, 0x52, 0xde, 0x31,
	0x1c, 0x15, 0xf8, 0x50, 0xb2, 0x88, 0x25, 0x9a, 0x07, 0xb1, 0xf2, 0x0e, 0xe1, 0xa0, 0xc0, 0xa7,
	0x3a, 0xd0, 0xca, 0x7b, 0xb7, 0xa1, 0x51, 0x40, 0xe8, 0x1a, 0x6c, 0x4a, 0xb1, 0xe5, 0x5a, 0x7e,
	0xb3, 0x7f, 0xf1, 0x63, 0xd0, 0x28, 0xd0, 0x41, 0x1a, 0x07, 0x09, 0x1b, 0x9a, 0xe0, 0x43, 0x91,
	0xbc, 0xf2, 0x39, 0xb1, 0x29, 0x45, 0x27, 0xd0, 0x28, 0x62, 0x62, 0xdb, 0xb5, 0xfc, 0x36, 0xd9,
	0xff, 0xca, 0x82, 0x2e, 0xa1, 0xb3, 0x21, 0x15, 0x5f, 0xb1, 0x59, 0xb8, 0xd4, 0x4c, 0xe1, 0xaa,
	0x6b, 0xf9, 0x35, 0xe2, 0xe4, 0xc4, 0x94, 0xaf, 0xd8, 0xed, 0x1a, 0x46, 0x77, 0xd0, 0xa4, 0xdb,
	0xb4, 0xb8, 0x66, 0x72, 0x9c, 0x97, 0x15, 0xde, 0x29, 0x46, 0x76, 0xdf, 0xa1, 0x01, 0xd4, 0xd5,
	0xba, 0x1c, 0xae, 0x1b, 0x83, 0xd3, 0x32, 0x03, 0xb3, 0x01, 0x92, 0x6b, 0xfb, 0x1f, 0x36, 0x38,
	0x24, 0x0a, 0x1f, 0xcd, 0x2f, 0xcb, 0xcb, 0xa1, 0x7b, 0x70, 0x36, 0xa3, 0xcc, 0x77, 0x26, 0x19,
	0x2a, 0x5d, 0x7f, 0xf7, 0x17, 0x93, 0x52, 0xc2, 0x54, 0x16, 0x6b, 0xaf, 0x82, 0xc6, 0xd0, 0xc9,
	0x55, 0xcf, 0x09, 0xfd, 0xa7, 0xd5, 0x08, 0x5a, 0xb9, 0x6a, 0xc4, 0x62, 0xa6, 0xff, 0xea, 0x32,
	0x01, 0xd8, 0x1e, 0x1b, 0x3a, 0x2b, 0xf3, 0x28, 0x4e, 0xb1, 0x5b, 0x3a, 0xc6, 0xab, 0x5c, 0x59,
	0xe1, 0x9e, 0xb9, 0xe2, 0xc1, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x9a, 0xb2, 0x26, 0xf5,
	0x02, 0x00, 0x00,
}