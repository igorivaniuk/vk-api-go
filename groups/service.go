package groups

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodIsMember                    = "groups.isMember"
	MethodGetById                     = "groups.getById"
	MethodGet                         = "groups.get"
	MethodGetMembers                  = "groups.getMembers"
	MethodJoin                        = "groups.join"
	MethodLeave                       = "groups.leave"
	MethodSearch                      = "groups.search"
	MethodGetCatalog                  = "groups.getCatalog"
	MethodGetCatalogInfo              = "groups.getCatalogInfo"
	MethodGetInvites                  = "groups.getInvites"
	MethodGetInvitedUsers             = "groups.getInvitedUsers"
	MethodBanUser                     = "groups.banUser"
	MethodUnbanUser                   = "groups.unbanUser"
	MethodGetBanned                   = "groups.getBanned"
	MethodCreate                      = "groups.create"
	MethodEdit                        = "groups.edit"
	MethodEditPlace                   = "groups.editPlace"
	MethodGetSettings                 = "groups.getSettings"
	MethodGetRequests                 = "groups.getRequests"
	MethodEditManager                 = "groups.editManager"
	MethodInvite                      = "groups.invite"
	MethodAddLink                     = "groups.addLink"
	MethodDeleteLink                  = "groups.deleteLink"
	MethodEditLink                    = "groups.editLink"
	MethodReorderLink                 = "groups.reorderLink"
	MethodRemoveUser                  = "groups.removeUser"
	MethodApproveRequest              = "groups.approveRequest"
	MethodGetCallbackConfirmationCode = "groups.getCallbackConfirmationCode"
	MethodGetCallbackServerSettings   = "groups.getCallbackServerSettings"
	MethodGetCallbackSettings         = "groups.getCallbackSettings"
	MethodSetCallbackServer           = "groups.setCallbackServer"
	MethodSetCallbackServerSettings   = "groups.setCallbackServerSettings"
	MethodSetCallbackSettings         = "groups.setCallbackSettings"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) IsMember(r IsMemberRequest) ([]byte, error) {
	return s.Client.Do(MethodIsMember, r)
}

func (s *Service) GetById(r GetByIdRequest) ([]byte, error) {
	return s.Client.Do(MethodGetById, r)
}

func (s *Service) Get(r GetRequest) ([]byte, error) {
	return s.Client.Do(MethodGet, r)
}

func (s *Service) GetMembers(r GetMembersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetMembers, r)
}

func (s *Service) Join(r JoinRequest) ([]byte, error) {
	return s.Client.Do(MethodJoin, r)
}

func (s *Service) Leave(r LeaveRequest) ([]byte, error) {
	return s.Client.Do(MethodLeave, r)
}

func (s *Service) Search(r SearchRequest) ([]byte, error) {
	return s.Client.Do(MethodSearch, r)
}

func (s *Service) GetCatalog(r GetCatalogRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCatalog, r)
}

func (s *Service) GetCatalogInfo(r GetCatalogInfoRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCatalogInfo, r)
}

func (s *Service) GetInvites(r GetInvitesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetInvites, r)
}

func (s *Service) GetInvitedUsers(r GetInvitedUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetInvitedUsers, r)
}

func (s *Service) BanUser(r BanUserRequest) ([]byte, error) {
	return s.Client.Do(MethodBanUser, r)
}

func (s *Service) UnbanUser(r UnbanUserRequest) ([]byte, error) {
	return s.Client.Do(MethodUnbanUser, r)
}

func (s *Service) GetBanned(r GetBannedRequest) ([]byte, error) {
	return s.Client.Do(MethodGetBanned, r)
}

func (s *Service) Create(r CreateRequest) ([]byte, error) {
	return s.Client.Do(MethodCreate, r)
}

func (s *Service) Edit(r EditRequest) ([]byte, error) {
	return s.Client.Do(MethodEdit, r)
}

func (s *Service) EditPlace(r EditPlaceRequest) ([]byte, error) {
	return s.Client.Do(MethodEditPlace, r)
}

func (s *Service) GetSettings(r GetSettingsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetSettings, r)
}

func (s *Service) GetRequests(r GetRequestsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetRequests, r)
}

func (s *Service) EditManager(r EditManagerRequest) ([]byte, error) {
	return s.Client.Do(MethodEditManager, r)
}

func (s *Service) Invite(r InviteRequest) ([]byte, error) {
	return s.Client.Do(MethodInvite, r)
}

func (s *Service) AddLink(r AddLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodAddLink, r)
}

func (s *Service) DeleteLink(r DeleteLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteLink, r)
}

func (s *Service) EditLink(r EditLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodEditLink, r)
}

func (s *Service) ReorderLink(r ReorderLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodReorderLink, r)
}

func (s *Service) RemoveUser(r RemoveUserRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveUser, r)
}

func (s *Service) ApproveRequest(r ApproveRequestRequest) ([]byte, error) {
	return s.Client.Do(MethodApproveRequest, r)
}

func (s *Service) GetCallbackConfirmationCode(r GetCallbackConfirmationCodeRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCallbackConfirmationCode, r)
}

func (s *Service) GetCallbackServerSettings(r GetCallbackServerSettingsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCallbackServerSettings, r)
}

func (s *Service) GetCallbackSettings(r GetCallbackSettingsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCallbackSettings, r)
}

func (s *Service) SetCallbackServer(r SetCallbackServerRequest) ([]byte, error) {
	return s.Client.Do(MethodSetCallbackServer, r)
}

func (s *Service) SetCallbackServerSettings(r SetCallbackServerSettingsRequest) ([]byte, error) {
	return s.Client.Do(MethodSetCallbackServerSettings, r)
}

func (s *Service) SetCallbackSettings(r SetCallbackSettingsRequest) ([]byte, error) {
	return s.Client.Do(MethodSetCallbackSettings, r)
}
