# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: directfs.proto

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
  name='directfs.proto',
  package='storageos_rpc',
  syntax='proto3',
  serialized_pb=_b('\n\x0e\x64irectfs.proto\x12\rstorageos_rpc\x1a\x0c\x63ommon.proto\"\x14\n\x12\x44\x66sHostCredentials\"\xa4\x01\n\x07\x44\x66sHost\x12\x30\n\x02\x63\x63\x18\x01 \x01(\x0b\x32$.storageos_rpc.DataplaneCommonConfig\x12\x0f\n\x07host_id\x18\x02 \x01(\r\x12\x10\n\x08hostname\x18\x03 \x01(\t\x12\x0c\n\x04port\x18\x04 \x01(\r\x12\x36\n\x0b\x63redentials\x18\x05 \x01(\x0b\x32!.storageos_rpc.DfsHostCredentials\"4\n\x0b\x44\x66sHostList\x12%\n\x05hosts\x18\x01 \x03(\x0b\x32\x16.storageos_rpc.DfsHost\"<\n\x10\x44\x66sHostListQuery\x12(\n\x08host_ids\x18\x01 \x03(\x0b\x32\x16.storageos_rpc.DfsHost\"\x16\n\x14\x44\x66sVolumeCredentials\"\x10\n\x0e\x44\x66sVolumeStats\"\xc9\x01\n\tDfsVolume\x12\x30\n\x02\x63\x63\x18\x01 \x01(\x0b\x32$.storageos_rpc.DataplaneCommonConfig\x12\x11\n\tvolume_id\x18\x02 \x01(\r\x12\x0f\n\x07host_id\x18\x03 \x01(\r\x12\x38\n\x0b\x63redentials\x18\x04 \x01(\x0b\x32#.storageos_rpc.DfsVolumeCredentials\x12,\n\x05stats\x18\x05 \x01(\x0b\x32\x1d.storageos_rpc.DfsVolumeStats\">\n\rDfsVolumeList\x12-\n\x07volumes\x18\x01 \x03(\x0b\x32\x1c.storageos_rpc.DfsVolumeList\"B\n\x12\x44\x66sVolumeListQuery\x12,\n\nvolume_ids\x18\x01 \x03(\x0b\x32\x18.storageos_rpc.DfsVolume2\xc8\x03\n\x14\x44irectfsClientConfig\x12\x42\n\x0cServerCreate\x12\x16.storageos_rpc.DfsHost\x1a\x18.storageos_rpc.RpcResult\"\x00\x12\x42\n\x0cServerDelete\x12\x16.storageos_rpc.DfsHost\x1a\x18.storageos_rpc.RpcResult\"\x00\x12K\n\nServerList\x12\x1f.storageos_rpc.DfsHostListQuery\x1a\x1a.storageos_rpc.DfsHostList\"\x00\x12\x44\n\x0cVolumeCreate\x12\x18.storageos_rpc.DfsVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12\x44\n\x0cVolumeDelete\x12\x18.storageos_rpc.DfsVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12O\n\nVolumeList\x12!.storageos_rpc.DfsVolumeListQuery\x1a\x1c.storageos_rpc.DfsVolumeList\"\x00\x32\xf1\x01\n\x14\x44irectfsServerConfig\x12\x44\n\x0cVolumeCreate\x12\x18.storageos_rpc.DfsVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12\x44\n\x0cVolumeDelete\x12\x18.storageos_rpc.DfsVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12M\n\nVolumeList\x12!.storageos_rpc.DfsVolumeListQuery\x1a\x18.storageos_rpc.DfsVolume\"\x00\x30\x01\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_DFSHOSTCREDENTIALS = _descriptor.Descriptor(
  name='DfsHostCredentials',
  full_name='storageos_rpc.DfsHostCredentials',
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
  serialized_start=47,
  serialized_end=67,
)


_DFSHOST = _descriptor.Descriptor(
  name='DfsHost',
  full_name='storageos_rpc.DfsHost',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='storageos_rpc.DfsHost.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='host_id', full_name='storageos_rpc.DfsHost.host_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='hostname', full_name='storageos_rpc.DfsHost.hostname', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='port', full_name='storageos_rpc.DfsHost.port', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='credentials', full_name='storageos_rpc.DfsHost.credentials', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=70,
  serialized_end=234,
)


