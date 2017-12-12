package pages

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet        = "pages.get"
	MethodSave       = "pages.save"
	MethodSaveAccess = "pages.saveAccess"
	MethodGetHistory = "pages.getHistory"
	MethodGetTitles  = "pages.getTitles"
	MethodGetVersion = "pages.getVersion"
	MethodParseWiki  = "pages.parseWiki"
	MethodClearCache = "pages.clearCache"
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

func (s *Service) Save(r SaveRequest) ([]byte, error) {
	return s.Client.Do(MethodSave, r)
}

func (s *Service) SaveAccess(r SaveAccessRequest) ([]byte, error) {
	return s.Client.Do(MethodSaveAccess, r)
}

func (s *Service) GetHistory(r GetHistoryRequest) ([]byte, error) {
	return s.Client.Do(MethodGetHistory, r)
}

func (s *Service) GetTitles(r GetTitlesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetTitles, r)
}

func (s *Service) GetVersion(r GetVersionRequest) ([]byte, error) {
	return s.Client.Do(MethodGetVersion, r)
}

func (s *Service) ParseWiki(r ParseWikiRequest) ([]byte, error) {
	return s.Client.Do(MethodParseWiki, r)
}

func (s *Service) ClearCache(r ClearCacheRequest) ([]byte, error) {
	return s.Client.Do(MethodClearCache, r)
}
