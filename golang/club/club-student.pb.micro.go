// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: club-student.proto

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

// Client API for ClubStudent service

type ClubStudentService interface {
	GetClubsSortByUpdateTime(ctx context.Context, in *GetClubsSortByUpdateTimeRequest, opts ...client.CallOption) (*GetClubsSortByUpdateTimeResponse, error)
	GetRecruitmentsSortByCreateTime(ctx context.Context, in *GetRecruitmentsSortByCreateTimeRequest, opts ...client.CallOption) (*GetRecruitmentsSortByCreateTimeResponse, error)
	GetClubInformWithUUID(ctx context.Context, in *GetClubInformWithUUIDRequest, opts ...client.CallOption) (*GetClubInformWithUUIDResponse, error)
	GetRecruitmentInformWithUUID(ctx context.Context, in *GetRecruitmentInformWithUUIDRequest, opts ...client.CallOption) (*GetRecruitmentInformWithUUIDResponse, error)
	GetAllClubFields(ctx context.Context, in *GetAllClubFieldsRequest, opts ...client.CallOption) (*GetAllClubFieldsResponse, error)
	GetTotalCountOfClubs(ctx context.Context, in *GetTotalCountOfClubsRequest, opts ...client.CallOption) (*GetTotalCountOfClubsResponse, error)
	GetTotalCountOfRecruitments(ctx context.Context, in *GetTotalCountOfRecruitmentsRequest, opts ...client.CallOption) (*GetTotalCountOfRecruitmentsResponse, error)
	CheckIfClubLeader(ctx context.Context, in *CheckIfClubLeaderRequest, opts ...client.CallOption) (*CheckIfClubLeaderResponse, error)
}

type clubStudentService struct {
	c    client.Client
	name string
}

func NewClubStudentService(name string, c client.Client) ClubStudentService {
	return &clubStudentService{
		c:    c,
		name: name,
	}
}

