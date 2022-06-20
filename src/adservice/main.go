package main

import (
	"github.com/go-micro/demo/adservice/config"
	"github.com/go-micro/demo/adservice/handler"
	pb "github.com/go-micro/demo/adservice/proto"
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "adservice"
	version = "1.0.0"
)

func main() {
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}
	srv := micro.NewService(micro.Server(grpcs.NewServer()), micro.Client(grpcc.NewClient()))
	srv.Init(micro.Name(service), micro.Version(version), micro.Address(config.Address()))
	if err := pb.RegisterAdServiceHandler(srv.Server(), new(handler.AdService)); err != nil {
		logger.Fatal(err)
	}
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
