// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: outing-parents.proto

package outingproto

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

// Client API for OutingParents service

type OutingParentsService interface {
	ApproveOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, opts ...client.CallOption) (*ConfirmOutingByOCodeResponse, error)
	RejectOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, opts ...client.CallOption) (*ConfirmOutingByOCodeResponse, error)
	GetOutingByOCode(ctx context.Context, in *GetOutingByOCodeRequest, opts ...client.CallOption) (*GetOutingByOCodeResponse, error)
}

type outingParentsService struct {
	c    client.Client
	name string
}

func NewOutingParentsService(name string, c client.Client) OutingParentsService {
	return &outingParentsService{
		c:    c,
		name: name,
	}
}

func (c *outingParentsService) ApproveOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, opts ...client.CallOption) (*ConfirmOutingByOCodeResponse, error) {
	req := c.c.NewRequest(c.name, "OutingParents.ApproveOutingByOCode", in)
	out := new(ConfirmOutingByOCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outingParentsService) RejectOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, opts ...client.CallOption) (*ConfirmOutingByOCodeResponse, error) {
	req := c.c.NewRequest(c.name, "OutingParents.RejectOutingByOCode", in)
	out := new(ConfirmOutingByOCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *outingParentsService) GetOutingByOCode(ctx context.Context, in *GetOutingByOCodeRequest, opts ...client.CallOption) (*GetOutingByOCodeResponse, error) {
	req := c.c.NewRequest(c.name, "OutingParents.GetOutingByOCode", in)
	out := new(GetOutingByOCodeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for OutingParents service

type OutingParentsHandler interface {
	ApproveOutingByOCode(context.Context, *ConfirmOutingByOCodeRequest, *ConfirmOutingByOCodeResponse) error
	RejectOutingByOCode(context.Context, *ConfirmOutingByOCodeRequest, *ConfirmOutingByOCodeResponse) error
	GetOutingByOCode(context.Context, *GetOutingByOCodeRequest, *GetOutingByOCodeResponse) error
}

func RegisterOutingParentsHandler(s server.Server, hdlr OutingParentsHandler, opts ...server.HandlerOption) error {
	type outingParents interface {
		ApproveOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, out *ConfirmOutingByOCodeResponse) error
		RejectOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, out *ConfirmOutingByOCodeResponse) error
		GetOutingByOCode(ctx context.Context, in *GetOutingByOCodeRequest, out *GetOutingByOCodeResponse) error
	}
	type OutingParents struct {
		outingParents
	}
	h := &outingParentsHandler{hdlr}
	return s.Handle(s.NewHandler(&OutingParents{h}, opts...))
}

type outingParentsHandler struct {
	OutingParentsHandler
}

func (h *outingParentsHandler) ApproveOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, out *ConfirmOutingByOCodeResponse) error {
	return h.OutingParentsHandler.ApproveOutingByOCode(ctx, in, out)
}

func (h *outingParentsHandler) RejectOutingByOCode(ctx context.Context, in *ConfirmOutingByOCodeRequest, out *ConfirmOutingByOCodeResponse) error {
	return h.OutingParentsHandler.RejectOutingByOCode(ctx, in, out)
}

func (h *outingParentsHandler) GetOutingByOCode(ctx context.Context, in *GetOutingByOCodeRequest, out *GetOutingByOCodeResponse) error {
	return h.OutingParentsHandler.GetOutingByOCode(ctx, in, out)
}
