# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: director.proto

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
  name='director.proto',
  package='storageos_rpc',
  syntax='proto3',
  serialized_pb=_b('\n\x0e\x64irector.proto\x12\rstorageos_rpc\x1a\x0c\x63ommon.proto\"-\n\x17\x44irectorVolumeListQuery\x12\x12\n\nvolume_ids\x18\x01 \x03(\r\"\x15\n\x13\x44irectorVolumeStats\"\xd1\x01\n\x0e\x44irectorVolume\x12\x30\n\x02\x63\x63\x18\x01 \x01(\x0b\x32$.storageos_rpc.DataplaneCommonConfig\x12\x11\n\tvolume_id\x18\x02 \x01(\r\x12\x12\n\nwrite_pipe\x18\x03 \x01(\r\x12\x11\n\tread_pipe\x18\x04 \x01(\r\x12\x0b\n\x03qos\x18\x05 \x01(\x04\x12\x13\n\x0breplica_ids\x18\x06 \x03(\r\x12\x31\n\x05stats\x18\x07 \x01(\x0b\x32\".storageos_rpc.DirectorVolumeStats\"D\n\x12\x44irectorVolumeList\x12.\n\x07volumes\x18\x01 \x03(\x0b\x32\x1d.storageos_rpc.DirectorVolume\"-\n\x19\x44irectorRedirectListQuery\x12\x10\n\x08query_id\x18\x01 \x03(\r\"j\n\x10\x44irectorRedirect\x12\x30\n\x02\x63\x63\x18\x01 \x01(\x0b\x32$.storageos_rpc.DataplaneCommonConfig\x12\x11\n\tsource_id\x18\x02 \x01(\r\x12\x11\n\ttarget_id\x18\x03 \x01(\r\"J\n\x14\x44irectorRedirectList\x12\x32\n\tredirects\x18\x01 \x03(\x0b\x32\x1f.storageos_rpc.DirectorRedirect2\x90\x04\n\x0e\x44irectorConfig\x12L\n\x0fVolumeConfigure\x12\x1d.storageos_rpc.DirectorVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12N\n\x11VolumeUnconfigure\x12\x1d.storageos_rpc.DirectorVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12Y\n\nVolumeList\x12&.storageos_rpc.DirectorVolumeListQuery\x1a!.storageos_rpc.DirectorVolumeList\"\x00\x12P\n\x11RedirectConfigure\x12\x1f.storageos_rpc.DirectorRedirect\x1a\x18.storageos_rpc.RpcResult\"\x00\x12R\n\x13RedirectUnconfigure\x12\x1f.storageos_rpc.DirectorRedirect\x1a\x18.storageos_rpc.RpcResult\"\x00\x12_\n\x0cRedirectList\x12(.storageos_rpc.DirectorRedirectListQuery\x1a#.storageos_rpc.DirectorRedirectList\"\x00\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_DIRECTORVOLUMELISTQUERY = _descriptor.Descriptor(
  name='DirectorVolumeListQuery',
  full_name='storageos_rpc.DirectorVolumeListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volume_ids', full_name='storageos_rpc.DirectorVolumeListQuery.volume_ids', index=0,
      number=1, type=13, cpp_type=3, label=3,
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
  serialized_start=47,
  serialized_end=92,
)


_DIRECTORVOLUMESTATS = _descriptor.Descriptor(
  name='DirectorVolumeStats',
  full_name='storageos_rpc.DirectorVolumeStats',
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
  serialized_start=94,
  serialized_end=115,
)


_DIRECTORVOLUME = _descriptor.Descriptor(
  name='DirectorVolume',
  full_name='storageos_rpc.DirectorVolume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='storageos_rpc.DirectorVolume.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_id', full_name='storageos_rpc.DirectorVolume.volume_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='write_pipe', full_name='storageos_rpc.DirectorVolume.write_pipe', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='read_pipe', full_name='storageos_rpc.DirectorVolume.read_pipe', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='qos', full_name='storageos_rpc.DirectorVolume.qos', index=4,
      number=5, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='replica_ids', full_name='storageos_rpc.DirectorVolume.replica_ids', index=5,
      number=6, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stats', full_name='storageos_rpc.DirectorVolume.stats', index=6,
      number=7, type=11, cpp_type=10, label=1,
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
  serialized_start=118,
  serialized_end=327,
)


_DIRECTORVOLUMELIST = _descriptor.Descriptor(
  name='DirectorVolumeList',
  full_name='storageos_rpc.DirectorVolumeList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volumes', full_name='storageos_rpc.DirectorVolumeList.volumes', index=0,
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
  serialized_start=329,
  serialized_end=397,
)


_DIRECTORREDIRECTLISTQUERY = _descriptor.Descriptor(
  name='DirectorRedirectListQuery',
  full_name='storageos_rpc.DirectorRedirectListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_id', full_name='storageos_rpc.DirectorRedirectListQuery.query_id', index=0,
      number=1, type=13, cpp_type=3, label=3,
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
  serialized_start=399,
  serialized_end=444,
)


