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
  package='director.v1',
  syntax='proto3',
  serialized_pb=_b('\n\x0e\x64irector.proto\x12\x0b\x64irector.v1\x1a\x0c\x63ommon.proto\"\x17\n\x15\x44irectorStatusRequest\"&\n\x0e\x44irectorStatus\x12\x14\n\x0cversion_info\x18\x01 \x01(\t\"-\n\x17\x44irectorVolumeListQuery\x12\x12\n\nvolume_ids\x18\x01 \x03(\r\"\x1a\n\x18\x44irectorVolumeStatistics\"\x16\n\x14\x44irectorVolumeStatus\"\xfd\x01\n\x0e\x44irectorVolume\x12&\n\x02\x63\x63\x18\x01 \x01(\x0b\x32\x1a.common.v1.DataplaneCommon\x12\x11\n\tvolume_id\x18\x02 \x01(\r\x12\x12\n\nwrite_pipe\x18\x03 \x01(\r\x12\x11\n\tread_pipe\x18\x04 \x01(\r\x12\x0b\n\x03qos\x18\x05 \x01(\x04\x12\x13\n\x0breplica_ids\x18\x06 \x03(\r\x12\x34\n\x05stats\x18\x07 \x01(\x0b\x32%.director.v1.DirectorVolumeStatistics\x12\x31\n\x06status\x18\x08 \x01(\x0b\x32!.director.v1.DirectorVolumeStatus\"B\n\x12\x44irectorVolumeList\x12,\n\x07volumes\x18\x01 \x03(\x0b\x32\x1b.director.v1.DirectorVolume\"9\n\x1d\x44irectorPresentationListQuery\x12\x18\n\x10presentation_ids\x18\x01 \x03(\r\"j\n\x14\x44irectorPresentation\x12&\n\x02\x63\x63\x18\x01 \x01(\x0b\x32\x1a.common.v1.DataplaneCommon\x12\x17\n\x0fpresentation_id\x18\x02 \x01(\r\x12\x11\n\ttarget_id\x18\x03 \x01(\r\"T\n\x18\x44irectorPresentationList\x12\x38\n\rpresentations\x18\x01 \x03(\x0b\x32!.director.v1.DirectorPresentation2\xd9\x05\n\x08\x44irector\x12K\n\x06Status\x12\".director.v1.DirectorStatusRequest\x1a\x1b.director.v1.DirectorStatus\"\x00\x12\x43\n\x0cVolumeCreate\x12\x1b.director.v1.DirectorVolume\x1a\x14.common.v1.RpcResult\"\x00\x12\x43\n\x0cVolumeUpdate\x12\x1b.director.v1.DirectorVolume\x1a\x14.common.v1.RpcResult\"\x00\x12\x43\n\x0cVolumeDelete\x12\x1b.director.v1.DirectorVolume\x1a\x14.common.v1.RpcResult\"\x00\x12U\n\nVolumeList\x12$.director.v1.DirectorVolumeListQuery\x1a\x1f.director.v1.DirectorVolumeList\"\x00\x12O\n\x12PresentationCreate\x12!.director.v1.DirectorPresentation\x1a\x14.common.v1.RpcResult\"\x00\x12O\n\x12PresentationUpdate\x12!.director.v1.DirectorPresentation\x1a\x14.common.v1.RpcResult\"\x00\x12O\n\x12PresentationDelete\x12!.director.v1.DirectorPresentation\x1a\x14.common.v1.RpcResult\"\x00\x12g\n\x10PresentationList\x12*.director.v1.DirectorPresentationListQuery\x1a%.director.v1.DirectorPresentationList\"\x00\x62\x06proto3')
  ,
  dependencies=[common__pb2.DESCRIPTOR,])




_DIRECTORSTATUSREQUEST = _descriptor.Descriptor(
  name='DirectorStatusRequest',
  full_name='director.v1.DirectorStatusRequest',
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
  serialized_end=68,
)


_DIRECTORSTATUS = _descriptor.Descriptor(
  name='DirectorStatus',
  full_name='director.v1.DirectorStatus',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version_info', full_name='director.v1.DirectorStatus.version_info', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_end=108,
)


_DIRECTORVOLUMELISTQUERY = _descriptor.Descriptor(
  name='DirectorVolumeListQuery',
  full_name='director.v1.DirectorVolumeListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volume_ids', full_name='director.v1.DirectorVolumeListQuery.volume_ids', index=0,
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
  serialized_start=110,
  serialized_end=155,
)


