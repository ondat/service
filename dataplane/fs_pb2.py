# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: fs.proto

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
  name='fs.proto',
  package='storageos_rpc',
  syntax='proto3',
  serialized_pb=_b('\n\x08\x66s.proto\x12\rstorageos_rpc\x1a\x0c\x63ommon.proto\"\'\n\x11\x46sVolumeListQuery\x12\x12\n\nvolume_ids\x18\x01 \x03(\r\"\x14\n\x12\x46sVolumeStatistics\"\x10\n\x0e\x46sVolumeStatus\"\x89\x03\n\x08\x46sVolume\x12*\n\x02\x63\x63\x18\x01 \x01(\x0b\x32\x1e.storageos_rpc.DataplaneCommon\x12\x11\n\tvolume_id\x18\x02 \x01(\r\x12;\n\tnode_type\x18\x03 \x01(\x0e\x32(.storageos_rpc.FsVolume.VolumeDeviceType\x12\x15\n\rdevice_number\x18\x04 \x01(\r\x12\x10\n\x08\x66ilename\x18\x05 \x01(\t\x12\x15\n\rlinked_volume\x18\x06 \x01(\x08\x12\x18\n\x10target_volume_id\x18\x07 \x01(\r\x12\x19\n\x11volume_size_bytes\x18\x08 \x01(\x04\x12\x30\n\x05stats\x18\t \x01(\x0b\x32!.storageos_rpc.FsVolumeStatistics\x12-\n\x06status\x18\n \x01(\x0b\x32\x1d.storageos_rpc.FsVolumeStatus\"+\n\x10VolumeDeviceType\x12\x08\n\x04\x46ILE\x10\x00\x12\r\n\tNBD_BLOCK\x10\x01\"8\n\x0c\x46sVolumeList\x12(\n\x07volumes\x18\x01 \x03(\x0b\x32\x17.storageos_rpc.FsVolume\"2\n\x17\x46sPresentationListQuery\x12\x17\n\x0fpresentation_id\x18\x01 \x03(\r\"b\n\x0e\x46sPresentation\x12*\n\x02\x63\x63\x18\x01 \x01(\x0b\x32\x1e.storageos_rpc.DataplaneCommon\x12\x11\n\tsource_id\x18\x02 \x01(\r\x12\x11\n\ttarget_id\x18\x03 \x01(\r\"J\n\x12\x46sPresentationList\x12\x34\n\rpresentations\x18\x01 \x03(\x0b\x32\x1d.storageos_rpc.FsPresentation2\xf6\x04\n\x02\x46s\x12\x43\n\x0cVolumeCreate\x12\x17.storageos_rpc.FsVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12\x43\n\x0cVolumeUpdate\x12\x17.storageos_rpc.FsVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12\x43\n\x0cVolumeDelete\x12\x17.storageos_rpc.FsVolume\x1a\x18.storageos_rpc.RpcResult\"\x00\x12M\n\nVolumeList\x12 .storageos_rpc.FsVolumeListQuery\x1a\x1b.storageos_rpc.FsVolumeList\"\x00\x12O\n\x12PresentationCreate\x12\x1d.storageos_rpc.FsPresentation\x1a\x18.storageos_rpc.RpcResult\"\x00\x12O\n\x12PresentationUpdate\x12\x1d.storageos_rpc.FsPresentation\x1a\x18.storageos_rpc.RpcResult\"\x00\x12O\n\x12PresentationDelete\x12\x1d.storageos_rpc.FsPresentation\x1a\x18.storageos_rpc.RpcResult\"\x00\x12_\n\x10PresentationList\x12&.storageos_rpc.FsPresentationListQuery\x1a!.storageos_rpc.FsPresentationList\"\x00\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])



_FSVOLUME_VOLUMEDEVICETYPE = _descriptor.EnumDescriptor(
  name='VolumeDeviceType',
  full_name='storageos_rpc.FsVolume.VolumeDeviceType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='FILE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NBD_BLOCK', index=1, number=1,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=473,
  serialized_end=516,
)
_sym_db.RegisterEnumDescriptor(_FSVOLUME_VOLUMEDEVICETYPE)


_FSVOLUMELISTQUERY = _descriptor.Descriptor(
  name='FsVolumeListQuery',
  full_name='storageos_rpc.FsVolumeListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volume_ids', full_name='storageos_rpc.FsVolumeListQuery.volume_ids', index=0,
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
  serialized_start=41,
  serialized_end=80,
)


_FSVOLUMESTATISTICS = _descriptor.Descriptor(
  name='FsVolumeStatistics',
  full_name='storageos_rpc.FsVolumeStatistics',
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
  serialized_start=82,
  serialized_end=102,
)


_FSVOLUMESTATUS = _descriptor.Descriptor(
  name='FsVolumeStatus',
  full_name='storageos_rpc.FsVolumeStatus',
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
  serialized_start=104,
  serialized_end=120,
)


_FSVOLUME = _descriptor.Descriptor(
  name='FsVolume',
  full_name='storageos_rpc.FsVolume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='storageos_rpc.FsVolume.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_id', full_name='storageos_rpc.FsVolume.volume_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='node_type', full_name='storageos_rpc.FsVolume.node_type', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='device_number', full_name='storageos_rpc.FsVolume.device_number', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='filename', full_name='storageos_rpc.FsVolume.filename', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='linked_volume', full_name='storageos_rpc.FsVolume.linked_volume', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='target_volume_id', full_name='storageos_rpc.FsVolume.target_volume_id', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_size_bytes', full_name='storageos_rpc.FsVolume.volume_size_bytes', index=7,
      number=8, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stats', full_name='storageos_rpc.FsVolume.stats', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='storageos_rpc.FsVolume.status', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _FSVOLUME_VOLUMEDEVICETYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=123,
  serialized_end=516,
)


_FSVOLUMELIST = _descriptor.Descriptor(
  name='FsVolumeList',
  full_name='storageos_rpc.FsVolumeList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volumes', full_name='storageos_rpc.FsVolumeList.volumes', index=0,
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
  serialized_start=518,
  serialized_end=574,
)


_FSPRESENTATIONLISTQUERY = _descriptor.Descriptor(
  name='FsPresentationListQuery',
  full_name='storageos_rpc.FsPresentationListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='presentation_id', full_name='storageos_rpc.FsPresentationListQuery.presentation_id', index=0,
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
  serialized_start=576,
  serialized_end=626,
)


_FSPRESENTATION = _descriptor.Descriptor(
  name='FsPresentation',
  full_name='storageos_rpc.FsPresentation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='storageos_rpc.FsPresentation.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='source_id', full_name='storageos_rpc.FsPresentation.source_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='target_id', full_name='storageos_rpc.FsPresentation.target_id', index=2,
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
  serialized_start=628,
  serialized_end=726,
)


_FSPRESENTATIONLIST = _descriptor.Descriptor(
  name='FsPresentationList',
  full_name='storageos_rpc.FsPresentationList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='presentations', full_name='storageos_rpc.FsPresentationList.presentations', index=0,
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
  serialized_start=728,
  serialized_end=802,
)

_FSVOLUME.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMON
_FSVOLUME.fields_by_name['node_type'].enum_type = _FSVOLUME_VOLUMEDEVICETYPE
_FSVOLUME.fields_by_name['stats'].message_type = _FSVOLUMESTATISTICS
_FSVOLUME.fields_by_name['status'].message_type = _FSVOLUMESTATUS
_FSVOLUME_VOLUMEDEVICETYPE.containing_type = _FSVOLUME
_FSVOLUMELIST.fields_by_name['volumes'].message_type = _FSVOLUME
_FSPRESENTATION.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMON
_FSPRESENTATIONLIST.fields_by_name['presentations'].message_type = _FSPRESENTATION
DESCRIPTOR.message_types_by_name['FsVolumeListQuery'] = _FSVOLUMELISTQUERY
DESCRIPTOR.message_types_by_name['FsVolumeStatistics'] = _FSVOLUMESTATISTICS
DESCRIPTOR.message_types_by_name['FsVolumeStatus'] = _FSVOLUMESTATUS
DESCRIPTOR.message_types_by_name['FsVolume'] = _FSVOLUME
DESCRIPTOR.message_types_by_name['FsVolumeList'] = _FSVOLUMELIST
DESCRIPTOR.message_types_by_name['FsPresentationListQuery'] = _FSPRESENTATIONLISTQUERY
DESCRIPTOR.message_types_by_name['FsPresentation'] = _FSPRESENTATION
DESCRIPTOR.message_types_by_name['FsPresentationList'] = _FSPRESENTATIONLIST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

FsVolumeListQuery = _reflection.GeneratedProtocolMessageType('FsVolumeListQuery', (_message.Message,), dict(
  DESCRIPTOR = _FSVOLUMELISTQUERY,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsVolumeListQuery)
  ))
