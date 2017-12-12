package newsfeed

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet                 = "newsfeed.get"
	MethodGetRecommended      = "newsfeed.getRecommended"
	MethodGetComments         = "newsfeed.getComments"
	MethodGetMentions         = "newsfeed.getMentions"
	MethodGetBanned           = "newsfeed.getBanned"
	MethodAddBan              = "newsfeed.addBan"
	MethodDeleteBan           = "newsfeed.deleteBan"
	MethodIgnoreItem          = "newsfeed.ignoreItem"
	MethodUnignoreItem        = "newsfeed.unignoreItem"
	MethodSearch              = "newsfeed.search"
	MethodGetLists            = "newsfeed.getLists"
	MethodSaveList            = "newsfeed.saveList"
	MethodDeleteList          = "newsfeed.deleteList"
	MethodUnsubscribe         = "newsfeed.unsubscribe"
	MethodGetSuggestedSources = "newsfeed.getSuggestedSources"
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

func (s *Service) GetRecommended(r GetRecommendedRequest) ([]byte, error) {
	return s.Client.Do(MethodGetRecommended, r)
}

func (s *Service) GetComments(r GetCommentsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetComments, r)
}

func (s *Service) GetMentions(r GetMentionsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetMentions, r)
}

func (s *Service) GetBanned(r GetBannedRequest) ([]byte, error) {
	return s.Client.Do(MethodGetBanned, r)
}

func (s *Service) AddBan(r AddBanRequest) ([]byte, error) {
	return s.Client.Do(MethodAddBan, r)
}

func (s *Service) DeleteBan(r DeleteBanRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteBan, r)
}

func (s *Service) IgnoreItem(r IgnoreItemRequest) ([]byte, error) {
	return s.Client.Do(MethodIgnoreItem, r)
}

func (s *Service) UnignoreItem(r UnignoreItemRequest) ([]byte, error) {
	return s.Client.Do(MethodUnignoreItem, r)
}

func (s *Service) Search(r SearchRequest) ([]byte, error) {
	return s.Client.Do(MethodSearch, r)
}

func (s *Service) GetLists(r GetListsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLists, r)
}

func (s *Service) SaveList(r SaveListRequest) ([]byte, error) {
	return s.Client.Do(MethodSaveList, r)
}

func (s *Service) DeleteList(r DeleteListRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteList, r)
}

func (s *Service) Unsubscribe(r UnsubscribeRequest) ([]byte, error) {
	return s.Client.Do(MethodUnsubscribe, r)
}

func (s *Service) GetSuggestedSources(r GetSuggestedSourcesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetSuggestedSources, r)
}
