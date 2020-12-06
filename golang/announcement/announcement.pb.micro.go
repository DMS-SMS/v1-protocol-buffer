// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: announcement.proto

package announcementproto

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

// Client API for AnnouncementService service

type AnnouncementService interface {
	GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...client.CallOption) (*GetAnnouncementsResponse, error)
	GetAnnouncementDetail(ctx context.Context, in *GetAnnouncementDetailRequest, opts ...client.CallOption) (*GetAnnouncementDetailResponse, error)
	CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...client.CallOption) (*DefaultAnnouncementResponse, error)
	UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...client.CallOption) (*DefaultAnnouncementResponse, error)
	DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...client.CallOption) (*DefaultAnnouncementResponse, error)
}

type announcementService struct {
	c    client.Client
	name string
}

func NewAnnouncementService(name string, c client.Client) AnnouncementService {
	return &announcementService{
		c:    c,
		name: name,
	}
}

func (c *announcementService) GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, opts ...client.CallOption) (*GetAnnouncementsResponse, error) {
	req := c.c.NewRequest(c.name, "AnnouncementService.GetAnnouncements", in)
	out := new(GetAnnouncementsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementService) GetAnnouncementDetail(ctx context.Context, in *GetAnnouncementDetailRequest, opts ...client.CallOption) (*GetAnnouncementDetailResponse, error) {
	req := c.c.NewRequest(c.name, "AnnouncementService.GetAnnouncementDetail", in)
	out := new(GetAnnouncementDetailResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementService) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, opts ...client.CallOption) (*DefaultAnnouncementResponse, error) {
	req := c.c.NewRequest(c.name, "AnnouncementService.CreateAnnouncement", in)
	out := new(DefaultAnnouncementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementService) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, opts ...client.CallOption) (*DefaultAnnouncementResponse, error) {
	req := c.c.NewRequest(c.name, "AnnouncementService.UpdateAnnouncement", in)
	out := new(DefaultAnnouncementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *announcementService) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, opts ...client.CallOption) (*DefaultAnnouncementResponse, error) {
	req := c.c.NewRequest(c.name, "AnnouncementService.DeleteAnnouncement", in)
	out := new(DefaultAnnouncementResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AnnouncementService service

type AnnouncementServiceHandler interface {
	GetAnnouncements(context.Context, *GetAnnouncementsRequest, *GetAnnouncementsResponse) error
	GetAnnouncementDetail(context.Context, *GetAnnouncementDetailRequest, *GetAnnouncementDetailResponse) error
	CreateAnnouncement(context.Context, *CreateAnnouncementRequest, *DefaultAnnouncementResponse) error
	UpdateAnnouncement(context.Context, *UpdateAnnouncementRequest, *DefaultAnnouncementResponse) error
	DeleteAnnouncement(context.Context, *DeleteAnnouncementRequest, *DefaultAnnouncementResponse) error
}

func RegisterAnnouncementServiceHandler(s server.Server, hdlr AnnouncementServiceHandler, opts ...server.HandlerOption) error {
	type announcementService interface {
		GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, out *GetAnnouncementsResponse) error
		GetAnnouncementDetail(ctx context.Context, in *GetAnnouncementDetailRequest, out *GetAnnouncementDetailResponse) error
		CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, out *DefaultAnnouncementResponse) error
		UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, out *DefaultAnnouncementResponse) error
		DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, out *DefaultAnnouncementResponse) error
	}
	type AnnouncementService struct {
		announcementService
	}
	h := &announcementServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AnnouncementService{h}, opts...))
}

type announcementServiceHandler struct {
	AnnouncementServiceHandler
}

func (h *announcementServiceHandler) GetAnnouncements(ctx context.Context, in *GetAnnouncementsRequest, out *GetAnnouncementsResponse) error {
	return h.AnnouncementServiceHandler.GetAnnouncements(ctx, in, out)
}

func (h *announcementServiceHandler) GetAnnouncementDetail(ctx context.Context, in *GetAnnouncementDetailRequest, out *GetAnnouncementDetailResponse) error {
	return h.AnnouncementServiceHandler.GetAnnouncementDetail(ctx, in, out)
}

func (h *announcementServiceHandler) CreateAnnouncement(ctx context.Context, in *CreateAnnouncementRequest, out *DefaultAnnouncementResponse) error {
	return h.AnnouncementServiceHandler.CreateAnnouncement(ctx, in, out)
}

func (h *announcementServiceHandler) UpdateAnnouncement(ctx context.Context, in *UpdateAnnouncementRequest, out *DefaultAnnouncementResponse) error {
	return h.AnnouncementServiceHandler.UpdateAnnouncement(ctx, in, out)
}

func (h *announcementServiceHandler) DeleteAnnouncement(ctx context.Context, in *DeleteAnnouncementRequest, out *DefaultAnnouncementResponse) error {
	return h.AnnouncementServiceHandler.DeleteAnnouncement(ctx, in, out)
}