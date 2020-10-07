package authproto

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/mock"
)

func MockAuthParentService(mock *mock.Mock) mockAuthParentService {
	return mockAuthParentService{mock: mock}
}

type mockAuthParentService struct {
	mock *mock.Mock
}

func (m mockAuthParentService) LoginParentAuth(ctx context.Context, in *LoginParentAuthRequest, opts ...client.CallOption) (*LoginParentAuthResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*LoginParentAuthResponse), args.Error(1)
}

func (m mockAuthParentService) ChangeParentPW(ctx context.Context, in *ChangeParentPWRequest, opts ...client.CallOption) (*ChangeParentPWResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*ChangeParentPWResponse), args.Error(1)
}

func (m mockAuthParentService) GetParentInformWithUUID(ctx context.Context, in *GetParentInformWithUUIDRequest, opts ...client.CallOption) (*GetParentInformWithUUIDResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*GetParentInformWithUUIDResponse), args.Error(1)
}

func (m mockAuthParentService) GetParentUUIDsWithInform(ctx context.Context, in *GetParentUUIDsWithInformRequest, opts ...client.CallOption) (*GetParentUUIDsWithInformResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*GetParentUUIDsWithInformResponse), args.Error(1)
}
