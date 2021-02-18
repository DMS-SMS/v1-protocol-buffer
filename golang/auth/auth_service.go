package authproto

import (
	"github.com/micro/go-micro/v2/client"
)

type AuthServiceClient struct {
	AuthAdminService
	AuthStudentService
	AuthTeacherService
	AuthParentService
	AuthEventService
}

func NewAuthService(name string, c client.Client) AuthServiceClient {
	return AuthServiceClient{
		AuthAdminService:   NewAuthAdminService(name, c),
		AuthStudentService: NewAuthStudentService(name, c),
		AuthTeacherService: NewAuthTeacherService(name, c),
		AuthParentService:  NewAuthParentService(name, c),
		AuthEventService:   NewAuthEventService(name, c),
	}
}
