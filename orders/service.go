package orders

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet         = "orders.get"
	MethodGetById     = "orders.getById"
	MethodChangeState = "orders.changeState"
	MethodGetAmount   = "orders.getAmount"
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

func (s *Service) GetById(r GetByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetById, r)
}

func (s *Service) ChangeState(r ChangeStateRequest) ([]byte, error) {
	return s.Client.Do(MethodChangeState, r)
}

func (s *Service) GetAmount(r GetAmountRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAmount, r)
}
