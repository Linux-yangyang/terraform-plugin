package util

import (
	"os"
	"path/filepath"
	"tuyoops/go/terraform-plugin/config"
)

var directories = map[string]map[string]string{
	"var": {},
	"huawei/ecs": {
		"main.tf":      huaweiEcsRootMainTFContent,
		"variables.tf": huaweiEcsRootVariablesTFContent,
	},
	"huawei/data": {},
	"huawei/modules/eip": {
		"main.tf":      huaweiEcsEipMainTFContent,
		"variables.tf": huaweiEcsEipVariablesTFContent,
		"outputs.tf":   huaweiEcsEipOutputsTFContent,
	},
	"huawei/modules/image": {
		"main.tf":      huaweiEcsImageMainTFContent,
		"variables.tf": huaweiEcsImageVariablesTFContent,
		"outputs.tf":   huaweiEcsImageOutputsTFContent,
	},
	"huawei/modules/instance": {
		"main.tf":      huaweiEcsInstanceMainTFContent,
		"variables.tf": huaweiEcsInstanceVariablesTFContent,
		"outputs.tf":   huaweiEcsInstanceOutputsTFContent,
	},
	"huawei/modules/project": {
		"main.tf":      huaweiEcsProjectMainTFContent,
		"variables.tf": huaweiEcsProjectVariablesTFContent,
		"outputs.tf":   huaweiEcsProjectOutputsTFContent,
	},
	"huawei/modules/secgroup": {
		"main.tf":      huaweiEcsSecgroupMainTFContent,
		"variables.tf": huaweiEcsSecgroupVariablesTFContent,
		"outputs.tf":   huaweiEcsSecgroupOutputsTFContent,
	},
	"huawei/modules/subnet": {
		"main.tf":      huaweiEcsSubnetMainTFContent,
		"variables.tf": huaweiEcsSubnetVariablesTFContent,
		"outputs.tf":   huaweiEcsSubnetOutputsTFContent,
	},

	"aliyun/ecs": {
		"main.tf":      aliyunEcsRootMainTFContent,
		"variables.tf": aliyunEcsRootVariablesTFContent,
	},
	"aliyun/data": {},
	"aliyun/modules/eip": {
		"main.tf":      aliyunEcsEipMainTFContent,
		"variables.tf": aliyunEcsEipVariablesTFContent,
		"outputs.tf":   aliyunEcsEipOutputsTFContent,
	},
	"aliyun/modules/image": {
		"main.tf":      aliyunEcsImageMainTFContent,
		"variables.tf": aliyunEcsImageVariablesTFContent,
		"outputs.tf":   aliyunEcsImageOutputsTFContent,
	},
	"aliyun/modules/instance": {
		"main.tf":      aliyunEcsInstanceMainTFContent,
		"variables.tf": aliyunEcsInstanceVariablesTFContent,
		"outputs.tf":   aliyunEcsInstanceOutputsTFContent,
	},
	"aliyun/modules/project": {
		"main.tf":      aliyunEcsProjectMainTFContent,
		"variables.tf": aliyunEcsProjectVariablesTFContent,
		"outputs.tf":   aliyunEcsProjectOutputsTFContent,
	},
	"aliyun/modules/secgroup": {
		"main.tf":      aliyunEcsSecgroupMainTFContent,
		"variables.tf": aliyunEcsSecgroupVariablesTFContent,
		"outputs.tf":   aliyunEcsSecgroupOutputsTFContent,
	},
	"aliyun/modules/subnet": {
		"main.tf":      aliyunEcsSubnetMainTFContent,
		"variables.tf": aliyunEcsSubnetVariablesTFContent,
		"outputs.tf":   aliyunEcsSubnetOutputsTFContent,
	},

	"volcengine/ecs": {
		"main.tf":      volcengineEcsRootMainTFContent,
		"variables.tf": volcengineEcsRootVariablesTFContent,
	},
	"volcengine/data": {},
	"volcengine/modules/eip": {
		"main.tf":      volcengineEcsEipMainTFContent,
		"variables.tf": volcengineEcsEipVariablesTFContent,
		"outputs.tf":   volcengineEcsEipOutputsTFContent,
	},
	"volcengine/modules/image": {
		"main.tf":      volcengineEcsImageMainTFContent,
		"variables.tf": volcengineEcsImageVariablesTFContent,
		"outputs.tf":   volcengineEcsImageOutputsTFContent,
	},
	"volcengine/modules/instance": {
		"main.tf":      volcengineEcsInstanceMainTFContent,
		"variables.tf": volcengineEcsInstanceVariablesTFContent,
		"outputs.tf":   volcengineEcsInstanceOutputsTFContent,
	},
	"volcengine/modules/project": {
		"main.tf":      volcengineEcsProjectMainTFContent,
		"variables.tf": volcengineEcsProjectVariablesTFContent,
		"outputs.tf":   volcengineEcsProjectOutputsTFContent,
	},
	"volcengine/modules/secgroup": {
		"main.tf":      volcengineEcsSecgroupMainTFContent,
		"variables.tf": volcengineEcsSecgroupVariablesTFContent,
		"outputs.tf":   volcengineEcsSecgroupOutputsTFContent,
	},
	"volcengine/modules/subnet": {
		"main.tf":      volcengineEcsSubnetMainTFContent,
		"variables.tf": volcengineEcsSubnetVariablesTFContent,
		"outputs.tf":   volcengineEcsSubnetOutputsTFContent,
	},
}

