package apps

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetScore = "apps.getScore"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetScore(r GetScoreRequest) ([]byte, error) {
	return s.Client.Do(MethodGetScore, r)
}
