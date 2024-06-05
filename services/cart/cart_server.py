# Copyright 2015 gRPC authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
"""The Python implementation of the GRPC helloworld.Greeter server."""

from concurrent import futures
import logging

import grpc
import cart_pb2
import cart_pb2_grpc

import auth_pb2
import auth_pb2_grpc


class CartServicer(cart_pb2_grpc.CartServicer):
    def GetCart(self, request, context):
        metadata = dict(context.invocation_metadata())
        token = metadata.get("authorization", "")[7:]

        with grpc.insecure_channel("localhost:50050") as channel:

            auth_stub = auth_pb2_grpc.AuthStub(channel)
            response = auth_stub.IsTokenValid(auth_pb2.IsTokenValidRequest(token=token))
            print("Greeter client received: ", response.token_valid)

        return cart_pb2.GetCartResponse(
            products=[],
        )


def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    cart_pb2_grpc.add_CartServicer_to_server(CartServicer(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)
    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()
