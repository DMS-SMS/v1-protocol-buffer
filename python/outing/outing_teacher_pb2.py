# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: outing-teacher.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='outing-teacher.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x14outing-teacher.proto\"v\n\x1aGetOutingWithFilterRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t\x12\r\n\x05grade\x18\x03 \x01(\x05\x12\r\n\x05group\x18\x04 \x01(\x05\x12\r\n\x05start\x18\x05 \x01(\x05\x12\r\n\x05\x63ount\x18\x06 \x01(\x05\"T\n\x0eOutingResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t\x12\x17\n\x06outing\x18\x04 \x03(\x0b\x32\x07.Outing\"\xcb\x01\n\x06Outing\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\r\n\x05grade\x18\x02 \x01(\x05\x12\r\n\x05group\x18\x03 \x01(\x05\x12\x0e\n\x06number\x18\x04 \x01(\x05\x12\r\n\x05place\x18\x05 \x01(\t\x12\x0e\n\x06reason\x18\x06 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x07 \x01(\t\x12\x12\n\nstart_time\x18\x08 \x01(\t\x12\x10\n\x08\x65nd_time\x18\t \x01(\t\x12\x0e\n\x06status\x18\n \x01(\t\x12\x11\n\tsituation\x18\x0b \x01(\t\x12\x0f\n\x07is_late\x18\x0c \x01(\x08\"1\n\x14\x43onfirmOutingRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0b\n\x03oid\x18\x02 \x01(\t\"B\n\x15\x43onfirmOutingResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t2\x9b\x02\n\rOutingTeacher\x12\x45\n\x13GetOutingWithFilter\x12\x1b.GetOutingWithFilterRequest\x1a\x0f.OutingResponse\"\x00\x12@\n\rApproveOuting\x12\x15.ConfirmOutingRequest\x1a\x16.ConfirmOutingResponse\"\x00\x12?\n\x0cRejectOuting\x12\x15.ConfirmOutingRequest\x1a\x16.ConfirmOutingResponse\"\x00\x12@\n\rCertifyOuting\x12\x15.ConfirmOutingRequest\x1a\x16.ConfirmOutingResponse\"\x00\x62\x06proto3'
)




_GETOUTINGWITHFILTERREQUEST = _descriptor.Descriptor(
  name='GetOutingWithFilterRequest',
  full_name='GetOutingWithFilterRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='GetOutingWithFilterRequest.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='GetOutingWithFilterRequest.status', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='grade', full_name='GetOutingWithFilterRequest.grade', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='group', full_name='GetOutingWithFilterRequest.group', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start', full_name='GetOutingWithFilterRequest.start', index=4,
      number=5, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='count', full_name='GetOutingWithFilterRequest.count', index=5,
      number=6, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=24,
  serialized_end=142,
)


_OUTINGRESPONSE = _descriptor.Descriptor(
  name='OutingResponse',
  full_name='OutingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='OutingResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='OutingResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='OutingResponse.msg', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='outing', full_name='OutingResponse.outing', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=144,
  serialized_end=228,
)


_OUTING = _descriptor.Descriptor(
  name='Outing',
  full_name='Outing',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='Outing.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='grade', full_name='Outing.grade', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='group', full_name='Outing.group', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='number', full_name='Outing.number', index=3,
      number=4, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='place', full_name='Outing.place', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='reason', full_name='Outing.reason', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='date', full_name='Outing.date', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='Outing.start_time', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='Outing.end_time', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='Outing.status', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='situation', full_name='Outing.situation', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='is_late', full_name='Outing.is_late', index=11,
      number=12, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=231,
  serialized_end=434,
)


_CONFIRMOUTINGREQUEST = _descriptor.Descriptor(
  name='ConfirmOutingRequest',
  full_name='ConfirmOutingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='ConfirmOutingRequest.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='oid', full_name='ConfirmOutingRequest.oid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=436,
  serialized_end=485,
)


