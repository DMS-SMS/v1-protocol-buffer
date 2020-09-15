# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: outing-student.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='outing-student.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x14outing-student.proto\"\x89\x01\n\x13\x43reateOutingRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x02 \x01(\t\x12\x12\n\nstart_time\x18\x03 \x01(\t\x12\x10\n\x08\x65nd_time\x18\x04 \x01(\t\x12\r\n\x05place\x18\x05 \x01(\t\x12\x0e\n\x06reason\x18\x06 \x01(\t\x12\x11\n\tsituation\x18\x07 \x01(\t\"N\n\x14\x43reateOutingResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t\x12\x0b\n\x03oid\x18\x04 \x01(\t\"S\n\x18GetStudentOutingsRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0b\n\x03sid\x18\x02 \x01(\t\x12\r\n\x05start\x18\x03 \x01(\x05\x12\r\n\x05\x63ount\x18\x04 \x01(\x05\"f\n\x19GetStudentOutingsResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t\x12\x1e\n\x06outing\x18\x04 \x03(\x0b\x32\x0e.StudentOuting\"\x85\x01\n\rStudentOuting\x12\r\n\x05place\x18\x01 \x01(\t\x12\x0e\n\x06reason\x18\x02 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x03 \x01(\t\x12\x12\n\nstart_time\x18\x04 \x01(\t\x12\x10\n\x08\x65nd_time\x18\x05 \x01(\t\x12\x11\n\tsituation\x18\x06 \x01(\t\x12\x0e\n\x06status\x18\x07 \x01(\t\"3\n\x16GetOutingInformRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0b\n\x03oid\x18\x02 \x01(\t\"\xbe\x01\n\x17GetOutingInformResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t\x12\r\n\x05place\x18\x04 \x01(\t\x12\x0e\n\x06reason\x18\x05 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x06 \x01(\t\x12\x12\n\nstart_time\x18\x07 \x01(\t\x12\x10\n\x08\x65nd_time\x18\x08 \x01(\t\x12\x13\n\x0bo_situation\x18\t \x01(\t\x12\x10\n\x08o_status\x18\n \x01(\t\"6\n\x19GetCardAboutOutingRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0b\n\x03oid\x18\x02 \x01(\t\"\xf4\x01\n\x1aGetCardAboutOutingResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t\x12\r\n\x05place\x18\x04 \x01(\t\x12\x0c\n\x04\x64\x61te\x18\x05 \x01(\t\x12\x12\n\nstart_time\x18\x06 \x01(\t\x12\x10\n\x08\x65nd_time\x18\x07 \x01(\t\x12\x10\n\x08o_status\x18\x08 \x01(\t\x12\x0c\n\x04name\x18\t \x01(\t\x12\r\n\x05grade\x18\n \x01(\x05\x12\x0e\n\x06\x63lass_\x18\x0b \x01(\x05\x12\x0e\n\x06number\x18\x0c \x01(\x05\x12\x19\n\x11profile_image_uri\x18\r \x01(\t\")\n\x0cGoOutRequest\x12\x0c\n\x04uuid\x18\x01 \x01(\t\x12\x0b\n\x03oid\x18\x02 \x01(\t\":\n\rGoOutResponse\x12\x0e\n\x06status\x18\x01 \x01(\x05\x12\x0c\n\x04\x63ode\x18\x02 \x01(\x05\x12\x0b\n\x03msg\x18\x03 \x01(\t2\x8f\x03\n\rOutingStudent\x12=\n\x0c\x43reateOuting\x12\x14.CreateOutingRequest\x1a\x15.CreateOutingResponse\"\x00\x12L\n\x11GetStudentOutings\x12\x19.GetStudentOutingsRequest\x1a\x1a.GetStudentOutingsResponse\"\x00\x12\x46\n\x0fGetOutingInform\x12\x17.GetOutingInformRequest\x1a\x18.GetOutingInformResponse\"\x00\x12O\n\x12GetCardAboutOuting\x12\x1a.GetCardAboutOutingRequest\x1a\x1b.GetCardAboutOutingResponse\"\x00\x12(\n\x05GoOut\x12\r.GoOutRequest\x1a\x0e.GoOutResponse\"\x00\x12.\n\x0b\x46inishGoOut\x12\r.GoOutRequest\x1a\x0e.GoOutResponse\"\x00\x62\x06proto3'
)




