package provider

import (
	"fmt"
	"os"
	"sync"
	"tuyoops/go/terraform-plugin/config"
	"tuyoops/go/terraform-plugin/pkg/util"

	pb "tuyoops/go/terraform-plugin/api"
)

type CloudECSProvider struct{}

var workspaceMutex sync.Mutex

func (p *CloudECSProvider) GenerateECSCode(params map[string]interface{}) error {
	workspaceMutex.Lock()
	defer workspaceMutex.Unlock()
	if params["cloud_type"] == "HUAWEI" {
		util.ChangeWorkingDirectory(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, "huawei", "ecs"))
	} else if params["cloud_type"] == "ALI" {
		util.ChangeWorkingDirectory(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, "aliyun", "ecs"))
	} else if params["cloud_type"] == "VOLCENGINE" {
		util.ChangeWorkingDirectory(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, "volcengine", "ecs"))
	}
	err := util.RunCommand("terraform", "workspace", "select", params["approval_id"].(string))
	if err != nil {
		err = util.RunCommand("terraform", "workspace", "new", params["approval_id"].(string))
		if err != nil {
			util.ChangeWorkingDirectory("/app")
			return err
		}
	}
	util.ChangeWorkingDirectory("/app")
	err = util.AppendJSONToFile(params, fmt.Sprintf("%s/var.json", fmt.Sprintf("%s/%s", config.Conf.TerraformDir.Root, "var")), "01234567890123456789012345678901")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	util.ChangeWorkingDirectory("/app")
	return nil
}

