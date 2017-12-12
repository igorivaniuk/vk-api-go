package places

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodAdd         = "places.add"
	MethodGetById     = "places.getById"
	MethodSearch      = "places.search"
	MethodCheckin     = "places.checkin"
	MethodGetCheckins = "places.getCheckins"
	MethodGetTypes    = "places.getTypes"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) Add(r AddRequest) ([]byte, error) {
	return s.Client.Do(MethodAdd, r)
}

func (s *Service) GetById(r GetByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetById, r)
}

func (s *Service) Search(r SearchRequest) ([]byte, error) {
	return s.Client.Do(MethodSearch, r)
}

func (s *Service) Checkin(r CheckinRequest) ([]byte, error) {
	return s.Client.Do(MethodCheckin, r)
}

func (s *Service) GetCheckins(r GetCheckinsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCheckins, r)
}

func (s *Service) GetTypes(r GetTypesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetTypes, r)
}
