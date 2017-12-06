#!/usr/bin/env python
#
# Filesystem gRPC API mock implementation.

# pylint: disable=C0111,E1101

import logging

import common.v1.common_pb2 as common_pb2
import filesystem.v1.filesystem_pb2 as filesystem_pb2
import filesystem.v1.filesystem_pb2_grpc as filesystem_pb2_grpc

import mock_common

# pylint: disable=C0103
log = logging.getLogger()


def fsvolume_copy(request):
    """Return a copy of an FsVolume message."""

    # Copy scalar fields.
    msg = filesystem_pb2.FsVolume(
        volume_id=request.volume_id,
        node_type=request.node_type,
        device_number=request.device_number,
        filename=request.filename,
        linked_volume=request.linked_volume,
        target_volume_id=request.target_volume_id,
        volume_size_bytes=request.volume_size_bytes,
    )
    mock_common.copy_submessage(msg.cc, request, 'cc')
    mock_common.copy_submessage(msg.stats, request, 'stats')
    mock_common.copy_submessage(msg.status, request, 'status')
    return msg


class FilesystemServicer(filesystem_pb2_grpc.FsServicer):

    def __init__(self):
        self.volumes = {}
        self.version_string = 'notset'

    def get_volumes(self):
        return self.volumes

    def Status(self, _request, _context):
        return filesystem_pb2.FsStatus(version_info=self.version_string)

    def VolumeCreate(self, request, _context):
        self.volumes[request.volume_id] = fsvolume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeUpdate(self, request, _context):
        self.volumes[request.volume_id] = fsvolume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeDelete(self, request, _context):
        del self.volumes[request.volume_id]
        return common_pb2.RpcResult(success=True)

    def VolumeList(self, request, _context):
        vlist = []
        for _k, vol in self.volumes.items():
            vlist.append(vol)
        response = filesystem_pb2.FsVolumeList(volumes=vlist)
        return response
