#!/usr/bin/env python
#
# DirectFS gRPC API mock implementation.

# pylint: disable=C0111

import logging

import common.v1.common_pb2 as common_pb2
import directfs.v1.directfs_pb2 as directfs_pb2
import directfs.v1.directfs_pb2_grpc as directfs_pb2_grpc

from mock_common import copy_submessage

# pylint: disable=C0103
log = logging.getLogger()


def dfshost_copy(src):
    """Copy a DfsHost message."""
    msg = directfs_pb2.DfsHost(
        host_id=src.host_id,
        hostname=src.hostname,
        port=src.port
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.credentials, src, 'credentials')
    return msg


def dfsvolume_copy(src):
    """Copy a DfsVolume message."""
    # Copy scalar fields.
    msg = directfs_pb2.DfsVolume(
        volume_id=src.volume_id,
        host_id=src.host_id,
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.credentials, src, 'credentials')
    copy_submessage(msg.stats, src, 'stats')
    copy_submessage(msg.status, src, 'status')
    return msg


class DirectfsClientServicer(directfs_pb2_grpc.DfsClientServicer):

    def __init__(self):
        self.servers = {}
        self.volumes = {}
        self.version_string = 'notset'

    def Status(self, _request, _context):
        return directfs_pb2.DfsClientStatus(version_info=self.version_string)

    def ServerCreate(self, request, _context):
        self.servers[request.host_id] = dfshost_copy(request)
        return common_pb2.RpcResult(success=True)

    def ServerUpdate(self, request, _context):
        self.servers[request.host_id] = dfshost_copy(request)
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
        newvol = dfsvolume_copy(request)
        return self.checked_volume_update(newvol)

    def VolumeUpdate(self, request, _context):
        newvol = dfsvolume_copy(request)
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
        self.volumes[request.volume_id] = dfsvolume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeUpdate(self, request, _context):
        self.volumes[request.volume_id] = dfsvolume_copy(request)
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
