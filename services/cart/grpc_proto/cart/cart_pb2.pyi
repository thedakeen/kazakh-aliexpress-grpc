from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class GetCartRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ProductCartEntry(_message.Message):
    __slots__ = ("quantity", "item_id")
    QUANTITY_FIELD_NUMBER: _ClassVar[int]
    ITEM_ID_FIELD_NUMBER: _ClassVar[int]
    quantity: int
    item_id: str
    def __init__(self, quantity: _Optional[int] = ..., item_id: _Optional[str] = ...) -> None: ...

class GetCartResponse(_message.Message):
    __slots__ = ("cart",)
    CART_FIELD_NUMBER: _ClassVar[int]
    cart: _containers.RepeatedCompositeFieldContainer[ProductCartEntry]
    def __init__(self, cart: _Optional[_Iterable[_Union[ProductCartEntry, _Mapping]]] = ...) -> None: ...

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
