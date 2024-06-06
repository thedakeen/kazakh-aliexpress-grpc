from concurrent import futures
import logging
import jwt

import grpc

import grpc_proto.cart.cart_pb2 as cart_pb2
import grpc_proto.cart.cart_pb2_grpc as cart_pb2_grpc

import grpc_proto.auth.auth_pb2 as auth_pb2
import grpc_proto.auth.auth_pb2_grpc as auth_pb2_grpc

import grpc_proto.product.product_pb2 as product_pb2
import grpc_proto.product.product_pb2_grpc as product_pb2_grpc

from config import get_settings
from config import CommonSettings

from db.mongodb import MongoDB
from repositories.cart import CartRepository


class CartServicer(cart_pb2_grpc.CartServicer):
    def __init__(self, config: CommonSettings, cart_repository: CartRepository) -> None:
        self.__config = config
        self.__cart_repository = cart_repository
        super().__init__()

    def validate_user_token(self, token: str) -> str | None:
        with grpc.insecure_channel(
            f"{self.__config.auth_service_host}:{self.__config.auth_service_port}"
        ) as channel:
            auth_stub = auth_pb2_grpc.AuthStub(channel)
            response = auth_stub.IsTokenValid(auth_pb2.IsTokenValidRequest(token=token))

            if not response.token_valid:
                return None

            try:
                payload = jwt.decode(token, options={"verify_signature": False})
                return payload.get("email")
            except jwt.InvalidTokenError:
                logging.info(f"Invalid token provided: {token}")
                return None

    def get_product(self, item_id: str):
        with grpc.insecure_channel(
            f"{self.__config.product_service_host}:{self.__config.product_service_port}"
        ) as channel:
            product_stub = product_pb2_grpc.ProductStub(channel)
            response = product_stub.GetProduct(product_pb2.ProductRequest(id=item_id))
            logging.info(f"get product with id: {item_id}. result: {response.product}")
            return response.product

    def GetCart(self, request, context):
        metadata = dict(context.invocation_metadata())
        token = metadata.get("authorization", "")[7:]

        email = self.validate_user_token(token=token)

        cart_items = self.__cart_repository.get(email=email)
        if cart_items:
            product_cart_entries = [
                cart_pb2.ProductCartEntry(
                    quantity=item.quantity,
                    item=cart_pb2.ProductEntry(
                        id=item.item.id,
                        price=item.item.price,
                        item_name=item.item.item_name,
                        categories=[
                            cart_pb2.Category(
                                id=category.id,
                                category_name=category.category_name,
                            )
                            for category in item.item.categories
                        ],
                        item_photos=item.item.item_photos,
                        reviews=item.item.reviews,
                        info=[
                            cart_pb2.Info(
                                info_title=info.info_title,
                                info_content=info.info_content,
                            )
                            for info in item.item.info
                        ],
                        options=(
                            [
                                cart_pb2.Option(
                                    option_title=option.option_title,
                                    option_options=option.option_options,
                                )
                                for option in item.item.options
                            ]
                            if item.item.options
                            else None
                        ),
                    ),
                )
                for item in cart_items
            ]
            return cart_pb2.GetCartResponse(cart=product_cart_entries)
        return cart_pb2.GetCartResponse(cart=[])

    def AddToCart(self, request, context):
        metadata = dict(context.invocation_metadata())
        token = metadata.get("authorization", "")[7:]

        email = self.validate_user_token(token)

        product_id = request.product_id
        quantity = request.quantity

        product = self.get_product(item_id=product_id)
        logging.info(product)
        if not product:
            context.set_code(grpc.StatusCode.INVALID_ARGUMENT)
            context.set_details("Failed to add item to cart. Item does not exists")
            return cart_pb2.AddToCartResponse()

        success = self.__cart_repository.add(email, product_id, quantity)

        if success:
            return cart_pb2.AddToCartResponse()
        else:
            context.set_code(grpc.StatusCode.INTERNAL)
            context.set_details("Failed to add item to cart")
            return cart_pb2.AddToCartResponse()


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
    logging.basicConfig(level=logging.INFO)
    serve()
