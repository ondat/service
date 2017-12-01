# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: rdbplugin.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import common_pb2 as common__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='rdbplugin.proto',
  package='storageos_rpc',
  syntax='proto3',
  serialized_pb=_b('\n\x0frdbplugin.proto\x12\rstorageos_rpc\x1a\x0c\x63ommon.proto\"B\n\x12RdbVolumeListQuery\x12,\n\nvolume_ids\x18\x01 \x03(\x0b\x32\x18.storageos_rpc.RdbVolume\"\x16\n\x14RdbVolumeCredentials\"\x15\n\x13RdbVolumeStatistics\"\x11\n\x0fRdbVolumeStatus\"\x82\x02\n\tRdbVolume\x12*\n\x02\x63\x63\x18\x01 \x01(\x0b\x32\x1e.storageos_rpc.DataplaneCommon\x12\x11\n\tvolume_id\x18\x02 \x01(\r\x12\x19\n\x11volume_size_bytes\x18\x03 \x01(\x04\x12\x38\n\x0b\x63redentials\x18\x04 \x01(\x0b\x32#.storageos_rpc.RdbVolumeCredentials\x12\x31\n\x05stats\x18\x05 \x01(\x0b\x32\".storageos_rpc.RdbVolumeStatistics\x12.\n\x06status\x18\x06 \x01(\x0b\x32\x1e.storageos_rpc.RdbVolumeStatus\":\n\rRdbVolumeList\x12)\n\x07volumes\x18\x01 \x03(\x0b\x32\x18.storageos_rpc.RdbVolume2\xae\x02\n\tRdbPlugin\x12\x44\n\x0cVolumeCreate\x12\x18.storageos_rpc.RdbVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12\x44\n\x0cVolumeUpdate\x12\x18.storageos_rpc.RdbVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12\x44\n\x0cVolumeDelete\x12\x18.storageos_rpc.RdbVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12O\n\nVolumeList\x12!.storageos_rpc.RdbVolumeListQuery\x1a\x1c.storageos_rpc.RdbVolumeList\"\x00\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_RDBVOLUMELISTQUERY = _descriptor.Descriptor(
  name='RdbVolumeListQuery',
  full_name='storageos_rpc.RdbVolumeListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volume_ids', full_name='storageos_rpc.RdbVolumeListQuery.volume_ids', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=48,
  serialized_end=114,
)


_RDBVOLUMECREDENTIALS = _descriptor.Descriptor(
  name='RdbVolumeCredentials',
  full_name='storageos_rpc.RdbVolumeCredentials',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=116,
  serialized_end=138,
)


_RDBVOLUMESTATISTICS = _descriptor.Descriptor(
  name='RdbVolumeStatistics',
  full_name='storageos_rpc.RdbVolumeStatistics',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=140,
  serialized_end=161,
)


_RDBVOLUMESTATUS = _descriptor.Descriptor(
  name='RdbVolumeStatus',
  full_name='storageos_rpc.RdbVolumeStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=163,
  serialized_end=180,
)


_RDBVOLUME = _descriptor.Descriptor(
  name='RdbVolume',
  full_name='storageos_rpc.RdbVolume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='storageos_rpc.RdbVolume.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_id', full_name='storageos_rpc.RdbVolume.volume_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_size_bytes', full_name='storageos_rpc.RdbVolume.volume_size_bytes', index=2,
      number=3, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='credentials', full_name='storageos_rpc.RdbVolume.credentials', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stats', full_name='storageos_rpc.RdbVolume.stats', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='storageos_rpc.RdbVolume.status', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=183,
  serialized_end=441,
)


_RDBVOLUMELIST = _descriptor.Descriptor(
  name='RdbVolumeList',
  full_name='storageos_rpc.RdbVolumeList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volumes', full_name='storageos_rpc.RdbVolumeList.volumes', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=443,
  serialized_end=501,
)

