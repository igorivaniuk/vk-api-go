package account

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetCounters       = "account.getCounters"
	MethodSetNameInMenu     = "account.setNameInMenu"
	MethodSetOnline         = "account.setOnline"
	MethodSetOffline        = "account.setOffline"
	MethodLookupContacts    = "account.lookupContacts"
	MethodRegisterDevice    = "account.registerDevice"
	MethodUnregisterDevice  = "account.unregisterDevice"
	MethodSetSilenceMode    = "account.setSilenceMode"
	MethodGetPushSettings   = "account.getPushSettings"
	MethodSetPushSettings   = "account.setPushSettings"
	MethodGetAppPermissions = "account.getAppPermissions"
	MethodGetActiveOffers   = "account.getActiveOffers"
	MethodBanUser           = "account.banUser"
	MethodUnbanUser         = "account.unbanUser"
	MethodGetBanned         = "account.getBanned"
	MethodGetInfo           = "account.getInfo"
	MethodSetInfo           = "account.setInfo"
	MethodChangePassword    = "account.changePassword"
	MethodGetProfileInfo    = "account.getProfileInfo"
	MethodSaveProfileInfo   = "account.saveProfileInfo"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetCounters(r GetCountersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCounters, r)
}

func (s *Service) SetNameInMenu(r SetNameInMenuRequest) ([]byte, error) {
	return s.Client.Do(MethodSetNameInMenu, r)
}

func (s *Service) SetOnline(r SetOnlineRequest) ([]byte, error) {
	return s.Client.Do(MethodSetOnline, r)
}

func (s *Service) SetOffline(r SetOfflineRequest) ([]byte, error) {
	return s.Client.Do(MethodSetOffline, r)
}

func (s *Service) LookupContacts(r LookupContactsRequest) ([]byte, error) {
	return s.Client.Do(MethodLookupContacts, r)
}

func (s *Service) RegisterDevice(r RegisterDeviceRequest) ([]byte, error) {
	return s.Client.Do(MethodRegisterDevice, r)
}

func (s *Service) UnregisterDevice(r UnregisterDeviceRequest) ([]byte, error) {
	return s.Client.Do(MethodUnregisterDevice, r)
}

func (s *Service) SetSilenceMode(r SetSilenceModeRequest) ([]byte, error) {
	return s.Client.Do(MethodSetSilenceMode, r)
}

func (s *Service) GetPushSettings(r GetPushSettingsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetPushSettings, r)
}

func (s *Service) SetPushSettings(r SetPushSettingsRequest) ([]byte, error) {
	return s.Client.Do(MethodSetPushSettings, r)
}

func (s *Service) GetAppPermissions(r GetAppPermissionsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAppPermissions, r)
}

func (s *Service) GetActiveOffers(r GetActiveOffersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetActiveOffers, r)
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

func (s *Service) GetInfo(r GetInfoRequest) ([]byte, error) {
	return s.Client.Do(MethodGetInfo, r)
}

func (s *Service) SetInfo(r SetInfoRequest) ([]byte, error) {
	return s.Client.Do(MethodSetInfo, r)
}

func (s *Service) ChangePassword(r ChangePasswordRequest) ([]byte, error) {
	return s.Client.Do(MethodChangePassword, r)
}

func (s *Service) GetProfileInfo(r GetProfileInfoRequest) ([]byte, error) {
	return s.Client.Do(MethodGetProfileInfo, r)
}

func (s *Service) SaveProfileInfo(r SaveProfileInfoRequest) ([]byte, error) {
	return s.Client.Do(MethodSaveProfileInfo, r)
}
