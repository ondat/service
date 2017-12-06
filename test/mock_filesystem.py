#!/usr/bin/env python
#
# Filesystem gRPC API mock implementation.

# pylint: disable=C0111,E1101

import logging

import common.v1.common_pb2 as common_pb2
import filesystem.v1.filesystem_pb2 as filesystem_pb2
import filesystem.v1.filesystem_pb2_grpc as filesystem_pb2_grpc

from mock_common import copy_submessage

# pylint: disable=C0103
log = logging.getLogger()


def fsvolume_copy(src):
    """Return a copy of an FsVolume message."""

    # Copy scalar fields.
    msg = filesystem_pb2.FsVolume(
        volume_id=src.volume_id,
        node_type=src.node_type,
        device_number=src.device_number,
        filename=src.filename,
        linked_volume=src.linked_volume,
        target_volume_id=src.target_volume_id,
        volume_size_bytes=src.volume_size_bytes,
    )
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.stats, src, 'stats')
    copy_submessage(msg.status, src, 'status')
    return msg


def fspresentation_copy(src):
    """Return a copy of an FsPresentation message."""
    msg = filesystem_pb2.FsPresentation(
        presentation_id=src.presentation_id,
        target_id=src.target_id
    )
    copy_submessage(msg.cc, src, 'cc')
    return msg


class FilesystemServicer(filesystem_pb2_grpc.FsServicer):

    def __init__(self):
        self.volumes = {}
        self.presentations = {}
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

    def checked_presentation_update(self, new_presentation):
        if new_presentation.target_id in self.volumes:
            self.presentations[new_presentation.presentation_id] = new_presentation
            return common_pb2.RpcResult(success=True)
        # Missing target - fail.
        err = "Presentation {}, missing target {}".format(
            new_presentation.presentation_id,
            new_presentation.target_id
        )
        return common_pb2.RpcResult(success=False, reason=err)

    def PresentationCreate(self, request, _context):
        return self.checked_presentation_update(fspresentation_copy(request))

    def PresentationUpdate(self, request, _context):
        return self.checked_presentation_update(fspresentation_copy(request))

    def PresentationDelete(self, request, _context):
        del self.presentations[request.presentation_id]
        return common_pb2.RpcResult(success=True)

    def PresentationList(self, request, _context):
        plist = []
        for _k, vol in self.presentations.items():
            plist.append(vol)
        response = filesystem_pb2.FsPresentationList(presentations=plist)
        return response
