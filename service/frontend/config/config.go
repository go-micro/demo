package config

import (
	"github.com/pkg/errors"
	"go-micro.dev/v5/config"
	"go-micro.dev/v5/config/source/env"
)

type Config struct {
	Address               string
	Tracing               TracingConfig
	AdService             string
	CartService           string
	CheckoutService       string
	CurrencyService       string
	ProductCatalogService string
	RecommendationService string
	ShippingService       string
}

type TracingConfig struct {
	Enable bool
	Jaeger JaegerConfig
}

type JaegerConfig struct {
	URL string
}

var cfg *Config = &Config{
	Address:               ":8090",
	AdService:             "ad",
	CartService:           "cart",
	CheckoutService:       "checkout",
	CurrencyService:       "currency",
	ProductCatalogService: "productcatalog",
	RecommendationService: "recommendation",
	ShippingService:       "shipping",
}

func Get() Config {
	return *cfg
}

func Address() string {
	return cfg.Address
}

func Tracing() TracingConfig {
	return cfg.Tracing
}

func Load() error {
	configor, err := config.NewConfig(config.WithSource(env.NewSource()))
	if err != nil {
		return errors.Wrap(err, "configor.New")
	}
	if err := configor.Load(); err != nil {
		return errors.Wrap(err, "configor.Load")
	}
	if err := configor.Scan(cfg); err != nil {
		return errors.Wrap(err, "configor.Scan")
	}
	return nil
}