_CREATEOUTINGREQUEST = _descriptor.Descriptor(
  name='CreateOutingRequest',
  full_name='CreateOutingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='CreateOutingRequest.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='date', full_name='CreateOutingRequest.date', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='CreateOutingRequest.start_time', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='CreateOutingRequest.end_time', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='place', full_name='CreateOutingRequest.place', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='reason', full_name='CreateOutingRequest.reason', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='situation', full_name='CreateOutingRequest.situation', index=6,
      number=7, type=9, cpp_type=9, label=1,
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
  serialized_start=25,
  serialized_end=162,
)


_CREATEOUTINGRESPONSE = _descriptor.Descriptor(
  name='CreateOutingResponse',
  full_name='CreateOutingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='CreateOutingResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='CreateOutingResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='CreateOutingResponse.msg', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='oid', full_name='CreateOutingResponse.oid', index=3,
      number=4, type=9, cpp_type=9, label=1,
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
  serialized_start=164,
  serialized_end=242,
)


_GETSTUDENTOUTINGSREQUEST = _descriptor.Descriptor(
  name='GetStudentOutingsRequest',
  full_name='GetStudentOutingsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='GetStudentOutingsRequest.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='sid', full_name='GetStudentOutingsRequest.sid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start', full_name='GetStudentOutingsRequest.start', index=2,
      number=3, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='count', full_name='GetStudentOutingsRequest.count', index=3,
      number=4, type=5, cpp_type=1, label=1,
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
  serialized_start=244,
  serialized_end=327,
)


_GETSTUDENTOUTINGSRESPONSE = _descriptor.Descriptor(
  name='GetStudentOutingsResponse',
  full_name='GetStudentOutingsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='GetStudentOutingsResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='GetStudentOutingsResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='GetStudentOutingsResponse.msg', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='outing', full_name='GetStudentOutingsResponse.outing', index=3,
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
  serialized_start=329,
  serialized_end=431,
)


_STUDENTOUTING = _descriptor.Descriptor(
  name='StudentOuting',
  full_name='StudentOuting',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='place', full_name='StudentOuting.place', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='reason', full_name='StudentOuting.reason', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='date', full_name='StudentOuting.date', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='StudentOuting.start_time', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='StudentOuting.end_time', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='situation', full_name='StudentOuting.situation', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='status', full_name='StudentOuting.status', index=6,
      number=7, type=9, cpp_type=9, label=1,
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
  serialized_start=434,
  serialized_end=567,
)


_GETOUTINGINFORMREQUEST = _descriptor.Descriptor(
  name='GetOutingInformRequest',
  full_name='GetOutingInformRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='GetOutingInformRequest.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='oid', full_name='GetOutingInformRequest.oid', index=1,
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
  serialized_start=569,
  serialized_end=620,
)


_GETOUTINGINFORMRESPONSE = _descriptor.Descriptor(
  name='GetOutingInformResponse',
  full_name='GetOutingInformResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='GetOutingInformResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='GetOutingInformResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='GetOutingInformResponse.msg', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='place', full_name='GetOutingInformResponse.place', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='reason', full_name='GetOutingInformResponse.reason', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='date', full_name='GetOutingInformResponse.date', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='GetOutingInformResponse.start_time', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='GetOutingInformResponse.end_time', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='o_situation', full_name='GetOutingInformResponse.o_situation', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='o_status', full_name='GetOutingInformResponse.o_status', index=9,
      number=10, type=9, cpp_type=9, label=1,
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
  serialized_start=623,
  serialized_end=813,
)


_GETCARDABOUTOUTINGREQUEST = _descriptor.Descriptor(
  name='GetCardAboutOutingRequest',
  full_name='GetCardAboutOutingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='GetCardAboutOutingRequest.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='oid', full_name='GetCardAboutOutingRequest.oid', index=1,
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
  serialized_start=815,
  serialized_end=869,
)