var huaweiEcsEipMainTFContent = `terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.61.1"
    }
  }
}

provider "huaweicloud" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Project模块
module "project" {
  source     = "./../project/"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-project    = var.my-project
  my-region     = var.my-region
}

resource "huaweicloud_vpc_eip" "tf_eip" {
  provider = huaweicloud.my_region
  count = var.create_eip ? var.my-instance-count : 0
  enterprise_project_id = module.project.project_id
  auto_renew = true
  charging_mode = "postPaid"
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "${var.my-eip-name}-${count.index + 1}"
    size        = 300
    share_type  = "PER"
    charge_mode = "traffic"
  }
}
`
var huaweiEcsEipVariablesTFContent = `variable "my-access-key" {
  description = "华为云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "华为云 Secret Key"
  type        = string
}

variable "my-project" {
  description = "华为云 Project Name"
  type        = string
}

variable "my-region" {
  description = "华为云 Region"
  type        = string
}

variable "create_eip" {
  description = "华为云 EIP"
  type        = bool
  default     = false
}

variable "my-instance-count" {
  description = "华为云 设备数量"
  type        = string
}

variable "my-eip-name" {
  description = "华为云 Eip Name"
  type        = string
}
`
var huaweiEcsEipOutputsTFContent = `output "eip_ip" {
  value = can(huaweicloud_vpc_eip.tf_eip) ? [for eip in huaweicloud_vpc_eip.tf_eip : eip.address] : []
}

output "eip_id" {
  value = can(huaweicloud_vpc_eip.tf_eip) ? [for eip in huaweicloud_vpc_eip.tf_eip : eip.id] : []
}`

var huaweiEcsImageMainTFContent = `terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.61.1"
    }
  }
}

provider "huaweicloud" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "huaweicloud_images_image" "tf_image" {
  provider = huaweicloud.my_region
  name        = var.my-image
  most_recent = true
}
`
var huaweiEcsImageVariablesTFContent = `variable "my-access-key" {
  description = "华为云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "华为云 Secret Key"
  type        = string
}

variable "my-image" {
  description = "华为云 Image Name "
  type        = string
}

variable "my-region" {
  description = "华为云 Region"
  type        = string
}`
var huaweiEcsImageOutputsTFContent = `output "image_id" {
  value = data.huaweicloud_images_image.tf_image.id
}`

var huaweiEcsInstanceMainTFContent = `terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.61.1"
    }
  }
}

provider "huaweicloud" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Project模块
module "project" {
  source     = "./../project"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-project    = var.my-project
  my-region     = var.my-region
}

# 引入Image模块
module "image" {
  source     = "./../image"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-image = var.my-image
  my-region     = var.my-region
}

# 引入Secgroup模块
module "secgroup" {
  source     = "./../secgroup"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-secgroup   = var.my-secgroup
  my-region     = var.my-region
}

# 引入Subnet模块
module "subnet" {
  source     = "./../subnet"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-subnet     = var.my-subnet
  my-region     = var.my-region
}

# 引入Eip模块
module "eip" {
  source     = "./../eip"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-project     = var.my-project
  my-region     = var.my-region
  my-eip-name   = var.my-eip-name
  create_eip    = var.create_eip
  my-instance-count = var.my-instance-count
}

resource "huaweicloud_compute_instance" "tf-instance" {
  count    = var.my-instance-count
  provider = huaweicloud.my_region
  name = "${var.my-instance-name}-${count.index + 1}"
  image_id           = module.image.image_id
  flavor_id          = var.my-flavor-id
  availability_zone = var.my-availability-zone
  system_disk_type = var.my-sys-disk-type
  system_disk_size = var.my-sys-disk-size
  delete_disks_on_termination = true
  delete_eip_on_termination = true
  charging_mode = "prePaid"
  security_group_ids = [module.secgroup.secgroup_id]
  enterprise_project_id = module.project.project_id
  period = 1
  eip_id = length(module.eip.eip_id) > 0 ? element(module.eip.eip_id, count.index) : null
  period_unit = "month"
  auto_renew = true
  user_data = "#!/bin/bash\nwget https://ops-oss.obs.cn-north-4.myhuaweicloud.com/ops_sync_obs/init_system/init_master.sh -U fe9daf961b2b2190aeb57dff823c0649 -O /tmp/init_master.sh\nbash /tmp/init_master.sh ${var.my-instance-name} ${var.my-env}"
  dynamic "data_disks" {
    for_each = var.my-data-disk-size != "0" ? [1] : []
    content {
      type = var.my-data-disk-type
      size = var.my-data-disk-size
    }
  }
  network {
    uuid = module.subnet.subnet_id
  }
}

#resource "huaweicloud_compute_eip_associate" "associated" {
#  count    = var.create_eip ? var.my-instance-count : 0
#  provider = huaweicloud.my_region
#  public_ip   = module.eip.eip_ip[count.index]
#  instance_id = huaweicloud_compute_instance.tf-instance[count.index].id
#}`
var huaweiEcsInstanceVariablesTFContent = `variable "my-access-key" {
  description = "华为云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "华为云 Secret Key"
  type        = string
}

variable "my-region" {
  description = "华为云 Region"
  type        = string
}

variable "my-image" {
  description = "华为云 Image Name "
  type        = string
}

variable "my-project" {
  description = "华为云 Project Name"
  type        = string
}

variable "my-secgroup" {
  description = "华为云 Secgroup Name"
  type        = string
}

variable "my-subnet" {
  description = "华为云 Subnet Name"
  type        = string
}

variable "my-env" {
  description = "华为云 环境"
  type        = string
}

variable "my-data-disk-type" {
  description = "华为云 磁盘类型"
  type        = string
}

variable "my-data-disk-size" {
  description = "华为云 磁盘大小"
  type        = string
}

variable "create_eip" {
  description = "华为云 EIP"
  type        = bool
  default     = false
}
variable "my-eip-name" {
  description = "华为云 Eip Name"
  type        = string
}
variable "my-instance-count" {
  description = "华为云 设备数量"
  type        = string
}
variable "my-sys-disk-type" {
  description = "华为云 系统磁盘类型"
  type        = string
}
variable "my-sys-disk-size" {
  description = "华为云 系统磁盘大小"
  type        = string
}
variable "my-flavor-id" {
  description = "华为云  设备配置ID"
  type        = string
}
variable "my-availability-zone" {
  description = "华为云 可用区"
  type        = string
}
variable "my-instance-name" {
  description = "华为云 Instance Name"
  type        = string
}`
var huaweiEcsInstanceOutputsTFContent = ""

