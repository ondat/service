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


def dfsclient_host_copy(src):
    """Copy a DfsInitiatorHost message."""
    msg = directfs_pb2.DfsInitiatorHost(
        host_id=src.host_id,
        hostname=src.hostname,
        port=src.port
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.credentials, src, 'credentials')
    return msg


def dfsclient_volume_copy(src):
    """Copy a DfsInitiatorVolume message."""
    # Copy scalar fields.
    msg = directfs_pb2.DfsInitiatorVolume(
        volume_id=src.volume_id,
        host_id=src.host_id,
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.credentials, src, 'credentials')
    copy_submessage(msg.stats, src, 'stats')
    copy_submessage(msg.status, src, 'status')
    return msg


def dfsserver_volume_copy(src):
    """Copy a DfsResponderVolume message."""
    # Copy scalar fields.
    msg = directfs_pb2.DfsResponderVolume(
        volume_id=src.volume_id,
        host_id=src.host_id,
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.credentials, src, 'credentials')
    copy_submessage(msg.stats, src, 'stats')
    copy_submessage(msg.status, src, 'status')
    return msg


class DirectfsInitiatorServicer(directfs_pb2_grpc.DfsInitiatorServicer):

    def __init__(self):
        self.servers = {}
        self.volumes = {}
        self.version_string = 'notset'

    def Status(self, _request, _context):
        return directfs_pb2.DfsInitiatorStatus(version_info=self.version_string)

    def ServerCreate(self, request, _context):
        self.servers[request.host_id] = dfsclient_host_copy(request)
        return common_pb2.RpcResult(success=True)

    def ServerUpdate(self, request, _context):
        self.servers[request.host_id] = dfsclient_host_copy(request)
        return common_pb2.RpcResult(success=True)

    def ServerDelete(self, request, _context):
        del self.servers[request.host_id]
        return common_pb2.RpcResult(success=True)

    def ServerList(self, request, _context):
        slist = []
        for _k, server in self.servers.items():
            slist.append(server)
        response = directfs_pb2.DfsInitiatorHostList(hosts=slist)
        return response

    def checked_volume_update(self, new_volume):
        if new_volume.host_id in self.servers:
            self.volumes[new_volume.volume_id] = new_volume
            return common_pb2.RpcResult(success=True)
        # host_id not known - failure.
        return common_pb2.RpcResult(success=False,
                                    reason="Unknown host_id {}".format(new_volume.host_id))

    def VolumeCreate(self, request, _context):
        newvol = dfsclient_volume_copy(request)
        return self.checked_volume_update(newvol)

    def VolumeUpdate(self, request, _context):
        newvol = dfsclient_volume_copy(request)
        return self.checked_volume_update(newvol)

    def VolumeDelete(self, request, _context):
        del self.volumes[request.volume_id]
        return common_pb2.RpcResult(success=True)

    def VolumeList(self, request, _context):
        vlist = []
        for _k, vol in self.volumes.items():
            vlist.append(vol)
        response = directfs_pb2.DfsInitiatorVolumeList(volumes=vlist)
        return response


class DirectfsResponderServicer(directfs_pb2_grpc.DfsResponderServicer):

    volumes = {}
    version_string = 'notset'

    def get_volumes(self):
        return self.volumes

    def Status(self, _request, _context):
        return directfs_pb2.DfsResponderStatus(version_info=self.version_string)

    def VolumeCreate(self, request, _context):
        self.volumes[request.volume_id] = dfsserver_volume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeUpdate(self, request, _context):
        self.volumes[request.volume_id] = dfsserver_volume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeDelete(self, request, _context):
        del self.volumes[request.volume_id]
        return common_pb2.RpcResult(success=True)

    def VolumeList(self, request, _context):
        vlist = []
        for _k, vol in self.volumes.items():
            vlist.append(vol)
        response = directfs_pb2.DfsResponderVolumeList(volumes=vlist)
        return response
