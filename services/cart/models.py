from pydantic import BaseModel, Field


class Category(BaseModel):
    id: str = Field(alias="_id")
    category_name: str


class Info(BaseModel):
    info_title: str
    info_content: str


class Option(BaseModel):
    option_title: str
    option_options: list[str]


class CartItem(BaseModel):
    id: str = Field(alias="_id")
    price: float
    item_name: str
    categories: list[Category]
    item_photos: list[str] | None = None
    info: list[Info]
    options: list[Option] | None = None


class CartItemEntry(BaseModel):
    quantity: int  # Ensures the quantity is at least 1
    item: CartItem


# class BookUpdate(BaseModel):
#     title: Optional[str]
#     author: Optional[str]
#     synopsis: Optional[str]

#     class Config:
#         schema_extra = {
#             "example": {
#                 "title": "Don Quixote",
#                 "author": "Miguel de Cervantes",
#                 "synopsis": "Don Quixote is a Spanish novel by Miguel de Cervantes..."
#             }
#         }
