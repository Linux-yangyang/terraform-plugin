package config

import (
	"os"

	"log"

	"gopkg.in/yaml.v3"
)

// Config 整个项目的配置
type Config struct {
	Mode          string `yaml:"mode"`
	Port          string `yaml:"port"`
	*LogConfig    `yaml:"log"`
	*TerraformDir `yaml:"terraformDir"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `yaml:"level"`
	Filename   string `yaml:"filename"`
	MaxSize    int    `yaml:"maxsize"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"`
}

// TerraformDir Terraform生成配置
type TerraformDir struct {
	Root     string   `yaml:"root"`
	Cloud    []string `yaml:"cloud"`
	Plugging []string `yaml:"plugging"`
}

// Conf 全局配置变量
var Conf = new(Config)

// Init 初始化配置；从指定文件加载配置文件
func Init() error {
	filePath := "conf/config.yaml"
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return yaml.Unmarshal(b, Conf)
}
