# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\nauth.proto\x12\x04\x61uth\"@\n\x0fRegisterRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\"#\n\x10RegisterResponse\x12\x0f\n\x07user_id\x18\x01 \x01(\t\"/\n\x0cLoginRequest\x12\r\n\x05\x65mail\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x1e\n\rLoginResponse\x12\r\n\x05token\x18\x01 \x01(\t\"$\n\x13IsTokenValidRequest\x12\r\n\x05token\x18\x01 \x01(\t\"+\n\x14IsTokenValidResponse\x12\x13\n\x0btoken_valid\x18\x01 \x01(\x08\x32\xba\x01\n\x04\x41uth\x12\x39\n\x08Register\x12\x15.auth.RegisterRequest\x1a\x16.auth.RegisterResponse\x12\x30\n\x05Login\x12\x12.auth.LoginRequest\x1a\x13.auth.LoginResponse\x12\x45\n\x0cIsTokenValid\x12\x19.auth.IsTokenValidRequest\x1a\x1a.auth.IsTokenValidResponseB\x18Z\x16kazali.auth.v1.;authv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'auth_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\026kazali.auth.v1.;authv1'
  _globals['_REGISTERREQUEST']._serialized_start=20
  _globals['_REGISTERREQUEST']._serialized_end=84
  _globals['_REGISTERRESPONSE']._serialized_start=86
  _globals['_REGISTERRESPONSE']._serialized_end=121
  _globals['_LOGINREQUEST']._serialized_start=123
  _globals['_LOGINREQUEST']._serialized_end=170
  _globals['_LOGINRESPONSE']._serialized_start=172
  _globals['_LOGINRESPONSE']._serialized_end=202
  _globals['_ISTOKENVALIDREQUEST']._serialized_start=204
  _globals['_ISTOKENVALIDREQUEST']._serialized_end=240
  _globals['_ISTOKENVALIDRESPONSE']._serialized_start=242
  _globals['_ISTOKENVALIDRESPONSE']._serialized_end=285
  _globals['_AUTH']._serialized_start=288
  _globals['_AUTH']._serialized_end=474
# @@protoc_insertion_point(module_scope)
