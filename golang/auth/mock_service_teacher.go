package authproto

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/mock"
)

func MockAuthTeacherService(mock *mock.Mock) mockAuthTeacherService {
	return mockAuthTeacherService{mock: mock}
}

type mockAuthTeacherService struct {
	mock *mock.Mock
}

func (m mockAuthTeacherService) LoginTeacherAuth(ctx context.Context, in *LoginTeacherAuthRequest, opts ...client.CallOption) (*LoginTeacherAuthResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*LoginTeacherAuthResponse), args.Error(1)
}

func (m mockAuthTeacherService) ChangeTeacherPW(ctx context.Context, in *ChangeTeacherPWRequest, opts ...client.CallOption) (*ChangeTeacherPWResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*ChangeTeacherPWResponse), args.Error(1)
}

func (m mockAuthTeacherService) GetTeacherInformWithUUID(ctx context.Context, in *GetTeacherInformWithUUIDRequest, opts ...client.CallOption) (*GetTeacherInformWithUUIDResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*GetTeacherInformWithUUIDResponse), args.Error(1)
}

func (m mockAuthTeacherService) GetTeacherUUIDsWithInform(ctx context.Context, in *GetTeacherUUIDsWithInformRequest, opts ...client.CallOption) (*GetTeacherUUIDsWithInformResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*GetTeacherUUIDsWithInformResponse), args.Error(1)
}
