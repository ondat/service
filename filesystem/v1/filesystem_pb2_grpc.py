# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import common_pb2 as common__pb2
import filesystem_pb2 as filesystem__pb2


class FsStub(object):
  """*
  Filesystem configuration and status service.
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.VolumeCreate = channel.unary_unary(
        '/filesystem.v1.Fs/VolumeCreate',
        request_serializer=filesystem__pb2.FsVolume.SerializeToString,
        response_deserializer=common__pb2.RpcResult.FromString,
        )
    self.VolumeUpdate = channel.unary_unary(
        '/filesystem.v1.Fs/VolumeUpdate',
        request_serializer=filesystem__pb2.FsVolume.SerializeToString,
        response_deserializer=common__pb2.RpcResult.FromString,
        )
    self.VolumeDelete = channel.unary_unary(
        '/filesystem.v1.Fs/VolumeDelete',
        request_serializer=filesystem__pb2.FsVolume.SerializeToString,
        response_deserializer=common__pb2.RpcResult.FromString,
        )
    self.VolumeList = channel.unary_unary(
        '/filesystem.v1.Fs/VolumeList',
        request_serializer=filesystem__pb2.FsVolumeListQuery.SerializeToString,
        response_deserializer=filesystem__pb2.FsVolumeList.FromString,
        )
    self.PresentationCreate = channel.unary_unary(
        '/filesystem.v1.Fs/PresentationCreate',
        request_serializer=filesystem__pb2.FsPresentation.SerializeToString,
        response_deserializer=common__pb2.RpcResult.FromString,
        )
    self.PresentationUpdate = channel.unary_unary(
        '/filesystem.v1.Fs/PresentationUpdate',
        request_serializer=filesystem__pb2.FsPresentation.SerializeToString,
        response_deserializer=common__pb2.RpcResult.FromString,
        )
    self.PresentationDelete = channel.unary_unary(
        '/filesystem.v1.Fs/PresentationDelete',
        request_serializer=filesystem__pb2.FsPresentation.SerializeToString,
        response_deserializer=common__pb2.RpcResult.FromString,
        )
    self.PresentationList = channel.unary_unary(
        '/filesystem.v1.Fs/PresentationList',
        request_serializer=filesystem__pb2.FsPresentationListQuery.SerializeToString,
        response_deserializer=filesystem__pb2.FsPresentationList.FromString,
        )


class FsServicer(object):
  """*
  Filesystem configuration and status service.
  """

  def VolumeCreate(self, request, context):
    """*
    Create the specified FsVolume.

    returns RpcResult
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def VolumeUpdate(self, request, context):
    """*
    Update the specified FsVolume.

    returns RpcResult
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def VolumeDelete(self, request, context):
    """*
    Delete the specified FsVolume.

    returns RpcResult
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def VolumeList(self, request, context):
    """*
    Return a list of FsVolume messages, optionally filtered using the supplied
    FsVolumeListQuery message.

    returns A FsVolumeList message containing FsVolume objects,
    if any are found that match the filter.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PresentationCreate(self, request, context):
    """*
    Add configuration for a Presentation volume specified in the FsPresentation message.

    returns RpcResult
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PresentationUpdate(self, request, context):
    """*
    Update configuration for a Presentation volume specified in the FsPresentation message.

    returns RpcResult
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PresentationDelete(self, request, context):
    """*
    Remove configuration for the Presentation volume specified in the FsPresentation message.

    returns RpcResult
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PresentationList(self, request, context):
    """*
    List configured Presentation volumes, optionally filtered using a FsPresentationListQuery
    message.

    returns A FsPresentationList message containing FsPresentation mesages,
    if any are found matching the filter.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_FsServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'VolumeCreate': grpc.unary_unary_rpc_method_handler(
          servicer.VolumeCreate,
          request_deserializer=filesystem__pb2.FsVolume.FromString,
          response_serializer=common__pb2.RpcResult.SerializeToString,
      ),
      'VolumeUpdate': grpc.unary_unary_rpc_method_handler(
          servicer.VolumeUpdate,
          request_deserializer=filesystem__pb2.FsVolume.FromString,
          response_serializer=common__pb2.RpcResult.SerializeToString,
      ),
      'VolumeDelete': grpc.unary_unary_rpc_method_handler(
          servicer.VolumeDelete,
          request_deserializer=filesystem__pb2.FsVolume.FromString,
          response_serializer=common__pb2.RpcResult.SerializeToString,
      ),
      'VolumeList': grpc.unary_unary_rpc_method_handler(
          servicer.VolumeList,
          request_deserializer=filesystem__pb2.FsVolumeListQuery.FromString,
          response_serializer=filesystem__pb2.FsVolumeList.SerializeToString,
      ),
      'PresentationCreate': grpc.unary_unary_rpc_method_handler(
          servicer.PresentationCreate,
          request_deserializer=filesystem__pb2.FsPresentation.FromString,
          response_serializer=common__pb2.RpcResult.SerializeToString,
      ),
      'PresentationUpdate': grpc.unary_unary_rpc_method_handler(
          servicer.PresentationUpdate,
          request_deserializer=filesystem__pb2.FsPresentation.FromString,
          response_serializer=common__pb2.RpcResult.SerializeToString,
      ),
      'PresentationDelete': grpc.unary_unary_rpc_method_handler(
          servicer.PresentationDelete,
          request_deserializer=filesystem__pb2.FsPresentation.FromString,
          response_serializer=common__pb2.RpcResult.SerializeToString,
      ),
      'PresentationList': grpc.unary_unary_rpc_method_handler(
          servicer.PresentationList,
          request_deserializer=filesystem__pb2.FsPresentationListQuery.FromString,
          response_serializer=filesystem__pb2.FsPresentationList.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'filesystem.v1.Fs', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
