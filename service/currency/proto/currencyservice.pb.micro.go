// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/currency.proto

package hipstershop

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v5/api"
	client "go-micro.dev/v5/client"
	server "go-micro.dev/v5/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CurrencyService service

func NewCurrencyServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CurrencyService service

type CurrencyService interface {
	GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...client.CallOption) (*GetSupportedCurrenciesResponse, error)
	Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...client.CallOption) (*Money, error)
}

type currencyService struct {
	c    client.Client
	name string
}

func NewCurrencyService(name string, c client.Client) CurrencyService {
	return &currencyService{
		c:    c,
		name: name,
	}
}

func (c *currencyService) GetSupportedCurrencies(ctx context.Context, in *Empty, opts ...client.CallOption) (*GetSupportedCurrenciesResponse, error) {
	req := c.c.NewRequest(c.name, "CurrencyService.GetSupportedCurrencies", in)
	out := new(GetSupportedCurrenciesResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *currencyService) Convert(ctx context.Context, in *CurrencyConversionRequest, opts ...client.CallOption) (*Money, error) {
	req := c.c.NewRequest(c.name, "CurrencyService.Convert", in)
	out := new(Money)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CurrencyService service

type CurrencyServiceHandler interface {
	GetSupportedCurrencies(context.Context, *Empty, *GetSupportedCurrenciesResponse) error
	Convert(context.Context, *CurrencyConversionRequest, *Money) error
}

func RegisterCurrencyServiceHandler(s server.Server, hdlr CurrencyServiceHandler, opts ...server.HandlerOption) error {
	type currencyService interface {
		GetSupportedCurrencies(ctx context.Context, in *Empty, out *GetSupportedCurrenciesResponse) error
		Convert(ctx context.Context, in *CurrencyConversionRequest, out *Money) error
	}
	type CurrencyService struct {
		currencyService
	}
	h := &currencyServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CurrencyService{h}, opts...))
}

type currencyServiceHandler struct {
	CurrencyServiceHandler
}

func (h *currencyServiceHandler) GetSupportedCurrencies(ctx context.Context, in *Empty, out *GetSupportedCurrenciesResponse) error {
	return h.CurrencyServiceHandler.GetSupportedCurrencies(ctx, in, out)
}

func (h *currencyServiceHandler) Convert(ctx context.Context, in *CurrencyConversionRequest, out *Money) error {
	return h.CurrencyServiceHandler.Convert(ctx, in, out)
}
