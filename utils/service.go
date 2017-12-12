package utils

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodCheckLink               = "utils.checkLink"
	MethodDeleteFromLastShortened = "utils.deleteFromLastShortened"
	MethodGetLastShortenedLinks   = "utils.getLastShortenedLinks"
	MethodGetLinkStats            = "utils.getLinkStats"
	MethodGetShortLink            = "utils.getShortLink"
	MethodResolveScreenName       = "utils.resolveScreenName"
	MethodGetServerTime           = "utils.getServerTime"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) CheckLink(r CheckLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodCheckLink, r)
}

func (s *Service) DeleteFromLastShortened(r DeleteFromLastShortenedRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteFromLastShortened, r)
}

func (s *Service) GetLastShortenedLinks(r GetLastShortenedLinksRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLastShortenedLinks, r)
}

func (s *Service) GetLinkStats(r GetLinkStatsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLinkStats, r)
}

func (s *Service) GetShortLink(r GetShortLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodGetShortLink, r)
}

func (s *Service) ResolveScreenName(r ResolveScreenNameRequest) ([]byte, error) {
	return s.Client.Do(MethodResolveScreenName, r)
}

func (s *Service) GetServerTime(r GetServerTimeRequest) ([]byte, error) {
	return s.Client.Do(MethodGetServerTime, r)
}
