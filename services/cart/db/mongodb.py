from pymongo import MongoClient
from pymongo.database import Database


class MongoDB:
    def __init__(self, uri: str) -> None:
        self.client: MongoClient = MongoClient(uri)

    def get_database(self, db: str) -> Database:
        return self.client.get_database(db)