_DIRECTORVOLUMESTATISTICS = _descriptor.Descriptor(
  name='DirectorVolumeStatistics',
  full_name='director.v1.DirectorVolumeStatistics',
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
  serialized_start=157,
  serialized_end=183,
)


_DIRECTORVOLUMESTATUS = _descriptor.Descriptor(
  name='DirectorVolumeStatus',
  full_name='director.v1.DirectorVolumeStatus',
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
  serialized_start=185,
  serialized_end=207,
)


_DIRECTORVOLUME = _descriptor.Descriptor(
  name='DirectorVolume',
  full_name='director.v1.DirectorVolume',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='director.v1.DirectorVolume.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='volume_id', full_name='director.v1.DirectorVolume.volume_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='write_pipe', full_name='director.v1.DirectorVolume.write_pipe', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='read_pipe', full_name='director.v1.DirectorVolume.read_pipe', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='qos', full_name='director.v1.DirectorVolume.qos', index=4,
      number=5, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='replica_ids', full_name='director.v1.DirectorVolume.replica_ids', index=5,
      number=6, type=13, cpp_type=3, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='stats', full_name='director.v1.DirectorVolume.stats', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='status', full_name='director.v1.DirectorVolume.status', index=7,
      number=8, type=11, cpp_type=10, label=1,
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
  serialized_start=210,
  serialized_end=463,
)


_DIRECTORVOLUMELIST = _descriptor.Descriptor(
  name='DirectorVolumeList',
  full_name='director.v1.DirectorVolumeList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='volumes', full_name='director.v1.DirectorVolumeList.volumes', index=0,
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
  serialized_start=465,
  serialized_end=531,
)


_DIRECTORPRESENTATIONLISTQUERY = _descriptor.Descriptor(
  name='DirectorPresentationListQuery',
  full_name='director.v1.DirectorPresentationListQuery',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='presentation_ids', full_name='director.v1.DirectorPresentationListQuery.presentation_ids', index=0,
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
  serialized_start=533,
  serialized_end=590,
)


_DIRECTORPRESENTATION = _descriptor.Descriptor(
  name='DirectorPresentation',
  full_name='director.v1.DirectorPresentation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cc', full_name='director.v1.DirectorPresentation.cc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='presentation_id', full_name='director.v1.DirectorPresentation.presentation_id', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='target_id', full_name='director.v1.DirectorPresentation.target_id', index=2,
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
  serialized_start=592,
  serialized_end=698,
)


_DIRECTORPRESENTATIONLIST = _descriptor.Descriptor(
  name='DirectorPresentationList',
  full_name='director.v1.DirectorPresentationList',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='presentations', full_name='director.v1.DirectorPresentationList.presentations', index=0,
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
  serialized_start=700,
  serialized_end=784,
)

_DIRECTORVOLUME.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMON
_DIRECTORVOLUME.fields_by_name['stats'].message_type = _DIRECTORVOLUMESTATISTICS
_DIRECTORVOLUME.fields_by_name['status'].message_type = _DIRECTORVOLUMESTATUS
_DIRECTORVOLUMELIST.fields_by_name['volumes'].message_type = _DIRECTORVOLUME
_DIRECTORPRESENTATION.fields_by_name['cc'].message_type = common__pb2._DATAPLANECOMMON
_DIRECTORPRESENTATIONLIST.fields_by_name['presentations'].message_type = _DIRECTORPRESENTATION
DESCRIPTOR.message_types_by_name['DirectorStatusRequest'] = _DIRECTORSTATUSREQUEST
DESCRIPTOR.message_types_by_name['DirectorStatus'] = _DIRECTORSTATUS
DESCRIPTOR.message_types_by_name['DirectorVolumeListQuery'] = _DIRECTORVOLUMELISTQUERY
DESCRIPTOR.message_types_by_name['DirectorVolumeStatistics'] = _DIRECTORVOLUMESTATISTICS
DESCRIPTOR.message_types_by_name['DirectorVolumeStatus'] = _DIRECTORVOLUMESTATUS
DESCRIPTOR.message_types_by_name['DirectorVolume'] = _DIRECTORVOLUME
DESCRIPTOR.message_types_by_name['DirectorVolumeList'] = _DIRECTORVOLUMELIST
DESCRIPTOR.message_types_by_name['DirectorPresentationListQuery'] = _DIRECTORPRESENTATIONLISTQUERY
DESCRIPTOR.message_types_by_name['DirectorPresentation'] = _DIRECTORPRESENTATION
DESCRIPTOR.message_types_by_name['DirectorPresentationList'] = _DIRECTORPRESENTATIONLIST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DirectorStatusRequest = _reflection.GeneratedProtocolMessageType('DirectorStatusRequest', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORSTATUSREQUEST,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorStatusRequest)
  ))
