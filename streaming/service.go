package streaming

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetServerUrl = "streaming.getServerUrl"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetServerUrl(r GetServerUrlRequest) ([]byte, error) {
	return s.Client.Do(MethodGetServerUrl, r)
}
