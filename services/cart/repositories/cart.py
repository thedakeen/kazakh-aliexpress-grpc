import logging
from pymongo.database import Database
from pymongo.collection import Collection

from models import CartItemEntry
from models import CartItem


class CartRepository:
    def __init__(
        self,
        db: Database,
        collection: str,
    ):
        self.collection: Collection = db.get_collection(collection)

    def get(self, email: str) -> list[CartItemEntry] | None:
        projection = {
            "_id": 0,
            "cart": 1,
        }
        result = self.collection.find_one(
            {"email": email},
            projection=projection,
        )
        logging.info(f"find cart result for email: {email}. Result: {result}")
        if result:
            return [
                CartItemEntry.model_validate(**item) for item in result.get("cart", [])
            ]
        return None

    def add(
        self,
        email: str,
        product_id: str,
        quantity: int,
    ) -> bool:
        try:
            # Check if the product already exists in the user's cart
            user = self.collection.find_one(
                {"email": email, "cart.item_id": product_id}
            )

            if user:
                # If the product exists, increment the quantity
                self.collection.update_one(
                    {"email": email, "cart.item_id": product_id},
                    {"$inc": {"cart.$.quantity": quantity}},
                )
            else:
                # If the product does not exist, add the product to the cart
                self.collection.update_one(
                    {"email": email},
                    {"$push": {"cart": {"item_id": product_id, "quantity": quantity}}},
                )

            return True
        except Exception as e:
            logging.error(f"Failed to add item to cart: {e}")
            return False

    # def add(
    #     self,
    #     email: str,
    #     product: CartItem,
    #     quantity: int,
    # ) -> bool:
    #     try:
    #         self.collection.update_one(
    #             {"email": email, "cart.product_id": product_id},
    #             {"$inc": {"cart.$.quantity": quantity}},
    #             upsert=True,
    #         )
    #         return True
    #     except Exception as e:
    #         logging.error(f"Failed to add item to cart: {e}")
    #         return False

    # async def create_chat(
    #     self,
    #     user_card_id: int,
    #     recipient_card_id: int,
    # ) -> InsertOneResult:
    #     return await self.collection.insert_one(
    #         {
    #             "participants": [user_card_id, recipient_card_id],
    #             "messages": [],
    #         }
    #     )

    # async def get_user_chats(self, user_card_id: int) -> list[Chat]:
    #     projection = {"_id": 0}
    #     chat_cursor = self.collection.find(
    #         {
    #             "participants": {
    #                 "$all": [user_card_id],
    #             },
    #         },
    #         projection=projection,
    #     )
    #     chats = await chat_cursor.to_list(length=None)
    #     return [
    #         Chat.model_validate(
    #             {
    #                 "participants": chat["participants"],
    #                 "messages": list(
    #                     sorted(
    #                         [
    #                             ChatMessage.model_validate(message)
    #                             for message in chat["messages"]
    #                         ],
    #                         key=lambda message: message.created_at,
    #                     )
    #                 ),
    #             }
    #         )
    #         for chat in chats
    #     ]

    # async def insert_chat(self): ...

    # async def add_message(
    #     self,
    #     message_id: int,
    #     message_text: str,
    #     author_id: int,
    #     recipient_id: int,
    #     created_at: datetime,
    #     message_type: str,
    # ) -> UpdateResult:
    #     return await self.collection.update_one(
    #         {
    #             "participants": {
    #                 "$all": [
    #                     {"$elemMatch": {"$eq": author_id}},
    #                     {"$elemMatch": {"$eq": recipient_id}},
    #                 ]
    #             }
    #         },
    #         {
    #             "$setOnInsert": {
    #                 "participants": [
    #                     author_id,
    #                     recipient_id,
    #                 ],
    #             },
    #             "$push": {
    #                 "messages": {
    #                     "message_id": message_id,
    #                     "owner_card_id": author_id,
    #                     "text": message_text,
    #                     "created_at": created_at,
    #                     "message_type": message_type,
    #                     "has_been_read_at": None,
    #                 }
    #             },
    #         },
    #         upsert=True,
    #     )
