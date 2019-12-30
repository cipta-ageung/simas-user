// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for AuthService service

type AuthService interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error)
	GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "auth"
	}
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuthService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Delete(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuthService.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Get(ctx context.Context, in *User, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuthService.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) GetAll(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "AuthService.GetAll", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "AuthService.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.name, "AuthService.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthService service

type AuthServiceHandler interface {
	Create(context.Context, *User, *Response) error
	Delete(context.Context, *User, *Response) error
	Get(context.Context, *User, *Response) error
	GetAll(context.Context, *Request, *Response) error
	Auth(context.Context, *User, *Token) error
	ValidateToken(context.Context, *Token, *Token) error
}

func RegisterAuthServiceHandler(s server.Server, hdlr AuthServiceHandler, opts ...server.HandlerOption) error {
	type authService interface {
		Create(ctx context.Context, in *User, out *Response) error
		Delete(ctx context.Context, in *User, out *Response) error
		Get(ctx context.Context, in *User, out *Response) error
		GetAll(ctx context.Context, in *Request, out *Response) error
		Auth(ctx context.Context, in *User, out *Token) error
		ValidateToken(ctx context.Context, in *Token, out *Token) error
	}
	type AuthService struct {
		authService
	}
	h := &authServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthService{h}, opts...))
}

type authServiceHandler struct {
	AuthServiceHandler
}

func (h *authServiceHandler) Create(ctx context.Context, in *User, out *Response) error {
	return h.AuthServiceHandler.Create(ctx, in, out)
}

func (h *authServiceHandler) Delete(ctx context.Context, in *User, out *Response) error {
	return h.AuthServiceHandler.Delete(ctx, in, out)
}

func (h *authServiceHandler) Get(ctx context.Context, in *User, out *Response) error {
	return h.AuthServiceHandler.Get(ctx, in, out)
}

func (h *authServiceHandler) GetAll(ctx context.Context, in *Request, out *Response) error {
	return h.AuthServiceHandler.GetAll(ctx, in, out)
}

func (h *authServiceHandler) Auth(ctx context.Context, in *User, out *Token) error {
	return h.AuthServiceHandler.Auth(ctx, in, out)
}

func (h *authServiceHandler) ValidateToken(ctx context.Context, in *Token, out *Token) error {
	return h.AuthServiceHandler.ValidateToken(ctx, in, out)
}
