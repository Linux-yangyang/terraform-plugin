package provider

import (
	pb "tuyoops/go/terraform-plugin/api"
	"tuyoops/go/terraform-plugin/pkg/util"
)

type ECSProvider interface {
	GenerateECSCode(params map[string]interface{}) error
	ExecuteECSCode(approvalId string, exec string, stream pb.TerraformService_ExecuteTerraformCmdServer) error
	ResourceInfo(approvalId string) ([]util.InstanceInfo, error)
}

type LBProvider interface {
	GenerateLBCode(params map[string]interface{}) error
	ExecuteLBCode(exec string, stream pb.TerraformService_ExecuteTerraformCmdServer) error
}
