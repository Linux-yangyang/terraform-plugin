package service

import (
	"context"
	pb "tuyoops/go/terraform-plugin/api"
)

type Service interface {
	RenderHandleRequest(ctx context.Context, req *pb.RenderTerraformCodeRequest) (*pb.RenderTerraformCodeResponse, error)
	ExecuteHandleRequest(req *pb.ExecuteTerraformCmdRequest) error
	ResourceInfoRequest(ctx context.Context, req *pb.ResourceInfoRequest) (*pb.ResourceInfoResponse, error)
}
