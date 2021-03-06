# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth-admin.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='auth-admin.proto',
  package='',
  syntax='proto3',
  serialized_options=b'Z\013.;authproto',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n\x10\x61uth-admin.proto\"\xc8\x01\n\x17\x43reateNewStudentRequest\x12\x0c\n\x04UUID\x18\x01 \x01(\t\x12\x11\n\tStudentID\x18\x02 \x01(\t\x12\x11\n\tStudentPW\x18\x03 \x01(\t\x12\x12\n\nParentUUID\x18\x04 \x01(\t\x12\r\n\x05Grade\x18\x05 \x01(\r\x12\r\n\x05Group\x18\x06 \x01(\r\x12\x15\n\rStudentNumber\x18\x07 \x01(\r\x12\x0c\n\x04Name\x18\x08 \x01(\t\x12\x13\n\x0bPhoneNumber\x18\t \x01(\t\x12\r\n\x05Image\x18\n \x01(\x0c\"e\n\x18\x43reateNewStudentResponse\x12\x0e\n\x06Status\x18\x01 \x01(\r\x12\x0c\n\x04\x43ode\x18\x02 \x01(\x11\x12\x0f\n\x07Message\x18\x03 \x01(\t\x12\x1a\n\x12\x43reatedStudentUUID\x18\x04 \x01(\t\"\x8e\x01\n\x17\x43reateNewTeacherRequest\x12\x0c\n\x04UUID\x18\x01 \x01(\t\x12\x11\n\tTeacherID\x18\x02 \x01(\t\x12\x11\n\tTeacherPW\x18\x03 \x01(\t\x12\r\n\x05Grade\x18\x04 \x01(\r\x12\r\n\x05Group\x18\x05 \x01(\r\x12\x0c\n\x04Name\x18\x07 \x01(\t\x12\x13\n\x0bPhoneNumber\x18\x08 \x01(\t\"e\n\x18\x43reateNewTeacherResponse\x12\x0e\n\x06Status\x18\x01 \x01(\r\x12\x0c\n\x04\x43ode\x18\x02 \x01(\x11\x12\x0f\n\x07Message\x18\x03 \x01(\t\x12\x1a\n\x12\x43reatedTeacherUUID\x18\x04 \x01(\t\"m\n\x16\x43reateNewParentRequest\x12\x0c\n\x04UUID\x18\x01 \x01(\t\x12\x10\n\x08ParentID\x18\x02 \x01(\t\x12\x10\n\x08ParentPW\x18\x03 \x01(\t\x12\x0c\n\x04Name\x18\x05 \x01(\t\x12\x13\n\x0bPhoneNumber\x18\x06 \x01(\t\"c\n\x17\x43reateNewParentResponse\x12\x0e\n\x06Status\x18\x01 \x01(\r\x12\x0c\n\x04\x43ode\x18\x02 \x01(\x11\x12\x0f\n\x07Message\x18\x03 \x01(\t\x12\x19\n\x11\x43reatedParentUUID\x18\x04 \x01(\t\"9\n\x15LoginAdminAuthRequest\x12\x0f\n\x07\x41\x64minID\x18\x01 \x01(\t\x12\x0f\n\x07\x41\x64minPW\x18\x02 \x01(\t\"w\n\x16LoginAdminAuthResponse\x12\x0e\n\x06Status\x18\x01 \x01(\r\x12\x0c\n\x04\x43ode\x18\x02 \x01(\x11\x12\x0f\n\x07Message\x18\x03 \x01(\t\x12\x13\n\x0b\x41\x63\x63\x65ssToken\x18\x04 \x01(\t\x12\x19\n\x11LoggedInAdminUUID\x18\x05 \x01(\t2\xae\x02\n\tAuthAdmin\x12I\n\x10\x43reateNewStudent\x12\x18.CreateNewStudentRequest\x1a\x19.CreateNewStudentResponse\"\x00\x12I\n\x10\x43reateNewTeacher\x12\x18.CreateNewTeacherRequest\x1a\x19.CreateNewTeacherResponse\"\x00\x12\x46\n\x0f\x43reateNewParent\x12\x17.CreateNewParentRequest\x1a\x18.CreateNewParentResponse\"\x00\x12\x43\n\x0eLoginAdminAuth\x12\x16.LoginAdminAuthRequest\x1a\x17.LoginAdminAuthResponse\"\x00\x42\rZ\x0b.;authprotob\x06proto3'
)