func (p *CloudECSProvider) ExecuteECSCode(approvalId string, exec string, stream pb.TerraformService_ExecuteTerraformCmdServer) error {
	workspaceMutex.Lock()
	defer workspaceMutex.Unlock()
	requestData, err := util.ReadJSONFromFile(fmt.Sprintf("%s/var.json", fmt.Sprintf("%s/%s", config.Conf.TerraformDir.Root, "var")))
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	if requestData[approvalId].CloudType == "HUAWEI" {
		util.ChangeWorkingDirectory(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, "huawei", "ecs"))
	} else if requestData[approvalId].CloudType == "ALI" {
		util.ChangeWorkingDirectory(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, "aliyun", "ecs"))
	} else if requestData[approvalId].CloudType == "VOLCENGINE" {
		util.ChangeWorkingDirectory(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, "volcengine", "ecs"))
	}
	if exec == "plan" {
		var createEip = false
		if requestData[approvalId].EIPBandwidthName != "" {
			createEip = true
		}
		err := util.RunTerraformCommand(stream, "plan",
			"-var", fmt.Sprintf("my-access-key=%s", requestData[approvalId].AccessKey),
			"-var", fmt.Sprintf("my-secret-key=%s", requestData[approvalId].SecretKey),
			"-var", fmt.Sprintf("my-region=%s", requestData[approvalId].Region),
			"-var", fmt.Sprintf("my-image=%s", requestData[approvalId].ImageName),
			"-var", fmt.Sprintf("my-project=%s", requestData[approvalId].ProjectName),
			"-var", fmt.Sprintf("my-secgroup=%s", requestData[approvalId].SecGroupName),
			"-var", fmt.Sprintf("my-subnet=%s", requestData[approvalId].SubnetName),
			"-var", fmt.Sprintf("my-env=%s", os.Getenv("OPSHELPER_ENV")),
			"-var", fmt.Sprintf("my-data-disk-type=%s", requestData[approvalId].DataDiskType),
			"-var", fmt.Sprintf("my-data-disk-size=%s", requestData[approvalId].DataDiskSize),
			"-var", fmt.Sprintf("create_eip=%v", createEip),
			"-var", fmt.Sprintf("my-eip-name=%s", requestData[approvalId].EIPBandwidthName),
			"-var", fmt.Sprintf("my-instance-name=%s", requestData[approvalId].InstanceName),
			"-var", fmt.Sprintf("my-instance-count=%s", requestData[approvalId].InstanceCount),
			"-var", fmt.Sprintf("my-sys-disk-type=%s", requestData[approvalId].SysDiskType),
			"-var", fmt.Sprintf("my-sys-disk-size=%s", requestData[approvalId].SysDiskSize),
			"-var", fmt.Sprintf("my-flavor-id=%s", requestData[approvalId].FlavorID),
			"-var", fmt.Sprintf("my-availability-zone=%s", requestData[approvalId].AvailabilityZone),
		)
		if err != nil {
			util.ChangeWorkingDirectory("/app")
			return err
		}
		util.ChangeWorkingDirectory("/app")
		return nil
	} else if exec == "apply" {
		var createEip = false
		if requestData[approvalId].EIPBandwidthName != "" {
			createEip = true
		}
		err := util.RunTerraformCommand(stream, exec, "-auto-approve",
			"-var", fmt.Sprintf("my-access-key=%s", requestData[approvalId].AccessKey),
			"-var", fmt.Sprintf("my-secret-key=%s", requestData[approvalId].SecretKey),
			"-var", fmt.Sprintf("my-region=%s", requestData[approvalId].Region),
			"-var", fmt.Sprintf("my-image=%s", requestData[approvalId].ImageName),
			"-var", fmt.Sprintf("my-project=%s", requestData[approvalId].ProjectName),
			"-var", fmt.Sprintf("my-secgroup=%s", requestData[approvalId].SecGroupName),
			"-var", fmt.Sprintf("my-env=%s", os.Getenv("OPSHELPER_ENV")),
			"-var", fmt.Sprintf("my-subnet=%s", requestData[approvalId].SubnetName),
			"-var", fmt.Sprintf("my-data-disk-type=%s", requestData[approvalId].DataDiskType),
			"-var", fmt.Sprintf("my-data-disk-size=%s", requestData[approvalId].DataDiskSize),
			"-var", fmt.Sprintf("create_eip=%v", createEip),
			"-var", fmt.Sprintf("my-eip-name=%s", requestData[approvalId].EIPBandwidthName),
			"-var", fmt.Sprintf("my-instance-name=%s", requestData[approvalId].InstanceName),
			"-var", fmt.Sprintf("my-instance-count=%s", requestData[approvalId].InstanceCount),
			"-var", fmt.Sprintf("my-sys-disk-type=%s", requestData[approvalId].SysDiskType),
			"-var", fmt.Sprintf("my-sys-disk-size=%s", requestData[approvalId].SysDiskSize),
			"-var", fmt.Sprintf("my-flavor-id=%s", requestData[approvalId].FlavorID),
			"-var", fmt.Sprintf("my-availability-zone=%s", requestData[approvalId].AvailabilityZone),
		)
		if err != nil {
			util.ChangeWorkingDirectory("/app")
			return err
		}
	}
	util.ChangeWorkingDirectory("/app")
	return nil
}

func (p *CloudECSProvider) ResourceInfo(approvalId string) ([]util.InstanceInfo, error) {
	workspaceMutex.Lock()
	defer workspaceMutex.Unlock()
	util.ChangeWorkingDirectory("/app")
	requestData, err := util.ReadJSONFromFile(fmt.Sprintf("%s/var.json", fmt.Sprintf("%s/%s", config.Conf.TerraformDir.Root, "var")))
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	if requestData[approvalId].CloudType == "HUAWEI" {
		res, err := util.HuaweiResultParse(approvalId, "huawei", "ecs")
		if err != nil {
			return nil, err
		}
		return res, nil
	} else if requestData[approvalId].CloudType == "ALI" {
		res, err := util.AliyunResultParse(approvalId, "aliyun", "ecs")
		if err != nil {
			return nil, err
		}
		return res, nil
	} else if requestData[approvalId].CloudType == "VOLCENGINE" {
		res, err := util.VolcengineResultParse(approvalId, "volcengine", "ecs")
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	return nil, nil
}
