package service

import (
	"log"
	"strings"
	pb "tuyoops/go/terraform-plugin/api"
	"tuyoops/go/terraform-plugin/internal/provider"
	mysqlutil "tuyoops/go/terraform-plugin/pkg/util"
)

type ECSService struct {
	CloudProvider provider.ECSProvider
}

func NewECSService(cloudProvider provider.ECSProvider) *ECSService {
	return &ECSService{
		CloudProvider: cloudProvider,
	}
}

func (s *ECSService) RenderHandleRequest(req *pb.RenderTerraformCodeRequest) (*pb.RenderTerraformCodeResponse, error) {
	region := req.GetRegion()
	dataDiskSize := req.GetDataDiskSize()
	instanceName := req.GetInstanceName()
	imageName := req.GetImageName()
	flavorID := req.GetFlavorId()
	availabilityZone := req.GetAvailabilityZone()
	eipBandwidthName := req.GetEipBandwidthName()
	dataDiskType := req.GetDataDiskType()
	subnetName := req.GetSubnetName()
	secGroup := req.GetSecGroup()
	cloud := req.GetCloud()
	approvalId := strings.TrimLeft(req.GetApprovalId(), "-")
	projectName := req.GetProjectName()
	instanceCount := req.GetInstanceCount()
	sysDiskType := req.GetSysDiskType()
	sysDiskSize := req.GetSysDiskSize()

	// 创建 MySQLUtil 实例
	mysqlUtil, err := mysqlutil.NewMySQLUtil()
	if err != nil {
		log.Fatal("Error creating MySQLUtil:", err)
		return nil, err
	}
	defer mysqlUtil.Close()

	tableName := "ops_ecs_cloud_info"
	rows, err := mysqlUtil.QueryTable(tableName)
	if err != nil {
		log.Fatal("Error querying table:", err)

		return nil, err
	}
	defer rows.Close()

	// 定义变量用于存储每一行的数据
	var cloudTypeB string
	var accountName string
	var ak string
	var sk string
	var cloudType string
	var accessKey string
	var secretKey string

	// 遍历结果集中的每一行
	for rows.Next() {
		// 将当前行的列值扫描到相应的变量中
		err = rows.Scan(&cloudTypeB, &accountName, &ak, &sk)
		if err != nil {
			log.Fatal(err)
		}
		if accountName == cloud {
			accessKey = ak
			secretKey = sk
			cloudType = cloudTypeB
		}
	}
	// 检查迭代过程中是否有错误
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	params := map[string]interface{}{
		"access_key":         accessKey,
		"secret_key":         secretKey,
		"region":             region,
		"image_name":         imageName,
		"project_name":       projectName,
		"secgroup_name":      secGroup,
		"subnet_name":        subnetName,
		"data_disk_type":     dataDiskType,
		"data_disk_size":     dataDiskSize,
		"eip_bandwidth_name": eipBandwidthName,
		"instance_name":      instanceName,
		"instance_count":     instanceCount,
		"sys_disk_type":      sysDiskType,
		"sys_disk_size":      sysDiskSize,
		"flavor_id":          flavorID,
		"availability_zone":  availabilityZone,
		"cloud":              cloud,
		"approval_id":        approvalId,
		"cloud_type":         cloudType,
	}

	// 生成和执行 Terraform 代码
	err = s.CloudProvider.GenerateECSCode(params)
	if err != nil {
		response := &pb.RenderTerraformCodeResponse{
			SuccessCode:  "ERROR",
			ErrorDetails: err.Error(),
		}
		return response, nil
	}
	response := &pb.RenderTerraformCodeResponse{
		SuccessCode:  "OK",
		ErrorDetails: "NO ERROR",
	}
	return response, nil
}

func (s *ECSService) ExecuteHandleRequest(stream pb.TerraformService_ExecuteTerraformCmdServer, req *pb.ExecuteTerraformCmdRequest) error {
	todo := req.GetTodo()
	approvalId := strings.TrimLeft(req.GetApprovalId(), "-")

	err := s.CloudProvider.ExecuteECSCode(approvalId, todo, stream)
	if err != nil {
		return err
	}
	return nil
}

func (s *ECSService) ResourceInfoRequest(req *pb.ResourceInfoRequest) (*pb.ResourceInfoResponse, error) {
	approvalId := strings.TrimLeft(req.GetApprovalId(), "-")
	res, err := s.CloudProvider.ResourceInfo(approvalId)
	if err != nil {
		return nil, err
	}

	// 创建一个 map 存储资源信息
	resourceInfoMap := make(map[string]*pb.ResourceInfoStruct)
	for _, r := range res {
		resource := &pb.ResourceInfoStruct{
			Id:           r.InstanceID,
			InstanceName: r.InstanceName,
			PublicIp:     r.PublicIP,
			PrivateIp:    r.PrivateIP,
			Provider:     r.Cloud,
		}
		resourceInfoMap[r.InstanceID] = resource
	}

	// 构建 ResourceInfoResponse 消息并返回
	response := &pb.ResourceInfoResponse{
		ResourceInfo: resourceInfoMap,
	}
	return response, nil
}
