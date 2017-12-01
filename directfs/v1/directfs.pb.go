// Code generated by protoc-gen-go. DO NOT EDIT.
// source: directfs.proto

/*
Package directfs_v1 is a generated protocol buffer package.

It is generated from these files:
	directfs.proto

It has these top-level messages:
	DfsHostCredentials
	DfsHost
	DfsHostList
	DfsHostListQuery
	DfsVolumeCredentials
	DfsVolumeStatistics
	DfsVolumeStatus
	DfsVolume
	DfsVolumeList
	DfsVolumeListQuery
*/
package directfs_v1

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

type DfsVolumeStatus_DfsConnectionState int32

const (
	DfsVolumeStatus_NONE          DfsVolumeStatus_DfsConnectionState = 0
	DfsVolumeStatus_CONNECTING    DfsVolumeStatus_DfsConnectionState = 1
	DfsVolumeStatus_CONNECTED     DfsVolumeStatus_DfsConnectionState = 2
	DfsVolumeStatus_DISCONNECTING DfsVolumeStatus_DfsConnectionState = 3
	DfsVolumeStatus_DISCONNECTED  DfsVolumeStatus_DfsConnectionState = 4
)

var DfsVolumeStatus_DfsConnectionState_name = map[int32]string{
	0: "NONE",
	1: "CONNECTING",
	2: "CONNECTED",
	3: "DISCONNECTING",
	4: "DISCONNECTED",
}
var DfsVolumeStatus_DfsConnectionState_value = map[string]int32{
	"NONE":          0,
	"CONNECTING":    1,
	"CONNECTED":     2,
	"DISCONNECTING": 3,
	"DISCONNECTED":  4,
}

func (x DfsVolumeStatus_DfsConnectionState) String() string {
	return proto.EnumName(DfsVolumeStatus_DfsConnectionState_name, int32(x))
}
func (DfsVolumeStatus_DfsConnectionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 0}
}

type DfsVolumeStatus_DfsAddressFamily int32

const (
	DfsVolumeStatus_IPV4 DfsVolumeStatus_DfsAddressFamily = 0
	DfsVolumeStatus_IPV6 DfsVolumeStatus_DfsAddressFamily = 1
)

var DfsVolumeStatus_DfsAddressFamily_name = map[int32]string{
	0: "IPV4",
	1: "IPV6",
}
var DfsVolumeStatus_DfsAddressFamily_value = map[string]int32{
	"IPV4": 0,
	"IPV6": 1,
}

func (x DfsVolumeStatus_DfsAddressFamily) String() string {
	return proto.EnumName(DfsVolumeStatus_DfsAddressFamily_name, int32(x))
}
func (DfsVolumeStatus_DfsAddressFamily) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor0, []int{6, 1}
}

type DfsHostCredentials struct {
}