var huaweiEcsProjectMainTFContent = `terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.61.1"
    }
  }
}

provider "huaweicloud" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "huaweicloud_enterprise_project" "tf_project" {
  provider = huaweicloud.my_region
  name     = var.my-project
}`
var huaweiEcsProjectVariablesTFContent = `variable "my-access-key" {
  description = "华为云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "华为云 Secret Key"
  type        = string
}

variable "my-project" {
  description = "华为云 Project Name"
  type        = string
}

variable "my-region" {
  description = "华为云 Region"
  type        = string
}`
var huaweiEcsProjectOutputsTFContent = `output "project_id" {
  value = data.huaweicloud_enterprise_project.tf_project.id
}`

var huaweiEcsSecgroupMainTFContent = `terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.61.1"
    }
  }
}

provider "huaweicloud" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "huaweicloud_networking_secgroup" "tf_secgroup" {
  provider = huaweicloud.my_region
  name = var.my-secgroup
}`
var huaweiEcsSecgroupVariablesTFContent = `variable "my-access-key" {
  description = "华为云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "华为云 Secret Key"
  type        = string
}

variable "my-secgroup" {
  description = "华为云 Secgroup Name"
  type        = string
}

variable "my-region" {
  description = "华为云 Region"
  type        = string
}`
var huaweiEcsSecgroupOutputsTFContent = `output "secgroup_id" {
  value = data.huaweicloud_networking_secgroup.tf_secgroup.id
}`

var huaweiEcsSubnetMainTFContent = `terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.61.1"
    }
  }
}

provider "huaweicloud" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "huaweicloud_vpc_subnet" "tf_subnet" {
  provider = huaweicloud.my_region
  name = var.my-subnet
}`
var huaweiEcsSubnetVariablesTFContent = `variable "my-access-key" {
  description = "华为云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "华为云 Secret Key"
  type        = string
}

variable "my-subnet" {
  description = "华为云 Subnet Name"
  type        = string
}

variable "my-region" {
  description = "华为云 Region"
  type        = string
}`
var huaweiEcsSubnetOutputsTFContent = `output "subnet_id" {
  value = data.huaweicloud_vpc_subnet.tf_subnet.id
}`

var huaweiEcsRootMainTFContent = `terraform {
  backend "local" {
    path = "./../data/huawei.tfstate"
  }
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.61.1"
    }
  }
}

provider "huaweicloud" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Instance模块
module "instance" {
  source     = "./../modules/instance"
  my-region     = var.my-region
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-image      = var.my-image
  my-project    = var.my-project
  my-secgroup   = var.my-secgroup
  my-subnet     = var.my-subnet
  my-eip-name   = var.my-eip-name
  my-env        = var.my-env
  create_eip    = var.create_eip
  my-instance-name = var.my-instance-name
  my-data-disk-type = var.my-data-disk-type
  my-data-disk-size = var.my-data-disk-size
  my-instance-count = var.my-instance-count
  my-sys-disk-type = var.my-sys-disk-type
  my-sys-disk-size = var.my-sys-disk-size
  my-flavor-id = var.my-flavor-id
  my-availability-zone = var.my-availability-zone
}`
var huaweiEcsRootVariablesTFContent = `variable "my-access-key" {
  description = "华为云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "华为云 Secret Key"
  type        = string
}

variable "my-region" {
  description = "华为云 Region"
  type        = string
}

variable "my-image" {
  description = "华为云 Image Name "
  type        = string
}

variable "my-project" {
  description = "华为云 Project Name"
  type        = string
}

variable "my-secgroup" {
  description = "华为云 Secgroup Name"
  type        = string
}

variable "my-subnet" {
  description = "华为云 Subnet Name"
  type        = string
}

variable "my-data-disk-type" {
  description = "华为云 磁盘类型"
  type        = string
}

variable "my-data-disk-size" {
  description = "华为云 磁盘大小"
  type        = string
}

variable "my-env" {
  description = "华为云 环境"
  type        = string
}

variable "create_eip" {
  description = "华为云 EIP"
  type        = bool
  default     = false
}
variable "my-eip-name" {
  description = "华为云 Eip Name"
  type        = string
}

variable "my-instance-name" {
  description = "华为云 Instance Name"
  type        = string
}
variable "my-instance-count" {
  description = "华为云 设备数量"
  type        = string
}
variable "my-sys-disk-type" {
  description = "华为云 系统磁盘类型"
  type        = string
}
variable "my-sys-disk-size" {
  description = "华为云 系统磁盘大小"
  type        = string
}
variable "my-flavor-id" {
  description = "华为云  设备配置ID"
  type        = string
}
variable "my-availability-zone" {
  description = "华为云 可用区"
  type        = string
}`