_GETCARDABOUTOUTINGRESPONSE = _descriptor.Descriptor(
  name='GetCardAboutOutingResponse',
  full_name='GetCardAboutOutingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='GetCardAboutOutingResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='GetCardAboutOutingResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='GetCardAboutOutingResponse.msg', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='place', full_name='GetCardAboutOutingResponse.place', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='date', full_name='GetCardAboutOutingResponse.date', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='GetCardAboutOutingResponse.start_time', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='GetCardAboutOutingResponse.end_time', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='o_status', full_name='GetCardAboutOutingResponse.o_status', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='name', full_name='GetCardAboutOutingResponse.name', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='grade', full_name='GetCardAboutOutingResponse.grade', index=9,
      number=10, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='class_', full_name='GetCardAboutOutingResponse.class_', index=10,
      number=11, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='number', full_name='GetCardAboutOutingResponse.number', index=11,
      number=12, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='profile_image_uri', full_name='GetCardAboutOutingResponse.profile_image_uri', index=12,
      number=13, type=9, cpp_type=9, label=1,
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
  serialized_start=872,
  serialized_end=1116,
)


_GOOUTREQUEST = _descriptor.Descriptor(
  name='GoOutRequest',
  full_name='GoOutRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='uuid', full_name='GoOutRequest.uuid', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='oid', full_name='GoOutRequest.oid', index=1,
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
  serialized_start=1118,
  serialized_end=1159,
)


_GOOUTRESPONSE = _descriptor.Descriptor(
  name='GoOutResponse',
  full_name='GoOutResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='GoOutResponse.status', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='code', full_name='GoOutResponse.code', index=1,
      number=2, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='msg', full_name='GoOutResponse.msg', index=2,
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
  serialized_start=1161,
  serialized_end=1219,
)

_GETSTUDENTOUTINGSRESPONSE.fields_by_name['outing'].message_type = _STUDENTOUTING
DESCRIPTOR.message_types_by_name['CreateOutingRequest'] = _CREATEOUTINGREQUEST
DESCRIPTOR.message_types_by_name['CreateOutingResponse'] = _CREATEOUTINGRESPONSE
DESCRIPTOR.message_types_by_name['GetStudentOutingsRequest'] = _GETSTUDENTOUTINGSREQUEST
DESCRIPTOR.message_types_by_name['GetStudentOutingsResponse'] = _GETSTUDENTOUTINGSRESPONSE
DESCRIPTOR.message_types_by_name['StudentOuting'] = _STUDENTOUTING
DESCRIPTOR.message_types_by_name['GetOutingInformRequest'] = _GETOUTINGINFORMREQUEST
DESCRIPTOR.message_types_by_name['GetOutingInformResponse'] = _GETOUTINGINFORMRESPONSE
DESCRIPTOR.message_types_by_name['GetCardAboutOutingRequest'] = _GETCARDABOUTOUTINGREQUEST
DESCRIPTOR.message_types_by_name['GetCardAboutOutingResponse'] = _GETCARDABOUTOUTINGRESPONSE
DESCRIPTOR.message_types_by_name['GoOutRequest'] = _GOOUTREQUEST
DESCRIPTOR.message_types_by_name['GoOutResponse'] = _GOOUTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateOutingRequest = _reflection.GeneratedProtocolMessageType('CreateOutingRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEOUTINGREQUEST,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:CreateOutingRequest)
  })
_sym_db.RegisterMessage(CreateOutingRequest)

CreateOutingResponse = _reflection.GeneratedProtocolMessageType('CreateOutingResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEOUTINGRESPONSE,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:CreateOutingResponse)
  })
_sym_db.RegisterMessage(CreateOutingResponse)

GetStudentOutingsRequest = _reflection.GeneratedProtocolMessageType('GetStudentOutingsRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETSTUDENTOUTINGSREQUEST,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GetStudentOutingsRequest)
  })
