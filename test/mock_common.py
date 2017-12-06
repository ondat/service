#!/usr/bin/env python
#
# Common features of the dataplane gRPC API: Mock implementation.

"""
Support functions for the common.v1 gRPC/Protobuf services and messages.

Currently nothing to do.
"""

# pylint: disable=C0111

import logging

# import common.v1.common_pb2 as common_pb2

# pylint: disable=C0103
log = logging.getLogger()


def copy_submessage(dstfield, src, name):
    """
    Copy a message field from another message.

    If message src has the field 'name', set dstfield to the value of src.name.
    Otherwise, do nothing.
    """
    if src.HasField(name):
        dstfield.CopyFrom(getattr(src, name))