_sym_db.RegisterMessage(DirectorStatusRequest)

DirectorStatus = _reflection.GeneratedProtocolMessageType('DirectorStatus', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORSTATUS,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorStatus)
  ))
_sym_db.RegisterMessage(DirectorStatus)

DirectorVolumeListQuery = _reflection.GeneratedProtocolMessageType('DirectorVolumeListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUMELISTQUERY,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorVolumeListQuery)
  ))
_sym_db.RegisterMessage(DirectorVolumeListQuery)

DirectorVolumeStatistics = _reflection.GeneratedProtocolMessageType('DirectorVolumeStatistics', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUMESTATISTICS,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorVolumeStatistics)
  ))
_sym_db.RegisterMessage(DirectorVolumeStatistics)

DirectorVolumeStatus = _reflection.GeneratedProtocolMessageType('DirectorVolumeStatus', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUMESTATUS,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorVolumeStatus)
  ))
_sym_db.RegisterMessage(DirectorVolumeStatus)

DirectorVolume = _reflection.GeneratedProtocolMessageType('DirectorVolume', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUME,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorVolume)
  ))
_sym_db.RegisterMessage(DirectorVolume)

DirectorVolumeList = _reflection.GeneratedProtocolMessageType('DirectorVolumeList', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORVOLUMELIST,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorVolumeList)
  ))
_sym_db.RegisterMessage(DirectorVolumeList)

DirectorPresentationListQuery = _reflection.GeneratedProtocolMessageType('DirectorPresentationListQuery', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORPRESENTATIONLISTQUERY,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorPresentationListQuery)
  ))
_sym_db.RegisterMessage(DirectorPresentationListQuery)

DirectorPresentation = _reflection.GeneratedProtocolMessageType('DirectorPresentation', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORPRESENTATION,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorPresentation)
  ))
_sym_db.RegisterMessage(DirectorPresentation)

DirectorPresentationList = _reflection.GeneratedProtocolMessageType('DirectorPresentationList', (_message.Message,), dict(
  DESCRIPTOR = _DIRECTORPRESENTATIONLIST,
  __module__ = 'director_pb2'
  # @@protoc_insertion_point(class_scope:director.v1.DirectorPresentationList)
  ))
_sym_db.RegisterMessage(DirectorPresentationList)



_DIRECTOR = _descriptor.ServiceDescriptor(
  name='Director',
  full_name='director.v1.Director',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=787,
  serialized_end=1516,
  methods=[
  _descriptor.MethodDescriptor(
    name='Status',
    full_name='director.v1.Director.Status',
    index=0,
    containing_service=None,
    input_type=_DIRECTORSTATUSREQUEST,
    output_type=_DIRECTORSTATUS,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeCreate',
    full_name='director.v1.Director.VolumeCreate',
    index=1,
    containing_service=None,
    input_type=_DIRECTORVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeUpdate',
    full_name='director.v1.Director.VolumeUpdate',
    index=2,
    containing_service=None,
    input_type=_DIRECTORVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeDelete',
    full_name='director.v1.Director.VolumeDelete',
    index=3,
    containing_service=None,
    input_type=_DIRECTORVOLUME,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VolumeList',
    full_name='director.v1.Director.VolumeList',
    index=4,
    containing_service=None,
    input_type=_DIRECTORVOLUMELISTQUERY,
    output_type=_DIRECTORVOLUMELIST,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationCreate',
    full_name='director.v1.Director.PresentationCreate',
    index=5,
    containing_service=None,
    input_type=_DIRECTORPRESENTATION,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationUpdate',
    full_name='director.v1.Director.PresentationUpdate',
    index=6,
    containing_service=None,
    input_type=_DIRECTORPRESENTATION,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationDelete',
    full_name='director.v1.Director.PresentationDelete',
    index=7,
    containing_service=None,
    input_type=_DIRECTORPRESENTATION,
    output_type=common__pb2._RPCRESULT,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PresentationList',
    full_name='director.v1.Director.PresentationList',
    index=8,
    containing_service=None,
    input_type=_DIRECTORPRESENTATIONLISTQUERY,
    output_type=_DIRECTORPRESENTATIONLIST,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_DIRECTOR)

DESCRIPTOR.services_by_name['Director'] = _DIRECTOR

# @@protoc_insertion_point(module_scope)
