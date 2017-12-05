#!/usr/bin/env python
#
# DirectFS gRPC API mock implementation.

# pylint: disable=C0111

import logging

import common.v1.common_pb2 as common_pb2
import directfs.v1.directfs_pb2 as directfs_pb2
import directfs.v1.directfs_pb2_grpc as directfs_pb2_grpc

# pylint: disable=C0103
log = logging.getLogger()


class DirectfsClientServicer(directfs_pb2_grpc.DfsClientServicer):

    servers = {}
    volumes = {}
    version_string = 'notset'

    def Status(self, _request, _context):
        return directfs_pb2.DfsClientStatus(version_info=self.version_string)

    def ServerCreate(self, request, _context):
        # NOTE: Create a new DfsHost object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        self.servers[request.host_id] = directfs_pb2.DfsHost(
            host_id=request.host_id, hostname=request.hostname, port=request.port)
        return common_pb2.RpcResult(success=True)

    def ServerUpdate(self, request, _context):
        # NOTE: Create a new DfsHost object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        self.servers[request.host_id] = directfs_pb2.DfsHost(
            host_id=request.host_id, hostname=request.hostname, port=request.port)
        return common_pb2.RpcResult(success=True)

    def ServerDelete(self, request, _context):
        del self.servers[request.host_id]
        return common_pb2.RpcResult(success=True)

    def ServerList(self, request, _context):
        slist = []
        for _k, server in self.servers.items():
            slist.append(server)
        response = directfs_pb2.DfsHostList(hosts=slist)
        return response

    def checked_volume_update(self, new_volume):
        if new_volume.host_id in self.servers:
            self.volumes[new_volume.volume_id] = new_volume
            return common_pb2.RpcResult(success=True)
        # host_id not known - failure.
        return common_pb2.RpcResult(success=False,
                                    reason="Unknown host_id {}".format(new_volume.host_id))

    def VolumeCreate(self, request, _context):
        # NOTE: Create a new DfsVolume object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        newvol = directfs_pb2.DfsVolume(volume_id=request.volume_id, host_id=request.host_id)
        return self.checked_volume_update(newvol)

    def VolumeUpdate(self, request, _context):
        # NOTE: Create a new DfsVolume object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        newvol = directfs_pb2.DfsVolume(volume_id=request.volume_id, host_id=request.host_id)
        return self.checked_volume_update(newvol)

    def VolumeDelete(self, request, _context):
        del self.volumes[request.volume_id]
        return common_pb2.RpcResult(success=True)

    def VolumeList(self, request, _context):
        vlist = []
        for _k, vol in self.volumes.items():
            vlist.append(vol)
        response = directfs_pb2.DfsVolumeList(volumes=vlist)
        return response


class DirectfsServerServicer(directfs_pb2_grpc.DfsServerServicer):

    volumes = {}
    version_string = 'notset'

    def get_volumes(self):
        return self.volumes

    def Status(self, _request, _context):
        return directfs_pb2.DfsServerStatus(version_info=self.version_string)

    def VolumeCreate(self, request, _context):
        # NOTE: Create a new DfsVolume object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        newvol = directfs_pb2.DfsVolume(volume_id=request.volume_id, host_id=request.host_id)
        self.volumes[request.volume_id] = newvol
        return common_pb2.RpcResult(success=True)

    def VolumeUpdate(self, request, _context):
        # NOTE: Create a new DfsVolume object. Do not use the request.
        # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
        newvol = directfs_pb2.DfsVolume(volume_id=request.volume_id, host_id=request.host_id)
        self.volumes[request.volume_id] = newvol
        return common_pb2.RpcResult(success=True)

    def VolumeDelete(self, request, _context):
        del self.volumes[request.volume_id]
        return common_pb2.RpcResult(success=True)

    def VolumeList(self, request, _context):
        vlist = []
        for _k, vol in self.volumes.items():
            vlist.append(vol)
        response = directfs_pb2.DfsVolumeList(volumes=vlist)
        return response
