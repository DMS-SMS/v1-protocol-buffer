// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth-event.proto

package authproto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Client API for AuthEvent service

type AuthEventService interface {
	ChangeAllServiceNodes(ctx context.Context, in *Empty, opts ...client.CallOption) (*Empty, error)
}

type authEventService struct {
	c    client.Client
	name string
}

func NewAuthEventService(name string, c client.Client) AuthEventService {
	return &authEventService{
		c:    c,
		name: name,
	}
}

func (c *authEventService) ChangeAllServiceNodes(ctx context.Context, in *Empty, opts ...client.CallOption) (*Empty, error) {
	req := c.c.NewRequest(c.name, "AuthEvent.ChangeAllServiceNodes", in)
	out := new(Empty)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthEvent service

type AuthEventHandler interface {
	ChangeAllServiceNodes(context.Context, *Empty, *Empty) error
}

func RegisterAuthEventHandler(s server.Server, hdlr AuthEventHandler, opts ...server.HandlerOption) error {
	type authEvent interface {
		ChangeAllServiceNodes(ctx context.Context, in *Empty, out *Empty) error
	}
	type AuthEvent struct {
		authEvent
	}
	h := &authEventHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthEvent{h}, opts...))
}

type authEventHandler struct {
	AuthEventHandler
}

func (h *authEventHandler) ChangeAllServiceNodes(ctx context.Context, in *Empty, out *Empty) error {
	return h.AuthEventHandler.ChangeAllServiceNodes(ctx, in, out)
}
