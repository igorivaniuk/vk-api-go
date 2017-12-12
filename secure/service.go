package secure

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodCheckToken = "secure.checkToken"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) CheckToken(r CheckTokenRequest) ([]byte, error) {
	return s.Client.Do(MethodCheckToken, r)
}