var aliyunEcsEipMainTFContent = `terraform {
  required_providers {
    aliyun = {
      source  = "aliyun/alicloud"
      version = ">= 1.217.0"
    }
  }
}

provider "aliyun" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Project模块
module "project" {
  source     = "./../project/"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-project    = var.my-project
  my-region     = var.my-region
}

resource "alicloud_eip_address" "tf_eip" {
  provider = aliyun.my_region
  count = var.create_eip ? var.my-instance-count : 0
  resource_group_id = module.project.project_id
  auto_pay = true
  internet_charge_type = "PayByTraffic"
  isp = "BGP"
  netmode       = "public"
  address_name  = "${var.my-eip-name}-${count.index + 1}"
  bandwidth     = 200
}
`
var aliyunEcsEipVariablesTFContent = `variable "my-access-key" {
  description = "阿里云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "阿里云 Secret Key"
  type        = string
}

variable "my-project" {
  description = "阿里云 Project Name"
  type        = string
}

variable "my-region" {
  description = "阿里云 Region"
  type        = string
}

variable "create_eip" {
  description = "阿里云 EIP"
  type        = bool
  default     = false
}

variable "my-instance-count" {
  description = "阿里云 设备数量"
  type        = string
}

variable "my-eip-name" {
  description = "阿里云 Eip Name"
  type        = string
}
`
var aliyunEcsEipOutputsTFContent = `output "eip_ip" {
  value = can(alicloud_eip_address.tf_eip) ? [for eip in alicloud_eip_address.tf_eip : eip.id] : []
}`

var aliyunEcsImageMainTFContent = `terraform {
  required_providers {
    aliyun = {
      source  = "aliyun/alicloud"
      version = ">= 1.217.0"
    }
  }
}

provider "aliyun" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "alicloud_images" "tf_image" {
  provider   = aliyun.my_region
  name_regex = var.my-image
}
`
var aliyunEcsImageVariablesTFContent = `variable "my-access-key" {
  description = "阿里云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "阿里云 Secret Key"
  type        = string
}

variable "my-image" {
  description = "阿里云 Image Name "
  type        = string
}

variable "my-region" {
  description = "阿里云 Region"
  type        = string
}`
var aliyunEcsImageOutputsTFContent = `output "image_id" {
  value = data.alicloud_images.tf_image.images.0.id
}`

