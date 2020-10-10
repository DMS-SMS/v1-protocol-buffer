// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: club-leader.proto

package clubproto

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

// Client API for ClubLeader service

type ClubLeaderService interface {
	AddClubMember(ctx context.Context, in *AddClubMemberRequest, opts ...client.CallOption) (*AddClubMemberResponse, error)
	DeleteClubMember(ctx context.Context, in *DeleteClubMemberRequest, opts ...client.CallOption) (*DeleteClubMemberResponse, error)
	ChangeClubLeader(ctx context.Context, in *ChangeClubLeaderRequest, opts ...client.CallOption) (*ChangeClubLeaderResponse, error)
	ModifyClubInform(ctx context.Context, in *ModifyClubInformRequest, opts ...client.CallOption) (*ModifyClubInformResponse, error)
	DeleteClubWithUUID(ctx context.Context, in *DeleteClubWithUUIDRequest, opts ...client.CallOption) (*DeleteClubWithUUIDResponse, error)
	RegisterRecruitment(ctx context.Context, in *RegisterRecruitmentRequest, opts ...client.CallOption) (*RegisterRecruitmentResponse, error)
	ModifyRecruitment(ctx context.Context, in *ModifyRecruitmentRequest, opts ...client.CallOption) (*ModifyRecruitmentResponse, error)
	DeleteRecruitmentWithUUID(ctx context.Context, in *DeleteRecruitmentWithUUIDRequest, opts ...client.CallOption) (*DeleteRecruitmentWithUUIDResponse, error)
}

type clubLeaderService struct {
	c    client.Client
	name string
}

func NewClubLeaderService(name string, c client.Client) ClubLeaderService {
	return &clubLeaderService{
		c:    c,
		name: name,
	}
}

