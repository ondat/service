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
  package='directfs.v1',
  syntax='proto3',
  serialized_pb=_b('\n\x0e\x64irectfs.proto\x12\x0b\x64irectfs.v1\x1a\x0c\x63ommon.proto\"\x14\n\x12\x44\x66sHostCredentials\"\x98\x01\n\x07\x44\x66sHost\x12&\n\x02\x63\x63\x18\x01 \x01(\x0b\x32\x1a.common.v1.DataplaneCommon\x12\x0f\n\x07host_id\x18\x02 \x01(\r\x12\x10\n\x08hostname\x18\x03 \x01(\t\x12\x0c\n\x04port\x18\x04 \x01(\r\x12\x34\n\x0b\x63redentials\x18\x05 \x01(\x0b\x32\x1f.directfs.v1.DfsHostCredentials\"2\n\x0b\x44\x66sHostList\x12#\n\x05hosts\x18\x01 \x03(\x0b\x32\x14.directfs.v1.DfsHost\":\n\x10\x44\x66sHostListQuery\x12&\n\x08host_ids\x18\x01 \x03(\x0b\x32\x14.directfs.v1.DfsHost\"\x16\n\x14\x44\x66sVolumeCredentials\"\x15\n\x13\x44\x66sVolumeStatistics\"\xb5\x02\n\x0f\x44\x66sVolumeStatus\x12\x43\n\nconn_state\x18\x01 \x01(\x0e\x32/.directfs.v1.DfsVolumeStatus.DfsConnectionState\x12\x11\n\tpeer_name\x18\x02 \x01(\t\x12>\n\x07peer_af\x18\x03 \x01(\x0e\x32-.directfs.v1.DfsVolumeStatus.DfsAddressFamily\"b\n\x12\x44\x66sConnectionState\x12\x08\n\x04NONE\x10\x00\x12\x0e\n\nCONNECTING\x10\x01\x12\r\n\tCONNECTED\x10\x02\x12\x11\n\rDISCONNECTING\x10\x03\x12\x10\n\x0c\x44ISCONNECTED\x10\x04\"&\n\x10\x44\x66sAddressFamily\x12\x08\n\x04IPV4\x10\x00\x12\x08\n\x04IPV6\x10\x01\"\xee\x01\n\tDfsVolume\x12&\n\x02\x63\x63\x18\x01 \x01(\x0b\x32\x1a.common.v1.DataplaneCommon\x12\x11\n\tvolume_id\x18\x02 \x01(\r\x12\x0f\n\x07host_id\x18\x03 \x01(\r\x12\x36\n\x0b\x63redentials\x18\x04 \x01(\x0b\x32!.directfs.v1.DfsVolumeCredentials\x12/\n\x05stats\x18\x05 \x01(\x0b\x32 .directfs.v1.DfsVolumeStatistics\x12,\n\x06status\x18\x06 \x01(\x0b\x32\x1c.directfs.v1.DfsVolumeStatus\"<\n\rDfsVolumeList\x12+\n\x07volumes\x18\x01 \x03(\x0b\x32\x1a.directfs.v1.DfsVolumeList\"@\n\x12\x44\x66sVolumeListQuery\x12*\n\nvolume_ids\x18\x01 \x03(\x0b\x32\x16.directfs.v1.DfsVolume2\x9b\x04\n\tDfsClient\x12<\n\x0cServerCreate\x12\x14.directfs.v1.DfsHost\x1a\x14.common.v1.RpcResult\"\x00\x12<\n\x0cServerUpdate\x12\x14.directfs.v1.DfsHost\x1a\x14.common.v1.RpcResult\"\x00\x12<\n\x0cServerDelete\x12\x14.directfs.v1.DfsHost\x1a\x14.common.v1.RpcResult\"\x00\x12G\n\nServerList\x12\x1d.directfs.v1.DfsHostListQuery\x1a\x18.directfs.v1.DfsHostList\"\x00\x12>\n\x0cVolumeCreate\x12\x16.directfs.v1.DfsVolume\x1a\x14.common.v1.RpcResult\"\x00\x12>\n\x0cVolumeUpdate\x12\x16.directfs.v1.DfsVolume\x1a\x14.common.v1.RpcResult\"\x00\x12>\n\x0cVolumeDelete\x12\x16.directfs.v1.DfsVolume\x1a\x14.common.v1.RpcResult\"\x00\x12K\n\nVolumeList\x12\x1f.directfs.v1.DfsVolumeListQuery\x1a\x1a.directfs.v1.DfsVolumeList\"\x00\x32\x96\x02\n\tDfsServer\x12>\n\x0cVolumeCreate\x12\x16.directfs.v1.DfsVolume\x1a\x14.common.v1.RpcResult\"\x00\x12>\n\x0cVolumeUpdate\x12\x16.directfs.v1.DfsVolume\x1a\x14.common.v1.RpcResult\"\x00\x12>\n\x0cVolumeDelete\x12\x16.directfs.v1.DfsVolume\x1a\x14.common.v1.RpcResult\"\x00\x12I\n\nVolumeList\x12\x1f.directfs.v1.DfsVolumeListQuery\x1a\x16.directfs.v1.DfsVolume\"\x00\x30\x01\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])



_DFSVOLUMESTATUS_DFSCONNECTIONSTATE = _descriptor.EnumDescriptor(
  name='DfsConnectionState',
  full_name='directfs.v1.DfsVolumeStatus.DfsConnectionState',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONNECTING', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONNECTED', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DISCONNECTING', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DISCONNECTED', index=4, number=4,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=553,
  serialized_end=651,
)
_sym_db.RegisterEnumDescriptor(_DFSVOLUMESTATUS_DFSCONNECTIONSTATE)

_DFSVOLUMESTATUS_DFSADDRESSFAMILY = _descriptor.EnumDescriptor(
  name='DfsAddressFamily',
  full_name='directfs.v1.DfsVolumeStatus.DfsAddressFamily',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='IPV4', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='IPV6', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=653,
  serialized_end=691,
)
_sym_db.RegisterEnumDescriptor(_DFSVOLUMESTATUS_DFSADDRESSFAMILY)


_DFSHOSTCREDENTIALS = _descriptor.Descriptor(
  name='DfsHostCredentials',
  full_name='directfs.v1.DfsHostCredentials',
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
  serialized_start=45,
  serialized_end=65,
)


_DFSHOST = _descriptor.Descriptor(
  name='DfsHost',
  full_name='directfs.v1.DfsHost',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='directfs.v1.DfsHost.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='host_id', full_name='directfs.v1.DfsHost.host_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='hostname', full_name='directfs.v1.DfsHost.hostname', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='port', full_name='directfs.v1.DfsHost.port', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='credentials', full_name='directfs.v1.DfsHost.credentials', index=4,
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
  serialized_start=68,
  serialized_end=220,
)


_DFSHOSTLIST = _descriptor.Descriptor(
  name='DfsHostList',
  full_name='directfs.v1.DfsHostList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='hosts', full_name='directfs.v1.DfsHostList.hosts', index=0,
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
  serialized_start=222,
  serialized_end=272,
)


_DFSHOSTLISTQUERY = _descriptor.Descriptor(
  name='DfsHostListQuery',
  full_name='directfs.v1.DfsHostListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='host_ids', full_name='directfs.v1.DfsHostListQuery.host_ids', index=0,
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
  serialized_start=274,
  serialized_end=332,
)


_DFSVOLUMECREDENTIALS = _descriptor.Descriptor(
  name='DfsVolumeCredentials',
  full_name='directfs.v1.DfsVolumeCredentials',
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
  serialized_start=334,
  serialized_end=356,
)


_DFSVOLUMESTATISTICS = _descriptor.Descriptor(
  name='DfsVolumeStatistics',
  full_name='directfs.v1.DfsVolumeStatistics',
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
  serialized_start=358,
  serialized_end=379,
)


_DFSVOLUMESTATUS = _descriptor.Descriptor(
  name='DfsVolumeStatus',
  full_name='directfs.v1.DfsVolumeStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='conn_state', full_name='directfs.v1.DfsVolumeStatus.conn_state', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='peer_name', full_name='directfs.v1.DfsVolumeStatus.peer_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='peer_af', full_name='directfs.v1.DfsVolumeStatus.peer_af', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _DFSVOLUMESTATUS_DFSCONNECTIONSTATE,
    _DFSVOLUMESTATUS_DFSADDRESSFAMILY,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=382,
  serialized_end=691,
)


_DFSVOLUME = _descriptor.Descriptor(
  name='DfsVolume',
  full_name='directfs.v1.DfsVolume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='directfs.v1.DfsVolume.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_id', full_name='directfs.v1.DfsVolume.volume_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='host_id', full_name='directfs.v1.DfsVolume.host_id', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='credentials', full_name='directfs.v1.DfsVolume.credentials', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stats', full_name='directfs.v1.DfsVolume.stats', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='directfs.v1.DfsVolume.status', index=5,
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
  serialized_start=694,
  serialized_end=932,
)


_DFSVOLUMELIST = _descriptor.Descriptor(
  name='DfsVolumeList',
  full_name='directfs.v1.DfsVolumeList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volumes', full_name='directfs.v1.DfsVolumeList.volumes', index=0,
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
  serialized_start=934,
  serialized_end=994,
)


_DFSVOLUMELISTQUERY = _descriptor.Descriptor(
  name='DfsVolumeListQuery',
  full_name='directfs.v1.DfsVolumeListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volume_ids', full_name='directfs.v1.DfsVolumeListQuery.volume_ids', index=0,
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
  serialized_start=996,
  serialized_end=1060,
)

_DFSHOST.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMON
_DFSHOST.fields_by_name['credentials'].message_type = _DFSHOSTCREDENTIALS
_DFSHOSTLIST.fields_by_name['hosts'].message_type = _DFSHOST
_DFSHOSTLISTQUERY.fields_by_name['host_ids'].message_type = _DFSHOST
_DFSVOLUMESTATUS.fields_by_name['conn_state'].enum_type = _DFSVOLUMESTATUS_DFSCONNECTIONSTATE
_DFSVOLUMESTATUS.fields_by_name['peer_af'].enum_type = _DFSVOLUMESTATUS_DFSADDRESSFAMILY
_DFSVOLUMESTATUS_DFSCONNECTIONSTATE.containing_type = _DFSVOLUMESTATUS
_DFSVOLUMESTATUS_DFSADDRESSFAMILY.containing_type = _DFSVOLUMESTATUS
_DFSVOLUME.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMON
_DFSVOLUME.fields_by_name['credentials'].message_type = _DFSVOLUMECREDENTIALS
_DFSVOLUME.fields_by_name['stats'].message_type = _DFSVOLUMESTATISTICS
_DFSVOLUME.fields_by_name['status'].message_type = _DFSVOLUMESTATUS
_DFSVOLUMELIST.fields_by_name['volumes'].message_type = _DFSVOLUMELIST
_DFSVOLUMELISTQUERY.fields_by_name['volume_ids'].message_type = _DFSVOLUME
DESCRIPTOR.message_types_by_name['DfsHostCredentials'] = _DFSHOSTCREDENTIALS
DESCRIPTOR.message_types_by_name['DfsHost'] = _DFSHOST
DESCRIPTOR.message_types_by_name['DfsHostList'] = _DFSHOSTLIST
DESCRIPTOR.message_types_by_name['DfsHostListQuery'] = _DFSHOSTLISTQUERY
DESCRIPTOR.message_types_by_name['DfsVolumeCredentials'] = _DFSVOLUMECREDENTIALS
DESCRIPTOR.message_types_by_name['DfsVolumeStatistics'] = _DFSVOLUMESTATISTICS
DESCRIPTOR.message_types_by_name['DfsVolumeStatus'] = _DFSVOLUMESTATUS
DESCRIPTOR.message_types_by_name['DfsVolume'] = _DFSVOLUME
DESCRIPTOR.message_types_by_name['DfsVolumeList'] = _DFSVOLUMELIST
DESCRIPTOR.message_types_by_name['DfsVolumeListQuery'] = _DFSVOLUMELISTQUERY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DfsHostCredentials = _reflection.GeneratedProtocolMessageType('DfsHostCredentials', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOSTCREDENTIALS,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsHostCredentials)
  ))
