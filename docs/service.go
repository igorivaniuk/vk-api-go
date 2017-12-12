package docs

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet                 = "docs.get"
	MethodGetById             = "docs.getById"
	MethodGetUploadServer     = "docs.getUploadServer"
	MethodGetWallUploadServer = "docs.getWallUploadServer"
	MethodSave                = "docs.save"
	MethodDelete              = "docs.delete"
	MethodAdd                 = "docs.add"
	MethodGetTypes            = "docs.getTypes"
	MethodSearch              = "docs.search"
	MethodEdit                = "docs.edit"
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

func (s *Service) GetUploadServer(r GetUploadServerRequest) ([]byte, error) {
	return s.Client.Do(MethodGetUploadServer, r)
}

func (s *Service) GetWallUploadServer(r GetWallUploadServerRequest) ([]byte, error) {
	return s.Client.Do(MethodGetWallUploadServer, r)
}

func (s *Service) Save(r SaveRequest) ([]byte, error) {
	return s.Client.Do(MethodSave, r)
}

func (s *Service) Delete(r DeleteRequest) ([]byte, error) {
	return s.Client.Do(MethodDelete, r)
}

func (s *Service) Add(r AddRequest) ([]byte, error) {
	return s.Client.Do(MethodAdd, r)
}

func (s *Service) GetTypes(r GetTypesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetTypes, r)
}

func (s *Service) Search(r SearchRequest) ([]byte, error) {
	return s.Client.Do(MethodSearch, r)
}

func (s *Service) Edit(r EditRequest) ([]byte, error) {
	return s.Client.Do(MethodEdit, r)
}
