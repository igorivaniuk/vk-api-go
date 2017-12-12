package auth

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodCheckPhone = "auth.checkPhone"
	MethodSignup     = "auth.signup"
	MethodConfirm    = "auth.confirm"
	MethodRestore    = "auth.restore"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) CheckPhone(r CheckPhoneRequest) ([]byte, error) {
	return s.Client.Do(MethodCheckPhone, r)
}

func (s *Service) Signup(r SignupRequest) ([]byte, error) {
	return s.Client.Do(MethodSignup, r)
}

func (s *Service) Confirm(r ConfirmRequest) ([]byte, error) {
	return s.Client.Do(MethodConfirm, r)
}

func (s *Service) Restore(r RestoreRequest) ([]byte, error) {
	return s.Client.Do(MethodRestore, r)
}
