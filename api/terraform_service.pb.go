// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: terraform_service.proto

package __

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RenderTerraformCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region           string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	DataDiskSize     string `protobuf:"bytes,2,opt,name=data_disk_size,json=dataDiskSize,proto3" json:"data_disk_size,omitempty"`
	InstanceName     string `protobuf:"bytes,3,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	ImageName        string `protobuf:"bytes,4,opt,name=image_name,json=imageName,proto3" json:"image_name,omitempty"`
	FlavorId         string `protobuf:"bytes,5,opt,name=flavor_id,json=flavorId,proto3" json:"flavor_id,omitempty"`
	AvailabilityZone string `protobuf:"bytes,6,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
	EipBandwidthName string `protobuf:"bytes,7,opt,name=eip_bandwidth_name,json=eipBandwidthName,proto3" json:"eip_bandwidth_name,omitempty"`
	DataDiskType     string `protobuf:"bytes,8,opt,name=data_disk_type,json=dataDiskType,proto3" json:"data_disk_type,omitempty"`
	SubnetName       string `protobuf:"bytes,9,opt,name=subnet_name,json=subnetName,proto3" json:"subnet_name,omitempty"`
	SecGroup         string `protobuf:"bytes,10,opt,name=sec_group,json=secGroup,proto3" json:"sec_group,omitempty"`
	Cloud            string `protobuf:"bytes,11,opt,name=cloud,proto3" json:"cloud,omitempty"`
	ApprovalId       string `protobuf:"bytes,12,opt,name=approval_id,json=approvalId,proto3" json:"approval_id,omitempty"`
	ProjectName      string `protobuf:"bytes,13,opt,name=project_name,json=projectName,proto3" json:"project_name,omitempty"`
	InstanceCount    string `protobuf:"bytes,14,opt,name=instance_count,json=instanceCount,proto3" json:"instance_count,omitempty"`
	SysDiskType      string `protobuf:"bytes,15,opt,name=sys_disk_type,json=sysDiskType,proto3" json:"sys_disk_type,omitempty"`
	SysDiskSize      string `protobuf:"bytes,16,opt,name=sys_disk_size,json=sysDiskSize,proto3" json:"sys_disk_size,omitempty"`
}

func (x *RenderTerraformCodeRequest) Reset() {
	*x = RenderTerraformCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderTerraformCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderTerraformCodeRequest) ProtoMessage() {}

