#!/usr/bin/env python

# pylint: disable=C0111

from concurrent import futures

import logging
import unittest

import grpc

import director.v1.director_pb2 as director_pb2
import director.v1.director_pb2_grpc as director_pb2_grpc

import mock_director

# pylint: disable=C0103
log = logging.getLogger()


class TestDirector(unittest.TestCase):

    url = '127.0.0.1:6668'
    version_string = 'TestClient'

    def setUp(self):
        self.server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
        servicer = mock_director.DirectorServicer()
        servicer.version_string = self.version_string
        director_pb2_grpc.add_DirectorServicer_to_server(servicer,
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
        stub = director_pb2_grpc.DirectorStub(channel)
        request = director_pb2.DirectorStatusRequest()
        response = stub.Status(request)
        self.assertEqual(self.version_string, response.version_info)

    def test_volume(self):
        '''
        Check we can add and remove volumes.
        '''
        channel = grpc.insecure_channel(self.url)
        stub = director_pb2_grpc.DirectorStub(channel)

        # Start empty.
        request = director_pb2.DirectorVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

        # Add item.
        vol = director_pb2.DirectorVolume(volume_id=666)
        response = stub.VolumeCreate(vol)
        self.assertEqual(True, response.success)

        request = director_pb2.DirectorVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Update item.
        vol = director_pb2.DirectorVolume(volume_id=666)
        response = stub.VolumeUpdate(vol)
        self.assertEqual(True, response.success)

        request = director_pb2.DirectorVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Delete item.
        response = stub.VolumeDelete(director_pb2.DirectorVolume(volume_id=666))
        self.assertEqual(True, response.success)

        request = director_pb2.DirectorVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(0, len(response.volumes))

    def test_presentation(self):
        '''
        Check we can add and remove volumes.
        '''
        channel = grpc.insecure_channel(self.url)
        stub = director_pb2_grpc.DirectorStub(channel)

        # Start empty.
        request = director_pb2.DirectorPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(0, len(response.presentations))

        # Add item without volume target, should fail.
        vol = director_pb2.DirectorPresentation(presentation_id=1666, target_id=666)
        response = stub.PresentationCreate(vol)
        self.assertEqual(False, response.success)

        # Add the target volume.
        vol = director_pb2.DirectorVolume(volume_id=666)
        response = stub.VolumeCreate(vol)
        self.assertEqual(True, response.success)

        request = director_pb2.DirectorVolumeListQuery(volume_ids=[])
        response = stub.VolumeList(request)
        self.assertEqual(1, len(response.volumes))

        # Add item - should now succeed.
        vol = director_pb2.DirectorPresentation(presentation_id=1666, target_id=666)
        response = stub.PresentationCreate(vol)
        self.assertEqual(True, response.success)

        request = director_pb2.DirectorPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(1, len(response.presentations))

        # Update item to missing host, should fail.
        vol = director_pb2.DirectorPresentation(presentation_id=1666, target_id=667)
        response = stub.PresentationUpdate(vol)
        self.assertEqual(False, response.success)

        request = director_pb2.DirectorPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(1, len(response.presentations))

        # Update item to correct host (no change), should succeed.
        vol = director_pb2.DirectorPresentation(presentation_id=1666, target_id=666)
        response = stub.PresentationUpdate(vol)
        self.assertEqual(True, response.success)

        request = director_pb2.DirectorPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(1, len(response.presentations))

        # Delete item.
        response = stub.PresentationDelete(
            director_pb2.DirectorPresentation(presentation_id=1666, target_id=666))
        self.assertEqual(True, response.success)

        request = director_pb2.DirectorPresentationListQuery(presentation_ids=[])
        response = stub.PresentationList(request)
        self.assertEqual(0, len(response.presentations))
