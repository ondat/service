// Code generated by protoc-gen-go. DO NOT EDIT.
// source: supervisor.proto

package supervisor_v1

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

type SupervisorStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupervisorStatusRequest) Reset()         { *m = SupervisorStatusRequest{} }
func (m *SupervisorStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SupervisorStatusRequest) ProtoMessage()    {}
func (*SupervisorStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{0}
}

func (m *SupervisorStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupervisorStatusRequest.Unmarshal(m, b)
}
func (m *SupervisorStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupervisorStatusRequest.Marshal(b, m, deterministic)
}
func (m *SupervisorStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupervisorStatusRequest.Merge(m, src)
}
func (m *SupervisorStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SupervisorStatusRequest.Size(m)
}
func (m *SupervisorStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SupervisorStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SupervisorStatusRequest proto.InternalMessageInfo

type SupervisorStatus struct {
	// Generic daemon status.
	Status               *v1.DaemonStatus `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SupervisorStatus) Reset()         { *m = SupervisorStatus{} }
func (m *SupervisorStatus) String() string { return proto.CompactTextString(m) }
func (*SupervisorStatus) ProtoMessage()    {}
func (*SupervisorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8b9452d77b1c7d2, []int{1}
}

func (m *SupervisorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupervisorStatus.Unmarshal(m, b)
}
func (m *SupervisorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupervisorStatus.Marshal(b, m, deterministic)
}
func (m *SupervisorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupervisorStatus.Merge(m, src)
}
func (m *SupervisorStatus) XXX_Size() int {
	return xxx_messageInfo_SupervisorStatus.Size(m)
}
func (m *SupervisorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SupervisorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SupervisorStatus proto.InternalMessageInfo

func (m *SupervisorStatus) GetStatus() *v1.DaemonStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterType((*SupervisorStatusRequest)(nil), "supervisor.v1.SupervisorStatusRequest")
	proto.RegisterType((*SupervisorStatus)(nil), "supervisor.v1.SupervisorStatus")
}

func init() { proto.RegisterFile("supervisor.proto", fileDescriptor_b8b9452d77b1c7d2) }

var fileDescriptor_b8b9452d77b1c7d2 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4f, 0xc2, 0x40,
	0x14, 0xc4, 0xab, 0x21, 0x1c, 0x9e, 0xa2, 0x64, 0xa3, 0xa9, 0xac, 0x1a, 0xcc, 0x1e, 0x8c, 0xa7,
	0x1a, 0xf0, 0xe0, 0x5d, 0x8c, 0xc6, 0x3f, 0x24, 0xda, 0xc6, 0x0f, 0x50, 0xf5, 0x49, 0x1a, 0x69,
	0x5f, 0xed, 0x6e, 0x9b, 0xf4, 0xeb, 0xfa, 0x49, 0x4c, 0xbb, 0x0b, 0x14, 0x4a, 0xe5, 0x46, 0x66,
	0xde, 0xfc, 0x76, 0x86, 0x14, 0xba, 0x32, 0x8d, 0x31, 0xc9, 0x02, 0x49, 0x89, 0x13, 0x27, 0xa4,
	0x88, 0x75, 0x2a, 0x4a, 0x36, 0xe0, 0xbb, 0x1f, 0x14, 0x86, 0x14, 0x69, 0x53, 0xf4, 0xc0, 0xf6,
	0xe6, 0xb6, 0xa7, 0x7c, 0x95, 0x4a, 0x17, 0x7f, 0x52, 0x94, 0x4a, 0x8c, 0xa0, 0xbb, 0x6a, 0xb1,
	0x4b, 0x68, 0xcb, 0xf2, 0xd7, 0xd1, 0xf6, 0xd9, 0xd6, 0xc5, 0xce, 0xd0, 0x76, 0x0c, 0x2d, 0x1b,
	0x38, 0xb7, 0x3e, 0x86, 0x14, 0x19, 0x86, 0x39, 0x1b, 0xfe, 0xb6, 0x00, 0x16, 0x14, 0xe6, 0x41,
	0xdb, 0x90, 0xce, 0x9d, 0xa5, 0x5a, 0x4e, 0x43, 0x0b, 0xde, 0xdf, 0x70, 0x27, 0x2c, 0x76, 0x0d,
	0xad, 0x17, 0xfa, 0x46, 0x76, 0x50, 0x29, 0x53, 0x08, 0xaf, 0x29, 0x26, 0x39, 0xb7, 0x57, 0x54,
	0x17, 0x65, 0x4c, 0x91, 0x44, 0x61, 0xb1, 0x3b, 0xe8, 0x8c, 0x28, 0xfa, 0x0a, 0x26, 0xf7, 0xa8,
	0x6e, 0x88, 0xa6, 0x4b, 0x04, 0xed, 0x3c, 0x61, 0xce, 0x4f, 0x6b, 0xaa, 0xb9, 0x77, 0x31, 0x9e,
	0xe6, 0xc2, 0x62, 0x0f, 0xd0, 0xd5, 0xfa, 0x5b, 0xfc, 0xe9, 0x2b, 0x2c, 0x51, 0x87, 0xb5, 0x50,
	0x21, 0xf3, 0x93, 0x9a, 0xac, 0x33, 0x0b, 0xd4, 0x9e, 0x96, 0x9f, 0x03, 0xa9, 0x3b, 0xf1, 0x5a,
	0xa2, 0xb0, 0xf4, 0xb6, 0xde, 0xda, 0x47, 0x0a, 0x5f, 0x58, 0xec, 0x11, 0xf6, 0xe7, 0x6d, 0x3d,
	0x95, 0x04, 0xd1, 0xa4, 0x61, 0x5f, 0x7f, 0xdd, 0x3e, 0x9d, 0x98, 0xd5, 0x1a, 0x03, 0xab, 0xb6,
	0x35, 0x38, 0xbb, 0x16, 0xd4, 0xc6, 0xc6, 0x95, 0xe3, 0xd9, 0x1f, 0x56, 0x54, 0x35, 0xb0, 0xff,
	0x76, 0x1e, 0x37, 0x3c, 0xa4, 0x97, 0xbe, 0xb7, 0xcb, 0x6f, 0xf9, 0xea, 0x2f, 0x00, 0x00, 0xff,
	0xff, 0x8d, 0xb6, 0xe7, 0xfd, 0xfc, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SupervisorClient is the client API for Supervisor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SupervisorClient interface {
	//*
	// Get program status.
	Status(ctx context.Context, in *SupervisorStatusRequest, opts ...grpc.CallOption) (*SupervisorStatus, error)
	//*
	// Perform actions that really only make sense for very low-level testing.
	Poke(ctx context.Context, in *v1.PokeQuery, opts ...grpc.CallOption) (*v1.PokeResponse, error)
	// Config services, from common.v1.
	ConfigGetBool(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetBoolReply, error)
	ConfigUpdateBool(ctx context.Context, in *v1.ConfigBool, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error)
	ConfigListBool(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigBoolList, error)
	ConfigGetString(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetStringReply, error)
	ConfigUpdateString(ctx context.Context, in *v1.ConfigString, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error)
	ConfigListString(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigStringList, error)
}

type supervisorClient struct {
	cc *grpc.ClientConn
}

func NewSupervisorClient(cc *grpc.ClientConn) SupervisorClient {
	return &supervisorClient{cc}
}

func (c *supervisorClient) Status(ctx context.Context, in *SupervisorStatusRequest, opts ...grpc.CallOption) (*SupervisorStatus, error) {
	out := new(SupervisorStatus)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) Poke(ctx context.Context, in *v1.PokeQuery, opts ...grpc.CallOption) (*v1.PokeResponse, error) {
	out := new(v1.PokeResponse)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/Poke", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) ConfigGetBool(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetBoolReply, error) {
	out := new(v1.ConfigGetBoolReply)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/ConfigGetBool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) ConfigUpdateBool(ctx context.Context, in *v1.ConfigBool, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
	out := new(v1.ConfigUpdateReply)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/ConfigUpdateBool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) ConfigListBool(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigBoolList, error) {
	out := new(v1.ConfigBoolList)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/ConfigListBool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) ConfigGetString(ctx context.Context, in *v1.ConfigKey, opts ...grpc.CallOption) (*v1.ConfigGetStringReply, error) {
	out := new(v1.ConfigGetStringReply)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/ConfigGetString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) ConfigUpdateString(ctx context.Context, in *v1.ConfigString, opts ...grpc.CallOption) (*v1.ConfigUpdateReply, error) {
	out := new(v1.ConfigUpdateReply)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/ConfigUpdateString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supervisorClient) ConfigListString(ctx context.Context, in *v1.ConfigListQuery, opts ...grpc.CallOption) (*v1.ConfigStringList, error) {
	out := new(v1.ConfigStringList)
	err := c.cc.Invoke(ctx, "/supervisor.v1.Supervisor/ConfigListString", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupervisorServer is the server API for Supervisor service.
type SupervisorServer interface {
	//*
	// Get program status.
	Status(context.Context, *SupervisorStatusRequest) (*SupervisorStatus, error)
	//*
	// Perform actions that really only make sense for very low-level testing.
	Poke(context.Context, *v1.PokeQuery) (*v1.PokeResponse, error)
	// Config services, from common.v1.
	ConfigGetBool(context.Context, *v1.ConfigKey) (*v1.ConfigGetBoolReply, error)
	ConfigUpdateBool(context.Context, *v1.ConfigBool) (*v1.ConfigUpdateReply, error)
	ConfigListBool(context.Context, *v1.ConfigListQuery) (*v1.ConfigBoolList, error)
	ConfigGetString(context.Context, *v1.ConfigKey) (*v1.ConfigGetStringReply, error)
	ConfigUpdateString(context.Context, *v1.ConfigString) (*v1.ConfigUpdateReply, error)
	ConfigListString(context.Context, *v1.ConfigListQuery) (*v1.ConfigStringList, error)
}

// UnimplementedSupervisorServer can be embedded to have forward compatible implementations.
type UnimplementedSupervisorServer struct {
}

func (*UnimplementedSupervisorServer) Status(ctx context.Context, req *SupervisorStatusRequest) (*SupervisorStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (*UnimplementedSupervisorServer) Poke(ctx context.Context, req *v1.PokeQuery) (*v1.PokeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Poke not implemented")
}
func (*UnimplementedSupervisorServer) ConfigGetBool(ctx context.Context, req *v1.ConfigKey) (*v1.ConfigGetBoolReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigGetBool not implemented")
}
func (*UnimplementedSupervisorServer) ConfigUpdateBool(ctx context.Context, req *v1.ConfigBool) (*v1.ConfigUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigUpdateBool not implemented")
}
func (*UnimplementedSupervisorServer) ConfigListBool(ctx context.Context, req *v1.ConfigListQuery) (*v1.ConfigBoolList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigListBool not implemented")
}
func (*UnimplementedSupervisorServer) ConfigGetString(ctx context.Context, req *v1.ConfigKey) (*v1.ConfigGetStringReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigGetString not implemented")
}
func (*UnimplementedSupervisorServer) ConfigUpdateString(ctx context.Context, req *v1.ConfigString) (*v1.ConfigUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigUpdateString not implemented")
}
func (*UnimplementedSupervisorServer) ConfigListString(ctx context.Context, req *v1.ConfigListQuery) (*v1.ConfigStringList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfigListString not implemented")
}

func RegisterSupervisorServer(s *grpc.Server, srv SupervisorServer) {
	s.RegisterService(&_Supervisor_serviceDesc, srv)
}

func _Supervisor_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupervisorStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Status(ctx, req.(*SupervisorStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_Poke_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.PokeQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).Poke(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/Poke",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).Poke(ctx, req.(*v1.PokeQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_ConfigGetBool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ConfigGetBool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/ConfigGetBool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ConfigGetBool(ctx, req.(*v1.ConfigKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_ConfigUpdateBool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigBool)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ConfigUpdateBool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/ConfigUpdateBool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ConfigUpdateBool(ctx, req.(*v1.ConfigBool))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_ConfigListBool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ConfigListBool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/ConfigListBool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ConfigListBool(ctx, req.(*v1.ConfigListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_ConfigGetString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ConfigGetString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/ConfigGetString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ConfigGetString(ctx, req.(*v1.ConfigKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_ConfigUpdateString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigString)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ConfigUpdateString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/ConfigUpdateString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ConfigUpdateString(ctx, req.(*v1.ConfigString))
	}
	return interceptor(ctx, in, info, handler)
}

func _Supervisor_ConfigListString_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.ConfigListQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServer).ConfigListString(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/supervisor.v1.Supervisor/ConfigListString",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServer).ConfigListString(ctx, req.(*v1.ConfigListQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _Supervisor_serviceDesc = grpc.ServiceDesc{
	ServiceName: "supervisor.v1.Supervisor",
	HandlerType: (*SupervisorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Status",
			Handler:    _Supervisor_Status_Handler,
		},
		{
			MethodName: "Poke",
			Handler:    _Supervisor_Poke_Handler,
		},
		{
			MethodName: "ConfigGetBool",
			Handler:    _Supervisor_ConfigGetBool_Handler,
		},
		{
			MethodName: "ConfigUpdateBool",
			Handler:    _Supervisor_ConfigUpdateBool_Handler,
		},
		{
			MethodName: "ConfigListBool",
			Handler:    _Supervisor_ConfigListBool_Handler,
		},
		{
			MethodName: "ConfigGetString",
			Handler:    _Supervisor_ConfigGetString_Handler,
		},
		{
			MethodName: "ConfigUpdateString",
			Handler:    _Supervisor_ConfigUpdateString_Handler,
		},
		{
			MethodName: "ConfigListString",
			Handler:    _Supervisor_ConfigListString_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor.proto",
}