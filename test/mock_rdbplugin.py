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


class RdbPluginServicer(rdbplugin_pb2_grpc.RdbPluginServicer):

    def __init__(self):
        self.volumes = {}
        self.version_string = 'notset'

    def Status(self, _request, _context):
        return rdbplugin_pb2.RdbStatus(version_info=self.version_string)
