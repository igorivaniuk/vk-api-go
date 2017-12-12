package leads

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodComplete  = "leads.complete"
	MethodStart     = "leads.start"
	MethodGetStats  = "leads.getStats"
	MethodGetUsers  = "leads.getUsers"
	MethodCheckUser = "leads.checkUser"
	MethodMetricHit = "leads.metricHit"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) Complete(r CompleteRequest) ([]byte, error) {
	return s.Client.Do(MethodComplete, r)
}

func (s *Service) Start(r StartRequest) ([]byte, error) {
	return s.Client.Do(MethodStart, r)
}

func (s *Service) GetStats(r GetStatsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetStats, r)
}

func (s *Service) GetUsers(r GetUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetUsers, r)
}

func (s *Service) CheckUser(r CheckUserRequest) ([]byte, error) {
	return s.Client.Do(MethodCheckUser, r)
}

func (s *Service) MetricHit(r MetricHitRequest) ([]byte, error) {
	return s.Client.Do(MethodMetricHit, r)
}
