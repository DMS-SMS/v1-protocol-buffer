package authproto

import (
	"context"
	"github.com/micro/go-micro/client"
	"github.com/stretchr/testify/mock"
)

func MockAuthAdminService(mock *mock.Mock) mockAuthAdminService {
	return mockAuthAdminService{mock: mock}
}

type mockAuthAdminService struct {
	mock *mock.Mock
}

func (m mockAuthAdminService) CreateNewStudent(ctx context.Context, in *CreateNewStudentRequest, opts ...client.CallOption) (*CreateNewStudentResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*CreateNewStudentResponse), args.Error(1)
}

func (m mockAuthAdminService) CreateNewTeacher(ctx context.Context, in *CreateNewTeacherRequest, opts ...client.CallOption) (*CreateNewTeacherResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*CreateNewTeacherResponse), args.Error(1)
}

func (m mockAuthAdminService) CreateNewParent(ctx context.Context, in *CreateNewParentRequest, opts ...client.CallOption) (*CreateNewParentResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*CreateNewParentResponse), args.Error(1)
}

func (m mockAuthAdminService) LoginAdminAuth(ctx context.Context, in *LoginAdminAuthRequest, opts ...client.CallOption) (*LoginAdminAuthResponse, error) {
	args := m.mock.Called(in)
	return args.Get(0).(*LoginAdminAuthResponse), args.Error(1)
}