func (m *DfsHostCredentials) Reset()                    { *m = DfsHostCredentials{} }
func (m *DfsHostCredentials) String() string            { return proto.CompactTextString(m) }
func (*DfsHostCredentials) ProtoMessage()               {}
func (*DfsHostCredentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// *
// A host used by DirectFS.
type DfsHost struct {
	Cc *common.DataplaneCommon `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The unique host identifier.
	HostId uint32 `protobuf:"varint,2,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	// The remote hostname.
	Hostname string `protobuf:"bytes,3,opt,name=hostname" json:"hostname,omitempty"`
	// The remote port.
	Port uint32 `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
	// Host credentials.
	Credentials *DfsHostCredentials `protobuf:"bytes,5,opt,name=credentials" json:"credentials,omitempty"`
}

func (m *DfsHost) Reset()                    { *m = DfsHost{} }
func (m *DfsHost) String() string            { return proto.CompactTextString(m) }
func (*DfsHost) ProtoMessage()               {}
func (*DfsHost) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DfsHost) GetCc() *common.DataplaneCommon {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *DfsHost) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *DfsHost) GetHostname() string {
	if m != nil {
		return m.Hostname
	}
	return ""
}

func (m *DfsHost) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *DfsHost) GetCredentials() *DfsHostCredentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

type DfsHostList struct {
	Hosts []*DfsHost `protobuf:"bytes,1,rep,name=hosts" json:"hosts,omitempty"`
}

func (m *DfsHostList) Reset()                    { *m = DfsHostList{} }
func (m *DfsHostList) String() string            { return proto.CompactTextString(m) }
func (*DfsHostList) ProtoMessage()               {}
func (*DfsHostList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DfsHostList) GetHosts() []*DfsHost {
	if m != nil {
		return m.Hosts
	}
	return nil
}

type DfsHostListQuery struct {
	// An optional list of hosts to query.
	HostIds []*DfsHost `protobuf:"bytes,1,rep,name=host_ids,json=hostIds" json:"host_ids,omitempty"`
}

func (m *DfsHostListQuery) Reset()                    { *m = DfsHostListQuery{} }
func (m *DfsHostListQuery) String() string            { return proto.CompactTextString(m) }
func (*DfsHostListQuery) ProtoMessage()               {}
func (*DfsHostListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *DfsHostListQuery) GetHostIds() []*DfsHost {
	if m != nil {
		return m.HostIds
	}
	return nil
}

type DfsVolumeCredentials struct {
}

func (m *DfsVolumeCredentials) Reset()                    { *m = DfsVolumeCredentials{} }
func (m *DfsVolumeCredentials) String() string            { return proto.CompactTextString(m) }
func (*DfsVolumeCredentials) ProtoMessage()               {}
func (*DfsVolumeCredentials) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type DfsVolumeStatistics struct {
}

func (m *DfsVolumeStatistics) Reset()                    { *m = DfsVolumeStatistics{} }
func (m *DfsVolumeStatistics) String() string            { return proto.CompactTextString(m) }
func (*DfsVolumeStatistics) ProtoMessage()               {}
func (*DfsVolumeStatistics) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type DfsVolumeStatus struct {
	ConnState DfsVolumeStatus_DfsConnectionState `protobuf:"varint,1,opt,name=conn_state,json=connState,enum=directfs.v1.DfsVolumeStatus_DfsConnectionState" json:"conn_state,omitempty"`
	// The remote peer, in string form for simplicity.
	PeerName string `protobuf:"bytes,2,opt,name=peer_name,json=peerName" json:"peer_name,omitempty"`
	// The address family we're using to connect to the peer.
	PeerAf DfsVolumeStatus_DfsAddressFamily `protobuf:"varint,3,opt,name=peer_af,json=peerAf,enum=directfs.v1.DfsVolumeStatus_DfsAddressFamily" json:"peer_af,omitempty"`
}

func (m *DfsVolumeStatus) Reset()                    { *m = DfsVolumeStatus{} }
func (m *DfsVolumeStatus) String() string            { return proto.CompactTextString(m) }
func (*DfsVolumeStatus) ProtoMessage()               {}
func (*DfsVolumeStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DfsVolumeStatus) GetConnState() DfsVolumeStatus_DfsConnectionState {
	if m != nil {
		return m.ConnState
	}
	return DfsVolumeStatus_NONE
}

func (m *DfsVolumeStatus) GetPeerName() string {
	if m != nil {
		return m.PeerName
	}
	return ""
}

func (m *DfsVolumeStatus) GetPeerAf() DfsVolumeStatus_DfsAddressFamily {
	if m != nil {
		return m.PeerAf
	}
	return DfsVolumeStatus_IPV4
}

// *
// A volume used by DirectFS.
//
// The remote host (for the client) isn't directly included here. However, the
// client clearly can't connect without a properly configured host_id item.
//
// The split between host and volume objects was decided based on now-irrelevant
// implementation details of ConfigFS v1. However, there's still some logic to
// it if the authentication (especially TLS) is per-host rather than per-volume.
type DfsVolume struct {
	Cc *common.DataplaneCommon `protobuf:"bytes,1,opt,name=cc" json:"cc,omitempty"`
	// The volume ID.
	VolumeId uint32 `protobuf:"varint,2,opt,name=volume_id,json=volumeId" json:"volume_id,omitempty"`
	// The host for this volume.
	HostId uint32 `protobuf:"varint,3,opt,name=host_id,json=hostId" json:"host_id,omitempty"`
	// Volume credentials.
	Credentials *DfsVolumeCredentials `protobuf:"bytes,4,opt,name=credentials" json:"credentials,omitempty"`
	// Volume statistics.
	Stats *DfsVolumeStatistics `protobuf:"bytes,5,opt,name=stats" json:"stats,omitempty"`
	// Server status for the volume.
	Status *DfsVolumeStatus `protobuf:"bytes,6,opt,name=status" json:"status,omitempty"`
}

func (m *DfsVolume) Reset()                    { *m = DfsVolume{} }
func (m *DfsVolume) String() string            { return proto.CompactTextString(m) }
func (*DfsVolume) ProtoMessage()               {}
func (*DfsVolume) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DfsVolume) GetCc() *common.DataplaneCommon {
	if m != nil {
		return m.Cc
	}
	return nil
}

func (m *DfsVolume) GetVolumeId() uint32 {
	if m != nil {
		return m.VolumeId
	}
	return 0
}

func (m *DfsVolume) GetHostId() uint32 {
	if m != nil {
		return m.HostId
	}
	return 0
}

func (m *DfsVolume) GetCredentials() *DfsVolumeCredentials {
	if m != nil {
		return m.Credentials
	}
	return nil
}

func (m *DfsVolume) GetStats() *DfsVolumeStatistics {
	if m != nil {
		return m.Stats
	}
	return nil
}

func (m *DfsVolume) GetStatus() *DfsVolumeStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

type DfsVolumeList struct {
	Volumes []*DfsVolumeList `protobuf:"bytes,1,rep,name=volumes" json:"volumes,omitempty"`
}

func (m *DfsVolumeList) Reset()                    { *m = DfsVolumeList{} }
func (m *DfsVolumeList) String() string            { return proto.CompactTextString(m) }
func (*DfsVolumeList) ProtoMessage()               {}
func (*DfsVolumeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DfsVolumeList) GetVolumes() []*DfsVolumeList {
	if m != nil {
		return m.Volumes
	}
	return nil
}

type DfsVolumeListQuery struct {
	// An optional list of volumes to query.
	VolumeIds []*DfsVolume `protobuf:"bytes,1,rep,name=volume_ids,json=volumeIds" json:"volume_ids,omitempty"`
}

func (m *DfsVolumeListQuery) Reset()                    { *m = DfsVolumeListQuery{} }
func (m *DfsVolumeListQuery) String() string            { return proto.CompactTextString(m) }
func (*DfsVolumeListQuery) ProtoMessage()               {}
func (*DfsVolumeListQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DfsVolumeListQuery) GetVolumeIds() []*DfsVolume {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

func init() {
	proto.RegisterType((*DfsHostCredentials)(nil), "directfs.v1.DfsHostCredentials")
	proto.RegisterType((*DfsHost)(nil), "directfs.v1.DfsHost")
	proto.RegisterType((*DfsHostList)(nil), "directfs.v1.DfsHostList")
	proto.RegisterType((*DfsHostListQuery)(nil), "directfs.v1.DfsHostListQuery")
	proto.RegisterType((*DfsVolumeCredentials)(nil), "directfs.v1.DfsVolumeCredentials")
	proto.RegisterType((*DfsVolumeStatistics)(nil), "directfs.v1.DfsVolumeStatistics")
	proto.RegisterType((*DfsVolumeStatus)(nil), "directfs.v1.DfsVolumeStatus")
	proto.RegisterType((*DfsVolume)(nil), "directfs.v1.DfsVolume")
	proto.RegisterType((*DfsVolumeList)(nil), "directfs.v1.DfsVolumeList")
	proto.RegisterType((*DfsVolumeListQuery)(nil), "directfs.v1.DfsVolumeListQuery")
	proto.RegisterEnum("directfs.v1.DfsVolumeStatus_DfsConnectionState", DfsVolumeStatus_DfsConnectionState_name, DfsVolumeStatus_DfsConnectionState_value)
	proto.RegisterEnum("directfs.v1.DfsVolumeStatus_DfsAddressFamily", DfsVolumeStatus_DfsAddressFamily_name, DfsVolumeStatus_DfsAddressFamily_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FsClient service

type FsClientClient interface {
	// *
	// Add a remote host entry to be used by volumes.
	//
	// returns RpcResult
	ServerCreate(ctx context.Context, in *DfsHost, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Update a remote host entry to be used by volumes.
	//
	// returns RpcResult
	ServerUpdate(ctx context.Context, in *DfsHost, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Remove a remote host entry.
	//
	// This will likely result in any volumes using this host to
	// become unconfigured. That has serious consequences.
	//
	// returns RpcResult
	ServerDelete(ctx context.Context, in *DfsHost, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// List configured host entries, optionally filtered via a DfsHostListQuery
	// message.
	//
	// returns A DfsHostList message containing DfsHost messages,
	//         if any are available matching the filter.
	ServerList(ctx context.Context, in *DfsHostListQuery, opts ...grpc.CallOption) (*DfsHostList, error)
	// *
	// Create a volume on a remote host. The DfsHost matching the host id
	// in the DfsVolume message must be configured for the volume to actually
	// be configured on the DirectFS client.
	//
	// returns RpcResult
	VolumeCreate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Update a volume on a remote host. The DfsHost matching the host id
	// in the DfsVolume message must be configured for the volume to actually
	// be configured on the DirectFS client.
	//
	// returns RpcResult
	VolumeUpdate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Delete a volume previously configured by VolumeCreate.
	//
	// returns RpcResult
	VolumeDelete(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// List configured volume entries, optionally filtered via a DfsVolumeListQuery
	// message.
	//
	// returns A DfsVolumeList message containing DfsVolume messages,
	//         if any are available matching the filter.
	VolumeList(ctx context.Context, in *DfsVolumeListQuery, opts ...grpc.CallOption) (*DfsVolumeList, error)
}

type fsClientClient struct {
	cc *grpc.ClientConn
}

func NewFsClientClient(cc *grpc.ClientConn) FsClientClient {
	return &fsClientClient{cc}
}

func (c *fsClientClient) ServerCreate(ctx context.Context, in *DfsHost, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/ServerCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClientClient) ServerUpdate(ctx context.Context, in *DfsHost, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/ServerUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClientClient) ServerDelete(ctx context.Context, in *DfsHost, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/ServerDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClientClient) ServerList(ctx context.Context, in *DfsHostListQuery, opts ...grpc.CallOption) (*DfsHostList, error) {
	out := new(DfsHostList)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/ServerList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClientClient) VolumeCreate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/VolumeCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClientClient) VolumeUpdate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/VolumeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClientClient) VolumeDelete(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/VolumeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsClientClient) VolumeList(ctx context.Context, in *DfsVolumeListQuery, opts ...grpc.CallOption) (*DfsVolumeList, error) {
	out := new(DfsVolumeList)
	err := grpc.Invoke(ctx, "/directfs.v1.FsClient/VolumeList", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FsClient service

type FsClientServer interface {
	// *
	// Add a remote host entry to be used by volumes.
	//
	// returns RpcResult
	ServerCreate(context.Context, *DfsHost) (*common.RpcResult, error)
	// *
	// Update a remote host entry to be used by volumes.
	//
	// returns RpcResult
	ServerUpdate(context.Context, *DfsHost) (*common.RpcResult, error)
	// *
	// Remove a remote host entry.
	//
	// This will likely result in any volumes using this host to
	// become unconfigured. That has serious consequences.
	//
	// returns RpcResult
	ServerDelete(context.Context, *DfsHost) (*common.RpcResult, error)
	// *
	// List configured host entries, optionally filtered via a DfsHostListQuery
	// message.
	//
	// returns A DfsHostList message containing DfsHost messages,
	//         if any are available matching the filter.
	ServerList(context.Context, *DfsHostListQuery) (*DfsHostList, error)
	// *
	// Create a volume on a remote host. The DfsHost matching the host id
	// in the DfsVolume message must be configured for the volume to actually
	// be configured on the DirectFS client.
	//
	// returns RpcResult
	VolumeCreate(context.Context, *DfsVolume) (*common.RpcResult, error)
	// *
	// Update a volume on a remote host. The DfsHost matching the host id
	// in the DfsVolume message must be configured for the volume to actually
	// be configured on the DirectFS client.
	//
	// returns RpcResult
	VolumeUpdate(context.Context, *DfsVolume) (*common.RpcResult, error)
	// *
	// Delete a volume previously configured by VolumeCreate.
	//
	// returns RpcResult
	VolumeDelete(context.Context, *DfsVolume) (*common.RpcResult, error)
	// *
	// List configured volume entries, optionally filtered via a DfsVolumeListQuery
	// message.
	//
	// returns A DfsVolumeList message containing DfsVolume messages,
	//         if any are available matching the filter.
	VolumeList(context.Context, *DfsVolumeListQuery) (*DfsVolumeList, error)
}

func RegisterFsClientServer(s *grpc.Server, srv FsClientServer) {
	s.RegisterService(&_FsClient_serviceDesc, srv)
}

func _FsClient_ServerCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsHost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).ServerCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/ServerCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).ServerCreate(ctx, req.(*DfsHost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsClient_ServerUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsHost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).ServerUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/ServerUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).ServerUpdate(ctx, req.(*DfsHost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsClient_ServerDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsHost)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).ServerDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/ServerDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).ServerDelete(ctx, req.(*DfsHost))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsClient_ServerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsHostListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).ServerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/ServerList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).ServerList(ctx, req.(*DfsHostListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsClient_VolumeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).VolumeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/VolumeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).VolumeCreate(ctx, req.(*DfsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsClient_VolumeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).VolumeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/VolumeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).VolumeUpdate(ctx, req.(*DfsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsClient_VolumeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).VolumeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/VolumeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).VolumeDelete(ctx, req.(*DfsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsClient_VolumeList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsVolumeListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsClientServer).VolumeList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsClient/VolumeList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsClientServer).VolumeList(ctx, req.(*DfsVolumeListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _FsClient_serviceDesc = grpc.ServiceDesc{
	ServiceName: "directfs.v1.FsClient",
	HandlerType: (*FsClientServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerCreate",
			Handler:    _FsClient_ServerCreate_Handler,
		},
		{
			MethodName: "ServerUpdate",
			Handler:    _FsClient_ServerUpdate_Handler,
		},
		{
			MethodName: "ServerDelete",
			Handler:    _FsClient_ServerDelete_Handler,
		},
		{
			MethodName: "ServerList",
			Handler:    _FsClient_ServerList_Handler,
		},
		{
			MethodName: "VolumeCreate",
			Handler:    _FsClient_VolumeCreate_Handler,
		},
		{
			MethodName: "VolumeUpdate",
			Handler:    _FsClient_VolumeUpdate_Handler,
		},
		{
			MethodName: "VolumeDelete",
			Handler:    _FsClient_VolumeDelete_Handler,
		},
		{
			MethodName: "VolumeList",
			Handler:    _FsClient_VolumeList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "directfs.proto",
}

// Client API for FsServer service

type FsServerClient interface {
	// *
	// Create a volume to be served via DFS. Currently does nothing, but should still
	// be performed as in the near future it will certainly be required.
	//
	// returns RpcResult
	VolumeCreate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Update a volume to be served via DFS. Currently does nothing, but should still
	// be performed as in the near future it will certainly be required.
	//
	// returns RpcResult
	VolumeUpdate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// Delete a volume to be served via DFS. Currently does nothing, but should still
	// be performed as in the near future it will certainly be required.
	//
	// returns RpcResult
	VolumeDelete(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error)
	// *
	// List configured volume entries, optionally filtered via a DfsVolumeListQuery
	// message.
	//
	// returns a stream of DfsVolume messages, if any are available matching the filter.
	VolumeList(ctx context.Context, in *DfsVolumeListQuery, opts ...grpc.CallOption) (FsServer_VolumeListClient, error)
}

type fsServerClient struct {
	cc *grpc.ClientConn
}

func NewFsServerClient(cc *grpc.ClientConn) FsServerClient {
	return &fsServerClient{cc}
}

func (c *fsServerClient) VolumeCreate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsServer/VolumeCreate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsServerClient) VolumeUpdate(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsServer/VolumeUpdate", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsServerClient) VolumeDelete(ctx context.Context, in *DfsVolume, opts ...grpc.CallOption) (*common.RpcResult, error) {
	out := new(common.RpcResult)
	err := grpc.Invoke(ctx, "/directfs.v1.FsServer/VolumeDelete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fsServerClient) VolumeList(ctx context.Context, in *DfsVolumeListQuery, opts ...grpc.CallOption) (FsServer_VolumeListClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_FsServer_serviceDesc.Streams[0], c.cc, "/directfs.v1.FsServer/VolumeList", opts...)
	if err != nil {
		return nil, err
	}
	x := &fsServerVolumeListClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type FsServer_VolumeListClient interface {
	Recv() (*DfsVolume, error)
	grpc.ClientStream
}

type fsServerVolumeListClient struct {
	grpc.ClientStream
}

func (x *fsServerVolumeListClient) Recv() (*DfsVolume, error) {
	m := new(DfsVolume)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for FsServer service

type FsServerServer interface {
	// *
	// Create a volume to be served via DFS. Currently does nothing, but should still
	// be performed as in the near future it will certainly be required.
	//
	// returns RpcResult
	VolumeCreate(context.Context, *DfsVolume) (*common.RpcResult, error)
	// *
	// Update a volume to be served via DFS. Currently does nothing, but should still
	// be performed as in the near future it will certainly be required.
	//
	// returns RpcResult
	VolumeUpdate(context.Context, *DfsVolume) (*common.RpcResult, error)
	// *
	// Delete a volume to be served via DFS. Currently does nothing, but should still
	// be performed as in the near future it will certainly be required.
	//
	// returns RpcResult
	VolumeDelete(context.Context, *DfsVolume) (*common.RpcResult, error)
	// *
	// List configured volume entries, optionally filtered via a DfsVolumeListQuery
	// message.
	//
	// returns a stream of DfsVolume messages, if any are available matching the filter.
	VolumeList(*DfsVolumeListQuery, FsServer_VolumeListServer) error
}

func RegisterFsServerServer(s *grpc.Server, srv FsServerServer) {
	s.RegisterService(&_FsServer_serviceDesc, srv)
}

func _FsServer_VolumeCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServerServer).VolumeCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsServer/VolumeCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServerServer).VolumeCreate(ctx, req.(*DfsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsServer_VolumeUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServerServer).VolumeUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsServer/VolumeUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServerServer).VolumeUpdate(ctx, req.(*DfsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsServer_VolumeDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DfsVolume)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FsServerServer).VolumeDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/directfs.v1.FsServer/VolumeDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FsServerServer).VolumeDelete(ctx, req.(*DfsVolume))
	}
	return interceptor(ctx, in, info, handler)
}

func _FsServer_VolumeList_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DfsVolumeListQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(FsServerServer).VolumeList(m, &fsServerVolumeListServer{stream})
}

type FsServer_VolumeListServer interface {
	Send(*DfsVolume) error
	grpc.ServerStream
}

type fsServerVolumeListServer struct {
	grpc.ServerStream
}

func (x *fsServerVolumeListServer) Send(m *DfsVolume) error {
	return x.ServerStream.SendMsg(m)
}

var _FsServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "directfs.v1.FsServer",
	HandlerType: (*FsServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VolumeCreate",
			Handler:    _FsServer_VolumeCreate_Handler,
		},
		{
			MethodName: "VolumeUpdate",
			Handler:    _FsServer_VolumeUpdate_Handler,
		},
		{
			MethodName: "VolumeDelete",
			Handler:    _FsServer_VolumeDelete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "VolumeList",
			Handler:       _FsServer_VolumeList_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "directfs.proto",
}

func init() { proto.RegisterFile("directfs.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 681 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x55, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0x8d, 0x9d, 0x34, 0x89, 0x6f, 0x1e, 0x98, 0xa1, 0x14, 0x2b, 0x05, 0x35, 0xcc, 0x02, 0x45,
	0x48, 0xa4, 0x60, 0x42, 0x25, 0x96, 0x95, 0x9d, 0x96, 0xa8, 0xc8, 0x05, 0x17, 0xba, 0x8d, 0x5c,
	0x7b, 0x22, 0x2c, 0x39, 0x76, 0xe4, 0x99, 0x54, 0xea, 0x8f, 0xf0, 0x2b, 0xfc, 0x07, 0x1f, 0xc0,
	0x6f, 0xb0, 0x45, 0x33, 0xe3, 0x38, 0x4e, 0x49, 0x4a, 0xb3, 0x64, 0x37, 0x73, 0xe7, 0x9e, 0xfb,
	0x38, 0xe7, 0x58, 0x86, 0x76, 0x10, 0xa6, 0xc4, 0x67, 0x13, 0xda, 0x9f, 0xa5, 0x09, 0x4b, 0x50,
	0x23, 0xbf, 0x5f, 0xbf, 0xe9, 0x34, 0xfd, 0x64, 0x3a, 0x4d, 0x62, 0xf9, 0x84, 0x77, 0x01, 0xd9,
	0x13, 0xfa, 0x21, 0xa1, 0xcc, 0x4a, 0x49, 0x40, 0x62, 0x16, 0x7a, 0x11, 0xc5, 0x3f, 0x14, 0xa8,
	0x65, 0x61, 0xd4, 0x05, 0xd5, 0xf7, 0x0d, 0xa5, 0xab, 0xf4, 0x1a, 0xa6, 0xde, 0xb7, 0x3d, 0xe6,
	0xcd, 0x22, 0x2f, 0x26, 0x96, 0xa8, 0xe2, 0xaa, 0xbe, 0x8f, 0x9e, 0x40, 0xed, 0x5b, 0x42, 0xd9,
	0x38, 0x0c, 0x0c, 0xb5, 0xab, 0xf4, 0x5a, 0x6e, 0x95, 0x5f, 0x47, 0x01, 0xea, 0x40, 0x9d, 0x9f,
	0x62, 0x6f, 0x4a, 0x8c, 0x72, 0x57, 0xe9, 0x69, 0x6e, 0x7e, 0x47, 0x08, 0x2a, 0xb3, 0x24, 0x65,
	0x46, 0x45, 0x20, 0xc4, 0x19, 0x1d, 0x43, 0xc3, 0x5f, 0x4e, 0x61, 0xec, 0x88, 0x9e, 0x07, 0xfd,
	0xc2, 0xf4, 0xfd, 0xbf, 0x87, 0x75, 0x8b, 0x18, 0xfc, 0x1e, 0x1a, 0x59, 0xca, 0xc7, 0x90, 0x32,
	0xf4, 0x12, 0x76, 0x78, 0x47, 0x6a, 0x28, 0xdd, 0x72, 0xaf, 0x61, 0xee, 0xae, 0xab, 0xe5, 0xca,
	0x14, 0x6c, 0x81, 0x5e, 0x80, 0x7e, 0x9e, 0x93, 0xf4, 0x06, 0x1d, 0xca, 0x0d, 0xc6, 0x61, 0x70,
	0x77, 0x89, 0x9a, 0xdc, 0x98, 0xe2, 0x3d, 0xd8, 0xb5, 0x27, 0xf4, 0x32, 0x89, 0xe6, 0x53, 0x52,
	0x64, 0xf4, 0x31, 0x3c, 0xca, 0xe3, 0x17, 0xcc, 0x63, 0x21, 0x65, 0xa1, 0x4f, 0xf1, 0x4f, 0x15,
	0x1e, 0xac, 0xc4, 0xe7, 0x14, 0x39, 0x00, 0x7e, 0x12, 0xc7, 0x63, 0xca, 0x3c, 0x46, 0x04, 0xf1,
	0x6d, 0xf3, 0xf0, 0x76, 0xd7, 0x22, 0x82, 0xdf, 0xad, 0x24, 0x8e, 0x89, 0xcf, 0xc2, 0x24, 0xe6,
	0x31, 0xe2, 0x6a, 0xbc, 0x84, 0x38, 0xa2, 0x7d, 0xd0, 0x66, 0x84, 0xa4, 0x63, 0x21, 0x83, 0x2a,
	0x65, 0xe0, 0x01, 0x87, 0xcb, 0x70, 0x02, 0x35, 0xf1, 0xe8, 0x4d, 0x84, 0x42, 0x6d, 0xf3, 0xd5,
	0xbf, 0x3a, 0x1d, 0x07, 0x41, 0x4a, 0x28, 0x3d, 0xf1, 0xa6, 0x61, 0x74, 0xe3, 0x56, 0x39, 0xfa,
	0x78, 0x82, 0xaf, 0x84, 0x8f, 0x6e, 0x4d, 0x81, 0xea, 0x50, 0x71, 0xce, 0x9d, 0xa1, 0x5e, 0x42,
	0x6d, 0x00, 0xeb, 0xdc, 0x71, 0x86, 0xd6, 0x97, 0x91, 0x73, 0xaa, 0x2b, 0xa8, 0x05, 0x5a, 0x76,
	0x1f, 0xda, 0xba, 0x8a, 0x1e, 0x42, 0xcb, 0x1e, 0x5d, 0x14, 0x32, 0xca, 0x48, 0x87, 0xe6, 0x32,
	0x34, 0xb4, 0xf5, 0x0a, 0x7e, 0x21, 0x04, 0x5a, 0xe9, 0xcf, 0x3b, 0x8c, 0x3e, 0x5d, 0x0e, 0xf4,
	0x52, 0x76, 0x3a, 0xd2, 0x15, 0xfc, 0x5d, 0x05, 0x2d, 0x1f, 0xfc, 0x1e, 0xfe, 0xdd, 0x07, 0xed,
	0x5a, 0xe4, 0x2e, 0x1d, 0x5c, 0x97, 0x81, 0x51, 0x50, 0x34, 0x77, 0x79, 0xc5, 0xdc, 0xd6, 0xaa,
	0x59, 0x2b, 0xa2, 0xc1, 0xf3, 0xf5, 0xec, 0x6d, 0xb2, 0x2b, 0x3a, 0x82, 0x1d, 0x2e, 0xf3, 0xc2,
	0xeb, 0xdd, 0xcd, 0xe4, 0x4b, 0xc3, 0xb8, 0x32, 0x1d, 0x0d, 0xa0, 0x4a, 0x85, 0x22, 0x46, 0x55,
	0x00, 0x9f, 0xde, 0xa5, 0x9a, 0x9b, 0xe5, 0xe2, 0x21, 0xb4, 0xf2, 0x27, 0xf1, 0x79, 0x0c, 0xa0,
	0x26, 0x17, 0x5d, 0xb8, 0xbb, 0xb3, 0xbe, 0x0e, 0x4f, 0x76, 0x17, 0xa9, 0xf8, 0x4c, 0x68, 0xbd,
	0x7c, 0x91, 0x9f, 0xca, 0x3b, 0x80, 0x9c, 0xc5, 0x45, 0xb9, 0xbd, 0xf5, 0xe5, 0x5c, 0x6d, 0x41,
	0x2f, 0x35, 0x7f, 0x95, 0xa1, 0x7e, 0x42, 0xad, 0x28, 0x24, 0x31, 0x43, 0x26, 0x34, 0x2f, 0x48,
	0x7a, 0x4d, 0x52, 0x2b, 0x25, 0xdc, 0x3f, 0x6b, 0x3f, 0xb6, 0x0e, 0xf4, 0xdd, 0x99, 0xef, 0x12,
	0x3a, 0x8f, 0x18, 0x2e, 0x2d, 0x31, 0x5f, 0x67, 0xc1, 0xd6, 0x18, 0x9b, 0x44, 0xe4, 0x9e, 0x98,
	0x53, 0x00, 0x89, 0x11, 0xcc, 0x3d, 0x5b, 0x87, 0xc8, 0xc9, 0xe8, 0x18, 0x9b, 0x9e, 0x71, 0x09,
	0x0d, 0xa0, 0x99, 0xbb, 0x82, 0x0f, 0xbc, 0x81, 0xa4, 0x5b, 0xed, 0x73, 0x54, 0xb6, 0xe6, 0x96,
	0xa8, 0x6c, 0xd1, 0xfb, 0xa1, 0xce, 0x00, 0x0a, 0x26, 0x39, 0xd8, 0xec, 0x09, 0xb9, 0xec, 0x1d,
	0xa6, 0xc1, 0x25, 0xf3, 0xb7, 0xc2, 0x05, 0x96, 0xd4, 0xfd, 0x07, 0xbb, 0x8f, 0xb6, 0xdb, 0x7d,
	0x43, 0x51, 0x5c, 0x7a, 0xad, 0x5c, 0x55, 0xc5, 0x2f, 0xf6, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xbb, 0x6b, 0xbd, 0x78, 0x8f, 0x07, 0x00, 0x00,
}