func (c *clubStudentService) GetClubsSortByUpdateTime(ctx context.Context, in *GetClubsSortByUpdateTimeRequest, opts ...client.CallOption) (*GetClubsSortByUpdateTimeResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.GetClubsSortByUpdateTime", in)
	out := new(GetClubsSortByUpdateTimeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubStudentService) GetRecruitmentsSortByCreateTime(ctx context.Context, in *GetRecruitmentsSortByCreateTimeRequest, opts ...client.CallOption) (*GetRecruitmentsSortByCreateTimeResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.GetRecruitmentsSortByCreateTime", in)
	out := new(GetRecruitmentsSortByCreateTimeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubStudentService) GetClubInformWithUUID(ctx context.Context, in *GetClubInformWithUUIDRequest, opts ...client.CallOption) (*GetClubInformWithUUIDResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.GetClubInformWithUUID", in)
	out := new(GetClubInformWithUUIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubStudentService) GetRecruitmentInformWithUUID(ctx context.Context, in *GetRecruitmentInformWithUUIDRequest, opts ...client.CallOption) (*GetRecruitmentInformWithUUIDResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.GetRecruitmentInformWithUUID", in)
	out := new(GetRecruitmentInformWithUUIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubStudentService) GetAllClubFields(ctx context.Context, in *GetAllClubFieldsRequest, opts ...client.CallOption) (*GetAllClubFieldsResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.GetAllClubFields", in)
	out := new(GetAllClubFieldsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubStudentService) GetTotalCountOfClubs(ctx context.Context, in *GetTotalCountOfClubsRequest, opts ...client.CallOption) (*GetTotalCountOfClubsResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.GetTotalCountOfClubs", in)
	out := new(GetTotalCountOfClubsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubStudentService) GetTotalCountOfRecruitments(ctx context.Context, in *GetTotalCountOfRecruitmentsRequest, opts ...client.CallOption) (*GetTotalCountOfRecruitmentsResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.GetTotalCountOfRecruitments", in)
	out := new(GetTotalCountOfRecruitmentsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubStudentService) CheckIfClubLeader(ctx context.Context, in *CheckIfClubLeaderRequest, opts ...client.CallOption) (*CheckIfClubLeaderResponse, error) {
	req := c.c.NewRequest(c.name, "ClubStudent.CheckIfClubLeader", in)
	out := new(CheckIfClubLeaderResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ClubStudent service

type ClubStudentHandler interface {
	GetClubsSortByUpdateTime(context.Context, *GetClubsSortByUpdateTimeRequest, *GetClubsSortByUpdateTimeResponse) error
	GetRecruitmentsSortByCreateTime(context.Context, *GetRecruitmentsSortByCreateTimeRequest, *GetRecruitmentsSortByCreateTimeResponse) error
	GetClubInformWithUUID(context.Context, *GetClubInformWithUUIDRequest, *GetClubInformWithUUIDResponse) error
	GetRecruitmentInformWithUUID(context.Context, *GetRecruitmentInformWithUUIDRequest, *GetRecruitmentInformWithUUIDResponse) error
	GetAllClubFields(context.Context, *GetAllClubFieldsRequest, *GetAllClubFieldsResponse) error
	GetTotalCountOfClubs(context.Context, *GetTotalCountOfClubsRequest, *GetTotalCountOfClubsResponse) error
	GetTotalCountOfRecruitments(context.Context, *GetTotalCountOfRecruitmentsRequest, *GetTotalCountOfRecruitmentsResponse) error
	CheckIfClubLeader(context.Context, *CheckIfClubLeaderRequest, *CheckIfClubLeaderResponse) error
}

func RegisterClubStudentHandler(s server.Server, hdlr ClubStudentHandler, opts ...server.HandlerOption) error {
	type clubStudent interface {
		GetClubsSortByUpdateTime(ctx context.Context, in *GetClubsSortByUpdateTimeRequest, out *GetClubsSortByUpdateTimeResponse) error
		GetRecruitmentsSortByCreateTime(ctx context.Context, in *GetRecruitmentsSortByCreateTimeRequest, out *GetRecruitmentsSortByCreateTimeResponse) error
		GetClubInformWithUUID(ctx context.Context, in *GetClubInformWithUUIDRequest, out *GetClubInformWithUUIDResponse) error
		GetRecruitmentInformWithUUID(ctx context.Context, in *GetRecruitmentInformWithUUIDRequest, out *GetRecruitmentInformWithUUIDResponse) error
		GetAllClubFields(ctx context.Context, in *GetAllClubFieldsRequest, out *GetAllClubFieldsResponse) error
		GetTotalCountOfClubs(ctx context.Context, in *GetTotalCountOfClubsRequest, out *GetTotalCountOfClubsResponse) error
		GetTotalCountOfRecruitments(ctx context.Context, in *GetTotalCountOfRecruitmentsRequest, out *GetTotalCountOfRecruitmentsResponse) error
		CheckIfClubLeader(ctx context.Context, in *CheckIfClubLeaderRequest, out *CheckIfClubLeaderResponse) error
	}
	type ClubStudent struct {
		clubStudent
	}
	h := &clubStudentHandler{hdlr}
	return s.Handle(s.NewHandler(&ClubStudent{h}, opts...))
}

type clubStudentHandler struct {
	ClubStudentHandler
}

func (h *clubStudentHandler) GetClubsSortByUpdateTime(ctx context.Context, in *GetClubsSortByUpdateTimeRequest, out *GetClubsSortByUpdateTimeResponse) error {
	return h.ClubStudentHandler.GetClubsSortByUpdateTime(ctx, in, out)
}

func (h *clubStudentHandler) GetRecruitmentsSortByCreateTime(ctx context.Context, in *GetRecruitmentsSortByCreateTimeRequest, out *GetRecruitmentsSortByCreateTimeResponse) error {
	return h.ClubStudentHandler.GetRecruitmentsSortByCreateTime(ctx, in, out)
}

func (h *clubStudentHandler) GetClubInformWithUUID(ctx context.Context, in *GetClubInformWithUUIDRequest, out *GetClubInformWithUUIDResponse) error {
	return h.ClubStudentHandler.GetClubInformWithUUID(ctx, in, out)
}

func (h *clubStudentHandler) GetRecruitmentInformWithUUID(ctx context.Context, in *GetRecruitmentInformWithUUIDRequest, out *GetRecruitmentInformWithUUIDResponse) error {
	return h.ClubStudentHandler.GetRecruitmentInformWithUUID(ctx, in, out)
}

func (h *clubStudentHandler) GetAllClubFields(ctx context.Context, in *GetAllClubFieldsRequest, out *GetAllClubFieldsResponse) error {
	return h.ClubStudentHandler.GetAllClubFields(ctx, in, out)
}

func (h *clubStudentHandler) GetTotalCountOfClubs(ctx context.Context, in *GetTotalCountOfClubsRequest, out *GetTotalCountOfClubsResponse) error {
	return h.ClubStudentHandler.GetTotalCountOfClubs(ctx, in, out)
}

func (h *clubStudentHandler) GetTotalCountOfRecruitments(ctx context.Context, in *GetTotalCountOfRecruitmentsRequest, out *GetTotalCountOfRecruitmentsResponse) error {
	return h.ClubStudentHandler.GetTotalCountOfRecruitments(ctx, in, out)
}

func (h *clubStudentHandler) CheckIfClubLeader(ctx context.Context, in *CheckIfClubLeaderRequest, out *CheckIfClubLeaderResponse) error {
	return h.ClubStudentHandler.CheckIfClubLeader(ctx, in, out)
}