func (x *RenderTerraformCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderTerraformCodeRequest.ProtoReflect.Descriptor instead.
func (*RenderTerraformCodeRequest) Descriptor() ([]byte, []int) {
	return file_terraform_service_proto_rawDescGZIP(), []int{0}
}

func (x *RenderTerraformCodeRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetDataDiskSize() string {
	if x != nil {
		return x.DataDiskSize
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetImageName() string {
	if x != nil {
		return x.ImageName
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetFlavorId() string {
	if x != nil {
		return x.FlavorId
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetAvailabilityZone() string {
	if x != nil {
		return x.AvailabilityZone
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetEipBandwidthName() string {
	if x != nil {
		return x.EipBandwidthName
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetDataDiskType() string {
	if x != nil {
		return x.DataDiskType
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetSubnetName() string {
	if x != nil {
		return x.SubnetName
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetSecGroup() string {
	if x != nil {
		return x.SecGroup
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetCloud() string {
	if x != nil {
		return x.Cloud
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetApprovalId() string {
	if x != nil {
		return x.ApprovalId
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetProjectName() string {
	if x != nil {
		return x.ProjectName
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetInstanceCount() string {
	if x != nil {
		return x.InstanceCount
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetSysDiskType() string {
	if x != nil {
		return x.SysDiskType
	}
	return ""
}

func (x *RenderTerraformCodeRequest) GetSysDiskSize() string {
	if x != nil {
		return x.SysDiskSize
	}
	return ""
}

type ExecuteTerraformCmdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Todo       string `protobuf:"bytes,1,opt,name=todo,proto3" json:"todo,omitempty"`
	ApprovalId string `protobuf:"bytes,2,opt,name=approval_id,json=approvalId,proto3" json:"approval_id,omitempty"`
}

func (x *ExecuteTerraformCmdRequest) Reset() {
	*x = ExecuteTerraformCmdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteTerraformCmdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteTerraformCmdRequest) ProtoMessage() {}

func (x *ExecuteTerraformCmdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteTerraformCmdRequest.ProtoReflect.Descriptor instead.
func (*ExecuteTerraformCmdRequest) Descriptor() ([]byte, []int) {
	return file_terraform_service_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteTerraformCmdRequest) GetTodo() string {
	if x != nil {
		return x.Todo
	}
	return ""
}

func (x *ExecuteTerraformCmdRequest) GetApprovalId() string {
	if x != nil {
		return x.ApprovalId
	}
	return ""
}

type RenderTerraformCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCode  string `protobuf:"bytes,1,opt,name=successCode,proto3" json:"successCode,omitempty"`
	ErrorDetails string `protobuf:"bytes,2,opt,name=errorDetails,proto3" json:"errorDetails,omitempty"`
}

func (x *RenderTerraformCodeResponse) Reset() {
	*x = RenderTerraformCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenderTerraformCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenderTerraformCodeResponse) ProtoMessage() {}

func (x *RenderTerraformCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenderTerraformCodeResponse.ProtoReflect.Descriptor instead.
func (*RenderTerraformCodeResponse) Descriptor() ([]byte, []int) {
	return file_terraform_service_proto_rawDescGZIP(), []int{2}
}

func (x *RenderTerraformCodeResponse) GetSuccessCode() string {
	if x != nil {
		return x.SuccessCode
	}
	return ""
}

func (x *RenderTerraformCodeResponse) GetErrorDetails() string {
	if x != nil {
		return x.ErrorDetails
	}
	return ""
}

type ExecuteTerraformCmdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessMessage string `protobuf:"bytes,1,opt,name=successMessage,proto3" json:"successMessage,omitempty"`
	ErrorMessage   string `protobuf:"bytes,2,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	StreamEOF      string `protobuf:"bytes,3,opt,name=streamEOF,proto3" json:"streamEOF,omitempty"`
}

func (x *ExecuteTerraformCmdResponse) Reset() {
	*x = ExecuteTerraformCmdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteTerraformCmdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteTerraformCmdResponse) ProtoMessage() {}

func (x *ExecuteTerraformCmdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteTerraformCmdResponse.ProtoReflect.Descriptor instead.
func (*ExecuteTerraformCmdResponse) Descriptor() ([]byte, []int) {
	return file_terraform_service_proto_rawDescGZIP(), []int{3}
}

func (x *ExecuteTerraformCmdResponse) GetSuccessMessage() string {
	if x != nil {
		return x.SuccessMessage
	}
	return ""
}

func (x *ExecuteTerraformCmdResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *ExecuteTerraformCmdResponse) GetStreamEOF() string {
	if x != nil {
		return x.StreamEOF
	}
	return ""
}

type ResourceInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApprovalId string `protobuf:"bytes,1,opt,name=approval_id,json=approvalId,proto3" json:"approval_id,omitempty"`
}

func (x *ResourceInfoRequest) Reset() {
	*x = ResourceInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInfoRequest) ProtoMessage() {}

func (x *ResourceInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInfoRequest.ProtoReflect.Descriptor instead.
func (*ResourceInfoRequest) Descriptor() ([]byte, []int) {
	return file_terraform_service_proto_rawDescGZIP(), []int{4}
}

func (x *ResourceInfoRequest) GetApprovalId() string {
	if x != nil {
		return x.ApprovalId
	}
	return ""
}

type ResourceInfoStruct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InstanceName string `protobuf:"bytes,2,opt,name=instance_name,json=instanceName,proto3" json:"instance_name,omitempty"`
	PublicIp     string `protobuf:"bytes,3,opt,name=public_ip,json=publicIp,proto3" json:"public_ip,omitempty"`
	PrivateIp    string `protobuf:"bytes,4,opt,name=private_ip,json=privateIp,proto3" json:"private_ip,omitempty"`
	Provider     string `protobuf:"bytes,5,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *ResourceInfoStruct) Reset() {
	*x = ResourceInfoStruct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInfoStruct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInfoStruct) ProtoMessage() {}

func (x *ResourceInfoStruct) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInfoStruct.ProtoReflect.Descriptor instead.
func (*ResourceInfoStruct) Descriptor() ([]byte, []int) {
	return file_terraform_service_proto_rawDescGZIP(), []int{5}
}

func (x *ResourceInfoStruct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResourceInfoStruct) GetInstanceName() string {
	if x != nil {
		return x.InstanceName
	}
	return ""
}

func (x *ResourceInfoStruct) GetPublicIp() string {
	if x != nil {
		return x.PublicIp
	}
	return ""
}

func (x *ResourceInfoStruct) GetPrivateIp() string {
	if x != nil {
		return x.PrivateIp
	}
	return ""
}

func (x *ResourceInfoStruct) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

type ResourceInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceInfo map[string]*ResourceInfoStruct `protobuf:"bytes,1,rep,name=resourceInfo,proto3" json:"resourceInfo,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *ResourceInfoResponse) Reset() {
	*x = ResourceInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terraform_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceInfoResponse) ProtoMessage() {}

func (x *ResourceInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_terraform_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceInfoResponse.ProtoReflect.Descriptor instead.
func (*ResourceInfoResponse) Descriptor() ([]byte, []int) {
	return file_terraform_service_proto_rawDescGZIP(), []int{6}
}

func (x *ResourceInfoResponse) GetResourceInfo() map[string]*ResourceInfoStruct {
	if x != nil {
		return x.ResourceInfo
	}
	return nil
}

var File_terraform_service_proto protoreflect.FileDescriptor

var file_terraform_service_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0xc3,
	0x04, 0x0a, 0x1a, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f,
	0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x64, 0x69,
	0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64,
	0x61, 0x74, 0x61, 0x44, 0x69, 0x73, 0x6b, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x6c, 0x61, 0x76, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11,
	0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x7a, 0x6f, 0x6e,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x2c, 0x0a, 0x12, 0x65, 0x69, 0x70,
	0x5f, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x65, 0x69, 0x70, 0x42, 0x61, 0x6e, 0x64, 0x77, 0x69,
	0x64, 0x74, 0x68, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0e, 0x64, 0x61, 0x74, 0x61, 0x5f,
	0x64, 0x69, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x64, 0x61, 0x74, 0x61, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x63, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x65, 0x63, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x79, 0x73, 0x44, 0x69, 0x73, 0x6b, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x73, 0x79, 0x73, 0x5f, 0x64, 0x69, 0x73, 0x6b, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x79, 0x73, 0x44, 0x69, 0x73, 0x6b,
	0x53, 0x69, 0x7a, 0x65, 0x22, 0x51, 0x0a, 0x1a, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x6f, 0x64, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70,
	0x72, 0x6f, 0x76, 0x61, 0x6c, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x1b, 0x52, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x87, 0x01, 0x0a,
	0x1b, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72,
	0x6d, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x0e,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x45, 0x4f, 0x46, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x45, 0x4f, 0x46, 0x22, 0x36, 0x0a, 0x13, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x49, 0x64, 0x22, 0xa1,
	0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53,
	0x74, 0x72, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x69, 0x6e,
	0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x49, 0x70, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x69, 0x76, 0x61,
	0x74, 0x65, 0x5f, 0x69, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69,
	0x76, 0x61, 0x74, 0x65, 0x49, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x22, 0xc1, 0x01, 0x0a, 0x14, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x58, 0x0a, 0x11,
	0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x2d, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x90, 0x02, 0x0a, 0x10, 0x54, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x13, 0x52,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f,
	0x64, 0x65, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x54,
	0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x13, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x54, 0x65, 0x72, 0x72, 0x61, 0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6d, 0x64, 0x12, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x65, 0x72, 0x72, 0x61, 0x66,
	0x6f, 0x72, 0x6d, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x54, 0x65, 0x72, 0x72, 0x61,
	0x66, 0x6f, 0x72, 0x6d, 0x43, 0x6d, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x46, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x47, 0x65, 0x74, 0x12, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_terraform_service_proto_rawDescOnce sync.Once
	file_terraform_service_proto_rawDescData = file_terraform_service_proto_rawDesc
)

func file_terraform_service_proto_rawDescGZIP() []byte {
	file_terraform_service_proto_rawDescOnce.Do(func() {
		file_terraform_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_terraform_service_proto_rawDescData)
	})
	return file_terraform_service_proto_rawDescData
}

var file_terraform_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_terraform_service_proto_goTypes = []interface{}{
	(*RenderTerraformCodeRequest)(nil),  // 0: api.RenderTerraformCodeRequest
	(*ExecuteTerraformCmdRequest)(nil),  // 1: api.ExecuteTerraformCmdRequest
	(*RenderTerraformCodeResponse)(nil), // 2: api.RenderTerraformCodeResponse
	(*ExecuteTerraformCmdResponse)(nil), // 3: api.ExecuteTerraformCmdResponse
	(*ResourceInfoRequest)(nil),         // 4: api.ResourceInfoRequest
	(*ResourceInfoStruct)(nil),          // 5: api.ResourceInfoStruct
	(*ResourceInfoResponse)(nil),        // 6: api.ResourceInfoResponse
	nil,                                 // 7: api.ResourceInfoResponse.ResourceInfoEntry
}
var file_terraform_service_proto_depIdxs = []int32{
	7, // 0: api.ResourceInfoResponse.resourceInfo:type_name -> api.ResourceInfoResponse.ResourceInfoEntry
	5, // 1: api.ResourceInfoResponse.ResourceInfoEntry.value:type_name -> api.ResourceInfoStruct
	0, // 2: api.TerraformService.RenderTerraformCode:input_type -> api.RenderTerraformCodeRequest
	1, // 3: api.TerraformService.ExecuteTerraformCmd:input_type -> api.ExecuteTerraformCmdRequest
	4, // 4: api.TerraformService.ResourceInfoGet:input_type -> api.ResourceInfoRequest
	2, // 5: api.TerraformService.RenderTerraformCode:output_type -> api.RenderTerraformCodeResponse
	3, // 6: api.TerraformService.ExecuteTerraformCmd:output_type -> api.ExecuteTerraformCmdResponse
	6, // 7: api.TerraformService.ResourceInfoGet:output_type -> api.ResourceInfoResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_terraform_service_proto_init() }
func file_terraform_service_proto_init() {
	if File_terraform_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_terraform_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderTerraformCodeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_terraform_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteTerraformCmdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_terraform_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenderTerraformCodeResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_terraform_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteTerraformCmdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_terraform_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInfoRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_terraform_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInfoStruct); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_terraform_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceInfoResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_terraform_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_terraform_service_proto_goTypes,
		DependencyIndexes: file_terraform_service_proto_depIdxs,
		MessageInfos:      file_terraform_service_proto_msgTypes,
	}.Build()
	File_terraform_service_proto = out.File
	file_terraform_service_proto_rawDesc = nil
	file_terraform_service_proto_goTypes = nil
	file_terraform_service_proto_depIdxs = nil
}
