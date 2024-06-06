from pymongo.database import Database
from pymongo.collection import Collection
from pymongo.results import InsertOneResult, UpdateResult


class CartRepository:
    def __init__(
        self,
        db: Database,
        collection: str,
    ):
        self.collection: Collection = db.get_collection(collection)

    def get(self, user_id: str):
        projection = {
            "_id": 0,
            "cart": 1,
        }
        return self.collection.find_one(
            {"_id": user_id},
            projection=projection,
        )

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
