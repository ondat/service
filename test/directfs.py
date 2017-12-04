#!/usr/bin/env python

from concurrent import futures

import copy
import logging
import os
import pprint
import sys
import unittest

lib_path = os.path.abspath(os.path.join('..'))
sys.path.append(lib_path)
sys.path.append(os.path.join(lib_path, 'common/v1'))
sys.path.append(os.path.join(lib_path, 'directfs/v1'))

import grpc

import common.v1.common_pb2 as common_pb2
import directfs.v1.directfs_pb2 as directfs_pb2
import directfs.v1.directfs_pb2_grpc as directfs_pb2_grpc

log = logging.getLogger()


class TestClient(unittest.TestCase):

    url = '127.0.0.1:6666'
    version_string = 'TestClient'

    class DirectfsClientServicer(directfs_pb2_grpc.DfsClientServicer):

        servers = {}
        volumes = {}

        def Status(self, _request, _context):
            return directfs_pb2.DfsClientStatus(version_info=TestClient.version_string)

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
            for _k, v in self.servers.items():
                slist.append(v)
            resp = directfs_pb2.DfsHostList(hosts=slist)
            return resp

        def VolumeCreate(self, request, _context):
            # NOTE: Create a new DfsVolume object. Do not use the request.
            # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
            self.volumes[request.volume_id] = directfs_pb2.DfsVolume(
                volume_id=request.volume_id)

            return common_pb2.RpcResult(success=True)

        def VolumeUpdate(self, request, _context):
            # NOTE: Create a new DfsVolume object. Do not use the request.
            # NOTE: I don't know why CopyFrom()/MergeFrom() don't work, but they don't.
            self.volumes[request.volume_id] = directfs_pb2.DfsVolume(
                volume_id=request.volume_id)
            return common_pb2.RpcResult(success=True)

        def VolumeDelete(self, request, _context):
            del self.volumes[request.volume_id]
            return common_pb2.RpcResult(success=True)

        def VolumeList(self, request, _context):
            vlist = []
            for _k, v in self.volumes.items():
                vlist.append(v)
            resp = directfs_pb2.DfsVolumeList(volumes=vlist)
            return resp

        def get_servers(self):
            return self.servers

        def get_volume(self):
            return self.volumes

    def setUp(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        directfs_pb2_grpc.add_DfsClientServicer_to_server(
            TestClient.DirectfsClientServicer(), self.server)
        self.server.add_insecure_port(self.url)
        log.debug('server start')
        self.server.start()

    def tearDown(self):
        log.debug('server stop')
        self.server.stop(0)

    def test_connect(self):
        channel = grpc.insecure_channel(self.url)
        stub = directfs_pb2_grpc.DfsClientStub(channel)
        request = directfs_pb2.DfsClientStatusRequest()
        response = stub.Status(request)
        self.assertEqual(self.version_string, response.version_info)

    def test_server(self):
        channel = grpc.insecure_channel(self.url)
        stub = directfs_pb2_grpc.DfsClientStub(channel)

        # Start empty.
        request = directfs_pb2.DfsHostListQuery(host_ids=[])
        response = stub.ServerList(request)
        self.assertEqual(0, len(response.hosts))

        # Add item.
        request = directfs_pb2.DfsHost(host_id=1000, hostname='rhubarb', port=17100)
        response = stub.ServerCreate(request)
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsHostListQuery(host_ids=[])
        response = stub.ServerList(request)
        self.assertEqual(1, len(response.hosts))
        self.assertEqual('rhubarb', response.hosts[0].hostname)

        # Update item.
        request = directfs_pb2.DfsHost(host_id=1000, hostname='custard', port=17100)
        response = stub.ServerUpdate(request)
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsHostListQuery(host_ids=[])
        response = stub.ServerList(request)
        self.assertEqual(1, len(response.hosts))
        self.assertEqual('custard', response.hosts[0].hostname)

        # Delete item.
        request = directfs_pb2.DfsHost(host_id=1000, hostname='custard', port=17100)
        response = stub.ServerDelete(request)
        self.assertEqual(True, response.success)

        # End empty.
        request = directfs_pb2.DfsHostListQuery(host_ids=[])
        response = stub.ServerList(request)
        self.assertEqual(0, len(response.hosts))

    def test_volume(self):
        channel = grpc.insecure_channel(self.url)
        stub = directfs_pb2_grpc.DfsClientStub(channel)

        # Start empty.
        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

        # XXX WIP


class TestServer(unittest.TestCase):

    url = '127.0.0.1:6666'
    version_string = 'TestClient'

    class DirectfsServerServicer(directfs_pb2_grpc.DfsServerServicer):

        def Status(self, _request, _context):
            return directfs_pb2.DfsServerStatus(version_info=TestServer.version_string)

    def setUp(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        directfs_pb2_grpc.add_DfsServerServicer_to_server(
            TestServer.DirectfsServerServicer(), self.server)
        self.server.add_insecure_port(self.url)
        self.server.start()

    def tearDown(self):
        self.server.stop(0)

    def test_connect(self):
        channel = grpc.insecure_channel(self.url)
        stub = directfs_pb2_grpc.DfsServerStub(channel)
        request = directfs_pb2.DfsServerStatusRequest()
        response = stub.Status(request)
        self.assertEqual(self.version_string, response.version_info)


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    unittest.main()