var aliyunEcsInstanceMainTFContent = `terraform {
  required_providers {
    aliyun = {
      source  = "aliyun/alicloud"
      version = ">= 1.217.0"
    }
  }
}

provider "aliyun" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Project模块
module "project" {
  source     = "./../project"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-project    = var.my-project
  my-region     = var.my-region
}

# 引入Image模块
module "image" {
  source     = "./../image"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-image = var.my-image
  my-region     = var.my-region
}

# 引入Secgroup模块
module "secgroup" {
  source     = "./../secgroup"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-secgroup   = var.my-secgroup
  my-region     = var.my-region
}

# 引入Subnet模块
module "subnet" {
  source     = "./../subnet"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-subnet     = var.my-subnet
  my-region     = var.my-region
}

# 引入Eip模块
# module "eip" {
#  source     = "./../eip"
#  my-secret-key = var.my-secret-key
#  my-access-key = var.my-access-key
#  my-project     = var.my-project
#  my-region     = var.my-region
#  my-eip-name   = var.my-eip-name
#  create_eip    = var.create_eip
#  my-instance-count = var.my-instance-count
# }

resource "alicloud_instance" "tf-instance" {
  count    = var.my-instance-count
  provider = aliyun.my_region
  instance_name = "${var.my-instance-name}-${count.index + 1}"
  host_name = "${var.my-instance-name}"
  instance_type = var.my-flavor-id
  image_id           = module.image.image_id
  instance_charge_type = "PrePaid"
  period = 1
  security_groups    = [module.secgroup.secgroup_id]
  availability_zone = var.my-availability-zone
  system_disk_size = var.my-sys-disk-size
  system_disk_category = var.my-sys-disk-type
  internet_charge_type = "PayByTraffic"
  internet_max_bandwidth_out = var.create_eip ? 100 : 0
  user_data = "#!/bin/bash\nwget https://ops-oss.obs.cn-north-4.myhuaweicloud.com/ops_sync_obs/init_system/init_master.sh -U fe9daf961b2b2190aeb57dff823c0649 -O /tmp/init_master.sh\nbash /tmp/init_master.sh ${var.my-instance-name} ${var.my-env}"
  dynamic "data_disks" {
    for_each = var.my-data-disk-size != "0" ? [1] : []
    content {
      category = var.my-data-disk-type
      size = var.my-data-disk-size
    }
  }
  vswitch_id = module.subnet.subnet_id
  resource_group_id = module.project.project_id
  renewal_status = "AutoRenewal"
}`
var aliyunEcsInstanceVariablesTFContent = `variable "my-access-key" {
  description = "阿里云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "阿里云 Secret Key"
  type        = string
}

variable "my-region" {
  description = "阿里云 Region"
  type        = string
}

variable "my-image" {
  description = "阿里云 Image Name "
  type        = string
}

variable "my-project" {
  description = "阿里云 Project Name"
  type        = string
}

variable "my-secgroup" {
  description = "阿里云 Secgroup Name"
  type        = string
}

variable "my-subnet" {
  description = "阿里云 Subnet Name"
  type        = string
}

variable "my-env" {
  description = "阿里云 环境"
  type        = string
}

variable "my-data-disk-type" {
  description = "阿里云 磁盘类型"
  type        = string
}

variable "my-data-disk-size" {
  description = "阿里云 磁盘大小"
  type        = string
}

variable "create_eip" {
  description = "阿里云 EIP"
  type        = bool
  default     = false
}
variable "my-eip-name" {
  description = "阿里云 Eip Name"
  type        = string
}
variable "my-instance-count" {
  description = "阿里云 设备数量"
  type        = string
}
variable "my-sys-disk-type" {
  description = "阿里云 系统磁盘类型"
  type        = string
}
variable "my-sys-disk-size" {
  description = "阿里云 系统磁盘大小"
  type        = string
}
variable "my-flavor-id" {
  description = "阿里云  设备配置ID"
  type        = string
}
variable "my-availability-zone" {
  description = "阿里云 可用区"
  type        = string
}
variable "my-instance-name" {
  description = "阿里云 Instance Name"
  type        = string
}`
var aliyunEcsInstanceOutputsTFContent = ""

var aliyunEcsProjectMainTFContent = `terraform {
  required_providers {
    aliyun = {
      source  = "aliyun/alicloud"
      version = ">= 1.217.0"
    }
  }
}

provider "aliyun" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "alicloud_resource_manager_resource_groups" "tf_project" {
  provider = aliyun.my_region
  name_regex  = var.my-project
}`
var aliyunEcsProjectVariablesTFContent = `variable "my-access-key" {
  description = "阿里云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "阿里云 Secret Key"
  type        = string
}

variable "my-project" {
  description = "阿里云 Project Name"
  type        = string
}

variable "my-region" {
  description = "阿里云 Region"
  type        = string
}`
var aliyunEcsProjectOutputsTFContent = `output "project_id" {
  value = data.alicloud_resource_manager_resource_groups.tf_project.ids.0
}`

var aliyunEcsSecgroupMainTFContent = `terraform {
  required_providers {
    aliyun = {
      source  = "aliyun/alicloud"
      version = ">= 1.217.0"
    }
  }
}

provider "aliyun" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "alicloud_security_groups" "tf_secgroup" {
  provider = aliyun.my_region
  name_regex = var.my-secgroup
}`
var aliyunEcsSecgroupVariablesTFContent = `variable "my-access-key" {
  description = "阿里云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "阿里云 Secret Key"
  type        = string
}

variable "my-secgroup" {
  description = "阿里云 Secgroup Name"
  type        = string
}

variable "my-region" {
  description = "阿里云 Region"
  type        = string
}`
var aliyunEcsSecgroupOutputsTFContent = `output "secgroup_id" {
  value = data.alicloud_security_groups.tf_secgroup.ids.0
}`

var aliyunEcsSubnetMainTFContent = `terraform {
  required_providers {
    aliyun = {
      source  = "aliyun/alicloud"
      version = ">= 1.217.0"
    }
  }
}

provider "aliyun" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "alicloud_vswitches" "tf_subnet" {
  provider = aliyun.my_region
  vswitch_name = var.my-subnet
}`
var aliyunEcsSubnetVariablesTFContent = `variable "my-access-key" {
  description = "阿里云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "阿里云 Secret Key"
  type        = string
}

variable "my-subnet" {
  description = "阿里云 Subnet Name"
  type        = string
}

variable "my-region" {
  description = "阿里云 Region"
  type        = string
}`
var aliyunEcsSubnetOutputsTFContent = `output "subnet_id" {
  value = data.alicloud_vswitches.tf_subnet.ids.0
}`