_DIRECTORREDIRECT = _descriptor.Descriptor(
  name='DirectorRedirect',
  full_name='storageos_rpc.DirectorRedirect',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='storageos_rpc.DirectorRedirect.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='source_id', full_name='storageos_rpc.DirectorRedirect.source_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='target_id', full_name='storageos_rpc.DirectorRedirect.target_id', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=446,
  serialized_end=552,
)


_DIRECTORREDIRECTLIST = _descriptor.Descriptor(
  name='DirectorRedirectList',
  full_name='storageos_rpc.DirectorRedirectList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='redirects', full_name='storageos_rpc.DirectorRedirectList.redirects', index=0,
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
  serialized_start=554,
  serialized_end=628,
)

_DIRECTORVOLUME.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMONCONFIG
_DIRECTORVOLUME.fields_by_name['stats'].message_type = _DIRECTORVOLUMESTATS
_DIRECTORVOLUMELIST.fields_by_name['volumes'].message_type = _DIRECTORVOLUME
_DIRECTORREDIRECT.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMONCONFIG
_DIRECTORREDIRECTLIST.fields_by_name['redirects'].message_type = _DIRECTORREDIRECT
DESCRIPTOR.message_types_by_name['DirectorVolumeListQuery'] = _DIRECTORVOLUMELISTQUERY
DESCRIPTOR.message_types_by_name['DirectorVolumeStats'] = _DIRECTORVOLUMESTATS
DESCRIPTOR.message_types_by_name['DirectorVolume'] = _DIRECTORVOLUME
DESCRIPTOR.message_types_by_name['DirectorVolumeList'] = _DIRECTORVOLUMELIST
DESCRIPTOR.message_types_by_name['DirectorRedirectListQuery'] = _DIRECTORREDIRECTLISTQUERY
DESCRIPTOR.message_types_by_name['DirectorRedirect'] = _DIRECTORREDIRECT
DESCRIPTOR.message_types_by_name['DirectorRedirectList'] = _DIRECTORREDIRECTLIST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DirectorVolumeListQuery = _reflection.GeneratedProtocolMessageType('DirectorVolumeListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUMELISTQUERY,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DirectorVolumeListQuery)
  ))
_sym_db.RegisterMessage(DirectorVolumeListQuery)

DirectorVolumeStats = _reflection.GeneratedProtocolMessageType('DirectorVolumeStats', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUMESTATS,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DirectorVolumeStats)
  ))
_sym_db.RegisterMessage(DirectorVolumeStats)

DirectorVolume = _reflection.GeneratedProtocolMessageType('DirectorVolume', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUME,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DirectorVolume)
  ))
_sym_db.RegisterMessage(DirectorVolume)

DirectorVolumeList = _reflection.GeneratedProtocolMessageType('DirectorVolumeList', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUMELIST,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DirectorVolumeList)
  ))
_sym_db.RegisterMessage(DirectorVolumeList)

DirectorRedirectListQuery = _reflection.GeneratedProtocolMessageType('DirectorRedirectListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORREDIRECTLISTQUERY,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DirectorRedirectListQuery)
  ))
_sym_db.RegisterMessage(DirectorRedirectListQuery)

DirectorRedirect = _reflection.GeneratedProtocolMessageType('DirectorRedirect', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORREDIRECT,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DirectorRedirect)
  ))
_sym_db.RegisterMessage(DirectorRedirect)

DirectorRedirectList = _reflection.GeneratedProtocolMessageType('DirectorRedirectList', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORREDIRECTLIST,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.DirectorRedirectList)
  ))
_sym_db.RegisterMessage(DirectorRedirectList)



_DIRECTORCONFIG = _descriptor.ServiceDescriptor(
  name='DirectorConfig',
  full_name='storageos_rpc.DirectorConfig',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=631,
  serialized_end=1159,
  methods=[
  _descriptor.MethodDescriptor(
    name='VolumeConfigure',
    full_name='storageos_rpc.DirectorConfig.VolumeConfigure',
    index=0,
    containing_service=None,
    input_type=_DIRECTORVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeUnconfigure',
    full_name='storageos_rpc.DirectorConfig.VolumeUnconfigure',
    index=1,
    containing_service=None,
    input_type=_DIRECTORVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='storageos_rpc.DirectorConfig.VolumeList',
    index=2,
    containing_service=None,
    input_type=_DIRECTORVOLUMELISTQUERY,
    output_type=_DIRECTORVOLUMELIST,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RedirectConfigure',
    full_name='storageos_rpc.DirectorConfig.RedirectConfigure',
    index=3,
    containing_service=None,
    input_type=_DIRECTORREDIRECT,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RedirectUnconfigure',
    full_name='storageos_rpc.DirectorConfig.RedirectUnconfigure',
    index=4,
    containing_service=None,
    input_type=_DIRECTORREDIRECT,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RedirectList',
    full_name='storageos_rpc.DirectorConfig.RedirectList',
    index=5,
    containing_service=None,
    input_type=_DIRECTORREDIRECTLISTQUERY,
    output_type=_DIRECTORREDIRECTLIST,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DIRECTORCONFIG)

DESCRIPTOR.services_by_name['DirectorConfig'] = _DIRECTORCONFIG

# @@protoc_insertion_point(module_scope)
