#!/usr/bin/env python

# pylint: disable=C0111
# flake8: noqa=D100

from concurrent import futures

import logging
import unittest

import grpc

import filesystem.v1.filesystem_pb2 as filesystem_pb2
import filesystem.v1.filesystem_pb2_grpc as filesystem_pb2_grpc

import mock_filesystem


# pylint: disable=C0103
log = logging.getLogger()


class TestServer(unittest.TestCase):

    url = '127.0.0.1:6666'
    version_string = 'TestClient'

    def setUp(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        servicer = mock_filesystem.FilesystemServicer()
        servicer.version_string = TestServer.version_string
        filesystem_pb2_grpc.add_FsServicer_to_server(servicer, self.server)
        self.server.add_insecure_port(self.url)
        self.server.start()

    def tearDown(self):
        self.server.stop(0)

    def test_connect(self):
        channel = grpc.insecure_channel(self.url)
        stub = filesystem_pb2_grpc.FsStub(channel)
        request = filesystem_pb2.FsStatusRequest()
        response = stub.Status(request)
        self.assertEqual(self.version_string, response.version_info)

    def test_volume(self):
        '''
        Check we can add and remove volumes.
        '''
        channel = grpc.insecure_channel(self.url)
        stub = filesystem_pb2_grpc.FsStub(channel)

        # Start empty.
        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

        # Add item.
        vol = filesystem_pb2.FsVolume(volume_id=666)
        response = stub.VolumeCreate(vol)
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Update item.
        vol = filesystem_pb2.FsVolume(volume_id=666)
        response = stub.VolumeUpdate(vol)
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Delete item.
        response = stub.VolumeDelete(filesystem_pb2.FsVolume(volume_id=666))
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    unittest.main()
