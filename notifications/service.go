package notifications

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet          = "notifications.get"
	MethodMarkAsViewed = "notifications.markAsViewed"
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

func (s *Service) MarkAsViewed(r MarkAsViewedRequest) ([]byte, error) {
	return s.Client.Do(MethodMarkAsViewed, r)
}
