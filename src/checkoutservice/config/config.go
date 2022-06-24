package config

import (
	"fmt"

	"github.com/pkg/errors"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/source/env"
)

type Config struct {
	Port                  int
	CartService           string
	CurrencyService       string
	EmailService          string
	PaymentService        string
	ProductCatalogService string
	ShippingService       string
}

var cfg *Config = &Config{
	Port:                  5050,
	CartService:           "cartservice",
	CurrencyService:       "currencyservice",
	EmailService:          "emailservice",
	PaymentService:        "paymentservice",
	ProductCatalogService: "productcatalogservice",
	ShippingService:       "shippingservice",
}

func Get() Config {
	return *cfg
}

func Address() string {
	return fmt.Sprintf(":%d", cfg.Port)
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
