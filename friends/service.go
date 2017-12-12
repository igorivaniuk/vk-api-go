package friends

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGet                 = "friends.get"
	MethodGetOnline           = "friends.getOnline"
	MethodGetMutual           = "friends.getMutual"
	MethodGetRecent           = "friends.getRecent"
	MethodGetRequests         = "friends.getRequests"
	MethodAdd                 = "friends.add"
	MethodEdit                = "friends.edit"
	MethodDelete              = "friends.delete"
	MethodGetLists            = "friends.getLists"
	MethodAddList             = "friends.addList"
	MethodEditList            = "friends.editList"
	MethodDeleteList          = "friends.deleteList"
	MethodGetAppUsers         = "friends.getAppUsers"
	MethodGetByPhones         = "friends.getByPhones"
	MethodDeleteAllRequests   = "friends.deleteAllRequests"
	MethodGetSuggestions      = "friends.getSuggestions"
	MethodAreFriends          = "friends.areFriends"
	MethodGetAvailableForCall = "friends.getAvailableForCall"
	MethodSearch              = "friends.search"
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

func (s *Service) GetOnline(r GetOnlineRequest) ([]byte, error) {
	return s.Client.Do(MethodGetOnline, r)
}

func (s *Service) GetMutual(r GetMutualRequest) ([]byte, error) {
	return s.Client.Do(MethodGetMutual, r)
}

func (s *Service) GetRecent(r GetRecentRequest) ([]byte, error) {
	return s.Client.Do(MethodGetRecent, r)
}

func (s *Service) GetRequests(r GetRequestsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetRequests, r)
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

func (s *Service) GetLists(r GetListsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLists, r)
}

func (s *Service) AddList(r AddListRequest) ([]byte, error) {
	return s.Client.Do(MethodAddList, r)
}

func (s *Service) EditList(r EditListRequest) ([]byte, error) {
	return s.Client.Do(MethodEditList, r)
}

func (s *Service) DeleteList(r DeleteListRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteList, r)
}

func (s *Service) GetAppUsers(r GetAppUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAppUsers, r)
}

func (s *Service) GetByPhones(r GetByPhonesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetByPhones, r)
}

func (s *Service) DeleteAllRequests(r DeleteAllRequestsRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteAllRequests, r)
}

func (s *Service) GetSuggestions(r GetSuggestionsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetSuggestions, r)
}

func (s *Service) AreFriends(r AreFriendsRequest) ([]byte, error) {
	return s.Client.Do(MethodAreFriends, r)
}

func (s *Service) GetAvailableForCall(r GetAvailableForCallRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAvailableForCall, r)
}

func (s *Service) Search(r SearchRequest) ([]byte, error) {
	return s.Client.Do(MethodSearch, r)
}
