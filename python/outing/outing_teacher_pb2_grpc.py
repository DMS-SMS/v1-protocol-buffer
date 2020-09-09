# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import outing_teacher_pb2 as outing__teacher__pb2


class OutingTeacherStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetOutingWithFilter = channel.unary_unary(
                '/OutingTeacher/GetOutingWithFilter',
                request_serializer=outing__teacher__pb2.GetOutingWithFilterRequest.SerializeToString,
                response_deserializer=outing__teacher__pb2.OutingResponse.FromString,
                )
        self.GetOutingStudentWithSN = channel.unary_unary(
                '/OutingTeacher/GetOutingStudentWithSN',
                request_serializer=outing__teacher__pb2.GetOutingStudentWithSNRequest.SerializeToString,
                response_deserializer=outing__teacher__pb2.OutingResponse.FromString,
                )
        self.GetOutingStudentWithFloor = channel.unary_unary(
                '/OutingTeacher/GetOutingStudentWithFloor',
                request_serializer=outing__teacher__pb2.GetOutingStudentWithFloorRequest.SerializeToString,
                response_deserializer=outing__teacher__pb2.OutingResponse.FromString,
                )
        self.ApproveOuting = channel.unary_unary(
                '/OutingTeacher/ApproveOuting',
                request_serializer=outing__teacher__pb2.ConfirmOutingRequest.SerializeToString,
                response_deserializer=outing__teacher__pb2.ConfirmOutingResponse.FromString,
                )
        self.RejectOuting = channel.unary_unary(
                '/OutingTeacher/RejectOuting',
                request_serializer=outing__teacher__pb2.ConfirmOutingRequest.SerializeToString,
                response_deserializer=outing__teacher__pb2.ConfirmOutingResponse.FromString,
                )
        self.CertifyOuting = channel.unary_unary(
                '/OutingTeacher/CertifyOuting',
                request_serializer=outing__teacher__pb2.ConfirmOutingRequest.SerializeToString,
                response_deserializer=outing__teacher__pb2.ConfirmOutingResponse.FromString,
                )


class OutingTeacherServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetOutingWithFilter(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOutingStudentWithSN(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetOutingStudentWithFloor(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ApproveOuting(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def RejectOuting(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CertifyOuting(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_OutingTeacherServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetOutingWithFilter': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOutingWithFilter,
                    request_deserializer=outing__teacher__pb2.GetOutingWithFilterRequest.FromString,
                    response_serializer=outing__teacher__pb2.OutingResponse.SerializeToString,
            ),
            'GetOutingStudentWithSN': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOutingStudentWithSN,
                    request_deserializer=outing__teacher__pb2.GetOutingStudentWithSNRequest.FromString,
                    response_serializer=outing__teacher__pb2.OutingResponse.SerializeToString,
            ),
            'GetOutingStudentWithFloor': grpc.unary_unary_rpc_method_handler(
                    servicer.GetOutingStudentWithFloor,
                    request_deserializer=outing__teacher__pb2.GetOutingStudentWithFloorRequest.FromString,
                    response_serializer=outing__teacher__pb2.OutingResponse.SerializeToString,
            ),
            'ApproveOuting': grpc.unary_unary_rpc_method_handler(
                    servicer.ApproveOuting,
                    request_deserializer=outing__teacher__pb2.ConfirmOutingRequest.FromString,
                    response_serializer=outing__teacher__pb2.ConfirmOutingResponse.SerializeToString,
            ),
            'RejectOuting': grpc.unary_unary_rpc_method_handler(
                    servicer.RejectOuting,
                    request_deserializer=outing__teacher__pb2.ConfirmOutingRequest.FromString,
                    response_serializer=outing__teacher__pb2.ConfirmOutingResponse.SerializeToString,
            ),
            'CertifyOuting': grpc.unary_unary_rpc_method_handler(
                    servicer.CertifyOuting,
                    request_deserializer=outing__teacher__pb2.ConfirmOutingRequest.FromString,
                    response_serializer=outing__teacher__pb2.ConfirmOutingResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'OutingTeacher', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class OutingTeacher(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetOutingWithFilter(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingTeacher/GetOutingWithFilter',
            outing__teacher__pb2.GetOutingWithFilterRequest.SerializeToString,
            outing__teacher__pb2.OutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetOutingStudentWithSN(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingTeacher/GetOutingStudentWithSN',
            outing__teacher__pb2.GetOutingStudentWithSNRequest.SerializeToString,
            outing__teacher__pb2.OutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetOutingStudentWithFloor(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingTeacher/GetOutingStudentWithFloor',
            outing__teacher__pb2.GetOutingStudentWithFloorRequest.SerializeToString,
            outing__teacher__pb2.OutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ApproveOuting(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingTeacher/ApproveOuting',
            outing__teacher__pb2.ConfirmOutingRequest.SerializeToString,
            outing__teacher__pb2.ConfirmOutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def RejectOuting(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingTeacher/RejectOuting',
            outing__teacher__pb2.ConfirmOutingRequest.SerializeToString,
            outing__teacher__pb2.ConfirmOutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def CertifyOuting(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/OutingTeacher/CertifyOuting',
            outing__teacher__pb2.ConfirmOutingRequest.SerializeToString,
            outing__teacher__pb2.ConfirmOutingResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
