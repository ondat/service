#!/usr/bin/env python
#
# rdbplugin gRPC API mock implementation.

# pylint: disable=C0111

import logging

import common.v1.common_pb2 as common_pb2
import rdbplugin.v1.rdbplugin_pb2 as rdbplugin_pb2
import rdbplugin.v1.rdbplugin_pb2_grpc as rdbplugin_pb2_grpc

from mock_common import copy_submessage

# pylint: disable=C0103
log = logging.getLogger()


def rdbvolume_copy(src):
    """Return a copy of an RdbVolume message."""
    msg = rdbplugin_pb2.RdbVolume(
        volume_id=src.volume_id,
        volume_size_bytes=src.volume_size_bytes,
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.credentials, src, 'credentials')
    copy_submessage(msg.stats, src, 'stats')
    copy_submessage(msg.status, src, 'status')
    return msg


class RdbPluginServicer(rdbplugin_pb2_grpc.RdbPluginServicer):

    def __init__(self):
        self.volumes = {}
        self.version_string = 'notset'

    def Status(self, _request, _context):
        return rdbplugin_pb2.RdbStatus(version_info=self.version_string)

    def VolumeCreate(self, request, _context):
        self.volumes[request.volume_id] = rdbvolume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeUpdate(self, request, _context):
        self.volumes[request.volume_id] = rdbvolume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeDelete(self, request, _context):
        del self.volumes[request.volume_id]
        return common_pb2.RpcResult(success=True)

    def VolumeList(self, request, _context):
        vlist = []
        for _k, vol in self.volumes.items():
            vlist.append(vol)
        response = rdbplugin_pb2.RdbVolumeList(volumes=vlist)
        return response