_CONFIRMOUTINGRESPONSE = _descriptor.Descriptor(
  name='ConfirmOutingResponse',
  full_name='ConfirmOutingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='ConfirmOutingResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='ConfirmOutingResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='ConfirmOutingResponse.msg', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=487,
  serialized_end=553,
)

_OUTINGRESPONSE.fields_by_name['outing'].message_type = _OUTING
DESCRIPTOR.message_types_by_name['GetOutingWithFilterRequest'] = _GETOUTINGWITHFILTERREQUEST
DESCRIPTOR.message_types_by_name['OutingResponse'] = _OUTINGRESPONSE
DESCRIPTOR.message_types_by_name['Outing'] = _OUTING
DESCRIPTOR.message_types_by_name['ConfirmOutingRequest'] = _CONFIRMOUTINGREQUEST
DESCRIPTOR.message_types_by_name['ConfirmOutingResponse'] = _CONFIRMOUTINGRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetOutingWithFilterRequest = _reflection.GeneratedProtocolMessageType('GetOutingWithFilterRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETOUTINGWITHFILTERREQUEST,
  '__module__' : 'outing_teacher_pb2'
  # @@protoc_insertion_point(class_scope:GetOutingWithFilterRequest)
  })
_sym_db.RegisterMessage(GetOutingWithFilterRequest)

OutingResponse = _reflection.GeneratedProtocolMessageType('OutingResponse', (_message.Message,), {
  'DESCRIPTOR' : _OUTINGRESPONSE,
  '__module__' : 'outing_teacher_pb2'
  # @@protoc_insertion_point(class_scope:OutingResponse)
  })
_sym_db.RegisterMessage(OutingResponse)

Outing = _reflection.GeneratedProtocolMessageType('Outing', (_message.Message,), {
  'DESCRIPTOR' : _OUTING,
  '__module__' : 'outing_teacher_pb2'
  # @@protoc_insertion_point(class_scope:Outing)
  })
_sym_db.RegisterMessage(Outing)

ConfirmOutingRequest = _reflection.GeneratedProtocolMessageType('ConfirmOutingRequest', (_message.Message,), {
  'DESCRIPTOR' : _CONFIRMOUTINGREQUEST,
  '__module__' : 'outing_teacher_pb2'
  # @@protoc_insertion_point(class_scope:ConfirmOutingRequest)
  })
_sym_db.RegisterMessage(ConfirmOutingRequest)

ConfirmOutingResponse = _reflection.GeneratedProtocolMessageType('ConfirmOutingResponse', (_message.Message,), {
  'DESCRIPTOR' : _CONFIRMOUTINGRESPONSE,
  '__module__' : 'outing_teacher_pb2'
  # @@protoc_insertion_point(class_scope:ConfirmOutingResponse)
  })
_sym_db.RegisterMessage(ConfirmOutingResponse)



_OUTINGTEACHER = _descriptor.ServiceDescriptor(
  name='OutingTeacher',
  full_name='OutingTeacher',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=556,
  serialized_end=839,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetOutingWithFilter',
    full_name='OutingTeacher.GetOutingWithFilter',
    index=0,
    containing_service=None,
    input_type=_GETOUTINGWITHFILTERREQUEST,
    output_type=_OUTINGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='ApproveOuting',
    full_name='OutingTeacher.ApproveOuting',
    index=1,
    containing_service=None,
    input_type=_CONFIRMOUTINGREQUEST,
    output_type=_CONFIRMOUTINGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RejectOuting',
    full_name='OutingTeacher.RejectOuting',
    index=2,
    containing_service=None,
    input_type=_CONFIRMOUTINGREQUEST,
    output_type=_CONFIRMOUTINGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CertifyOuting',
    full_name='OutingTeacher.CertifyOuting',
    index=3,
    containing_service=None,
    input_type=_CONFIRMOUTINGREQUEST,
    output_type=_CONFIRMOUTINGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_OUTINGTEACHER)

DESCRIPTOR.services_by_name['OutingTeacher'] = _OUTINGTEACHER

# @@protoc_insertion_point(module_scope)
