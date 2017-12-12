package notes

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet            = "notes.get"
	MethodGetById        = "notes.getById"
	MethodAdd            = "notes.add"
	MethodEdit           = "notes.edit"
	MethodDelete         = "notes.delete"
	MethodGetComments    = "notes.getComments"
	MethodCreateComment  = "notes.createComment"
	MethodEditComment    = "notes.editComment"
	MethodDeleteComment  = "notes.deleteComment"
	MethodRestoreComment = "notes.restoreComment"
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

func (s *Service) Add(r AddRequest) ([]byte, error) {
	return s.Client.Do(MethodAdd, r)
}

func (s *Service) Edit(r EditRequest) ([]byte, error) {
	return s.Client.Do(MethodEdit, r)
}

func (s *Service) Delete(r DeleteRequest) ([]byte, error) {
	return s.Client.Do(MethodDelete, r)
}

func (s *Service) GetComments(r GetCommentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetComments, r)
}

func (s *Service) CreateComment(r CreateCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateComment, r)
}

func (s *Service) EditComment(r EditCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodEditComment, r)
}

func (s *Service) DeleteComment(r DeleteCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteComment, r)
}

func (s *Service) RestoreComment(r RestoreCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodRestoreComment, r)
}
