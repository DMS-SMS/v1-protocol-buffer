package scheduleproto

import (
	"github.com/micro/go-micro/v2/client"
)

type ScheduleServiceClient struct {
	ScheduleService
	ScheduleEventService
}

func NewScheduleSrv(name string, c client.Client) ScheduleServiceClient {
	return ScheduleServiceClient{
		ScheduleService:      NewScheduleService(name, c),
		ScheduleEventService: NewScheduleEventService(name, c),
	}
}
