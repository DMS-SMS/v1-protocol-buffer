package authproto

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/mock"
)

func MockAuthStudentService(mock *mock.Mock) mockAuthStudentService {
	return mockAuthStudentService{mock: mock}
}

type mockAuthStudentService struct {
	mock *mock.Mock
}

func (m mockAuthStudentService) LoginStudentAuth(ctx context.Context, in *LoginStudentAuthRequest, opts ...client.CallOption) (*LoginStudentAuthResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*LoginStudentAuthResponse), args.Error(1)
}

func (m mockAuthStudentService) ChangeStudentPW(ctx context.Context, in *ChangeStudentPWRequest, opts ...client.CallOption) (*ChangeStudentPWResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*ChangeStudentPWResponse), args.Error(1)
}

func (m mockAuthStudentService) GetStudentInformWithUUID(ctx context.Context, in *GetStudentInformWithUUIDRequest, opts ...client.CallOption) (*GetStudentInformWithUUIDResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*GetStudentInformWithUUIDResponse), args.Error(1)
}

func (m mockAuthStudentService) GetStudentInformsWithUUIDs(ctx context.Context, in *GetStudentInformsWithUUIDsRequest, opts ...client.CallOption) (*GetStudentInformsWithUUIDsResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*GetStudentInformsWithUUIDsResponse), args.Error(1)
}

func (m mockAuthStudentService) GetStudentUUIDsWithInform(ctx context.Context, in *GetStudentUUIDsWithInformRequest, opts ...client.CallOption) (*GetStudentUUIDsWithInformResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*GetStudentUUIDsWithInformResponse), args.Error(1)
}
