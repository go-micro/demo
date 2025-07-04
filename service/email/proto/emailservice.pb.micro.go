// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/email.proto

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

// Api Endpoints for EmailService service

func NewEmailServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for EmailService service

type EmailService interface {
	SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...client.CallOption) (*Empty, error)
}

type emailService struct {
	c    client.Client
	name string
}

func NewEmailService(name string, c client.Client) EmailService {
	return &emailService{
		c:    c,
		name: name,
	}
}

func (c *emailService) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "EmailService.SendOrderConfirmation", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EmailService service

type EmailServiceHandler interface {
	SendOrderConfirmation(context.Context, *SendOrderConfirmationRequest, *Empty) error
}

func RegisterEmailServiceHandler(s server.Server, hdlr EmailServiceHandler, opts ...server.HandlerOption) error {
	type emailService interface {
		SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, out *Empty) error
	}
	type EmailService struct {
		emailService
	}
	h := &emailServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&EmailService{h}, opts...))
}

type emailServiceHandler struct {
	EmailServiceHandler
}

func (h *emailServiceHandler) SendOrderConfirmation(ctx context.Context, in *SendOrderConfirmationRequest, out *Empty) error {
	return h.EmailServiceHandler.SendOrderConfirmation(ctx, in, out)
}
