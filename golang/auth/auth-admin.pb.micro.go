// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth-admin.proto

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

// Client API for Auth service

type AuthService interface {
	CreateStudentAuth(ctx context.Context, in *CreateStudentAuthRequest, opts ...client.CallOption) (*CreateStudentAuthResponse, error)
	CreateTeacherAuth(ctx context.Context, in *CreateTeacherAuthRequest, opts ...client.CallOption) (*CreateTeacherAuthResponse, error)
	CreateParentAuth(ctx context.Context, in *CreateParentAuthRequest, opts ...client.CallOption) (*CreateParentAuthResponse, error)
}

type authService struct {
	c    client.Client
	name string
}

func NewAuthService(name string, c client.Client) AuthService {
	return &authService{
		c:    c,
		name: name,
	}
}

func (c *authService) CreateStudentAuth(ctx context.Context, in *CreateStudentAuthRequest, opts ...client.CallOption) (*CreateStudentAuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateStudentAuth", in)
	out := new(CreateStudentAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateTeacherAuth(ctx context.Context, in *CreateTeacherAuthRequest, opts ...client.CallOption) (*CreateTeacherAuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateTeacherAuth", in)
	out := new(CreateTeacherAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authService) CreateParentAuth(ctx context.Context, in *CreateParentAuthRequest, opts ...client.CallOption) (*CreateParentAuthResponse, error) {
	req := c.c.NewRequest(c.name, "Auth.CreateParentAuth", in)
	out := new(CreateParentAuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	CreateStudentAuth(context.Context, *CreateStudentAuthRequest, *CreateStudentAuthResponse) error
	CreateTeacherAuth(context.Context, *CreateTeacherAuthRequest, *CreateTeacherAuthResponse) error
	CreateParentAuth(context.Context, *CreateParentAuthRequest, *CreateParentAuthResponse) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) error {
	type auth interface {
		CreateStudentAuth(ctx context.Context, in *CreateStudentAuthRequest, out *CreateStudentAuthResponse) error
		CreateTeacherAuth(ctx context.Context, in *CreateTeacherAuthRequest, out *CreateTeacherAuthResponse) error
		CreateParentAuth(ctx context.Context, in *CreateParentAuthRequest, out *CreateParentAuthResponse) error
	}
	type Auth struct {
		auth
	}
	h := &authHandler{hdlr}
	return s.Handle(s.NewHandler(&Auth{h}, opts...))
}

type authHandler struct {
	AuthHandler
}

func (h *authHandler) CreateStudentAuth(ctx context.Context, in *CreateStudentAuthRequest, out *CreateStudentAuthResponse) error {
	return h.AuthHandler.CreateStudentAuth(ctx, in, out)
}

func (h *authHandler) CreateTeacherAuth(ctx context.Context, in *CreateTeacherAuthRequest, out *CreateTeacherAuthResponse) error {
	return h.AuthHandler.CreateTeacherAuth(ctx, in, out)
}

func (h *authHandler) CreateParentAuth(ctx context.Context, in *CreateParentAuthRequest, out *CreateParentAuthResponse) error {
	return h.AuthHandler.CreateParentAuth(ctx, in, out)
}
