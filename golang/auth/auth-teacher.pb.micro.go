// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth-teacher.proto

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

// Client API for AuthTeacher service

type AuthTeacherService interface {
	LoginTeacherAuth(ctx context.Context, in *LoginTeacherAuthRequest, opts ...client.CallOption) (*LoginTeacherAuthResponse, error)
	ChangeTeacherPW(ctx context.Context, in *ChangeTeacherPWRequest, opts ...client.CallOption) (*ChangeTeacherPWResponse, error)
	GetTeacherInformWithUUID(ctx context.Context, in *GetTeacherInformWithUUIDRequest, opts ...client.CallOption) (*GetTeacherInformWithUUIDResponse, error)
}

type authTeacherService struct {
	c    client.Client
	name string
}

func NewAuthTeacherService(name string, c client.Client) AuthTeacherService {
	return &authTeacherService{
		c:    c,
		name: name,
	}
}

func (c *authTeacherService) LoginTeacherAuth(ctx context.Context, in *LoginTeacherAuthRequest, opts ...client.CallOption) (*LoginTeacherAuthResponse, error) {
	req := c.c.NewRequest(c.name, "AuthTeacher.LoginTeacherAuth", in)
	out := new(LoginTeacherAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authTeacherService) ChangeTeacherPW(ctx context.Context, in *ChangeTeacherPWRequest, opts ...client.CallOption) (*ChangeTeacherPWResponse, error) {
	req := c.c.NewRequest(c.name, "AuthTeacher.ChangeTeacherPW", in)
	out := new(ChangeTeacherPWResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authTeacherService) GetTeacherInformWithUUID(ctx context.Context, in *GetTeacherInformWithUUIDRequest, opts ...client.CallOption) (*GetTeacherInformWithUUIDResponse, error) {
	req := c.c.NewRequest(c.name, "AuthTeacher.GetTeacherInformWithUUID", in)
	out := new(GetTeacherInformWithUUIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AuthTeacher service

type AuthTeacherHandler interface {
	LoginTeacherAuth(context.Context, *LoginTeacherAuthRequest, *LoginTeacherAuthResponse) error
	ChangeTeacherPW(context.Context, *ChangeTeacherPWRequest, *ChangeTeacherPWResponse) error
	GetTeacherInformWithUUID(context.Context, *GetTeacherInformWithUUIDRequest, *GetTeacherInformWithUUIDResponse) error
}

func RegisterAuthTeacherHandler(s server.Server, hdlr AuthTeacherHandler, opts ...server.HandlerOption) error {
	type authTeacher interface {
		LoginTeacherAuth(ctx context.Context, in *LoginTeacherAuthRequest, out *LoginTeacherAuthResponse) error
		ChangeTeacherPW(ctx context.Context, in *ChangeTeacherPWRequest, out *ChangeTeacherPWResponse) error
		GetTeacherInformWithUUID(ctx context.Context, in *GetTeacherInformWithUUIDRequest, out *GetTeacherInformWithUUIDResponse) error
	}
	type AuthTeacher struct {
		authTeacher
	}
	h := &authTeacherHandler{hdlr}
	return s.Handle(s.NewHandler(&AuthTeacher{h}, opts...))
}

type authTeacherHandler struct {
	AuthTeacherHandler
}

func (h *authTeacherHandler) LoginTeacherAuth(ctx context.Context, in *LoginTeacherAuthRequest, out *LoginTeacherAuthResponse) error {
	return h.AuthTeacherHandler.LoginTeacherAuth(ctx, in, out)
}

func (h *authTeacherHandler) ChangeTeacherPW(ctx context.Context, in *ChangeTeacherPWRequest, out *ChangeTeacherPWResponse) error {
	return h.AuthTeacherHandler.ChangeTeacherPW(ctx, in, out)
}

func (h *authTeacherHandler) GetTeacherInformWithUUID(ctx context.Context, in *GetTeacherInformWithUUIDRequest, out *GetTeacherInformWithUUIDResponse) error {
	return h.AuthTeacherHandler.GetTeacherInformWithUUID(ctx, in, out)
}
