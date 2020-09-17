# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: outing-parents.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='outing-parents.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x14outing-parents.proto\"-\n\x1b\x43onfirmOutingByOCodeRequest\x12\x0e\n\x06o_code\x18\x01 \x01(\t\"I\n\x1c\x43onfirmOutingByOCodeResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t2\xbc\x01\n\rOutingParents\x12U\n\x14\x41pproveOutingByOCode\x12\x1c.ConfirmOutingByOCodeRequest\x1a\x1d.ConfirmOutingByOCodeResponse\"\x00\x12T\n\x13RejectOutingByOCode\x12\x1c.ConfirmOutingByOCodeRequest\x1a\x1d.ConfirmOutingByOCodeResponse\"\x00\x62\x06proto3'
)




_CONFIRMOUTINGBYOCODEREQUEST = _descriptor.Descriptor(
  name='ConfirmOutingByOCodeRequest',
  full_name='ConfirmOutingByOCodeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='o_code', full_name='ConfirmOutingByOCodeRequest.o_code', index=0,
      number=1, type=9, cpp_type=9, label=1,
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
  serialized_start=24,
  serialized_end=69,
)


_CONFIRMOUTINGBYOCODERESPONSE = _descriptor.Descriptor(
  name='ConfirmOutingByOCodeResponse',
  full_name='ConfirmOutingByOCodeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='ConfirmOutingByOCodeResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='ConfirmOutingByOCodeResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='ConfirmOutingByOCodeResponse.msg', index=2,
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
  serialized_start=71,
  serialized_end=144,
)

DESCRIPTOR.message_types_by_name['ConfirmOutingByOCodeRequest'] = _CONFIRMOUTINGBYOCODEREQUEST
DESCRIPTOR.message_types_by_name['ConfirmOutingByOCodeResponse'] = _CONFIRMOUTINGBYOCODERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ConfirmOutingByOCodeRequest = _reflection.GeneratedProtocolMessageType('ConfirmOutingByOCodeRequest', (_message.Message,), {
  'DESCRIPTOR' : _CONFIRMOUTINGBYOCODEREQUEST,
  '__module__' : 'outing_parents_pb2'
  # @@protoc_insertion_point(class_scope:ConfirmOutingByOCodeRequest)
  })
_sym_db.RegisterMessage(ConfirmOutingByOCodeRequest)

ConfirmOutingByOCodeResponse = _reflection.GeneratedProtocolMessageType('ConfirmOutingByOCodeResponse', (_message.Message,), {
  'DESCRIPTOR' : _CONFIRMOUTINGBYOCODERESPONSE,
  '__module__' : 'outing_parents_pb2'
  # @@protoc_insertion_point(class_scope:ConfirmOutingByOCodeResponse)
  })
_sym_db.RegisterMessage(ConfirmOutingByOCodeResponse)



_OUTINGPARENTS = _descriptor.ServiceDescriptor(
  name='OutingParents',
  full_name='OutingParents',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=147,
  serialized_end=335,
  methods=[
  _descriptor.MethodDescriptor(
    name='ApproveOutingByOCode',
    full_name='OutingParents.ApproveOutingByOCode',
    index=0,
    containing_service=None,
    input_type=_CONFIRMOUTINGBYOCODEREQUEST,
    output_type=_CONFIRMOUTINGBYOCODERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='RejectOutingByOCode',
    full_name='OutingParents.RejectOutingByOCode',
    index=1,
    containing_service=None,
    input_type=_CONFIRMOUTINGBYOCODEREQUEST,
    output_type=_CONFIRMOUTINGBYOCODERESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_OUTINGPARENTS)

DESCRIPTOR.services_by_name['OutingParents'] = _OUTINGPARENTS

# @@protoc_insertion_point(module_scope)