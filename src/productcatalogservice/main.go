package main

import (
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/go-micro/demo/productcatalogservice/config"
	"github.com/go-micro/demo/productcatalogservice/handler"
	pb "github.com/go-micro/demo/productcatalogservice/proto"
)

var (
	name    = "productcatalogservice"
	version = "1.0.0"
)

func main() {
	// Load conigurations
	if err := config.Load(); err != nil {
		logger.Fatal(err)
	}

	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
	)
	srv.Init(
		micro.Name(name),
		micro.Version(version),
		micro.Address(config.Address()),
	)

	// Register handler
	if err := pb.RegisterProductCatalogServiceHandler(srv.Server(), new(handler.ProductCatalogService)); err != nil {
		logger.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		logger.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
