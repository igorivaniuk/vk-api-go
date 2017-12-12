package polls

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetById    = "polls.getById"
	MethodAddVote    = "polls.addVote"
	MethodDeleteVote = "polls.deleteVote"
	MethodGetVoters  = "polls.getVoters"
	MethodCreate     = "polls.create"
	MethodEdit       = "polls.edit"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetById(r GetByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetById, r)
}

func (s *Service) AddVote(r AddVoteRequest) ([]byte, error) {
	return s.Client.Do(MethodAddVote, r)
}

func (s *Service) DeleteVote(r DeleteVoteRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteVote, r)
}

func (s *Service) GetVoters(r GetVotersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetVoters, r)
}

func (s *Service) Create(r CreateRequest) ([]byte, error) {
	return s.Client.Do(MethodCreate, r)
}

func (s *Service) Edit(r EditRequest) ([]byte, error) {
	return s.Client.Do(MethodEdit, r)
}
