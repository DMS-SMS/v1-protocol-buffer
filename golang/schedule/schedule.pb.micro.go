// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: schedule.proto

package scheduleproto

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

// Client API for ScheduleService service

type ScheduleService interface {
	GetTimeTable(ctx context.Context, in *GetTimeTableRequest, opts ...client.CallOption) (*GetTimeTableResponse, error)
	GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...client.CallOption) (*GetScheduleResponse, error)
	CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...client.CallOption) (*DefaultScheduleResponse, error)
	UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...client.CallOption) (*DefaultScheduleResponse, error)
	DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...client.CallOption) (*DefaultScheduleResponse, error)
}

type scheduleService struct {
	c    client.Client
	name string
}

func NewScheduleService(name string, c client.Client) ScheduleService {
	return &scheduleService{
		c:    c,
		name: name,
	}
}

func (c *scheduleService) GetTimeTable(ctx context.Context, in *GetTimeTableRequest, opts ...client.CallOption) (*GetTimeTableResponse, error) {
	req := c.c.NewRequest(c.name, "ScheduleService.GetTimeTable", in)
	out := new(GetTimeTableResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleService) GetSchedule(ctx context.Context, in *GetScheduleRequest, opts ...client.CallOption) (*GetScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "ScheduleService.GetSchedule", in)
	out := new(GetScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleService) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, opts ...client.CallOption) (*DefaultScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "ScheduleService.CreateSchedule", in)
	out := new(DefaultScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleService) UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, opts ...client.CallOption) (*DefaultScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "ScheduleService.UpdateSchedule", in)
	out := new(DefaultScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *scheduleService) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, opts ...client.CallOption) (*DefaultScheduleResponse, error) {
	req := c.c.NewRequest(c.name, "ScheduleService.DeleteSchedule", in)
	out := new(DefaultScheduleResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ScheduleService service

type ScheduleServiceHandler interface {
	GetTimeTable(context.Context, *GetTimeTableRequest, *GetTimeTableResponse) error
	GetSchedule(context.Context, *GetScheduleRequest, *GetScheduleResponse) error
	CreateSchedule(context.Context, *CreateScheduleRequest, *DefaultScheduleResponse) error
	UpdateSchedule(context.Context, *UpdateScheduleRequest, *DefaultScheduleResponse) error
	DeleteSchedule(context.Context, *DeleteScheduleRequest, *DefaultScheduleResponse) error
}

func RegisterScheduleServiceHandler(s server.Server, hdlr ScheduleServiceHandler, opts ...server.HandlerOption) error {
	type scheduleService interface {
		GetTimeTable(ctx context.Context, in *GetTimeTableRequest, out *GetTimeTableResponse) error
		GetSchedule(ctx context.Context, in *GetScheduleRequest, out *GetScheduleResponse) error
		CreateSchedule(ctx context.Context, in *CreateScheduleRequest, out *DefaultScheduleResponse) error
		UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, out *DefaultScheduleResponse) error
		DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, out *DefaultScheduleResponse) error
	}
	type ScheduleService struct {
		scheduleService
	}
	h := &scheduleServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ScheduleService{h}, opts...))
}

type scheduleServiceHandler struct {
	ScheduleServiceHandler
}

func (h *scheduleServiceHandler) GetTimeTable(ctx context.Context, in *GetTimeTableRequest, out *GetTimeTableResponse) error {
	return h.ScheduleServiceHandler.GetTimeTable(ctx, in, out)
}

func (h *scheduleServiceHandler) GetSchedule(ctx context.Context, in *GetScheduleRequest, out *GetScheduleResponse) error {
	return h.ScheduleServiceHandler.GetSchedule(ctx, in, out)
}

func (h *scheduleServiceHandler) CreateSchedule(ctx context.Context, in *CreateScheduleRequest, out *DefaultScheduleResponse) error {
	return h.ScheduleServiceHandler.CreateSchedule(ctx, in, out)
}

func (h *scheduleServiceHandler) UpdateSchedule(ctx context.Context, in *UpdateScheduleRequest, out *DefaultScheduleResponse) error {
	return h.ScheduleServiceHandler.UpdateSchedule(ctx, in, out)
}

func (h *scheduleServiceHandler) DeleteSchedule(ctx context.Context, in *DeleteScheduleRequest, out *DefaultScheduleResponse) error {
	return h.ScheduleServiceHandler.DeleteSchedule(ctx, in, out)
}
