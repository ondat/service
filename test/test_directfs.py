#!/usr/bin/env python

# pylint: disable=C0111

from concurrent import futures

import logging
import unittest

import grpc

import directfs.v1.directfs_pb2 as directfs_pb2
import directfs.v1.directfs_pb2_grpc as directfs_pb2_grpc

import mock_directfs


# pylint: disable=C0103
log = logging.getLogger()


class TestClient(unittest.TestCase):

    url = '127.0.0.1:6666'
    version_string = 'TestClient'

    def setUp(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        servicer = mock_directfs.DirectfsClientServicer()
        servicer.version_string = TestClient.version_string
        directfs_pb2_grpc.add_DfsClientServicer_to_server(servicer, self.server)
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
        '''
        Check we can add and remove hosts.
        '''
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
        '''
        Check we can add and remove volumes.
        '''
        channel = grpc.insecure_channel(self.url)
        stub = directfs_pb2_grpc.DfsClientStub(channel)

        # Start empty.
        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

        # Add item without host present, should fail.
        vol = directfs_pb2.DfsVolume(volume_id=666, host_id=1000)
        reponse = stub.VolumeCreate(vol)
        self.assertEqual(False, reponse.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

        # Add the host.
        request = directfs_pb2.DfsHost(host_id=1000, hostname='rhubarb', port=17100)
        response = stub.ServerCreate(request)
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsHostListQuery(host_ids=[])
        response = stub.ServerList(request)
        self.assertEqual(1, len(response.hosts))
        self.assertEqual('rhubarb', response.hosts[0].hostname)

        # Add item with host present, should succeed.
        vol = directfs_pb2.DfsVolume(volume_id=666, host_id=1000)
        response = stub.VolumeCreate(vol)
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Update item to missing host, should fail.
        vol = directfs_pb2.DfsVolume(volume_id=666, host_id=1001)
        response = stub.VolumeUpdate(vol)
        self.assertEqual(False, response.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Update item to correct host (no change), should succeed.
        vol = directfs_pb2.DfsVolume(volume_id=666, host_id=1000)
        response = stub.VolumeUpdate(vol)
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Delete item.
        response = stub.VolumeDelete(directfs_pb2.DfsVolume(volume_id=666))
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))


class TestServer(unittest.TestCase):

    url = '127.0.0.1:6666'
    version_string = 'TestClient'

    def setUp(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        servicer = mock_directfs.DirectfsServerServicer()
        servicer.version_string = TestServer.version_string
        directfs_pb2_grpc.add_DfsServerServicer_to_server(servicer, self.server)
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

    def test_volume(self):
        '''
        Check we can add and remove volumes.
        '''
        channel = grpc.insecure_channel(self.url)
        stub = directfs_pb2_grpc.DfsServerStub(channel)

        # Start empty.
        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

        # Add item.
        vol = directfs_pb2.DfsVolume(volume_id=666)
        response = stub.VolumeCreate(vol)
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Update item.
        vol = directfs_pb2.DfsVolume(volume_id=666)
        response = stub.VolumeUpdate(vol)
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Delete item.
        response = stub.VolumeDelete(directfs_pb2.DfsVolume(volume_id=666))
        self.assertEqual(True, response.success)

        request = directfs_pb2.DfsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    unittest.main()
