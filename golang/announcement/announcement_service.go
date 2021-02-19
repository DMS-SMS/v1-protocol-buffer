package announcementproto

import (
	"github.com/micro/go-micro/v2/client"
)

type AnnouncementServiceClient struct {
	AnnouncementService
	AnnouncementEventService
}

func NewOutingService(name string, c client.Client) AnnouncementServiceClient {
	return AnnouncementServiceClient{
		AnnouncementService:      NewAnnouncementService(name, c),
		AnnouncementEventService: NewAnnouncementEventService(name, c),
	}
}
