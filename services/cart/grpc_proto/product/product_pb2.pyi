from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Category(_message.Message):
    __slots__ = ("id", "name")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class ProductInfo(_message.Message):
    __slots__ = ("info_title", "info_content")
    INFO_TITLE_FIELD_NUMBER: _ClassVar[int]
    INFO_CONTENT_FIELD_NUMBER: _ClassVar[int]
    info_title: str
    info_content: str
    def __init__(self, info_title: _Optional[str] = ..., info_content: _Optional[str] = ...) -> None: ...

class ProductVariant(_message.Message):
    __slots__ = ("variant_title", "variant_options")
    VARIANT_TITLE_FIELD_NUMBER: _ClassVar[int]
    VARIANT_OPTIONS_FIELD_NUMBER: _ClassVar[int]
    variant_title: str
    variant_options: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, variant_title: _Optional[str] = ..., variant_options: _Optional[_Iterable[str]] = ...) -> None: ...

class ProductEntry(_message.Message):
    __slots__ = ("id", "name", "price", "categories", "image_urls", "infos", "options")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    PRICE_FIELD_NUMBER: _ClassVar[int]
    CATEGORIES_FIELD_NUMBER: _ClassVar[int]
    IMAGE_URLS_FIELD_NUMBER: _ClassVar[int]
    INFOS_FIELD_NUMBER: _ClassVar[int]
    OPTIONS_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    price: float
    categories: _containers.RepeatedCompositeFieldContainer[Category]
    image_urls: _containers.RepeatedScalarFieldContainer[str]
    infos: _containers.RepeatedCompositeFieldContainer[ProductInfo]
    options: _containers.RepeatedCompositeFieldContainer[ProductVariant]
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ..., price: _Optional[float] = ..., categories: _Optional[_Iterable[_Union[Category, _Mapping]]] = ..., image_urls: _Optional[_Iterable[str]] = ..., infos: _Optional[_Iterable[_Union[ProductInfo, _Mapping]]] = ..., options: _Optional[_Iterable[_Union[ProductVariant, _Mapping]]] = ...) -> None: ...

class CategoryRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class CategoryResponse(_message.Message):
    __slots__ = ("categories",)
    CATEGORIES_FIELD_NUMBER: _ClassVar[int]
    categories: _containers.RepeatedCompositeFieldContainer[Category]
    def __init__(self, categories: _Optional[_Iterable[_Union[Category, _Mapping]]] = ...) -> None: ...

class CreateCategoryRequest(_message.Message):
    __slots__ = ("category",)
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    category: Category
    def __init__(self, category: _Optional[_Union[Category, _Mapping]] = ...) -> None: ...

class CreateCategoryResponse(_message.Message):
    __slots__ = ("category",)
    CATEGORY_FIELD_NUMBER: _ClassVar[int]
    category: Category
    def __init__(self, category: _Optional[_Union[Category, _Mapping]] = ...) -> None: ...

class UpdateCategoryRequest(_message.Message):
    __slots__ = ("id", "name")
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    id: str
    name: str
    def __init__(self, id: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class DeleteCategoryRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class Delete(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class ProductRequest(_message.Message):
    __slots__ = ("id",)
    ID_FIELD_NUMBER: _ClassVar[int]
    id: str
    def __init__(self, id: _Optional[str] = ...) -> None: ...

class ProductResponse(_message.Message):
    __slots__ = ("product",)
    PRODUCT_FIELD_NUMBER: _ClassVar[int]
    product: ProductEntry
    def __init__(self, product: _Optional[_Union[ProductEntry, _Mapping]] = ...) -> None: ...

class ProductsByCategoryRequest(_message.Message):
    __slots__ = ("category_id",)
    CATEGORY_ID_FIELD_NUMBER: _ClassVar[int]
    category_id: str
    def __init__(self, category_id: _Optional[str] = ...) -> None: ...

class ProductsByCategoryResponse(_message.Message):
    __slots__ = ("products",)
    PRODUCTS_FIELD_NUMBER: _ClassVar[int]
    products: _containers.RepeatedCompositeFieldContainer[ProductEntry]
    def __init__(self, products: _Optional[_Iterable[_Union[ProductEntry, _Mapping]]] = ...) -> None: ...
