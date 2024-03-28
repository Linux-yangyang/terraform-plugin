import grpc
import terraform_service_pb2_grpc
import terraform_service_pb2

def run(resource_type):
    with grpc.insecure_channel('127.0.0.1:4399') as channel:
        stub = terraform_service_pb2_grpc.TerraformServiceStub(channel)

        # Set the appropriate parameters based on the resource type
        if resource_type == "ecs":
            request = terraform_service_pb2.ExecuteTerraformCmdRequest(
                todo="plan",
                approval_id="1234567"
            )
        elif resource_type == "rds":
            pass
        elif resource_type == "lb":
            pass
        else:
            print(f"Unsupported resource type: {resource_type}")
            return

        try:
            responses = stub.ExecuteTerraformCmd(request)
            for response in responses:
                print(f"Received message: {response.successMessage}")
        except grpc.RpcError as e:
            if e.code() == grpc.StatusCode.OK and "Received trailing metadata with no trailing call status" in e.details():
                print("Stream closed")
            else:
                print(e)
                print(f"Stream closed with error: {e.details()}, {e.code()}")

if __name__ == '__main__':
    run("ecs")
