package users

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet              = "users.get"
	MethodGetFollowers     = "users.getFollowers"
	MethodGetNearby        = "users.getNearby"
	MethodGetSubscriptions = "users.getSubscriptions"
	MethodIsAppUser        = "users.isAppUser"
	MethodReport           = "users.report"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) Get(r GetRequest) ([]byte, error) {
	return s.Client.Do(MethodGet, r)
}

func (s *Service) GetFollowers(r GetFollowersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetFollowers, r)
}

func (s *Service) GetNearby(r GetNearbyRequest) ([]byte, error) {
	return s.Client.Do(MethodGetNearby, r)
}

func (s *Service) GetSubscriptions(r GetSubscriptionsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetSubscriptions, r)
}

func (s *Service) IsAppUser(r IsAppUserRequest) ([]byte, error) {
	return s.Client.Do(MethodIsAppUser, r)
}

func (s *Service) Report(r ReportRequest) ([]byte, error) {
	return s.Client.Do(MethodReport, r)
}
