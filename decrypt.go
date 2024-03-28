package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
)

type RequestData struct {
	AccessKey        string `json:"access_key"`
	ApprovalID       string `json:"approval_id"`
	AvailabilityZone string `json:"availability_zone"`
	Cloud            string `json:"cloud"`
	DataDiskSize     string `json:"data_disk_size"`
	EIPBandwidthName string `json:"eip_bandwidth_name"`
	DataDiskType     string `json:"data_disk_type"`
	FlavorID         string `json:"flavor_id"`
	ImageName        string `json:"image_name"`
	InstanceName     string `json:"instance_name"`
	InstanceCount    string `json:"instance_count"`
	Region           string `json:"region"`
	SecGroupName     string `json:"secgroup_name"`
	SecretKey        string `json:"secret_key"`
	SubnetName       string `json:"subnet_name"`
	ProjectName      string `json:"project_name"`
	SysDiskSize      string `json:"sys_disk_size"`
	SysDiskType      string `json:"sys_disk_type"`
	CloudType        string `json:"cloud_type"`
}

func decryptAndReadFromFile(inputFilename, outputFilename string, id string) error {
	ciphertext, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		return fmt.Errorf("打开文件失败：%v", err)
	}

	// 使用密钥初始化 AES 密码器
	block, err := aes.NewCipher([]byte("01234567890123456789012345678901"))
	if err != nil {
		return fmt.Errorf("创建 AES 密码失败：%v", err)
	}

	// 使用 AES-GCM 模式解密数据
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("创建 AES-GCM 密码失败：%v", err)
	}

	// 解密操作所需的 nonce 大小与加密时相同
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return fmt.Errorf("密文太短")
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// 使用 AES-GCM 解密密文
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return fmt.Errorf("解密失败：%v", err)
	}

	var data map[string]RequestData
	err = json.Unmarshal(plaintext, &data)
	if err != nil {
		return fmt.Errorf("解析 JSON 失败：%v", err)
	}

	if id != "" {
		if requestData, ok := data[id]; ok {
			jsonData, err := json.MarshalIndent(requestData, "", "  ")
			if err != nil {
				return fmt.Errorf("无法格式化输出 JSON：%v", err)
			}
			fmt.Println(string(jsonData))
		} else {
			return fmt.Errorf("未找到指定ID的数据")
		}
	}

	if outputFilename != "" {
		err = ioutil.WriteFile(outputFilename, plaintext, 0644)
		if err != nil {
			return fmt.Errorf("写入文件失败：%v", err)
		}
	}

	return nil
}

func main() {
	var inputFilename, outputFilename, id string

	flag.StringVar(&inputFilename, "input", "", "输入文件名")
	flag.StringVar(&outputFilename, "output", "", "输出文件名")
	flag.StringVar(&id, "id", "", "要提取的ID字段")
	flag.Parse()

	if inputFilename == "" {
		fmt.Println("错误：-h 显示帮助信息")
		return
	}

	if outputFilename == "" && id == "" {
		fmt.Println("错误：必须指定--output或--id参数之一")
		return
	}

	err := decryptAndReadFromFile(inputFilename, outputFilename, id)
	if err != nil {
		fmt.Println("解密并处理文件失败:", err)
		return
	}

	if outputFilename != "" && id != "" {
		fmt.Printf("解密成功并写入到文件 %s 中\n", outputFilename)
	}
}
