package outingproto

import (
	"github.com/micro/go-micro/v2/client"
)

type OutingServiceClient struct {
	OutingStudentService
	OutingTeacherService
	OutingParentsService
}

func NewOutingService(name string, c client.Client) OutingServiceClient {
	return OutingServiceClient{
		OutingStudentService: NewOutingStudentService(name, c),
		OutingTeacherService: NewOutingTeacherService(name, c),
		OutingParentsService: NewOutingParentsService(name, c),
	}
}
