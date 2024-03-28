package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
)

func ChangeWorkingDirectory(dirPath string) error {
	err := os.Chdir(dirPath)
	if err != nil {
		return err
	}
	return nil
}

func HasDir(directory string) (bool, error) {
	info, err := os.Stat(directory)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	if !info.IsDir() {
		return false, fmt.Errorf(".terraform is not a directory")
	}

	return true, nil
}

func RunCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	//cmd.Stdout = os.Stdout
	//cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command: %v", err)
	}

	return nil
}

func RunCommandResult(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), err
	}
	return string(output), nil
}

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

func encryptAndWriteToFile(data map[string]RequestData, filename, key string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("JSON 编码失败：%v", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return fmt.Errorf("创建 AES 密码失败：%v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("创建 AES-GCM 密码失败：%v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("生成 nonce 失败：%v", err)
	}

	ciphertext := gcm.Seal(nonce, nonce, jsonData, nil)

	err = ioutil.WriteFile(filename, ciphertext, 0777)
	if err != nil {
		return fmt.Errorf("写入加密数据失败：%v", err)
	}

	return nil
}

func decryptAndReadFromFile(filename, key string) (map[string]RequestData, error) {
	ciphertext, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("打开文件失败：%v", err)
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, fmt.Errorf("创建 AES 密码失败：%v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("创建 AES-GCM 密码失败：%v", err)
	}

	if len(ciphertext) < gcm.NonceSize() {
		return nil, fmt.Errorf("密文太短")
	}
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, fmt.Errorf("解密失败：%v", err)
	}

	var result map[string]RequestData
	err = json.Unmarshal(plaintext, &result)
	if err != nil {
		return nil, fmt.Errorf("JSON 解码失败：%v", err)
	}

	return result, nil
}

func AppendJSONToFile(newParams map[string]interface{}, filename, key string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("打开文件失败：%v", err)
	}
	defer file.Close()

	var existingData map[string]RequestData

	fileInfo, err := file.Stat()
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("获取文件信息失败：%v", err)
	}

	if fileInfo != nil && fileInfo.Size() > 0 {
		data, _ := decryptAndReadFromFile(filename, key)
		jsonData, err := json.Marshal(data)
		if err != nil {
			return fmt.Errorf("JSON 编码失败：%v", err)
		}
		decoder := json.NewDecoder(bytes.NewReader(jsonData))
		err = decoder.Decode(&existingData)
		if err != nil && err != io.EOF {
			return fmt.Errorf("解码 JSON 失败：%v", err)
		}
	} else {
		existingData = make(map[string]RequestData)
	}

	approvalID, ok := newParams["approval_id"].(string)
	if !ok {
		return fmt.Errorf("无法获取 approval_id 值")
	}

	existingData[approvalID] = RequestData{
		AccessKey:        newParams["access_key"].(string),
		ApprovalID:       newParams["approval_id"].(string),
		AvailabilityZone: newParams["availability_zone"].(string),
		Cloud:            newParams["cloud"].(string),
		DataDiskSize:     newParams["data_disk_size"].(string),
		EIPBandwidthName: newParams["eip_bandwidth_name"].(string),
		DataDiskType:     newParams["data_disk_type"].(string),
		FlavorID:         newParams["flavor_id"].(string),
		ImageName:        newParams["image_name"].(string),
		InstanceName:     newParams["instance_name"].(string),
		InstanceCount:    newParams["instance_count"].(string),
		Region:           newParams["region"].(string),
		SecGroupName:     newParams["secgroup_name"].(string),
		SecretKey:        newParams["secret_key"].(string),
		SubnetName:       newParams["subnet_name"].(string),
		ProjectName:      newParams["project_name"].(string),
		SysDiskSize:      newParams["sys_disk_size"].(string),
		SysDiskType:      newParams["sys_disk_type"].(string),
		CloudType:        newParams["cloud_type"].(string),
	}

	// 加密并将更新后的数据写入文件
	return encryptAndWriteToFile(existingData, filename, key)
}

func ReadJSONFromFile(filename string) (map[string]RequestData, error) {
	data, _ := decryptAndReadFromFile(filename, "01234567890123456789012345678901")
	Data, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("JSON 编码失败：%v", err)
	}

	var jsonData map[string]RequestData
	decoder := json.NewDecoder(bytes.NewReader(Data))
	if err := decoder.Decode(&jsonData); err != nil {
		return nil, fmt.Errorf("解码JSON失败：%v", err)
	}

	return jsonData, nil
}
