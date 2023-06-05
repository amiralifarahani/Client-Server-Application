import pb
import grpc
from concurrent import futures
import logging
import time

logger = logging.getLogger(__name__)


class AuthServiceServicer(pb.RegisterAuthenticationServiceServer):
    def Req_Pq(self, request, context):
        nonce = request.Nonce
        serverNonce = "12345678912345678900"
        messageId = request.MessageId + 1
        p = 23
        g = 11
        time.sleep(1)
        return pb.ReqPq_Request(nonce, serverNonce, messageId, p, g)

    def Req_Dh_Params(self, request, context):
        nonce = request.Nonce
        serverNonce = request.ServerNonce
        messageId = request.MessageId + 1
        b = 22
        time.sleep(1)
        return pb.ReqPq_Request(nonce, serverNonce, messageId, b)


def start_server():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb.add_RegisterAuthenticationServiceServer_to_server(
        RegisterAuthenticationServiceServer(), server
    )
    server.add_insecure_port("localhost:5052")
    server.start()
    logger.info("gRPC server started")
    server.wait_for_termination()


if __name__ == "__main__":
    start_server()