_CREATENEWSTUDENTREQUEST = _descriptor.Descriptor(
  name='CreateNewStudentRequest',
  full_name='CreateNewStudentRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='UUID', full_name='CreateNewStudentRequest.UUID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='StudentID', full_name='CreateNewStudentRequest.StudentID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='StudentPW', full_name='CreateNewStudentRequest.StudentPW', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ParentUUID', full_name='CreateNewStudentRequest.ParentUUID', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Grade', full_name='CreateNewStudentRequest.Grade', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Group', full_name='CreateNewStudentRequest.Group', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='StudentNumber', full_name='CreateNewStudentRequest.StudentNumber', index=6,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Name', full_name='CreateNewStudentRequest.Name', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='PhoneNumber', full_name='CreateNewStudentRequest.PhoneNumber', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Image', full_name='CreateNewStudentRequest.Image', index=9,
      number=10, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=b"",
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
  serialized_start=21,
  serialized_end=221,
)


_CREATENEWSTUDENTRESPONSE = _descriptor.Descriptor(
  name='CreateNewStudentResponse',
  full_name='CreateNewStudentResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Status', full_name='CreateNewStudentResponse.Status', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Code', full_name='CreateNewStudentResponse.Code', index=1,
      number=2, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Message', full_name='CreateNewStudentResponse.Message', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CreatedStudentUUID', full_name='CreateNewStudentResponse.CreatedStudentUUID', index=3,
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
  serialized_start=223,
  serialized_end=324,
)


_CREATENEWTEACHERREQUEST = _descriptor.Descriptor(
  name='CreateNewTeacherRequest',
  full_name='CreateNewTeacherRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='UUID', full_name='CreateNewTeacherRequest.UUID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='TeacherID', full_name='CreateNewTeacherRequest.TeacherID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='TeacherPW', full_name='CreateNewTeacherRequest.TeacherPW', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Grade', full_name='CreateNewTeacherRequest.Grade', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Group', full_name='CreateNewTeacherRequest.Group', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Name', full_name='CreateNewTeacherRequest.Name', index=5,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='PhoneNumber', full_name='CreateNewTeacherRequest.PhoneNumber', index=6,
      number=8, type=9, cpp_type=9, label=1,
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
  serialized_start=327,
  serialized_end=469,
)


_CREATENEWTEACHERRESPONSE = _descriptor.Descriptor(
  name='CreateNewTeacherResponse',
  full_name='CreateNewTeacherResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Status', full_name='CreateNewTeacherResponse.Status', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Code', full_name='CreateNewTeacherResponse.Code', index=1,
      number=2, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Message', full_name='CreateNewTeacherResponse.Message', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CreatedTeacherUUID', full_name='CreateNewTeacherResponse.CreatedTeacherUUID', index=3,
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
  serialized_start=471,
  serialized_end=572,
)


_CREATENEWPARENTREQUEST = _descriptor.Descriptor(
  name='CreateNewParentRequest',
  full_name='CreateNewParentRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='UUID', full_name='CreateNewParentRequest.UUID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ParentID', full_name='CreateNewParentRequest.ParentID', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='ParentPW', full_name='CreateNewParentRequest.ParentPW', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Name', full_name='CreateNewParentRequest.Name', index=3,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='PhoneNumber', full_name='CreateNewParentRequest.PhoneNumber', index=4,
      number=6, type=9, cpp_type=9, label=1,
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
  serialized_start=574,
  serialized_end=683,
)


_CREATENEWPARENTRESPONSE = _descriptor.Descriptor(
  name='CreateNewParentResponse',
  full_name='CreateNewParentResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Status', full_name='CreateNewParentResponse.Status', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Code', full_name='CreateNewParentResponse.Code', index=1,
      number=2, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Message', full_name='CreateNewParentResponse.Message', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='CreatedParentUUID', full_name='CreateNewParentResponse.CreatedParentUUID', index=3,
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
  serialized_start=685,
  serialized_end=784,
)


_LOGINADMINAUTHREQUEST = _descriptor.Descriptor(
  name='LoginAdminAuthRequest',
  full_name='LoginAdminAuthRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='AdminID', full_name='LoginAdminAuthRequest.AdminID', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='AdminPW', full_name='LoginAdminAuthRequest.AdminPW', index=1,
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
  serialized_start=786,
  serialized_end=843,
)


_LOGINADMINAUTHRESPONSE = _descriptor.Descriptor(
  name='LoginAdminAuthResponse',
  full_name='LoginAdminAuthResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='Status', full_name='LoginAdminAuthResponse.Status', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Code', full_name='LoginAdminAuthResponse.Code', index=1,
      number=2, type=17, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='Message', full_name='LoginAdminAuthResponse.Message', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='AccessToken', full_name='LoginAdminAuthResponse.AccessToken', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='LoggedInAdminUUID', full_name='LoginAdminAuthResponse.LoggedInAdminUUID', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=845,
  serialized_end=964,
)

