# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

import terraform_service_pb2 as terraform__service__pb2


class TerraformServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.RenderTerraformCode = channel.unary_unary(
                '/api.TerraformService/RenderTerraformCode',
                request_serializer=terraform__service__pb2.RenderTerraformCodeRequest.SerializeToString,
                response_deserializer=terraform__service__pb2.RenderTerraformCodeResponse.FromString,
                )
        self.ExecuteTerraformCmd = channel.unary_stream(
                '/api.TerraformService/ExecuteTerraformCmd',
                request_serializer=terraform__service__pb2.ExecuteTerraformCmdRequest.SerializeToString,
                response_deserializer=terraform__service__pb2.ExecuteTerraformCmdResponse.FromString,
                )
        self.ResourceInfoGet = channel.unary_unary(
                '/api.TerraformService/ResourceInfoGet',
                request_serializer=terraform__service__pb2.ResourceInfoRequest.SerializeToString,
                response_deserializer=terraform__service__pb2.ResourceInfoResponse.FromString,
                )


class TerraformServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def RenderTerraformCode(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ExecuteTerraformCmd(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ResourceInfoGet(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_TerraformServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'RenderTerraformCode': grpc.unary_unary_rpc_method_handler(
                    servicer.RenderTerraformCode,
                    request_deserializer=terraform__service__pb2.RenderTerraformCodeRequest.FromString,
                    response_serializer=terraform__service__pb2.RenderTerraformCodeResponse.SerializeToString,
            ),
            'ExecuteTerraformCmd': grpc.unary_stream_rpc_method_handler(
                    servicer.ExecuteTerraformCmd,
                    request_deserializer=terraform__service__pb2.ExecuteTerraformCmdRequest.FromString,
                    response_serializer=terraform__service__pb2.ExecuteTerraformCmdResponse.SerializeToString,
            ),
            'ResourceInfoGet': grpc.unary_unary_rpc_method_handler(
                    servicer.ResourceInfoGet,
                    request_deserializer=terraform__service__pb2.ResourceInfoRequest.FromString,
                    response_serializer=terraform__service__pb2.ResourceInfoResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.TerraformService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class TerraformService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def RenderTerraformCode(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.TerraformService/RenderTerraformCode',
            terraform__service__pb2.RenderTerraformCodeRequest.SerializeToString,
            terraform__service__pb2.RenderTerraformCodeResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ExecuteTerraformCmd(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/api.TerraformService/ExecuteTerraformCmd',
            terraform__service__pb2.ExecuteTerraformCmdRequest.SerializeToString,
            terraform__service__pb2.ExecuteTerraformCmdResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def ResourceInfoGet(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.TerraformService/ResourceInfoGet',
            terraform__service__pb2.ResourceInfoRequest.SerializeToString,
            terraform__service__pb2.ResourceInfoResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