func (c *clubLeaderService) AddClubMember(ctx context.Context, in *AddClubMemberRequest, opts ...client.CallOption) (*AddClubMemberResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.AddClubMember", in)
	out := new(AddClubMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubLeaderService) DeleteClubMember(ctx context.Context, in *DeleteClubMemberRequest, opts ...client.CallOption) (*DeleteClubMemberResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.DeleteClubMember", in)
	out := new(DeleteClubMemberResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubLeaderService) ChangeClubLeader(ctx context.Context, in *ChangeClubLeaderRequest, opts ...client.CallOption) (*ChangeClubLeaderResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.ChangeClubLeader", in)
	out := new(ChangeClubLeaderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubLeaderService) ModifyClubInform(ctx context.Context, in *ModifyClubInformRequest, opts ...client.CallOption) (*ModifyClubInformResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.ModifyClubInform", in)
	out := new(ModifyClubInformResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubLeaderService) DeleteClubWithUUID(ctx context.Context, in *DeleteClubWithUUIDRequest, opts ...client.CallOption) (*DeleteClubWithUUIDResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.DeleteClubWithUUID", in)
	out := new(DeleteClubWithUUIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubLeaderService) RegisterRecruitment(ctx context.Context, in *RegisterRecruitmentRequest, opts ...client.CallOption) (*RegisterRecruitmentResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.RegisterRecruitment", in)
	out := new(RegisterRecruitmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubLeaderService) ModifyRecruitment(ctx context.Context, in *ModifyRecruitmentRequest, opts ...client.CallOption) (*ModifyRecruitmentResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.ModifyRecruitment", in)
	out := new(ModifyRecruitmentResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubLeaderService) DeleteRecruitmentWithUUID(ctx context.Context, in *DeleteRecruitmentWithUUIDRequest, opts ...client.CallOption) (*DeleteRecruitmentWithUUIDResponse, error) {
	req := c.c.NewRequest(c.name, "ClubLeader.DeleteRecruitmentWithUUID", in)
	out := new(DeleteRecruitmentWithUUIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClubLeader service

type ClubLeaderHandler interface {
	AddClubMember(context.Context, *AddClubMemberRequest, *AddClubMemberResponse) error
	DeleteClubMember(context.Context, *DeleteClubMemberRequest, *DeleteClubMemberResponse) error
	ChangeClubLeader(context.Context, *ChangeClubLeaderRequest, *ChangeClubLeaderResponse) error
	ModifyClubInform(context.Context, *ModifyClubInformRequest, *ModifyClubInformResponse) error
	DeleteClubWithUUID(context.Context, *DeleteClubWithUUIDRequest, *DeleteClubWithUUIDResponse) error
	RegisterRecruitment(context.Context, *RegisterRecruitmentRequest, *RegisterRecruitmentResponse) error
	ModifyRecruitment(context.Context, *ModifyRecruitmentRequest, *ModifyRecruitmentResponse) error
	DeleteRecruitmentWithUUID(context.Context, *DeleteRecruitmentWithUUIDRequest, *DeleteRecruitmentWithUUIDResponse) error
}

func RegisterClubLeaderHandler(s server.Server, hdlr ClubLeaderHandler, opts ...server.HandlerOption) error {
	type clubLeader interface {
		AddClubMember(ctx context.Context, in *AddClubMemberRequest, out *AddClubMemberResponse) error
		DeleteClubMember(ctx context.Context, in *DeleteClubMemberRequest, out *DeleteClubMemberResponse) error
		ChangeClubLeader(ctx context.Context, in *ChangeClubLeaderRequest, out *ChangeClubLeaderResponse) error
		ModifyClubInform(ctx context.Context, in *ModifyClubInformRequest, out *ModifyClubInformResponse) error
		DeleteClubWithUUID(ctx context.Context, in *DeleteClubWithUUIDRequest, out *DeleteClubWithUUIDResponse) error
		RegisterRecruitment(ctx context.Context, in *RegisterRecruitmentRequest, out *RegisterRecruitmentResponse) error
		ModifyRecruitment(ctx context.Context, in *ModifyRecruitmentRequest, out *ModifyRecruitmentResponse) error
		DeleteRecruitmentWithUUID(ctx context.Context, in *DeleteRecruitmentWithUUIDRequest, out *DeleteRecruitmentWithUUIDResponse) error
	}
	type ClubLeader struct {
		clubLeader
	}
	h := &clubLeaderHandler{hdlr}
	return s.Handle(s.NewHandler(&ClubLeader{h}, opts...))
}

type clubLeaderHandler struct {
	ClubLeaderHandler
}

func (h *clubLeaderHandler) AddClubMember(ctx context.Context, in *AddClubMemberRequest, out *AddClubMemberResponse) error {
	return h.ClubLeaderHandler.AddClubMember(ctx, in, out)
}

func (h *clubLeaderHandler) DeleteClubMember(ctx context.Context, in *DeleteClubMemberRequest, out *DeleteClubMemberResponse) error {
	return h.ClubLeaderHandler.DeleteClubMember(ctx, in, out)
}

func (h *clubLeaderHandler) ChangeClubLeader(ctx context.Context, in *ChangeClubLeaderRequest, out *ChangeClubLeaderResponse) error {
	return h.ClubLeaderHandler.ChangeClubLeader(ctx, in, out)
}

func (h *clubLeaderHandler) ModifyClubInform(ctx context.Context, in *ModifyClubInformRequest, out *ModifyClubInformResponse) error {
	return h.ClubLeaderHandler.ModifyClubInform(ctx, in, out)
}

func (h *clubLeaderHandler) DeleteClubWithUUID(ctx context.Context, in *DeleteClubWithUUIDRequest, out *DeleteClubWithUUIDResponse) error {
	return h.ClubLeaderHandler.DeleteClubWithUUID(ctx, in, out)
}

func (h *clubLeaderHandler) RegisterRecruitment(ctx context.Context, in *RegisterRecruitmentRequest, out *RegisterRecruitmentResponse) error {
	return h.ClubLeaderHandler.RegisterRecruitment(ctx, in, out)
}

func (h *clubLeaderHandler) ModifyRecruitment(ctx context.Context, in *ModifyRecruitmentRequest, out *ModifyRecruitmentResponse) error {
	return h.ClubLeaderHandler.ModifyRecruitment(ctx, in, out)
}

func (h *clubLeaderHandler) DeleteRecruitmentWithUUID(ctx context.Context, in *DeleteRecruitmentWithUUIDRequest, out *DeleteRecruitmentWithUUIDResponse) error {
	return h.ClubLeaderHandler.DeleteRecruitmentWithUUID(ctx, in, out)
}