_RDBVOLUMELISTQUERY.fields_by_name['volume_ids'].message_type = _RDBVOLUME
_RDBVOLUME.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMON
_RDBVOLUME.fields_by_name['credentials'].message_type = _RDBVOLUMECREDENTIALS
_RDBVOLUME.fields_by_name['stats'].message_type = _RDBVOLUMESTATISTICS
_RDBVOLUME.fields_by_name['status'].message_type = _RDBVOLUMESTATUS
_RDBVOLUMELIST.fields_by_name['volumes'].message_type = _RDBVOLUME
DESCRIPTOR.message_types_by_name['RdbVolumeListQuery'] = _RDBVOLUMELISTQUERY
DESCRIPTOR.message_types_by_name['RdbVolumeCredentials'] = _RDBVOLUMECREDENTIALS
DESCRIPTOR.message_types_by_name['RdbVolumeStatistics'] = _RDBVOLUMESTATISTICS
DESCRIPTOR.message_types_by_name['RdbVolumeStatus'] = _RDBVOLUMESTATUS
DESCRIPTOR.message_types_by_name['RdbVolume'] = _RDBVOLUME
DESCRIPTOR.message_types_by_name['RdbVolumeList'] = _RDBVOLUMELIST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

RdbVolumeListQuery = _reflection.GeneratedProtocolMessageType('RdbVolumeListQuery', (_message.Message,), dict(
  DESCRIPTOR = _RDBVOLUMELISTQUERY,
  __module__ = 'rdbplugin_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.RdbVolumeListQuery)
  ))
_sym_db.RegisterMessage(RdbVolumeListQuery)

RdbVolumeCredentials = _reflection.GeneratedProtocolMessageType('RdbVolumeCredentials', (_message.Message,), dict(
  DESCRIPTOR = _RDBVOLUMECREDENTIALS,
  __module__ = 'rdbplugin_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.RdbVolumeCredentials)
  ))
_sym_db.RegisterMessage(RdbVolumeCredentials)

RdbVolumeStatistics = _reflection.GeneratedProtocolMessageType('RdbVolumeStatistics', (_message.Message,), dict(
  DESCRIPTOR = _RDBVOLUMESTATISTICS,
  __module__ = 'rdbplugin_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.RdbVolumeStatistics)
  ))
_sym_db.RegisterMessage(RdbVolumeStatistics)

RdbVolumeStatus = _reflection.GeneratedProtocolMessageType('RdbVolumeStatus', (_message.Message,), dict(
  DESCRIPTOR = _RDBVOLUMESTATUS,
  __module__ = 'rdbplugin_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.RdbVolumeStatus)
  ))
_sym_db.RegisterMessage(RdbVolumeStatus)

RdbVolume = _reflection.GeneratedProtocolMessageType('RdbVolume', (_message.Message,), dict(
  DESCRIPTOR = _RDBVOLUME,
  __module__ = 'rdbplugin_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.RdbVolume)
  ))
_sym_db.RegisterMessage(RdbVolume)

RdbVolumeList = _reflection.GeneratedProtocolMessageType('RdbVolumeList', (_message.Message,), dict(
  DESCRIPTOR = _RDBVOLUMELIST,
  __module__ = 'rdbplugin_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.RdbVolumeList)
  ))
_sym_db.RegisterMessage(RdbVolumeList)



_RDBPLUGIN = _descriptor.ServiceDescriptor(
  name='RdbPlugin',
  full_name='storageos_rpc.RdbPlugin',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=504,
  serialized_end=806,
  methods=[
  _descriptor.MethodDescriptor(
    name='VolumeCreate',
    full_name='storageos_rpc.RdbPlugin.VolumeCreate',
    index=0,
    containing_service=None,
    input_type=_RDBVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeUpdate',
    full_name='storageos_rpc.RdbPlugin.VolumeUpdate',
    index=1,
    containing_service=None,
    input_type=_RDBVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeDelete',
    full_name='storageos_rpc.RdbPlugin.VolumeDelete',
    index=2,
    containing_service=None,
    input_type=_RDBVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='storageos_rpc.RdbPlugin.VolumeList',
    index=3,
    containing_service=None,
    input_type=_RDBVOLUMELISTQUERY,
    output_type=_RDBVOLUMELIST,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_RDBPLUGIN)

DESCRIPTOR.services_by_name['RdbPlugin'] = _RDBPLUGIN

# @@protoc_insertion_point(module_scope)
