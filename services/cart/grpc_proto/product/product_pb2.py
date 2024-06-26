# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: product.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rproduct.proto\x12\x07product\"$\n\x08\x43\x61tegory\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"7\n\x0bProductInfo\x12\x12\n\ninfo_title\x18\x01 \x01(\t\x12\x14\n\x0cinfo_content\x18\x02 \x01(\t\"@\n\x0eProductVariant\x12\x15\n\rvariant_title\x18\x01 \x01(\t\x12\x17\n\x0fvariant_options\x18\x02 \x03(\t\"\xc1\x01\n\x0cProductEntry\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\r\n\x05price\x18\x03 \x01(\x02\x12%\n\ncategories\x18\x04 \x03(\x0b\x32\x11.product.Category\x12\x12\n\nimage_urls\x18\x05 \x03(\t\x12#\n\x05infos\x18\x06 \x03(\x0b\x32\x14.product.ProductInfo\x12(\n\x07options\x18\x07 \x03(\x0b\x32\x17.product.ProductVariant\"\x11\n\x0f\x43\x61tegoryRequest\"9\n\x10\x43\x61tegoryResponse\x12%\n\ncategories\x18\x01 \x03(\x0b\x32\x11.product.Category\"<\n\x15\x43reateCategoryRequest\x12#\n\x08\x63\x61tegory\x18\x01 \x01(\x0b\x32\x11.product.Category\"=\n\x16\x43reateCategoryResponse\x12#\n\x08\x63\x61tegory\x18\x01 \x01(\x0b\x32\x11.product.Category\"1\n\x15UpdateCategoryRequest\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"#\n\x15\x44\x65leteCategoryRequest\x12\n\n\x02id\x18\x01 \x01(\t\"\x08\n\x06\x44\x65lete\"\x1c\n\x0eProductRequest\x12\n\n\x02id\x18\x01 \x01(\t\"9\n\x0fProductResponse\x12&\n\x07product\x18\x01 \x01(\x0b\x32\x15.product.ProductEntry\"0\n\x19ProductsByCategoryRequest\x12\x13\n\x0b\x63\x61tegory_id\x18\x01 \x01(\t\"E\n\x1aProductsByCategoryResponse\x12\'\n\x08products\x18\x01 \x03(\x0b\x32\x15.product.ProductEntry2\xbf\x02\n\x07Product\x12\x41\n\nCategories\x12\x18.product.CategoryRequest\x1a\x19.product.CategoryResponse\x12Q\n\x0e\x43reateCategory\x12\x1e.product.CreateCategoryRequest\x1a\x1f.product.CreateCategoryResponse\x12]\n\x12ProductsByCategory\x12\".product.ProductsByCategoryRequest\x1a#.product.ProductsByCategoryResponse\x12?\n\nGetProduct\x12\x17.product.ProductRequest\x1a\x18.product.ProductResponseB\x1dZ\x1bkazali.product.v1;productv1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'product_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\033kazali.product.v1;productv1'
  _globals['_CATEGORY']._serialized_start=26
  _globals['_CATEGORY']._serialized_end=62
  _globals['_PRODUCTINFO']._serialized_start=64
  _globals['_PRODUCTINFO']._serialized_end=119
  _globals['_PRODUCTVARIANT']._serialized_start=121
  _globals['_PRODUCTVARIANT']._serialized_end=185
  _globals['_PRODUCTENTRY']._serialized_start=188
  _globals['_PRODUCTENTRY']._serialized_end=381
  _globals['_CATEGORYREQUEST']._serialized_start=383
  _globals['_CATEGORYREQUEST']._serialized_end=400
  _globals['_CATEGORYRESPONSE']._serialized_start=402
  _globals['_CATEGORYRESPONSE']._serialized_end=459
  _globals['_CREATECATEGORYREQUEST']._serialized_start=461
  _globals['_CREATECATEGORYREQUEST']._serialized_end=521
  _globals['_CREATECATEGORYRESPONSE']._serialized_start=523
  _globals['_CREATECATEGORYRESPONSE']._serialized_end=584
  _globals['_UPDATECATEGORYREQUEST']._serialized_start=586
  _globals['_UPDATECATEGORYREQUEST']._serialized_end=635
  _globals['_DELETECATEGORYREQUEST']._serialized_start=637
  _globals['_DELETECATEGORYREQUEST']._serialized_end=672
  _globals['_DELETE']._serialized_start=674
  _globals['_DELETE']._serialized_end=682
  _globals['_PRODUCTREQUEST']._serialized_start=684
  _globals['_PRODUCTREQUEST']._serialized_end=712
  _globals['_PRODUCTRESPONSE']._serialized_start=714
  _globals['_PRODUCTRESPONSE']._serialized_end=771
  _globals['_PRODUCTSBYCATEGORYREQUEST']._serialized_start=773
  _globals['_PRODUCTSBYCATEGORYREQUEST']._serialized_end=821
  _globals['_PRODUCTSBYCATEGORYRESPONSE']._serialized_start=823
  _globals['_PRODUCTSBYCATEGORYRESPONSE']._serialized_end=892
  _globals['_PRODUCT']._serialized_start=895
  _globals['_PRODUCT']._serialized_end=1214
# @@protoc_insertion_point(module_scope)