_DFSHOSTLIST = _descriptor.Descriptor(
  name='DfsHostList',
  full_name='storageos_rpc.DfsHostList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hosts', full_name='storageos_rpc.DfsHostList.hosts', index=0,
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
  serialized_start=236,
  serialized_end=288,
)


_DFSHOSTLISTQUERY = _descriptor.Descriptor(
  name='DfsHostListQuery',
  full_name='storageos_rpc.DfsHostListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='host_ids', full_name='storageos_rpc.DfsHostListQuery.host_ids', index=0,
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
  serialized_start=290,
  serialized_end=350,
)


_DFSVOLUMECREDENTIALS = _descriptor.Descriptor(
  name='DfsVolumeCredentials',
  full_name='storageos_rpc.DfsVolumeCredentials',
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
  serialized_start=352,
  serialized_end=374,
)


_DFSVOLUMESTATS = _descriptor.Descriptor(
  name='DfsVolumeStats',
  full_name='storageos_rpc.DfsVolumeStats',
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
  serialized_start=376,
  serialized_end=392,
)


_DFSVOLUME = _descriptor.Descriptor(
  name='DfsVolume',
  full_name='storageos_rpc.DfsVolume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='storageos_rpc.DfsVolume.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_id', full_name='storageos_rpc.DfsVolume.volume_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='host_id', full_name='storageos_rpc.DfsVolume.host_id', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='credentials', full_name='storageos_rpc.DfsVolume.credentials', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stats', full_name='storageos_rpc.DfsVolume.stats', index=4,
      number=5, type=11, cpp_type=10, label=1,
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
  serialized_start=395,
  serialized_end=596,
)


_DFSVOLUMELIST = _descriptor.Descriptor(
  name='DfsVolumeList',
  full_name='storageos_rpc.DfsVolumeList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volumes', full_name='storageos_rpc.DfsVolumeList.volumes', index=0,
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
  serialized_start=598,
  serialized_end=660,
)


_DFSVOLUMELISTQUERY = _descriptor.Descriptor(
  name='DfsVolumeListQuery',
  full_name='storageos_rpc.DfsVolumeListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volume_ids', full_name='storageos_rpc.DfsVolumeListQuery.volume_ids', index=0,
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
  serialized_start=662,
  serialized_end=728,
)

_DFSHOST.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMONCONFIG
_DFSHOST.fields_by_name['credentials'].message_type = _DFSHOSTCREDENTIALS
_DFSHOSTLIST.fields_by_name['hosts'].message_type = _DFSHOST
_DFSHOSTLISTQUERY.fields_by_name['host_ids'].message_type = _DFSHOST
_DFSVOLUME.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMONCONFIG
_DFSVOLUME.fields_by_name['credentials'].message_type = _DFSVOLUMECREDENTIALS
_DFSVOLUME.fields_by_name['stats'].message_type = _DFSVOLUMESTATS
_DFSVOLUMELIST.fields_by_name['volumes'].message_type = _DFSVOLUMELIST
_DFSVOLUMELISTQUERY.fields_by_name['volume_ids'].message_type = _DFSVOLUME
DESCRIPTOR.message_types_by_name['DfsHostCredentials'] = _DFSHOSTCREDENTIALS
DESCRIPTOR.message_types_by_name['DfsHost'] = _DFSHOST
DESCRIPTOR.message_types_by_name['DfsHostList'] = _DFSHOSTLIST
DESCRIPTOR.message_types_by_name['DfsHostListQuery'] = _DFSHOSTLISTQUERY
DESCRIPTOR.message_types_by_name['DfsVolumeCredentials'] = _DFSVOLUMECREDENTIALS
DESCRIPTOR.message_types_by_name['DfsVolumeStats'] = _DFSVOLUMESTATS
DESCRIPTOR.message_types_by_name['DfsVolume'] = _DFSVOLUME
DESCRIPTOR.message_types_by_name['DfsVolumeList'] = _DFSVOLUMELIST
DESCRIPTOR.message_types_by_name['DfsVolumeListQuery'] = _DFSVOLUMELISTQUERY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DfsHostCredentials = _reflection.GeneratedProtocolMessageType('DfsHostCredentials', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOSTCREDENTIALS,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsHostCredentials)
  ))
