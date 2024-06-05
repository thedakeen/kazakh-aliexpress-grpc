# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: cart.proto
# Protobuf Python Version: 5.26.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


import product_pb2 as product__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\ncart.proto\x12\x04\x63\x61rt\x1a\rproduct.proto\"\x10\n\x0eGetCartRequest\":\n\x0fGetCartResponse\x12\'\n\x08products\x18\x01 \x03(\x0b\x32\x15.product.ProductEntry\"8\n\x10\x41\x64\x64ToCartRequest\x12\x12\n\nproduct_id\x18\x02 \x01(\t\x12\x10\n\x08quantity\x18\x03 \x01(\x05\"\x13\n\x11\x41\x64\x64ToCartResponse\"+\n\x15\x44\x65leteFromCartRequest\x12\x12\n\nproduct_id\x18\x02 \x01(\t\"\x18\n\x16\x44\x65leteFromCartResponse\"\x12\n\x10\x43learCartRequest\"\x13\n\x11\x43learCartResponse\"9\n\x11UpdateCartRequest\x12\x12\n\nproduct_id\x18\x02 \x01(\t\x12\x10\n\x08quantity\x18\x03 \x01(\x05\"\x14\n\x12UpdateCartResponse2\xc8\x02\n\x04\x43\x61rt\x12\x36\n\x07GetCart\x12\x14.cart.GetCartRequest\x1a\x15.cart.GetCartResponse\x12<\n\tAddToCart\x12\x16.cart.AddToCartRequest\x1a\x17.cart.AddToCartResponse\x12K\n\x0e\x44\x65leteFromCart\x12\x1b.cart.DeleteFromCartRequest\x1a\x1c.cart.DeleteFromCartResponse\x12<\n\tClearCart\x12\x16.cart.ClearCartRequest\x1a\x17.cart.ClearCartResponse\x12?\n\nUpdateCart\x12\x17.cart.UpdateCartRequest\x1a\x18.cart.UpdateCartResponseb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'cart_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  DESCRIPTOR._loaded_options = None
  _globals['_GETCARTREQUEST']._serialized_start=35
  _globals['_GETCARTREQUEST']._serialized_end=51
  _globals['_GETCARTRESPONSE']._serialized_start=53
  _globals['_GETCARTRESPONSE']._serialized_end=111
  _globals['_ADDTOCARTREQUEST']._serialized_start=113
  _globals['_ADDTOCARTREQUEST']._serialized_end=169
  _globals['_ADDTOCARTRESPONSE']._serialized_start=171
  _globals['_ADDTOCARTRESPONSE']._serialized_end=190
  _globals['_DELETEFROMCARTREQUEST']._serialized_start=192
  _globals['_DELETEFROMCARTREQUEST']._serialized_end=235
  _globals['_DELETEFROMCARTRESPONSE']._serialized_start=237
  _globals['_DELETEFROMCARTRESPONSE']._serialized_end=261
  _globals['_CLEARCARTREQUEST']._serialized_start=263
  _globals['_CLEARCARTREQUEST']._serialized_end=281
  _globals['_CLEARCARTRESPONSE']._serialized_start=283
  _globals['_CLEARCARTRESPONSE']._serialized_end=302
  _globals['_UPDATECARTREQUEST']._serialized_start=304
  _globals['_UPDATECARTREQUEST']._serialized_end=361
  _globals['_UPDATECARTRESPONSE']._serialized_start=363
  _globals['_UPDATECARTRESPONSE']._serialized_end=383
  _globals['_CART']._serialized_start=386
  _globals['_CART']._serialized_end=714
# @@protoc_insertion_point(module_scope)
