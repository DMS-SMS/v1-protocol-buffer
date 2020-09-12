# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import proto.python.outing.outing_parent_pb2 as outing__parent__pb2


class OutingParentStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ApproveOutingByOCode = channel.unary_unary(
                '/OutingParent/ApproveOutingByOCode',
                request_serializer=outing__parent__pb2.ApproveOutingByOCodeRequest.SerializeToString,
                response_deserializer=outing__parent__pb2.ApproveOutingByOCodeResponse.FromString,
                )


class OutingParentServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ApproveOutingByOCode(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OutingParentServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ApproveOutingByOCode': grpc.unary_unary_rpc_method_handler(
                    servicer.ApproveOutingByOCode,
                    request_deserializer=outing__parent__pb2.ApproveOutingByOCodeRequest.FromString,
                    response_serializer=outing__parent__pb2.ApproveOutingByOCodeResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'OutingParent', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OutingParent(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ApproveOutingByOCode(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingParent/ApproveOutingByOCode',
            outing__parent__pb2.ApproveOutingByOCodeRequest.SerializeToString,
            outing__parent__pb2.ApproveOutingByOCodeResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)