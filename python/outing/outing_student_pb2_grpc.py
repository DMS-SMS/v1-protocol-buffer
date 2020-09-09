# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import outing_student_pb2 as outing__student__pb2


class OutingStudentStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CreateOuting = channel.unary_unary(
                '/OutingStudent/CreateOuting',
                request_serializer=outing__student__pb2.CreateOutingRequest.SerializeToString,
                response_deserializer=outing__student__pb2.CreateOutingResponse.FromString,
                )
        self.GetStudentOutings = channel.unary_unary(
                '/OutingStudent/GetStudentOutings',
                request_serializer=outing__student__pb2.GetStudentOutingsRequest.SerializeToString,
                response_deserializer=outing__student__pb2.GetStudentOutingsResponse.FromString,
                )
        self.GetOutingInform = channel.unary_unary(
                '/OutingStudent/GetOutingInform',
                request_serializer=outing__student__pb2.GetOutingInformRequest.SerializeToString,
                response_deserializer=outing__student__pb2.GetOutingInformResponse.FromString,
                )
        self.GetCardAboutOuting = channel.unary_unary(
                '/OutingStudent/GetCardAboutOuting',
                request_serializer=outing__student__pb2.GetCardAboutOutingRequest.SerializeToString,
                response_deserializer=outing__student__pb2.GetCardAboutOutingResponse.FromString,
                )


class OutingStudentServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CreateOuting(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetStudentOutings(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOutingInform(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetCardAboutOuting(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OutingStudentServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CreateOuting': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateOuting,
                    request_deserializer=outing__student__pb2.CreateOutingRequest.FromString,
                    response_serializer=outing__student__pb2.CreateOutingResponse.SerializeToString,
            ),
            'GetStudentOutings': grpc.unary_unary_rpc_method_handler(
                    servicer.GetStudentOutings,
                    request_deserializer=outing__student__pb2.GetStudentOutingsRequest.FromString,
                    response_serializer=outing__student__pb2.GetStudentOutingsResponse.SerializeToString,
            ),
            'GetOutingInform': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOutingInform,
                    request_deserializer=outing__student__pb2.GetOutingInformRequest.FromString,
                    response_serializer=outing__student__pb2.GetOutingInformResponse.SerializeToString,
            ),
            'GetCardAboutOuting': grpc.unary_unary_rpc_method_handler(
                    servicer.GetCardAboutOuting,
                    request_deserializer=outing__student__pb2.GetCardAboutOutingRequest.FromString,
                    response_serializer=outing__student__pb2.GetCardAboutOutingResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'OutingStudent', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OutingStudent(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CreateOuting(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingStudent/CreateOuting',
            outing__student__pb2.CreateOutingRequest.SerializeToString,
            outing__student__pb2.CreateOutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetStudentOutings(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingStudent/GetStudentOutings',
            outing__student__pb2.GetStudentOutingsRequest.SerializeToString,
            outing__student__pb2.GetStudentOutingsResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetOutingInform(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingStudent/GetOutingInform',
            outing__student__pb2.GetOutingInformRequest.SerializeToString,
            outing__student__pb2.GetOutingInformResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetCardAboutOuting(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingStudent/GetCardAboutOuting',
            outing__student__pb2.GetCardAboutOutingRequest.SerializeToString,
            outing__student__pb2.GetCardAboutOutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
