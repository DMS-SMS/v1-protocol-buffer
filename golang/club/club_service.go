package clubproto

import (
	"github.com/micro/go-micro/v2/client"
)

type ClubServiceClient struct {
	ClubAdminService
	ClubStudentService
	ClubLeaderService
}

func NewClubService(name string, c client.Client) ClubServiceClient {
	return ClubServiceClient{
		ClubAdminService:   NewClubAdminService(name, c),
		ClubStudentService: NewClubStudentService(name, c),
		ClubLeaderService:  NewClubLeaderService(name, c),
	}
}