DESCRIPTOR.message_types_by_name['CreateNewStudentRequest'] = _CREATENEWSTUDENTREQUEST
DESCRIPTOR.message_types_by_name['CreateNewStudentResponse'] = _CREATENEWSTUDENTRESPONSE
DESCRIPTOR.message_types_by_name['CreateNewTeacherRequest'] = _CREATENEWTEACHERREQUEST
DESCRIPTOR.message_types_by_name['CreateNewTeacherResponse'] = _CREATENEWTEACHERRESPONSE
DESCRIPTOR.message_types_by_name['CreateNewParentRequest'] = _CREATENEWPARENTREQUEST
DESCRIPTOR.message_types_by_name['CreateNewParentResponse'] = _CREATENEWPARENTRESPONSE
DESCRIPTOR.message_types_by_name['LoginAdminAuthRequest'] = _LOGINADMINAUTHREQUEST
DESCRIPTOR.message_types_by_name['LoginAdminAuthResponse'] = _LOGINADMINAUTHRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateNewStudentRequest = _reflection.GeneratedProtocolMessageType('CreateNewStudentRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATENEWSTUDENTREQUEST,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:CreateNewStudentRequest)
  })
_sym_db.RegisterMessage(CreateNewStudentRequest)

CreateNewStudentResponse = _reflection.GeneratedProtocolMessageType('CreateNewStudentResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATENEWSTUDENTRESPONSE,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:CreateNewStudentResponse)
  })
_sym_db.RegisterMessage(CreateNewStudentResponse)

CreateNewTeacherRequest = _reflection.GeneratedProtocolMessageType('CreateNewTeacherRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATENEWTEACHERREQUEST,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:CreateNewTeacherRequest)
  })
_sym_db.RegisterMessage(CreateNewTeacherRequest)

CreateNewTeacherResponse = _reflection.GeneratedProtocolMessageType('CreateNewTeacherResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATENEWTEACHERRESPONSE,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:CreateNewTeacherResponse)
  })
_sym_db.RegisterMessage(CreateNewTeacherResponse)

CreateNewParentRequest = _reflection.GeneratedProtocolMessageType('CreateNewParentRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATENEWPARENTREQUEST,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:CreateNewParentRequest)
  })
_sym_db.RegisterMessage(CreateNewParentRequest)

CreateNewParentResponse = _reflection.GeneratedProtocolMessageType('CreateNewParentResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATENEWPARENTRESPONSE,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:CreateNewParentResponse)
  })
_sym_db.RegisterMessage(CreateNewParentResponse)

LoginAdminAuthRequest = _reflection.GeneratedProtocolMessageType('LoginAdminAuthRequest', (_message.Message,), {
  'DESCRIPTOR' : _LOGINADMINAUTHREQUEST,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:LoginAdminAuthRequest)
  })
_sym_db.RegisterMessage(LoginAdminAuthRequest)

LoginAdminAuthResponse = _reflection.GeneratedProtocolMessageType('LoginAdminAuthResponse', (_message.Message,), {
  'DESCRIPTOR' : _LOGINADMINAUTHRESPONSE,
  '__module__' : 'auth_admin_pb2'
  # @@protoc_insertion_point(class_scope:LoginAdminAuthResponse)
  })
_sym_db.RegisterMessage(LoginAdminAuthResponse)


DESCRIPTOR._options = None

_AUTHADMIN = _descriptor.ServiceDescriptor(
  name='AuthAdmin',
  full_name='AuthAdmin',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  create_key=_descriptor._internal_create_key,
  serialized_start=967,
  serialized_end=1269,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateNewStudent',
    full_name='AuthAdmin.CreateNewStudent',
    index=0,
    containing_service=None,
    input_type=_CREATENEWSTUDENTREQUEST,
    output_type=_CREATENEWSTUDENTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateNewTeacher',
    full_name='AuthAdmin.CreateNewTeacher',
    index=1,
    containing_service=None,
    input_type=_CREATENEWTEACHERREQUEST,
    output_type=_CREATENEWTEACHERRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='CreateNewParent',
    full_name='AuthAdmin.CreateNewParent',
    index=2,
    containing_service=None,
    input_type=_CREATENEWPARENTREQUEST,
    output_type=_CREATENEWPARENTRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
  _descriptor.MethodDescriptor(
    name='LoginAdminAuth',
    full_name='AuthAdmin.LoginAdminAuth',
    index=3,
    containing_service=None,
    input_type=_LOGINADMINAUTHREQUEST,
    output_type=_LOGINADMINAUTHRESPONSE,
    serialized_options=None,
    create_key=_descriptor._internal_create_key,
  ),
])
_sym_db.RegisterServiceDescriptor(_AUTHADMIN)

DESCRIPTOR.services_by_name['AuthAdmin'] = _AUTHADMIN

# @@protoc_insertion_point(module_scope)
