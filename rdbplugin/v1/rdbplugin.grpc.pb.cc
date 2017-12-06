// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: rdbplugin.proto

#include "rdbplugin.pb.h"
#include "rdbplugin.grpc.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/channel_interface.h>
#include <grpc++/impl/codegen/client_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/rpc_service_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/sync_stream.h>
namespace rdbplugin {
namespace v1 {

static const char* RdbPlugin_method_names[] = {
  "/rdbplugin.v1.RdbPlugin/Status",
  "/rdbplugin.v1.RdbPlugin/VolumeCreate",
  "/rdbplugin.v1.RdbPlugin/VolumeUpdate",
  "/rdbplugin.v1.RdbPlugin/VolumeDelete",
  "/rdbplugin.v1.RdbPlugin/VolumeList",
};

std::unique_ptr< RdbPlugin::Stub> RdbPlugin::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< RdbPlugin::Stub> stub(new RdbPlugin::Stub(channel));
  return stub;
}

RdbPlugin::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_Status_(RdbPlugin_method_names[0], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_VolumeCreate_(RdbPlugin_method_names[1], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_VolumeUpdate_(RdbPlugin_method_names[2], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_VolumeDelete_(RdbPlugin_method_names[3], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_VolumeList_(RdbPlugin_method_names[4], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status RdbPlugin::Stub::Status(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbStatusRequest& request, ::rdbplugin::v1::RdbStatus* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_Status_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::rdbplugin::v1::RdbStatus>* RdbPlugin::Stub::AsyncStatusRaw(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbStatusRequest& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::rdbplugin::v1::RdbStatus>::Create(channel_.get(), cq, rpcmethod_Status_, context, request);
}

::grpc::Status RdbPlugin::Stub::VolumeCreate(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolume& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeCreate_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* RdbPlugin::Stub::AsyncVolumeCreateRaw(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolume& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_VolumeCreate_, context, request);
}

::grpc::Status RdbPlugin::Stub::VolumeUpdate(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolume& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeUpdate_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* RdbPlugin::Stub::AsyncVolumeUpdateRaw(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolume& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_VolumeUpdate_, context, request);
}

::grpc::Status RdbPlugin::Stub::VolumeDelete(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolume& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeDelete_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* RdbPlugin::Stub::AsyncVolumeDeleteRaw(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolume& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_VolumeDelete_, context, request);
}

::grpc::Status RdbPlugin::Stub::VolumeList(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolumeListQuery& request, ::rdbplugin::v1::RdbVolumeList* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeList_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::rdbplugin::v1::RdbVolumeList>* RdbPlugin::Stub::AsyncVolumeListRaw(::grpc::ClientContext* context, const ::rdbplugin::v1::RdbVolumeListQuery& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::rdbplugin::v1::RdbVolumeList>::Create(channel_.get(), cq, rpcmethod_VolumeList_, context, request);
}

RdbPlugin::Service::Service() {
  AddMethod(new ::grpc::RpcServiceMethod(
      RdbPlugin_method_names[0],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< RdbPlugin::Service, ::rdbplugin::v1::RdbStatusRequest, ::rdbplugin::v1::RdbStatus>(
          std::mem_fn(&RdbPlugin::Service::Status), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      RdbPlugin_method_names[1],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< RdbPlugin::Service, ::rdbplugin::v1::RdbVolume, ::common::v1::RpcResult>(
          std::mem_fn(&RdbPlugin::Service::VolumeCreate), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      RdbPlugin_method_names[2],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< RdbPlugin::Service, ::rdbplugin::v1::RdbVolume, ::common::v1::RpcResult>(
          std::mem_fn(&RdbPlugin::Service::VolumeUpdate), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      RdbPlugin_method_names[3],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< RdbPlugin::Service, ::rdbplugin::v1::RdbVolume, ::common::v1::RpcResult>(
          std::mem_fn(&RdbPlugin::Service::VolumeDelete), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      RdbPlugin_method_names[4],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< RdbPlugin::Service, ::rdbplugin::v1::RdbVolumeListQuery, ::rdbplugin::v1::RdbVolumeList>(
          std::mem_fn(&RdbPlugin::Service::VolumeList), this)));
}

RdbPlugin::Service::~Service() {
}

::grpc::Status RdbPlugin::Service::Status(::grpc::ServerContext* context, const ::rdbplugin::v1::RdbStatusRequest* request, ::rdbplugin::v1::RdbStatus* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status RdbPlugin::Service::VolumeCreate(::grpc::ServerContext* context, const ::rdbplugin::v1::RdbVolume* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status RdbPlugin::Service::VolumeUpdate(::grpc::ServerContext* context, const ::rdbplugin::v1::RdbVolume* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status RdbPlugin::Service::VolumeDelete(::grpc::ServerContext* context, const ::rdbplugin::v1::RdbVolume* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status RdbPlugin::Service::VolumeList(::grpc::ServerContext* context, const ::rdbplugin::v1::RdbVolumeListQuery* request, ::rdbplugin::v1::RdbVolumeList* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace rdbplugin
}  // namespace v1
