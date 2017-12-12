package board

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetTopics      = "board.getTopics"
	MethodGetComments    = "board.getComments"
	MethodAddTopic       = "board.addTopic"
	MethodCreateComment  = "board.createComment"
	MethodDeleteTopic    = "board.deleteTopic"
	MethodEditTopic      = "board.editTopic"
	MethodEditComment    = "board.editComment"
	MethodRestoreComment = "board.restoreComment"
	MethodDeleteComment  = "board.deleteComment"
	MethodOpenTopic      = "board.openTopic"
	MethodCloseTopic     = "board.closeTopic"
	MethodFixTopic       = "board.fixTopic"
	MethodUnfixTopic     = "board.unfixTopic"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetTopics(r GetTopicsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetTopics, r)
}

func (s *Service) GetComments(r GetCommentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetComments, r)
}

func (s *Service) AddTopic(r AddTopicRequest) ([]byte, error) {
	return s.Client.Do(MethodAddTopic, r)
}

func (s *Service) CreateComment(r CreateCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateComment, r)
}

func (s *Service) DeleteTopic(r DeleteTopicRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteTopic, r)
}

func (s *Service) EditTopic(r EditTopicRequest) ([]byte, error) {
	return s.Client.Do(MethodEditTopic, r)
}

func (s *Service) EditComment(r EditCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodEditComment, r)
}

func (s *Service) RestoreComment(r RestoreCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodRestoreComment, r)
}

func (s *Service) DeleteComment(r DeleteCommentRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteComment, r)
}

func (s *Service) OpenTopic(r OpenTopicRequest) ([]byte, error) {
	return s.Client.Do(MethodOpenTopic, r)
}

func (s *Service) CloseTopic(r CloseTopicRequest) ([]byte, error) {
	return s.Client.Do(MethodCloseTopic, r)
}

func (s *Service) FixTopic(r FixTopicRequest) ([]byte, error) {
	return s.Client.Do(MethodFixTopic, r)
}

func (s *Service) UnfixTopic(r UnfixTopicRequest) ([]byte, error) {
	return s.Client.Do(MethodUnfixTopic, r)
}