_sym_db.RegisterMessage(FsVolumeListQuery)

FsVolumeStatistics = _reflection.GeneratedProtocolMessageType('FsVolumeStatistics', (_message.Message,), dict(
  DESCRIPTOR = _FSVOLUMESTATISTICS,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsVolumeStatistics)
  ))
_sym_db.RegisterMessage(FsVolumeStatistics)

FsVolumeStatus = _reflection.GeneratedProtocolMessageType('FsVolumeStatus', (_message.Message,), dict(
  DESCRIPTOR = _FSVOLUMESTATUS,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsVolumeStatus)
  ))
_sym_db.RegisterMessage(FsVolumeStatus)

FsVolume = _reflection.GeneratedProtocolMessageType('FsVolume', (_message.Message,), dict(
  DESCRIPTOR = _FSVOLUME,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsVolume)
  ))
_sym_db.RegisterMessage(FsVolume)

FsVolumeList = _reflection.GeneratedProtocolMessageType('FsVolumeList', (_message.Message,), dict(
  DESCRIPTOR = _FSVOLUMELIST,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsVolumeList)
  ))
_sym_db.RegisterMessage(FsVolumeList)

FsPresentationListQuery = _reflection.GeneratedProtocolMessageType('FsPresentationListQuery', (_message.Message,), dict(
  DESCRIPTOR = _FSPRESENTATIONLISTQUERY,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsPresentationListQuery)
  ))
