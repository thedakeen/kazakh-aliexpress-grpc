from pydantic import BaseModel


class CartItemEntry(BaseModel):
    quantity: int
    item_id: str
