package stats

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet          = "stats.get"
	MethodTrackVisitor = "stats.trackVisitor"
	MethodGetPostReach = "stats.getPostReach"
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

func (s *Service) TrackVisitor(r TrackVisitorRequest) ([]byte, error) {
	return s.Client.Do(MethodTrackVisitor, r)
}

func (s *Service) GetPostReach(r GetPostReachRequest) ([]byte, error) {
	return s.Client.Do(MethodGetPostReach, r)
}
