import grpc
import terraform_service_pb2
import terraform_service_pb2_grpc

def run(resource_type):
    with grpc.insecure_channel('127.0.0.1:4399') as channel:
        stub = terraform_service_pb2_grpc.TerraformServiceStub(channel)

        # Set the appropriate parameters based on the resource type
        if resource_type == "ecs":
            request = terraform_service_pb2.RenderTerraformCodeRequest(
                region="cn-beijing",
                data_disk_size="50",
                instance_name="test",
                image_name="CentOS 7.9 64位",
                flavor_id="ecs.g3i.large",
                availability_zone="cn-beijing-b",
                eip_bandwidth_name="mybandwidth",
                data_disk_type="ESSD_PL0",
                subnet_name="mysubnet",
                sec_group="mysecgroup",
                cloud="volcengine",
                approval_id="1234567",
                project_name="myproject",
                instance_count="1",
                sys_disk_type="ESSD_PL0",
                sys_disk_size="50"
            )
        elif resource_type == "rds":
            pass
        elif resource_type == "lb":
            pass
        else:
            print(f"Unsupported resource type: {resource_type}")
            return

        try:
            response = stub.RenderTerraformCode(request)
            print(response.successCode)
            #print(response.errorDetails)
        except grpc.RpcError as e:
            print(f"{e.details()}")

if __name__ == '__main__':
    # Specify the resource type when calling the run function
    run("ecs")