_sym_db.RegisterMessage(GetStudentOutingsRequest)

GetStudentOutingsResponse = _reflection.GeneratedProtocolMessageType('GetStudentOutingsResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETSTUDENTOUTINGSRESPONSE,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GetStudentOutingsResponse)
  })
_sym_db.RegisterMessage(GetStudentOutingsResponse)

StudentOuting = _reflection.GeneratedProtocolMessageType('StudentOuting', (_message.Message,), {
  'DESCRIPTOR' : _STUDENTOUTING,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:StudentOuting)
  })
_sym_db.RegisterMessage(StudentOuting)

GetOutingInformRequest = _reflection.GeneratedProtocolMessageType('GetOutingInformRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETOUTINGINFORMREQUEST,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GetOutingInformRequest)
  })
_sym_db.RegisterMessage(GetOutingInformRequest)

GetOutingInformResponse = _reflection.GeneratedProtocolMessageType('GetOutingInformResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETOUTINGINFORMRESPONSE,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GetOutingInformResponse)
  })
_sym_db.RegisterMessage(GetOutingInformResponse)

GetCardAboutOutingRequest = _reflection.GeneratedProtocolMessageType('GetCardAboutOutingRequest', (_message.Message,), {
  'DESCRIPTOR' : _GETCARDABOUTOUTINGREQUEST,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GetCardAboutOutingRequest)
  })
_sym_db.RegisterMessage(GetCardAboutOutingRequest)

GetCardAboutOutingResponse = _reflection.GeneratedProtocolMessageType('GetCardAboutOutingResponse', (_message.Message,), {
  'DESCRIPTOR' : _GETCARDABOUTOUTINGRESPONSE,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GetCardAboutOutingResponse)
  })
_sym_db.RegisterMessage(GetCardAboutOutingResponse)

GoOutRequest = _reflection.GeneratedProtocolMessageType('GoOutRequest', (_message.Message,), {
  'DESCRIPTOR' : _GOOUTREQUEST,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GoOutRequest)
  })
_sym_db.RegisterMessage(GoOutRequest)

GoOutResponse = _reflection.GeneratedProtocolMessageType('GoOutResponse', (_message.Message,), {
  'DESCRIPTOR' : _GOOUTRESPONSE,
  '__module__' : 'outing_student_pb2'
  # @@protoc_insertion_point(class_scope:GoOutResponse)
  })
_sym_db.RegisterMessage(GoOutResponse)



_OUTINGSTUDENT = _descriptor.ServiceDescriptor(
  name='OutingStudent',
  full_name='OutingStudent',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=1222,
  serialized_end=1621,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateOuting',
    full_name='OutingStudent.CreateOuting',
    index=0,
    containing_service=None,
    input_type=_CREATEOUTINGREQUEST,
    output_type=_CREATEOUTINGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetStudentOutings',
    full_name='OutingStudent.GetStudentOutings',
    index=1,
    containing_service=None,
    input_type=_GETSTUDENTOUTINGSREQUEST,
    output_type=_GETSTUDENTOUTINGSRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetOutingInform',
    full_name='OutingStudent.GetOutingInform',
    index=2,
    containing_service=None,
    input_type=_GETOUTINGINFORMREQUEST,
    output_type=_GETOUTINGINFORMRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GetCardAboutOuting',
    full_name='OutingStudent.GetCardAboutOuting',
    index=3,
    containing_service=None,
    input_type=_GETCARDABOUTOUTINGREQUEST,
    output_type=_GETCARDABOUTOUTINGRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='GoOut',
    full_name='OutingStudent.GoOut',
    index=4,
    containing_service=None,
    input_type=_GOOUTREQUEST,
    output_type=_GOOUTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='FinishGoOut',
    full_name='OutingStudent.FinishGoOut',
    index=5,
    containing_service=None,
    input_type=_GOOUTREQUEST,
    output_type=_GOOUTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_OUTINGSTUDENT)

DESCRIPTOR.services_by_name['OutingStudent'] = _OUTINGSTUDENT

# @@protoc_insertion_point(module_scope)
