package main

import (
	grpcc "github.com/go-micro/plugins/v4/client/grpc"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	grpcs "github.com/go-micro/plugins/v4/server/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

	"github.com/go-micro/demo/checkoutservice/config"
	"github.com/go-micro/demo/checkoutservice/handler"
	pb "github.com/go-micro/demo/checkoutservice/proto"
)

var (
	name    = "checkoutservice"
	version = "latest"
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
	cfg, client := config.Get(), srv.Client()
	checkoutService := &handler.CheckoutService{
		CartService:           pb.NewCartService(cfg.CartService, client),
		CurrencyService:       pb.NewCurrencyService(cfg.CurrencyService, client),
		EmailService:          pb.NewEmailService(cfg.EmailService, client),
		PaymentService:        pb.NewPaymentService(cfg.PaymentService, client),
		ProductCatalogService: pb.NewProductCatalogService(cfg.ProductCatalogService, client),
		ShippingService:       pb.NewShippingService(cfg.ShippingService, client),
	}
	if err := pb.RegisterCheckoutServiceHandler(srv.Server(), checkoutService); err != nil {
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
