// Generated by the gRPC C++ plugin.
// If you make any local change, they will be lost.
// source: director.proto

#include "director.pb.h"
#include "director.grpc.pb.h"

#include <grpc++/impl/codegen/async_stream.h>
#include <grpc++/impl/codegen/async_unary_call.h>
#include <grpc++/impl/codegen/channel_interface.h>
#include <grpc++/impl/codegen/client_unary_call.h>
#include <grpc++/impl/codegen/method_handler_impl.h>
#include <grpc++/impl/codegen/rpc_service_method.h>
#include <grpc++/impl/codegen/service_type.h>
#include <grpc++/impl/codegen/sync_stream.h>
namespace director {
namespace v1 {

static const char* Director_method_names[] = {
  "/director.v1.Director/VolumeCreate",
  "/director.v1.Director/VolumeUpdate",
  "/director.v1.Director/VolumeDelete",
  "/director.v1.Director/VolumeList",
  "/director.v1.Director/PresentationCreate",
  "/director.v1.Director/PresentationUpdate",
  "/director.v1.Director/PresentationDelete",
  "/director.v1.Director/PresentationList",
};

std::unique_ptr< Director::Stub> Director::NewStub(const std::shared_ptr< ::grpc::ChannelInterface>& channel, const ::grpc::StubOptions& options) {
  std::unique_ptr< Director::Stub> stub(new Director::Stub(channel));
  return stub;
}

Director::Stub::Stub(const std::shared_ptr< ::grpc::ChannelInterface>& channel)
  : channel_(channel), rpcmethod_VolumeCreate_(Director_method_names[0], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_VolumeUpdate_(Director_method_names[1], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_VolumeDelete_(Director_method_names[2], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_VolumeList_(Director_method_names[3], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_PresentationCreate_(Director_method_names[4], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_PresentationUpdate_(Director_method_names[5], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_PresentationDelete_(Director_method_names[6], ::grpc::RpcMethod::NORMAL_RPC, channel)
  , rpcmethod_PresentationList_(Director_method_names[7], ::grpc::RpcMethod::NORMAL_RPC, channel)
  {}

::grpc::Status Director::Stub::VolumeCreate(::grpc::ClientContext* context, const ::director::v1::DirectorVolume& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeCreate_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* Director::Stub::AsyncVolumeCreateRaw(::grpc::ClientContext* context, const ::director::v1::DirectorVolume& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_VolumeCreate_, context, request);
}

::grpc::Status Director::Stub::VolumeUpdate(::grpc::ClientContext* context, const ::director::v1::DirectorVolume& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeUpdate_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* Director::Stub::AsyncVolumeUpdateRaw(::grpc::ClientContext* context, const ::director::v1::DirectorVolume& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_VolumeUpdate_, context, request);
}

::grpc::Status Director::Stub::VolumeDelete(::grpc::ClientContext* context, const ::director::v1::DirectorVolume& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeDelete_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* Director::Stub::AsyncVolumeDeleteRaw(::grpc::ClientContext* context, const ::director::v1::DirectorVolume& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_VolumeDelete_, context, request);
}

::grpc::Status Director::Stub::VolumeList(::grpc::ClientContext* context, const ::director::v1::DirectorVolumeListQuery& request, ::director::v1::DirectorVolumeList* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_VolumeList_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::director::v1::DirectorVolumeList>* Director::Stub::AsyncVolumeListRaw(::grpc::ClientContext* context, const ::director::v1::DirectorVolumeListQuery& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::director::v1::DirectorVolumeList>::Create(channel_.get(), cq, rpcmethod_VolumeList_, context, request);
}

::grpc::Status Director::Stub::PresentationCreate(::grpc::ClientContext* context, const ::director::v1::DirectorPresentation& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_PresentationCreate_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* Director::Stub::AsyncPresentationCreateRaw(::grpc::ClientContext* context, const ::director::v1::DirectorPresentation& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_PresentationCreate_, context, request);
}

::grpc::Status Director::Stub::PresentationUpdate(::grpc::ClientContext* context, const ::director::v1::DirectorPresentation& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_PresentationUpdate_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* Director::Stub::AsyncPresentationUpdateRaw(::grpc::ClientContext* context, const ::director::v1::DirectorPresentation& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_PresentationUpdate_, context, request);
}

::grpc::Status Director::Stub::PresentationDelete(::grpc::ClientContext* context, const ::director::v1::DirectorPresentation& request, ::common::v1::RpcResult* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_PresentationDelete_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>* Director::Stub::AsyncPresentationDeleteRaw(::grpc::ClientContext* context, const ::director::v1::DirectorPresentation& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::common::v1::RpcResult>::Create(channel_.get(), cq, rpcmethod_PresentationDelete_, context, request);
}

::grpc::Status Director::Stub::PresentationList(::grpc::ClientContext* context, const ::director::v1::DirectorPresentationListQuery& request, ::director::v1::DirectorPresentationList* response) {
  return ::grpc::BlockingUnaryCall(channel_.get(), rpcmethod_PresentationList_, context, request, response);
}

::grpc::ClientAsyncResponseReader< ::director::v1::DirectorPresentationList>* Director::Stub::AsyncPresentationListRaw(::grpc::ClientContext* context, const ::director::v1::DirectorPresentationListQuery& request, ::grpc::CompletionQueue* cq) {
  return ::grpc::ClientAsyncResponseReader< ::director::v1::DirectorPresentationList>::Create(channel_.get(), cq, rpcmethod_PresentationList_, context, request);
}

Director::Service::Service() {
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[0],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorVolume, ::common::v1::RpcResult>(
          std::mem_fn(&Director::Service::VolumeCreate), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[1],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorVolume, ::common::v1::RpcResult>(
          std::mem_fn(&Director::Service::VolumeUpdate), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[2],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorVolume, ::common::v1::RpcResult>(
          std::mem_fn(&Director::Service::VolumeDelete), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[3],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorVolumeListQuery, ::director::v1::DirectorVolumeList>(
          std::mem_fn(&Director::Service::VolumeList), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[4],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorPresentation, ::common::v1::RpcResult>(
          std::mem_fn(&Director::Service::PresentationCreate), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[5],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorPresentation, ::common::v1::RpcResult>(
          std::mem_fn(&Director::Service::PresentationUpdate), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[6],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorPresentation, ::common::v1::RpcResult>(
          std::mem_fn(&Director::Service::PresentationDelete), this)));
  AddMethod(new ::grpc::RpcServiceMethod(
      Director_method_names[7],
      ::grpc::RpcMethod::NORMAL_RPC,
      new ::grpc::RpcMethodHandler< Director::Service, ::director::v1::DirectorPresentationListQuery, ::director::v1::DirectorPresentationList>(
          std::mem_fn(&Director::Service::PresentationList), this)));
}

Director::Service::~Service() {
}

::grpc::Status Director::Service::VolumeCreate(::grpc::ServerContext* context, const ::director::v1::DirectorVolume* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Director::Service::VolumeUpdate(::grpc::ServerContext* context, const ::director::v1::DirectorVolume* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Director::Service::VolumeDelete(::grpc::ServerContext* context, const ::director::v1::DirectorVolume* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Director::Service::VolumeList(::grpc::ServerContext* context, const ::director::v1::DirectorVolumeListQuery* request, ::director::v1::DirectorVolumeList* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Director::Service::PresentationCreate(::grpc::ServerContext* context, const ::director::v1::DirectorPresentation* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Director::Service::PresentationUpdate(::grpc::ServerContext* context, const ::director::v1::DirectorPresentation* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Director::Service::PresentationDelete(::grpc::ServerContext* context, const ::director::v1::DirectorPresentation* request, ::common::v1::RpcResult* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}

::grpc::Status Director::Service::PresentationList(::grpc::ServerContext* context, const ::director::v1::DirectorPresentationListQuery* request, ::director::v1::DirectorPresentationList* response) {
  (void) context;
  (void) request;
  (void) response;
  return ::grpc::Status(::grpc::StatusCode::UNIMPLEMENTED, "");
}


}  // namespace director
}  // namespace v1

