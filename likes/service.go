package likes

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetList = "likes.getList"
	MethodAdd     = "likes.add"
	MethodDelete  = "likes.delete"
	MethodIsLiked = "likes.isLiked"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetList(r GetListRequest) ([]byte, error) {
	return s.Client.Do(MethodGetList, r)
}

func (s *Service) Add(r AddRequest) ([]byte, error) {
	return s.Client.Do(MethodAdd, r)
}

func (s *Service) Delete(r DeleteRequest) ([]byte, error) {
	return s.Client.Do(MethodDelete, r)
}

func (s *Service) IsLiked(r IsLikedRequest) ([]byte, error) {
	return s.Client.Do(MethodIsLiked, r)
}
