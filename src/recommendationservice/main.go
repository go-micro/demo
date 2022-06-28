package main

import (
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/go-micro/demo/recommendationservice/config"
	"github.com/go-micro/demo/recommendationservice/handler"
	pb "github.com/go-micro/demo/recommendationservice/proto"
)

var (
	service = "recommendationservice"
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
		micro.Name(service),
		micro.Version(version),
		micro.Address(config.Address()),
	)

	// Register handler
	cfg, client := config.Get(), srv.Client()
	recommendationservice := &handler.RecommendationService{
		ProductCatalogService: pb.NewProductCatalogService(cfg.ProductCatalogService, client),
	}
	if err := pb.RegisterRecommendationServiceHandler(srv.Server(), recommendationservice); err != nil {
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
