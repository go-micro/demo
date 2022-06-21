package main

import (
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/go-micro/demo/cartservice/handler"
	pb "github.com/go-micro/demo/cartservice/proto"
)

var (
	service = "cartservice"
	version = "1.0.0"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Server(grpcs.NewServer()),
		micro.Client(grpcc.NewClient()),
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterCartServiceHandler(srv.Server(), new(handler.CartService))
	pb.RegisterHealthHandler(srv.Server(), new(handler.Health))

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
