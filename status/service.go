package status

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet = "status.get"
	MethodSet = "status.set"
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

func (s *Service) Set(r SetRequest) ([]byte, error) {
	return s.Client.Do(MethodSet, r)
}
