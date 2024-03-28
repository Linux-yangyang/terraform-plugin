package provider

import (
	"bytes"
	"text/template"
	pb "tuyoops/go/terraform-plugin/api"
	"tuyoops/go/terraform-plugin/pkg/util"
)

type HuaweiCloudLBProvider struct{}

func (p *HuaweiCloudLBProvider) GenerateLBCode(params map[string]interface{}) error {
	tmpl, err := template.New("ecsTemplate").Parse(huaweiCloudLBTemplate)
	if err != nil {
		return err
	}

	var codeBuffer bytes.Buffer

	if err := tmpl.Execute(&codeBuffer, params); err != nil {
		return err
	}

	return nil
}

func (p *HuaweiCloudLBProvider) ExecuteLBCode(stream pb.TerraformService_ExecuteTerraformCmdServer) error {
	err := util.RunTerraformCommand(stream, "plan", "-var 'my-access-key=1H8UW5V7WDIQTHPMOEBQ'",
		"-var 'my-secret-key=iivKJtHFspZrL08k98XGhI8e8StUHOx0GlE2KW9F'",
		"-var 'my-region=cn-north-9'",
		"-var 'my-image=CentOS 7.9 64bit'",
		"-var 'my-project=sa101-terraform-project-test'",
		"-var 'my-secgroup=default'",
		"-var 'my-subnet=subnet-ac02'",
		"-var 'my-instance-name=test'",
	)
	if err != nil {
		return err
	}
	return nil
}

const huaweiCloudLBTemplate = `
terraform {
  required_providers {
    huaweicloud = {
      source  = "huaweicloud/huaweicloud"
      version = ">= 1.49.0"
    }
  }
}

provider "huaweicloud" {
  alias      = "region_{{.region}}"
  region     = "{{.region}}"
  access_key = "{{.access_key}}"
  secret_key = "{{.secret_key}}"
}

data "huaweicloud_vpc_subnet" "tf_net" {
  provider = huaweicloud.region_{{.region}}
  name = "{{.subnet_name}}"
}

data "huaweicloud_images_image" "tf_image" {
  provider = huaweicloud.region_{{.region}}
  name        = "{{.image_name}}"
  most_recent = true
}

data "huaweicloud_networking_secgroup" "mysecgroup" {
  provider = huaweicloud.region_{{.region}}
  name = "{{.secgroup_name}}"
}

resource "huaweicloud_compute_instance" "basic" {
  provider = huaweicloud.region_{{.region}}
  name               = "{{.instance_name}}"
  image_id           = data.huaweicloud_images_image.tf_image.id
  flavor_id          = "{{.flavor_id}}"
  availability_zone = "{{.availability_zone}}"
  security_group_ids = [data.huaweicloud_networking_secgroup.mysecgroup.id]
  system_disk_type = "SSD"
  system_disk_size = "50"
  delete_disks_on_termination = "true"
  delete_eip_on_termination = "true"
  charging_mode = "prePaid"
  data_disk_type = "{{.data_disk_type}}"
  period = 1
  period_unit = "month"
  auto_renew = "true"
  user_data = "#!/bin/bash\nwget http://op.aliyun.tuyoo.com/init/init_master.sh -O /tmp/init_master.sh\nbash /tmp/init_master.sh 133-slm-game"
  data_disks {
    type = "SSD"
    size = "{{.data_disk_size}}"
  }
  network {
    uuid = data.huaweicloud_vpc_subnet.tf_net.id
  }
}

resource "huaweicloud_vpc_eip" "tf_eip" {
  provider = huaweicloud.region_{{.region}}
  publicip {
    type = "5_bgp"
  }
  bandwidth {
    name        = "{{.eip_bandwidth_name}}"
    size        = 300
    share_type  = "PER"
    charge_mode = "traffic"
  }
}

resource "huaweicloud_compute_eip_associate" "associated" {
  provider = huaweicloud.region_{{.region}}
  public_ip   = huaweicloud_vpc_eip.tf_eip.address
  instance_id = huaweicloud_compute_instance.basic.id
}
`
