#!/usr/bin/env python
#
# Director gRPC API mock implementation.

# pylint: disable=C0111

import logging

import common.v1.common_pb2 as common_pb2
import director.v1.director_pb2 as director_pb2
import director.v1.director_pb2_grpc as director_pb2_grpc

from mock_common import copy_submessage

# pylint: disable=C0103
log = logging.getLogger()


def directorvolume_copy(src):
    """Copy a DirectorVolume message."""
    msg = director_pb2.DirectorVolume(
        volume_id=src.volume_id,
        write_pipe=src.write_pipe,
        read_pipe=src.read_pipe,
        qos=src.qos,
        replica_ids=src.replica_ids,
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    copy_submessage(msg.stats, src, 'stats')
    copy_submessage(msg.status, src, 'status')
    return msg


def directorpresentation_copy(src):
    """Copy a DirectorPresentation message."""
    msg = director_pb2.DirectorPresentation(
        presentation_id=src.presentation_id,
        target_id=src.target_id
    )
    # pylint: disable=E1101
    copy_submessage(msg.cc, src, 'cc')
    return msg


class DirectorServicer(director_pb2_grpc.DirectorServicer):

    version_string = 'notset'

    def __init__(self):
        self.volumes = {}
        self.presentation = {}

    def Status(self, _request, _context):
        return director_pb2.DirectorStatus(version_info=self.version_string)

    def VolumeCreate(self, request, _context):
        self.volumes[request.volume_id] = directorvolume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeUpdate(self, request, _context):
        self.volumes[request.volume_id] = directorvolume_copy(request)
        return common_pb2.RpcResult(success=True)

    def VolumeDelete(self, request, _context):
        del self.volumes[request.volume_id]
        return common_pb2.RpcResult(success=True)

    def VolumeList(self, request, _context):
        vlist = []
        for _k, vol in self.volumes.items():
            vlist.append(vol)
        response = director_pb2.DirectorVolumeList(volumes=vlist)
        return response

    def checked_presentation_update(self, new_presentation):
        if new_presentation.target_id in self.volumes:
            self.presentation[new_presentation.presentation_id] = new_presentation
            return common_pb2.RpcResult(success=True)
        # Missing target - fail.
        err = "Presentation {}, missing target {}".format(
            new_presentation.presentation_id,
            new_presentation.target_id
        )
        return common_pb2.RpcResult(success=False, reason=err)

    def PresentationCreate(self, request, _context):
        return self.checked_presentation_update(directorpresentation_copy(request))

    def PresentationUpdate(self, request, _context):
        return self.checked_presentation_update(directorpresentation_copy(request))

    def PresentationDelete(self, request, _context):
        del self.presentation[request.presentation_id]
        return common_pb2.RpcResult(success=True)

    def PresentationList(self, request, _context):
        plist = []
        for _k, vol in self.presentation.items():
            plist.append(vol)
        response = director_pb2.DirectorPresentationList(presentations=plist)
        return response
