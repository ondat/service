#!/usr/bin/env python

# pylint: disable=C0111

from concurrent import futures

import logging
import unittest

import grpc

import rdbplugin.v1.rdbplugin_pb2 as rdbplugin_pb2
import rdbplugin.v1.rdbplugin_pb2_grpc as rdbplugin_pb2_grpc

import mock_rdbplugin

# pylint: disable=C0103
log = logging.getLogger()


class TestRdbplugin(unittest.TestCase):

    url = '127.0.0.1:6670'
    version_string = 'TestClient'

    def setUp(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        servicer = mock_rdbplugin.RdbPluginServicer()
        servicer.version_string = self.version_string
        rdbplugin_pb2_grpc.add_RdbPluginServicer_to_server(servicer,
                                                           self.server)
        self.server.add_insecure_port(self.url)
        log.debug('server start')
        self.server.start()

    def tearDown(self):
        log.debug('server stop')
        self.server.stop(0)
        self.server = None

    def test_connect(self):
        channel = grpc.insecure_channel(self.url)
        stub = rdbplugin_pb2_grpc.RdbPluginStub(channel)
        request = rdbplugin_pb2.RdbStatusRequest()
        response = stub.Status(request)
        self.assertEqual(self.version_string, response.version_info)
