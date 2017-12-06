#!/usr/bin/env python

# C0111 is 'missing module docstring', we don't care here.
# E1101 is 'Instance/Class x has no member y', which isn't necessarily true -
#   protobuf uses metaclass magic a lot and pylint doesn't always keep up.
# pylint: disable=C0111,E1101

from concurrent import futures

import logging
import unittest

import grpc

import filesystem.v1.filesystem_pb2 as filesystem_pb2
import filesystem.v1.filesystem_pb2_grpc as filesystem_pb2_grpc

import mock_filesystem


# pylint: disable=C0103
log = logging.getLogger()

ONE_GiB_IN_BYTES = 1024 * 1024 * 1024


class TestServer(unittest.TestCase):

    url = '127.0.0.1:6667'
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
        vol = filesystem_pb2.FsVolume(
            volume_id=666,
            node_type=filesystem_pb2.FsVolume.FILE,
            device_number=1,
            filename='meh',
            volume_size_bytes=10 * ONE_GiB_IN_BYTES
        )
        # Set a label.
        vol.cc.labels['foo'] = 'bar'

        response = stub.VolumeCreate(vol)
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))
        # Check the label.
        self.assertEqual('bar', response.volumes[0].cc.labels['foo'])

        # Update item.
        vol = filesystem_pb2.FsVolume(
            volume_id=666,
            node_type=filesystem_pb2.FsVolume.FILE,
            device_number=1,
            filename='meh',
            volume_size_bytes=100 * ONE_GiB_IN_BYTES
        )
        vol.cc.labels['foo'] = 'baz'
        response = stub.VolumeUpdate(vol)
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))
        # Check the label update.
        self.assertEqual('baz', response.volumes[0].cc.labels['foo'])

        # Delete item.
        response = stub.VolumeDelete(filesystem_pb2.FsVolume(volume_id=666))
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

    def test_presentation(self):
        '''
        Check we can add and remove volumes.
        '''
        channel = grpc.insecure_channel(self.url)
        stub = filesystem_pb2_grpc.FsStub(channel)

        # Start empty.
        request = filesystem_pb2.FsPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(0, len(response.presentations))

        # Add item without volume target, should fail.
        vol = filesystem_pb2.FsPresentation(presentation_id=1666, target_id=666)
        response = stub.PresentationCreate(vol)
        self.assertEqual(False, response.success)

        # Add the target volume.
        vol = filesystem_pb2.FsVolume(volume_id=666)
        response = stub.VolumeCreate(vol)
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Add item - should now succeed.
        vol = filesystem_pb2.FsPresentation(presentation_id=1666, target_id=666)
        response = stub.PresentationCreate(vol)
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(1, len(response.presentations))

        # Update item to missing host, should fail.
        vol = filesystem_pb2.FsPresentation(presentation_id=1666, target_id=667)
        response = stub.PresentationUpdate(vol)
        self.assertEqual(False, response.success)

        request = filesystem_pb2.FsPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(1, len(response.presentations))

        # Update item to correct host (no change), should succeed.
        vol = filesystem_pb2.FsPresentation(presentation_id=1666, target_id=666)
        response = stub.PresentationUpdate(vol)
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(1, len(response.presentations))

        # Delete item.
        response = stub.PresentationDelete(
            filesystem_pb2.FsPresentation(presentation_id=1666, target_id=666))
        self.assertEqual(True, response.success)

        request = filesystem_pb2.FsPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(0, len(response.presentations))


if __name__ == '__main__':
    logging.basicConfig(level=logging.DEBUG)
    unittest.main()
