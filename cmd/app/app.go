package app

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"tuyoops/go/terraform-plugin/pkg/util"

	"google.golang.org/grpc"

	pb "tuyoops/go/terraform-plugin/api"
	"tuyoops/go/terraform-plugin/config"
	"tuyoops/go/terraform-plugin/internal/provider"
	"tuyoops/go/terraform-plugin/internal/service"
	"tuyoops/go/terraform-plugin/pkg/logger"
)

type Application struct {
	ECSService *service.ECSService
	//LBService  *service.LBService
}

type terraformServiceServer struct {
	pb.UnimplementedTerraformServiceServer
	App *Application
}

func (t *terraformServiceServer) RenderTerraformCode(
	ctx context.Context,
	req *pb.RenderTerraformCodeRequest,
) (*pb.RenderTerraformCodeResponse, error) {
	return t.App.ECSService.RenderHandleRequest(req)
}

func (t *terraformServiceServer) ExecuteTerraformCmd(
	req *pb.ExecuteTerraformCmdRequest, stream pb.TerraformService_ExecuteTerraformCmdServer,
) error {
	return t.App.ECSService.ExecuteHandleRequest(stream, req)
}

func (t *terraformServiceServer) ResourceInfoGet(
	ctx context.Context,
	req *pb.ResourceInfoRequest,
) (*pb.ResourceInfoResponse, error) {
	return t.App.ECSService.ResourceInfoRequest(req)
}

func NewApplication(ecsProvider provider.ECSProvider) *Application {
	return &Application{
		ECSService: service.NewECSService(ecsProvider),
	}
}

func (app *Application) Run() {
	lis, err := net.Listen("tcp", ":"+config.Conf.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterTerraformServiceServer(grpcServer, &terraformServiceServer{App: app})

	log.Printf("terraform-pluging is running on port %s", config.Conf.Port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// Init 初始化 API
func init() {
	// load config from config.json
	if len(os.Args) < 1 {
		return
	}
	if err := config.Init(); err != nil {
		log.Fatal("Config init error:", err)
		return
	}
	if os.Getenv("OPSHELPER_RDS_PASSWORD") == "" {
		log.Fatalf("Error: OPSHELPER_RDS_PASSWORD environment variable is not set. Exiting.")
	}
	mysqlUtil, err := util.NewMySQLUtil()
	if err != nil {
		log.Fatal("Error creating MySQLUtil:", err)
	}
	defer mysqlUtil.Close()
	// init logger
	if err := logger.InitLogger(config.Conf.LogConfig); err != nil {
		log.Fatal("Init logger failed, err:%v\n", err)
		return
	}

	err = util.CreateDirectoryStructure()
	if err != nil {
		log.Fatal("Error:", err)
		return
	} else {
		log.Printf("Directory structure created successfully in %s", config.Conf.TerraformDir.Root)
	}
	// 初始化 terraform 的配置 terraform init
	for _, i := range config.Conf.TerraformDir.Cloud {
		for _, j := range config.Conf.TerraformDir.Plugging {
			hasTerraform, err := util.HasDir(filepath.Join(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, i, j), ".terraform"))
			if err != nil {
				log.Fatal("Error:", err)
			}
			if !hasTerraform {
				log.Printf("Terraform environment is not initialized... (%s)\n", i)
				util.ChangeWorkingDirectory(fmt.Sprintf("%s/%s/%s", config.Conf.TerraformDir.Root, i, j))
				err = util.RunCommand("terraform", "init")
				if err != nil {
					log.Fatal("Error:", err)
					return
				}
			}
		}
	}

}
