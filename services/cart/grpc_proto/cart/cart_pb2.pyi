import product_pb2 as _product_pb2
from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetCartRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class GetCartResponse(_message.Message):
    __slots__ = ("products",)
    PRODUCTS_FIELD_NUMBER: _ClassVar[int]
    products: _containers.RepeatedCompositeFieldContainer[_product_pb2.ProductEntry]
    def __init__(self, products: _Optional[_Iterable[_Union[_product_pb2.ProductEntry, _Mapping]]] = ...) -> None: ...

class AddToCartRequest(_message.Message):
    __slots__ = ("product_id", "quantity")
    PRODUCT_ID_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    product_id: str
    quantity: int
    def __init__(self, product_id: _Optional[str] = ..., quantity: _Optional[int] = ...) -> None: ...

class AddToCartResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class DeleteFromCartRequest(_message.Message):
    __slots__ = ("product_id",)
    PRODUCT_ID_FIELD_NUMBER: _ClassVar[int]
    product_id: str
    def __init__(self, product_id: _Optional[str] = ...) -> None: ...

class DeleteFromCartResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ClearCartRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ClearCartResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class UpdateCartRequest(_message.Message):
    __slots__ = ("product_id", "quantity")
    PRODUCT_ID_FIELD_NUMBER: _ClassVar[int]
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    product_id: str
    quantity: int
    def __init__(self, product_id: _Optional[str] = ..., quantity: _Optional[int] = ...) -> None: ...

class UpdateCartResponse(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...
