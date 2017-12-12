package gifts

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet = "gifts.get"
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
