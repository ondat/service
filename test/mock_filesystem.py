#!/usr/bin/env python
#
# Filesystem gRPC API mock implementation.

# pylint: disable=C0111

import logging

import common.v1.common_pb2 as common_pb2
import filesystem.v1.filesystem_pb2 as filesystem_pb2
import filesystem.v1.filesystem_pb2_grpc as filesystem_pb2_grpc

# pylint: disable=C0103
log = logging.getLogger()


class FilesystemServicer(filesystem_pb2_grpc.FsServicer):
    volumes = {}
    version_string = 'notset'

    def get_volumes(self):
        return self.volumes

    def Status(self, _request, _context):
        return filesystem_pb2.FsStatus(version_info=self.version_string)

    def VolumeCreate(self, request, _context):
        # NOTE: Create a new FsVolume object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        newvol = filesystem_pb2.FsVolume(volume_id=request.volume_id)
        self.volumes[request.volume_id] = newvol
        return common_pb2.RpcResult(success=True)

    def VolumeUpdate(self, request, _context):
        # NOTE: Create a new FsVolume object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        newvol = filesystem_pb2.FsVolume(volume_id=request.volume_id)
        self.volumes[request.volume_id] = newvol
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
