# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import proto.python.auth.auth_parent_pb2 as auth__parent__pb2


class AuthParentStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.LoginParentAuth = channel.unary_unary(
                '/AuthParent/LoginParentAuth',
                request_serializer=auth__parent__pb2.LoginParentAuthRequest.SerializeToString,
                response_deserializer=auth__parent__pb2.LoginParentAuthResponse.FromString,
                )
        self.ChangeParentPW = channel.unary_unary(
                '/AuthParent/ChangeParentPW',
                request_serializer=auth__parent__pb2.ChangeParentPWRequest.SerializeToString,
                response_deserializer=auth__parent__pb2.ChangeParentPWResponse.FromString,
                )
        self.GetParentInformWithUUID = channel.unary_unary(
                '/AuthParent/GetParentInformWithUUID',
                request_serializer=auth__parent__pb2.GetParentInformWithUUIDRequest.SerializeToString,
                response_deserializer=auth__parent__pb2.GetParentInformWithUUIDResponse.FromString,
                )
        self.GetParentUUIDsWithInform = channel.unary_unary(
                '/AuthParent/GetParentUUIDsWithInform',
                request_serializer=auth__parent__pb2.GetParentUUIDsWithInformRequest.SerializeToString,
                response_deserializer=auth__parent__pb2.GetParentUUIDsWithInformResponse.FromString,
                )


class AuthParentServicer(object):
    """Missing associated documentation comment in .proto file."""

    def LoginParentAuth(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ChangeParentPW(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetParentInformWithUUID(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def GetParentUUIDsWithInform(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_AuthParentServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'LoginParentAuth': grpc.unary_unary_rpc_method_handler(
                    servicer.LoginParentAuth,
                    request_deserializer=auth__parent__pb2.LoginParentAuthRequest.FromString,
                    response_serializer=auth__parent__pb2.LoginParentAuthResponse.SerializeToString,
            ),
            'ChangeParentPW': grpc.unary_unary_rpc_method_handler(
                    servicer.ChangeParentPW,
                    request_deserializer=auth__parent__pb2.ChangeParentPWRequest.FromString,
                    response_serializer=auth__parent__pb2.ChangeParentPWResponse.SerializeToString,
            ),
            'GetParentInformWithUUID': grpc.unary_unary_rpc_method_handler(
                    servicer.GetParentInformWithUUID,
                    request_deserializer=auth__parent__pb2.GetParentInformWithUUIDRequest.FromString,
                    response_serializer=auth__parent__pb2.GetParentInformWithUUIDResponse.SerializeToString,
            ),
            'GetParentUUIDsWithInform': grpc.unary_unary_rpc_method_handler(
                    servicer.GetParentUUIDsWithInform,
                    request_deserializer=auth__parent__pb2.GetParentUUIDsWithInformRequest.FromString,
                    response_serializer=auth__parent__pb2.GetParentUUIDsWithInformResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'AuthParent', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class AuthParent(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def LoginParentAuth(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AuthParent/LoginParentAuth',
            auth__parent__pb2.LoginParentAuthRequest.SerializeToString,
            auth__parent__pb2.LoginParentAuthResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ChangeParentPW(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AuthParent/ChangeParentPW',
            auth__parent__pb2.ChangeParentPWRequest.SerializeToString,
            auth__parent__pb2.ChangeParentPWResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetParentInformWithUUID(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AuthParent/GetParentInformWithUUID',
            auth__parent__pb2.GetParentInformWithUUIDRequest.SerializeToString,
            auth__parent__pb2.GetParentInformWithUUIDResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def GetParentUUIDsWithInform(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/AuthParent/GetParentUUIDsWithInform',
            auth__parent__pb2.GetParentUUIDsWithInformRequest.SerializeToString,
            auth__parent__pb2.GetParentUUIDsWithInformResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