var aliyunEcsRootMainTFContent = `terraform {
  backend "local" {
    path = "./../data/aliyun.tfstate"
  }
  required_providers {
    aliyun = {
      source  = "aliyun/alicloud"
      version = ">= 1.217.0"
    }
  }
}

provider "aliyun" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Instance模块
module "instance" {
  source     = "./../modules/instance"
  my-region     = var.my-region
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-image      = var.my-image
  my-project    = var.my-project
  my-secgroup   = var.my-secgroup
  my-subnet     = var.my-subnet
  my-eip-name   = var.my-eip-name
  my-env        = var.my-env
  create_eip    = var.create_eip
  my-instance-name = var.my-instance-name
  my-data-disk-type = var.my-data-disk-type
  my-data-disk-size = var.my-data-disk-size
  my-instance-count = var.my-instance-count
  my-sys-disk-type = var.my-sys-disk-type
  my-sys-disk-size = var.my-sys-disk-size
  my-flavor-id = var.my-flavor-id
  my-availability-zone = var.my-availability-zone
}`
var aliyunEcsRootVariablesTFContent = `variable "my-access-key" {
  description = "阿里云 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "阿里云 Secret Key"
  type        = string
}

variable "my-region" {
  description = "阿里云 Region"
  type        = string
}

variable "my-image" {
  description = "阿里云 Image Name "
  type        = string
}

variable "my-project" {
  description = "阿里云 Project Name"
  type        = string
}

variable "my-secgroup" {
  description = "阿里云 Secgroup Name"
  type        = string
}

variable "my-env" {
  description = "阿里云 环境"
  type        = string
}

variable "my-subnet" {
  description = "阿里云 Subnet Name"
  type        = string
}

variable "my-data-disk-type" {
  description = "阿里云 磁盘类型"
  type        = string
}

variable "my-data-disk-size" {
  description = "阿里云 磁盘大小"
  type        = string
}

variable "create_eip" {
  description = "阿里云 EIP"
  type        = bool
  default     = false
}
variable "my-eip-name" {
  description = "阿里云 Eip Name"
  type        = string
}

variable "my-instance-name" {
  description = "阿里云 Instance Name"
  type        = string
}
variable "my-instance-count" {
  description = "阿里云 设备数量"
  type        = string
}
variable "my-sys-disk-type" {
  description = "阿里云 系统磁盘类型"
  type        = string
}
variable "my-sys-disk-size" {
  description = "阿里云 系统磁盘大小"
  type        = string
}
variable "my-flavor-id" {
  description = "阿里云  设备配置ID"
  type        = string
}
variable "my-availability-zone" {
  description = "阿里云 可用区"
  type        = string
}`

var volcengineEcsEipMainTFContent = `terraform {
  required_providers {
    volcengine = {
      source  = "volcengine/volcengine"
      version = ">= 0.0.140"
    }
  }
}

provider "volcengine" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Project模块
# module "project" {
#   source     = "./../project/"
#   my-secret-key = var.my-secret-key
#   my-access-key = var.my-access-key
#   my-project    = var.my-project
#   my-region     = var.my-region
# }

resource "volcengine_eip_address" "tf_eip" {
  provider = volcengine.my_region
  count = var.create_eip ? var.my-instance-count : 0
  billing_type = "PostPaidByTraffic"
  isp = "BGP"
  name  = var.my-eip-name
  bandwidth = 1000
  period = 1
  project_name = var.my-project
}
`
var volcengineEcsEipVariablesTFContent = `variable "my-access-key" {
  description = "火山引擎 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "火山引擎 Secret Key"
  type        = string
}

variable "my-project" {
  description = "火山引擎 Project Name"
  type        = string
}

variable "my-region" {
  description = "火山引擎 Region"
  type        = string
}

variable "create_eip" {
  description = "火山引擎 EIP"
  type        = bool
  default     = false
}

variable "my-instance-count" {
  description = "火山引擎 设备数量"
  type        = string
}

variable "my-eip-name" {
  description = "火山引擎 Eip Name"
  type        = string
}
`
var volcengineEcsEipOutputsTFContent = `output "eip_ip" {
  value = can(volcengine_eip_address.tf_eip) ? [for eip in volcengine_eip_address.tf_eip : eip.eip_address] : []
}
output "eip_id" {
  value = can(volcengine_eip_address.tf_eip) ? [for eip in volcengine_eip_address.tf_eip : eip.id] : []
}`

var volcengineEcsImageMainTFContent = `terraform {
  required_providers {
    volcengine = {
      source  = "volcengine/volcengine"
      version = ">= 0.0.140"
    }
  }
}

provider "volcengine" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "volcengine_images" "tf_image" {
  provider   = volcengine.my_region
  name_regex = var.my-image
}
`
var volcengineEcsImageVariablesTFContent = `variable "my-access-key" {
  description = "火山引擎 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "火山引擎 Secret Key"
  type        = string
}

variable "my-image" {
  description = "火山引擎 Image Name "
  type        = string
}

variable "my-region" {
  description = "火山引擎 Region"
  type        = string
}`
var volcengineEcsImageOutputsTFContent = `output "image_id" {
  value = data.volcengine_images.tf_image.images.0.image_id
}`

