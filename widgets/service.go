package widgets

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetComments = "widgets.getComments"
	MethodGetPages    = "widgets.getPages"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetComments(r GetCommentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetComments, r)
}

func (s *Service) GetPages(r GetPagesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetPages, r)
}
