package main

import (
	s "tuyoops/go/terraform-plugin/cmd/app"
	"tuyoops/go/terraform-plugin/internal/provider"
)

func main() {
	ecsProvider := &provider.CloudECSProvider{}
	//lbProvider := &provider.CloudLBProvider{}
	app := s.NewApplication(ecsProvider)
	app.Run()
}