_sym_db.RegisterMessage(DfsHostCredentials)

DfsHost = _reflection.GeneratedProtocolMessageType('DfsHost', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOST,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsHost)
  ))
_sym_db.RegisterMessage(DfsHost)

DfsHostList = _reflection.GeneratedProtocolMessageType('DfsHostList', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOSTLIST,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsHostList)
  ))
_sym_db.RegisterMessage(DfsHostList)

DfsHostListQuery = _reflection.GeneratedProtocolMessageType('DfsHostListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DFSHOSTLISTQUERY,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsHostListQuery)
  ))
_sym_db.RegisterMessage(DfsHostListQuery)

DfsVolumeCredentials = _reflection.GeneratedProtocolMessageType('DfsVolumeCredentials', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMECREDENTIALS,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsVolumeCredentials)
  ))
_sym_db.RegisterMessage(DfsVolumeCredentials)

DfsVolumeStatistics = _reflection.GeneratedProtocolMessageType('DfsVolumeStatistics', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMESTATISTICS,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsVolumeStatistics)
  ))
_sym_db.RegisterMessage(DfsVolumeStatistics)

DfsVolumeStatus = _reflection.GeneratedProtocolMessageType('DfsVolumeStatus', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMESTATUS,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsVolumeStatus)
  ))
