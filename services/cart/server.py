from concurrent import futures
import logging

import grpc

import grpc_proto.cart.cart_pb2 as cart_pb2
import grpc_proto.cart.cart_pb2_grpc as cart_pb2_grpc

import grpc_proto.auth.auth_pb2 as auth_pb2
import grpc_proto.auth.auth_pb2_grpc as auth_pb2_grpc

from config import get_settings
from config import CommonSettings

from db.mongodb import MongoDB
from repositories.cart import CartRepository


class CartServicer(cart_pb2_grpc.CartServicer):
    def __init__(self, config: CommonSettings, cart_repository: CartRepository) -> None:
        self.__config = config
        self.__cart_repository = cart_repository
        super().__init__()

    def GetCart(self, request, context):
        metadata = dict(context.invocation_metadata())
        token = metadata.get("authorization", "")[7:]

        with grpc.insecure_channel(
            f"{self.__config.auth_service_host}:{self.__config.auth_service_port}"
        ) as channel:

            logging.info(
                channel,
                f"{self.__config.auth_service_host}:{self.__config.auth_service_port}",
            )
            auth_stub = auth_pb2_grpc.AuthStub(channel)
            response = auth_stub.IsTokenValid(auth_pb2.IsTokenValidRequest(token=token))

            if response.token_valid:
                logging.info(
                    self.__cart_repository.get(user_id="65e84aa6375c87c257d0f6b6"),
                    "FRom repo",
                )
                logging.info("Greeter client received: ", response.token_valid)

        return cart_pb2.GetCartResponse(
            products=[],
        )


def serve():
    config = get_settings()
    port = str(config.cart_service_port)
    mongo_db = MongoDB(uri=config.mongo_uri)

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cart_pb2_grpc.add_CartServicer_to_server(
        CartServicer(
            config=config,
            cart_repository=CartRepository(
                db=mongo_db.get_database(db="Qazaq-Aliexpress"),
                collection="users",
            ),
        ),
        server,
    )
    server.add_insecure_port("[::]:" + port)
    server.start()
    logging.info("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()