var volcengineEcsInstanceMainTFContent = `terraform {
  required_providers {
    volcengine = {
      source  = "volcengine/volcengine"
      version = ">= 0.0.140"
    }
  }
}

provider "volcengine" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Project模块
# module "project" {
#   source     = "./../project"
#   my-secret-key = var.my-secret-key
#   my-access-key = var.my-access-key
#   my-project    = var.my-project
#   my-region     = var.my-region
# }

# 引入Image模块
module "image" {
  source     = "./../image"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-image = var.my-image
  my-region = var.my-region
}

# 引入Secgroup模块
module "secgroup" {
  source        = "./../secgroup"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-secgroup   = var.my-secgroup
  my-region     = var.my-region
}

# 引入Subnet模块
module "subnet" {
  source     = "./../subnet"
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-subnet     = var.my-subnet
  my-region     = var.my-region
}

# 引入Eip模块
module "eip" {
 source     = "./../eip"
 my-secret-key = var.my-secret-key
 my-access-key = var.my-access-key
 my-project     = var.my-project
 my-region     = var.my-region
 my-eip-name   = var.my-eip-name
 create_eip    = var.create_eip
 my-instance-count = var.my-instance-count
}

resource "volcengine_ecs_instance" "tf-instance" {
  count    = var.my-instance-count
  provider = volcengine.my_region
  instance_name = "${var.my-instance-name}-${count.index + 1}"
  host_name = "${var.my-instance-name}"
  instance_type = var.my-flavor-id
  image_id           = module.image.image_id
  instance_charge_type = "PrePaid"
  period = 1
  password = "Tuyou@setting"
  security_group_ids  = [module.secgroup.secgroup_id]
  zone_id = var.my-availability-zone
  system_volume_size = var.my-sys-disk-size
  system_volume_type = var.my-sys-disk-type
  user_data = "#!/bin/bash\nwget https://ops-oss.obs.cn-north-4.myhuaweicloud.com/ops_sync_obs/init_system/init_master.sh -U fe9daf961b2b2190aeb57dff823c0649 -O /tmp/init_master.sh\nbash /tmp/init_master.sh ${var.my-instance-name} ${var.my-env}"
  dynamic "data_volumes" {
    for_each = var.my-data-disk-size != "0" ? [1] : []
    content {
      volume_type = var.my-data-disk-type
      size = var.my-data-disk-size
      delete_with_instance = true
    }
  }
  subnet_id = module.subnet.subnet_id
  project_name = var.my-project
}

resource "volcengine_eip_associate" "associated" {
  count    = var.create_eip ? var.my-instance-count : 0
  provider = volcengine.my_region
  allocation_id   = module.eip.eip_id[count.index]
  instance_id = volcengine_ecs_instance.tf-instance[count.index].instance_id
  instance_type = "EcsInstance"
}`
var volcengineEcsInstanceVariablesTFContent = `variable "my-access-key" {
  description = "火山引擎 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "火山引擎 Secret Key"
  type        = string
}

variable "my-region" {
  description = "火山引擎 Region"
  type        = string
}

variable "my-image" {
  description = "火山引擎 Image Name "
  type        = string
}

variable "my-project" {
  description = "火山引擎 Project Name"
  type        = string
}

variable "my-secgroup" {
  description = "火山引擎 Secgroup Name"
  type        = string
}

variable "my-subnet" {
  description = "火山引擎 Subnet Name"
  type        = string
}

variable "my-env" {
  description = "火山引擎 环境"
  type        = string
}

variable "my-data-disk-type" {
  description = "火山引擎 磁盘类型"
  type        = string
}

variable "my-data-disk-size" {
  description = "火山引擎 磁盘大小"
  type        = string
}

variable "create_eip" {
  description = "火山引擎 EIP"
  type        = bool
  default     = false
}
variable "my-eip-name" {
  description = "火山引擎 Eip Name"
  type        = string
}
variable "my-instance-count" {
  description = "火山引擎 设备数量"
  type        = string
}
variable "my-sys-disk-type" {
  description = "火山引擎 系统磁盘类型"
  type        = string
}
variable "my-sys-disk-size" {
  description = "火山引擎 系统磁盘大小"
  type        = string
}
variable "my-flavor-id" {
  description = "火山引擎  设备配置ID"
  type        = string
}
variable "my-availability-zone" {
  description = "火山引擎 可用区"
  type        = string
}
variable "my-instance-name" {
  description = "火山引擎 Instance Name"
  type        = string
}`
var volcengineEcsInstanceOutputsTFContent = ""

var volcengineEcsProjectMainTFContent = `terraform {
  required_providers {
    volcengine = {
      source  = "volcengine/volcengine"
      version = ">= 0.0.140"
    }
  }
}

provider "volcengine" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "alicloud_resource_manager_resource_groups" "tf_project" {
  provider = volcengine.my_region
  name_regex  = var.my-project
}`
var volcengineEcsProjectVariablesTFContent = `variable "my-access-key" {
  description = "火山引擎 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "火山引擎 Secret Key"
  type        = string
}

variable "my-project" {
  description = "火山引擎 Project Name"
  type        = string
}

variable "my-region" {
  description = "火山引擎 Region"
  type        = string
}`
var volcengineEcsProjectOutputsTFContent = `output "project_id" {
  value = data.alicloud_resource_manager_resource_groups.tf_project.ids.0
}`

