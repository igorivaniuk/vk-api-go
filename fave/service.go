package fave

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetUsers       = "fave.getUsers"
	MethodGetPhotos      = "fave.getPhotos"
	MethodGetPosts       = "fave.getPosts"
	MethodGetVideos      = "fave.getVideos"
	MethodGetLinks       = "fave.getLinks"
	MethodGetMarketItems = "fave.getMarketItems"
	MethodAddUser        = "fave.addUser"
	MethodRemoveUser     = "fave.removeUser"
	MethodAddGroup       = "fave.addGroup"
	MethodRemoveGroup    = "fave.removeGroup"
	MethodAddLink        = "fave.addLink"
	MethodRemoveLink     = "fave.removeLink"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetUsers(r GetUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetUsers, r)
}

func (s *Service) GetPhotos(r GetPhotosRequest) ([]byte, error) {
	return s.Client.Do(MethodGetPhotos, r)
}

func (s *Service) GetPosts(r GetPostsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetPosts, r)
}

func (s *Service) GetVideos(r GetVideosRequest) ([]byte, error) {
	return s.Client.Do(MethodGetVideos, r)
}

func (s *Service) GetLinks(r GetLinksRequest) ([]byte, error) {
	return s.Client.Do(MethodGetLinks, r)
}

func (s *Service) GetMarketItems(r GetMarketItemsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetMarketItems, r)
}

func (s *Service) AddUser(r AddUserRequest) ([]byte, error) {
	return s.Client.Do(MethodAddUser, r)
}

func (s *Service) RemoveUser(r RemoveUserRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveUser, r)
}

func (s *Service) AddGroup(r AddGroupRequest) ([]byte, error) {
	return s.Client.Do(MethodAddGroup, r)
}

func (s *Service) RemoveGroup(r RemoveGroupRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveGroup, r)
}

func (s *Service) AddLink(r AddLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodAddLink, r)
}

func (s *Service) RemoveLink(r RemoveLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveLink, r)
}