_sym_db.RegisterMessage(FsPresentationListQuery)

FsPresentation = _reflection.GeneratedProtocolMessageType('FsPresentation', (_message.Message,), dict(
  DESCRIPTOR = _FSPRESENTATION,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsPresentation)
  ))
_sym_db.RegisterMessage(FsPresentation)

FsPresentationList = _reflection.GeneratedProtocolMessageType('FsPresentationList', (_message.Message,), dict(
  DESCRIPTOR = _FSPRESENTATIONLIST,
  __module__ = 'fs_pb2'
  # @@protoc_insertion_point(class_scope:storageos_rpc.FsPresentationList)
  ))
_sym_db.RegisterMessage(FsPresentationList)



_FS = _descriptor.ServiceDescriptor(
  name='Fs',
  full_name='storageos_rpc.Fs',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=805,
  serialized_end=1435,
  methods=[
  _descriptor.MethodDescriptor(
    name='VolumeCreate',
    full_name='storageos_rpc.Fs.VolumeCreate',
    index=0,
    containing_service=None,
    input_type=_FSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeUpdate',
    full_name='storageos_rpc.Fs.VolumeUpdate',
    index=1,
    containing_service=None,
    input_type=_FSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeDelete',
    full_name='storageos_rpc.Fs.VolumeDelete',
    index=2,
    containing_service=None,
    input_type=_FSVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='storageos_rpc.Fs.VolumeList',
    index=3,
    containing_service=None,
    input_type=_FSVOLUMELISTQUERY,
    output_type=_FSVOLUMELIST,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationCreate',
    full_name='storageos_rpc.Fs.PresentationCreate',
    index=4,
    containing_service=None,
    input_type=_FSPRESENTATION,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationUpdate',
    full_name='storageos_rpc.Fs.PresentationUpdate',
    index=5,
    containing_service=None,
    input_type=_FSPRESENTATION,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationDelete',
    full_name='storageos_rpc.Fs.PresentationDelete',
    index=6,
    containing_service=None,
    input_type=_FSPRESENTATION,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationList',
    full_name='storageos_rpc.Fs.PresentationList',
    index=7,
    containing_service=None,
    input_type=_FSPRESENTATIONLISTQUERY,
    output_type=_FSPRESENTATIONLIST,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_FS)

DESCRIPTOR.services_by_name['Fs'] = _FS

# @@protoc_insertion_point(module_scope)
