package util

import (
	"encoding/json"
	"fmt"
	"os"
	"tuyoops/go/terraform-plugin/config"
)

type ECS struct {
	Version          int         `json:"version"`
	TerraformVersion string      `json:"terraform_version"`
	Serial           int         `json:"serial"`
	Lineage          string      `json:"lineage"`
	Outputs          Outputs     `json:"outputs"`
	Resources        []Resources `json:"resources"`
	CheckResults     interface{} `json:"check_results"`
}

type Outputs struct{}

type Instances struct {
	IndexKey            int                    `json:"index_key"`
	SchemaVersion       int                    `json:"schema_version"`
	Attributes          map[string]interface{} `json:"attributes"`
	SensitiveAttributes []interface{}          `json:"sensitive_attributes"`
	Private             string                 `json:"private"`
	Dependencies        []string               `json:"dependencies"`
}

type Resources struct {
	Module    string      `json:"module"`
	Mode      string      `json:"mode"`
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	Provider  string      `json:"provider"`
	Instances []Instances `json:"instances"`
}

type InstanceInfo struct {
	InstanceID   string `json:"id"`
	InstanceName string `json:"instance_name"`
	PublicIP     string `json:"public_ip"`
	PrivateIP    string `json:"private_ip"`
	Cloud        string `json:"provider"`
}

func HuaweiResultParse(approvalId, cloud, resourcesType string) ([]InstanceInfo, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/terraform.tfstate.d/%s/terraform.tfstate", config.Conf.TerraformDir.Root, cloud, resourcesType, approvalId))
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return nil, err
	}

	var huaweiECS ECS
	err = json.Unmarshal(data, &huaweiECS)
	if err != nil {
		fmt.Println("无法解析 JSON:", err)
		return nil, err
	}

	instances := make(map[string]InstanceInfo)

	for _, resource := range huaweiECS.Resources {
		if resource.Name == "tf-instance" {
			for _, instance := range resource.Instances {
				var instanceInfo InstanceInfo
				for key, value := range instance.Attributes {
					switch key {
					case "id":
						if v, ok := value.(string); ok {
							instanceInfo.InstanceID = v
						} else {
							instanceInfo.InstanceID = ""
						}
					case "name":
						if v, ok := value.(string); ok {
							instanceInfo.InstanceName = v
						} else {
							instanceInfo.InstanceName = ""
						}
					case "public_ip":
						if v, ok := value.(string); ok {
							instanceInfo.PublicIP = v
						} else {
							instanceInfo.PublicIP = ""
						}
					case "access_ip_v4":
						if v, ok := value.(string); ok {
							instanceInfo.PrivateIP = v
						} else {
							instanceInfo.PrivateIP = ""
						}
					}
				}
				instanceInfo.Cloud = "huawei"
				if associatedInstance, ok := instances[instanceInfo.InstanceID]; ok {
					instanceInfo.PublicIP = associatedInstance.PublicIP
				}
				instances[instanceInfo.InstanceID] = instanceInfo
			}
		}
	}

	// 转换为切片形式
	var instanceSlice []InstanceInfo
	for _, v := range instances {
		instanceSlice = append(instanceSlice, v)
	}

	return instanceSlice, nil
}

func AliyunResultParse(approvalId, cloud, resourcesType string) ([]InstanceInfo, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/terraform.tfstate.d/%s/terraform.tfstate", config.Conf.TerraformDir.Root, cloud, resourcesType, approvalId))
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return nil, err
	}

	var aliyunECS ECS
	err = json.Unmarshal(data, &aliyunECS)
	if err != nil {
		fmt.Println("无法解析 JSON:", err)
		return nil, err
	}

	instances := make(map[string]InstanceInfo)

	for _, resource := range aliyunECS.Resources {
		if resource.Name == "tf-instance" {
			for _, instance := range resource.Instances {
				var instanceInfo InstanceInfo
				for key, value := range instance.Attributes {
					switch key {
					case "id":
						if v, ok := value.(string); ok {
							instanceInfo.InstanceID = v
						} else {
							instanceInfo.InstanceID = ""
						}
					case "instance_name":
						if v, ok := value.(string); ok {
							instanceInfo.InstanceName = v
						} else {
							instanceInfo.InstanceName = ""
						}
					case "public_ip":
						if v, ok := value.(string); ok {
							instanceInfo.PublicIP = v
						} else {
							instanceInfo.PublicIP = ""
						}
					case "private_ip":
						if v, ok := value.(string); ok {
							instanceInfo.PrivateIP = v
						} else {
							instanceInfo.PrivateIP = ""
						}
					}
				}
				instanceInfo.Cloud = "aliyun"
				instances[instanceInfo.InstanceID] = instanceInfo
			}
		}
	}

	// 转换为切片形式
	var instanceSlice []InstanceInfo
	for _, v := range instances {
		instanceSlice = append(instanceSlice, v)
	}

	return instanceSlice, nil
}

func VolcengineResultParse(approvalId, cloud, resourcesType string) ([]InstanceInfo, error) {
	data, err := os.ReadFile(fmt.Sprintf("%s/%s/%s/terraform.tfstate.d/%s/terraform.tfstate", config.Conf.TerraformDir.Root, cloud, resourcesType, approvalId))
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return nil, err
	}

	var volcengineECS ECS
	err = json.Unmarshal(data, &volcengineECS)
	if err != nil {
		fmt.Println("无法解析 JSON:", err)
		return nil, err
	}

	instances := make(map[string]InstanceInfo) // 将 instances 定义为 map

	for _, resource := range volcengineECS.Resources {
		if resource.Name == "associated" {
			for _, instance := range resource.Instances {
				var instanceInfo InstanceInfo
				for key, value := range instance.Attributes {
					switch key {
					case "instance_id":
						instanceInfo.InstanceID = value.(string)
					case "allocation_id":
						instanceInfo.PublicIP = value.(string)
					}
				}
				instances[instanceInfo.InstanceID] = instanceInfo
			}
		}
	}

	for _, resource := range volcengineECS.Resources {
		if resource.Name == "tf-instance" {
			for _, instance := range resource.Instances {
				var instanceInfo InstanceInfo
				for key, value := range instance.Attributes {
					switch key {
					case "instance_name":
						instanceInfo.InstanceName = value.(string)
					case "instance_id":
						instanceInfo.InstanceID = value.(string)
					case "primary_ip_address":
						instanceInfo.PrivateIP = value.(string)
					}
				}
				instanceInfo.Cloud = "volcengine"
				if associatedInstance, ok := instances[instanceInfo.InstanceID]; ok {
					instanceInfo.PublicIP = associatedInstance.PublicIP
				}
				instances[instanceInfo.InstanceID] = instanceInfo
			}
		}
	}

	eipMap := make(map[string]string)

	for _, resource := range volcengineECS.Resources {
		if resource.Name == "tf_eip" {
			for _, instance := range resource.Instances {
				eipMap[instance.Attributes["id"].(string)] = instance.Attributes["eip_address"].(string)
			}
		}
	}

	// 转换为切片形式
	var instanceSlice []InstanceInfo
	for _, v := range instances {
		v.PublicIP = eipMap[v.PublicIP]
		instanceSlice = append(instanceSlice, v)
	}

	return instanceSlice, nil
}
