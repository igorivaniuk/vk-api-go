package search

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetHints = "search.getHints"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetHints(r GetHintsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetHints, r)
}