_sym_db.RegisterMessage(DfsHostCredentials)

DfsHost = _reflection.GeneratedProtocolMessageType('DfsHost', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOST,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsHost)
  ))
_sym_db.RegisterMessage(DfsHost)

DfsHostList = _reflection.GeneratedProtocolMessageType('DfsHostList', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOSTLIST,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsHostList)
  ))
_sym_db.RegisterMessage(DfsHostList)

DfsHostListQuery = _reflection.GeneratedProtocolMessageType('DfsHostListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOSTLISTQUERY,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsHostListQuery)
  ))
_sym_db.RegisterMessage(DfsHostListQuery)

DfsVolumeCredentials = _reflection.GeneratedProtocolMessageType('DfsVolumeCredentials', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMECREDENTIALS,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsVolumeCredentials)
  ))
_sym_db.RegisterMessage(DfsVolumeCredentials)

DfsVolumeStats = _reflection.GeneratedProtocolMessageType('DfsVolumeStats', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMESTATS,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsVolumeStats)
  ))
_sym_db.RegisterMessage(DfsVolumeStats)

DfsVolume = _reflection.GeneratedProtocolMessageType('DfsVolume', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUME,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsVolume)
  ))
_sym_db.RegisterMessage(DfsVolume)

DfsVolumeList = _reflection.GeneratedProtocolMessageType('DfsVolumeList', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMELIST,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsVolumeList)
  ))
_sym_db.RegisterMessage(DfsVolumeList)

DfsVolumeListQuery = _reflection.GeneratedProtocolMessageType('DfsVolumeListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMELISTQUERY,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DfsVolumeListQuery)
  ))
_sym_db.RegisterMessage(DfsVolumeListQuery)



_DIRECTFSCLIENTCONFIG = _descriptor.ServiceDescriptor(
  name='DirectfsClientConfig',
  full_name='storageos_rpc.DirectfsClientConfig',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=731,
  serialized_end=1187,
  methods=[
  _descriptor.MethodDescriptor(
    name='ServerCreate',
    full_name='storageos_rpc.DirectfsClientConfig.ServerCreate',
    index=0,
    containing_service=None,
    input_type=_DFSHOST,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ServerDelete',
    full_name='storageos_rpc.DirectfsClientConfig.ServerDelete',
    index=1,
    containing_service=None,
    input_type=_DFSHOST,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ServerList',
    full_name='storageos_rpc.DirectfsClientConfig.ServerList',
    index=2,
    containing_service=None,
    input_type=_DFSHOSTLISTQUERY,
    output_type=_DFSHOSTLIST,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeCreate',
    full_name='storageos_rpc.DirectfsClientConfig.VolumeCreate',
    index=3,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeDelete',
    full_name='storageos_rpc.DirectfsClientConfig.VolumeDelete',
    index=4,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='storageos_rpc.DirectfsClientConfig.VolumeList',
    index=5,
    containing_service=None,
    input_type=_DFSVOLUMELISTQUERY,
    output_type=_DFSVOLUMELIST,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DIRECTFSCLIENTCONFIG)

DESCRIPTOR.services_by_name['DirectfsClientConfig'] = _DIRECTFSCLIENTCONFIG


_DIRECTFSSERVERCONFIG = _descriptor.ServiceDescriptor(
  name='DirectfsServerConfig',
  full_name='storageos_rpc.DirectfsServerConfig',
  file=DESCRIPTOR,
  index=1,
  options=None,
  serialized_start=1190,
  serialized_end=1431,
  methods=[
  _descriptor.MethodDescriptor(
    name='VolumeCreate',
    full_name='storageos_rpc.DirectfsServerConfig.VolumeCreate',
    index=0,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeDelete',
    full_name='storageos_rpc.DirectfsServerConfig.VolumeDelete',
    index=1,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='storageos_rpc.DirectfsServerConfig.VolumeList',
    index=2,
    containing_service=None,
    input_type=_DFSVOLUMELISTQUERY,
    output_type=_DFSVOLUME,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DIRECTFSSERVERCONFIG)

DESCRIPTOR.services_by_name['DirectfsServerConfig'] = _DIRECTFSSERVERCONFIG

# @@protoc_insertion_point(module_scope)
