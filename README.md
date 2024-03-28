# terraform-plugin

## 简介

`terraform-plugin` 是一个支持阿里云、火山引擎和华为云的 ECS 开通的 Terraform 插件，它使用 gRPC 进行调用。该插件可以帮助用户在 Terraform 中开通云上的 ECS 资源。

## 特性

- 支持阿里云、火山引擎和华为云的 ECS 资源管理
- 使用 gRPC 进行通信，提供高效的接口调用
- 具有灵活的配置选项，可以满足不同用户的需求
- AK/SK 从数据库中获取，无需在代码中配置

## 快速开始

1. 首先，确保MySQL数据库中有对应云服务商的账号和访问密钥。

2. 环境变量配置：

```
export OPSHELPER_RDS_HOST=xxx.xxx.xxx.xxx
export OPSHELPER_RDS_DBNAME=xxxx
export OPSHELPER_RDS_USERNAME=xxxx
export OPSHELPER_RDS_PASSWORD=12345678
```

3. 安装插件：

```
go run cmd/main.go 
```

4. gRPC 接口调用：

``` 
temp.py 示例文件是生成插件读取的配置文件，需要根据实际情况进行修改。
exec.py 示例文件是调用插件的接口，需要根据实际情况进行修改。
```