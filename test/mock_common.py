#!/usr/bin/env python
#
# Common features of the dataplane gRPC API: Mock implementation.

"""
Support functions for the common.v1 gRPC/Protobuf services and messages.

Currently nothing to do.
"""

# pylint: disable=C0111

import logging

import common.v1.common_pb2 as common_pb2

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


def new_dataplanecommon_message():
    """Return a fresh common.v1.DataplaneCommon message."""
    return common_pb2.DataplaneCommon()


def set_common_label(cmsg, key, value):
    """
    Set a label in the DataplaneCommon message.

    Any previous value of the key will be overwritten.
    """
    cmsg.labels[key] = value


def get_common_label(cmsg, key):
    """
    Retrieve a label in the DataplaneCommon message.

    Returns the value of the label if set, or None if the label is not set.
    """
    if key in cmsg.labels:
        return cmsg.labels[key]
    return None
