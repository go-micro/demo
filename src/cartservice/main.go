package main

import (
	"log"

	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/etcd"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/go-micro/demo/cartservice/cartstore"
	"github.com/go-micro/demo/cartservice/config"
	"github.com/go-micro/demo/cartservice/handler"
	pb "github.com/go-micro/demo/cartservice/proto"
)

var (
	name    = "cartservice"
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
	if err := pb.RegisterCartServiceHandler(srv.Server(), &handler.CartService{Store: cartstore.NewMemoryCartStore()}); err != nil {
		log.Fatal(err)
	}
	if err := pb.RegisterHealthHandler(srv.Server(), new(handler.Health)); err != nil {
		log.Fatal(err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