_sym_db.RegisterMessage(DfsVolumeStatus)

DfsVolume = _reflection.GeneratedProtocolMessageType('DfsVolume', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUME,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsVolume)
  ))
_sym_db.RegisterMessage(DfsVolume)

DfsVolumeList = _reflection.GeneratedProtocolMessageType('DfsVolumeList', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMELIST,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsVolumeList)
  ))
_sym_db.RegisterMessage(DfsVolumeList)

DfsVolumeListQuery = _reflection.GeneratedProtocolMessageType('DfsVolumeListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DFSVOLUMELISTQUERY,
  __module__ = 'directfs_pb2'
  # @@protoc_insertion_point(class_scope:directfs.v1.DfsVolumeListQuery)
  ))
_sym_db.RegisterMessage(DfsVolumeListQuery)



_DFSCLIENT = _descriptor.ServiceDescriptor(
  name='DfsClient',
  full_name='directfs.v1.DfsClient',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=1063,
  serialized_end=1602,
  methods=[
  _descriptor.MethodDescriptor(
    name='ServerCreate',
    full_name='directfs.v1.DfsClient.ServerCreate',
    index=0,
    containing_service=None,
    input_type=_DFSHOST,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ServerUpdate',
    full_name='directfs.v1.DfsClient.ServerUpdate',
    index=1,
    containing_service=None,
    input_type=_DFSHOST,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ServerDelete',
    full_name='directfs.v1.DfsClient.ServerDelete',
    index=2,
    containing_service=None,
    input_type=_DFSHOST,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ServerList',
    full_name='directfs.v1.DfsClient.ServerList',
    index=3,
    containing_service=None,
    input_type=_DFSHOSTLISTQUERY,
    output_type=_DFSHOSTLIST,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeCreate',
    full_name='directfs.v1.DfsClient.VolumeCreate',
    index=4,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeUpdate',
    full_name='directfs.v1.DfsClient.VolumeUpdate',
    index=5,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeDelete',
    full_name='directfs.v1.DfsClient.VolumeDelete',
    index=6,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='directfs.v1.DfsClient.VolumeList',
    index=7,
    containing_service=None,
    input_type=_DFSVOLUMELISTQUERY,
    output_type=_DFSVOLUMELIST,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DFSCLIENT)

DESCRIPTOR.services_by_name['DfsClient'] = _DFSCLIENT


_DFSSERVER = _descriptor.ServiceDescriptor(
  name='DfsServer',
  full_name='directfs.v1.DfsServer',
  file=DESCRIPTOR,
  index=1,
  options=None,
  serialized_start=1605,
  serialized_end=1883,
  methods=[
  _descriptor.MethodDescriptor(
    name='VolumeCreate',
    full_name='directfs.v1.DfsServer.VolumeCreate',
    index=0,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeUpdate',
    full_name='directfs.v1.DfsServer.VolumeUpdate',
    index=1,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeDelete',
    full_name='directfs.v1.DfsServer.VolumeDelete',
    index=2,
    containing_service=None,
    input_type=_DFSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='directfs.v1.DfsServer.VolumeList',
    index=3,
    containing_service=None,
    input_type=_DFSVOLUMELISTQUERY,
    output_type=_DFSVOLUME,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DFSSERVER)

DESCRIPTOR.services_by_name['DfsServer'] = _DFSSERVER

# @@protoc_insertion_point(module_scope)