var volcengineEcsSecgroupMainTFContent = `terraform {
  required_providers {
    volcengine = {
      source  = "volcengine/volcengine"
      version = ">= 0.0.140"
    }
  }
}

provider "volcengine" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "volcengine_security_groups" "tf_secgroup" {
  provider = volcengine.my_region
  name_regex = var.my-secgroup
}`
var volcengineEcsSecgroupVariablesTFContent = `variable "my-access-key" {
  description = "火山引擎 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "火山引擎 Secret Key"
  type        = string
}

variable "my-secgroup" {
  description = "火山引擎 Secgroup Name"
  type        = string
}

variable "my-region" {
  description = "火山引擎 Region"
  type        = string
}`
var volcengineEcsSecgroupOutputsTFContent = `output "secgroup_id" {
  value = data.volcengine_security_groups.tf_secgroup.security_groups.0.security_group_id
}`

var volcengineEcsSubnetMainTFContent = `terraform {
  required_providers {
    volcengine = {
      source  = "volcengine/volcengine"
      version = ">= 0.0.140"
    }
  }
}

provider "volcengine" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

data "volcengine_subnets" "tf_subnet" {
  provider = volcengine.my_region
  name_regex = var.my-subnet
}`
var volcengineEcsSubnetVariablesTFContent = `variable "my-access-key" {
  description = "火山引擎 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "火山引擎 Secret Key"
  type        = string
}

variable "my-subnet" {
  description = "火山引擎 Subnet Name"
  type        = string
}

variable "my-region" {
  description = "火山引擎 Region"
  type        = string
}`
var volcengineEcsSubnetOutputsTFContent = `output "subnet_id" {
  value = data.volcengine_subnets.tf_subnet.subnets.0.id
}`

var volcengineEcsRootMainTFContent = `terraform {
  backend "local" {
    path = "./../data/volcengine.tfstate"
  }
  required_providers {
    volcengine = {
      source  = "volcengine/volcengine"
      version = ">= 0.0.140"
    }
  }
}

provider "volcengine" {
  alias      = "my_region"
  region     = var.my-region
  access_key = var.my-access-key
  secret_key = var.my-secret-key
}

# 引入Instance模块
module "instance" {
  source     = "./../modules/instance"
  my-region     = var.my-region
  my-secret-key = var.my-secret-key
  my-access-key = var.my-access-key
  my-image      = var.my-image
  my-project    = var.my-project
  my-secgroup   = var.my-secgroup
  my-subnet     = var.my-subnet
  my-eip-name   = var.my-eip-name
  my-env        = var.my-env
  create_eip    = var.create_eip
  my-instance-name = var.my-instance-name
  my-data-disk-type = var.my-data-disk-type
  my-data-disk-size = var.my-data-disk-size
  my-instance-count = var.my-instance-count
  my-sys-disk-type = var.my-sys-disk-type
  my-sys-disk-size = var.my-sys-disk-size
  my-flavor-id = var.my-flavor-id
  my-availability-zone = var.my-availability-zone
}`
var volcengineEcsRootVariablesTFContent = `variable "my-access-key" {
  description = "火山引擎 Access Key"
  type        = string
}

variable "my-secret-key" {
  description = "火山引擎 Secret Key"
  type        = string
}

variable "my-region" {
  description = "火山引擎 Region"
  type        = string
}

variable "my-image" {
  description = "火山引擎 Image Name "
  type        = string
}

variable "my-project" {
  description = "火山引擎 Project Name"
  type        = string
}

variable "my-secgroup" {
  description = "火山引擎 Secgroup Name"
  type        = string
}

variable "my-env" {
  description = "火山引擎 环境"
  type        = string
}

variable "my-subnet" {
  description = "火山引擎 Subnet Name"
  type        = string
}

variable "my-data-disk-type" {
  description = "火山引擎 磁盘类型"
  type        = string
}

variable "my-data-disk-size" {
  description = "火山引擎 磁盘大小"
  type        = string
}

variable "create_eip" {
  description = "火山引擎 EIP"
  type        = bool
  default     = false
}
variable "my-eip-name" {
  description = "火山引擎 Eip Name"
  type        = string
}
variable "my-instance-name" {
  description = "火山引擎 Instance Name"
  type        = string
}
variable "my-instance-count" {
  description = "火山引擎 设备数量"
  type        = string
}
variable "my-sys-disk-type" {
  description = "火山引擎 系统磁盘类型"
  type        = string
}
variable "my-sys-disk-size" {
  description = "火山引擎 系统磁盘大小"
  type        = string
}
variable "my-flavor-id" {
  description = "火山引擎  设备配置ID"
  type        = string
}
variable "my-availability-zone" {
  description = "火山引擎 可用区"
  type        = string
}`

func CreateDirectoryStructure() error {
	rootDir := config.Conf.TerraformDir.Root
	for dir, filesContent := range directories {
		err := os.MkdirAll(filepath.Join(rootDir, dir), os.ModePerm)
		if err != nil {

			return err
		}

		for fileName, content := range filesContent {
			filePath := filepath.Join(rootDir, dir, fileName)
			err := os.WriteFile(filePath, []byte(content), os.ModePerm)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
