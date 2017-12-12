package wall

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetById        = "wall.getById"
	MethodPost           = "wall.post"
	MethodRepost         = "wall.repost"
	MethodGetReposts     = "wall.getReposts"
	MethodEdit           = "wall.edit"
	MethodDelete         = "wall.delete"
	MethodRestore        = "wall.restore"
	MethodPin            = "wall.pin"
	MethodUnpin          = "wall.unpin"
	MethodGetComments    = "wall.getComments"
	MethodCreateComment  = "wall.createComment"
	MethodEditComment    = "wall.editComment"
	MethodDeleteComment  = "wall.deleteComment"
	MethodRestoreComment = "wall.restoreComment"
	MethodReportPost     = "wall.reportPost"
	MethodReportComment  = "wall.reportComment"
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

func (s *Service) Post(r PostRequest) ([]byte, error) {
	return s.Client.Do(MethodPost, r)
}

func (s *Service) Repost(r RepostRequest) ([]byte, error) {
	return s.Client.Do(MethodRepost, r)
}

func (s *Service) GetReposts(r GetRepostsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetReposts, r)
}

func (s *Service) Edit(r EditRequest) ([]byte, error) {
	return s.Client.Do(MethodEdit, r)
}

func (s *Service) Delete(r DeleteRequest) ([]byte, error) {
	return s.Client.Do(MethodDelete, r)
}

func (s *Service) Restore(r RestoreRequest) ([]byte, error) {
	return s.Client.Do(MethodRestore, r)
}

func (s *Service) Pin(r PinRequest) ([]byte, error) {
	return s.Client.Do(MethodPin, r)
}

func (s *Service) Unpin(r UnpinRequest) ([]byte, error) {
	return s.Client.Do(MethodUnpin, r)
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

func (s *Service) ReportPost(r ReportPostRequest) ([]byte, error) {
	return s.Client.Do(MethodReportPost, r)
}

func (s *Service) ReportComment(r ReportCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodReportComment, r)
}
