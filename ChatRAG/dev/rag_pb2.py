# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: rag.proto
# Protobuf Python Version: 5.27.2
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    27,
    2,
    '',
    'rag.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\trag.proto\x12\x03rag\"\x1d\n\nRAGRequest\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t\"\x1e\n\x0bRAGResponse\x12\x0f\n\x07\x63ontent\x18\x01 \x01(\t2;\n\nRAGService\x12-\n\x08Generate\x12\x0f.rag.RAGRequest\x1a\x10.rag.RAGResponseB\x16Z\x14MMaxaMM.rag.v1;ragv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'rag_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\024MMaxaMM.rag.v1;ragv1'
  _globals['_RAGREQUEST']._serialized_start=18
  _globals['_RAGREQUEST']._serialized_end=47
  _globals['_RAGRESPONSE']._serialized_start=49
  _globals['_RAGRESPONSE']._serialized_end=79
  _globals['_RAGSERVICE']._serialized_start=81
  _globals['_RAGSERVICE']._serialized_end=140
# @@protoc_insertion_point(module_scope)