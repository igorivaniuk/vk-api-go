package ads

import (
	"github.com/drossha/vk-api-go"
)

const (
	MethodGetAccounts          = "ads.getAccounts"
	MethodGetClients           = "ads.getClients"
	MethodCreateClients        = "ads.createClients"
	MethodUpdateClients        = "ads.updateClients"
	MethodDeleteClients        = "ads.deleteClients"
	MethodGetCampaigns         = "ads.getCampaigns"
	MethodCreateCampaigns      = "ads.createCampaigns"
	MethodUpdateCampaigns      = "ads.updateCampaigns"
	MethodDeleteCampaigns      = "ads.deleteCampaigns"
	MethodGetAds               = "ads.getAds"
	MethodGetAdsLayout         = "ads.getAdsLayout"
	MethodGetAdsTargeting      = "ads.getAdsTargeting"
	MethodCreateAds            = "ads.createAds"
	MethodUpdateAds            = "ads.updateAds"
	MethodDeleteAds            = "ads.deleteAds"
	MethodCheckLink            = "ads.checkLink"
	MethodGetStatistics        = "ads.getStatistics"
	MethodGetDemographics      = "ads.getDemographics"
	MethodGetAdsPostsReach     = "ads.getAdsPostsReach"
	MethodGetBudget            = "ads.getBudget"
	MethodGetOfficeUsers       = "ads.getOfficeUsers"
	MethodAddOfficeUsers       = "ads.addOfficeUsers"
	MethodRemoveOfficeUsers    = "ads.removeOfficeUsers"
	MethodGetTargetingStats    = "ads.getTargetingStats"
	MethodGetSuggestions       = "ads.getSuggestions"
	MethodGetCategories        = "ads.getCategories"
	MethodGetUploadURL         = "ads.getUploadURL"
	MethodGetVideoUploadURL    = "ads.getVideoUploadURL"
	MethodGetFloodStats        = "ads.getFloodStats"
	MethodGetRejectionReason   = "ads.getRejectionReason"
	MethodCreateTargetGroup    = "ads.createTargetGroup"
	MethodUpdateTargetGroup    = "ads.updateTargetGroup"
	MethodDeleteTargetGroup    = "ads.deleteTargetGroup"
	MethodGetTargetGroups      = "ads.getTargetGroups"
	MethodImportTargetContacts = "ads.importTargetContacts"
)

type Service struct {
	Client *vk_api_go.Client
}

func New(Client *vk_api_go.Client) (s Service) {
	return Service{Client: Client}
}

func (s *Service) GetAccounts(r GetAccountsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAccounts, r)
}

func (s *Service) GetClients(r GetClientsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetClients, r)
}

func (s *Service) CreateClients(r CreateClientsRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateClients, r)
}

func (s *Service) UpdateClients(r UpdateClientsRequest) ([]byte, error) {
	return s.Client.Do(MethodUpdateClients, r)
}

func (s *Service) DeleteClients(r DeleteClientsRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteClients, r)
}

func (s *Service) GetCampaigns(r GetCampaignsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCampaigns, r)
}

func (s *Service) CreateCampaigns(r CreateCampaignsRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateCampaigns, r)
}

func (s *Service) UpdateCampaigns(r UpdateCampaignsRequest) ([]byte, error) {
	return s.Client.Do(MethodUpdateCampaigns, r)
}

func (s *Service) DeleteCampaigns(r DeleteCampaignsRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteCampaigns, r)
}

func (s *Service) GetAds(r GetAdsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAds, r)
}

func (s *Service) GetAdsLayout(r GetAdsLayoutRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAdsLayout, r)
}

func (s *Service) GetAdsTargeting(r GetAdsTargetingRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAdsTargeting, r)
}

func (s *Service) CreateAds(r CreateAdsRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateAds, r)
}

func (s *Service) UpdateAds(r UpdateAdsRequest) ([]byte, error) {
	return s.Client.Do(MethodUpdateAds, r)
}

func (s *Service) DeleteAds(r DeleteAdsRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteAds, r)
}

func (s *Service) CheckLink(r CheckLinkRequest) ([]byte, error) {
	return s.Client.Do(MethodCheckLink, r)
}

func (s *Service) GetStatistics(r GetStatisticsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetStatistics, r)
}

func (s *Service) GetDemographics(r GetDemographicsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetDemographics, r)
}

func (s *Service) GetAdsPostsReach(r GetAdsPostsReachRequest) ([]byte, error) {
	return s.Client.Do(MethodGetAdsPostsReach, r)
}

func (s *Service) GetBudget(r GetBudgetRequest) ([]byte, error) {
	return s.Client.Do(MethodGetBudget, r)
}

func (s *Service) GetOfficeUsers(r GetOfficeUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodGetOfficeUsers, r)
}

func (s *Service) AddOfficeUsers(r AddOfficeUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodAddOfficeUsers, r)
}

func (s *Service) RemoveOfficeUsers(r RemoveOfficeUsersRequest) ([]byte, error) {
	return s.Client.Do(MethodRemoveOfficeUsers, r)
}

func (s *Service) GetTargetingStats(r GetTargetingStatsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetTargetingStats, r)
}

func (s *Service) GetSuggestions(r GetSuggestionsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetSuggestions, r)
}

func (s *Service) GetCategories(r GetCategoriesRequest) ([]byte, error) {
	return s.Client.Do(MethodGetCategories, r)
}

func (s *Service) GetUploadURL(r GetUploadURLRequest) ([]byte, error) {
	return s.Client.Do(MethodGetUploadURL, r)
}

func (s *Service) GetVideoUploadURL(r GetVideoUploadURLRequest) ([]byte, error) {
	return s.Client.Do(MethodGetVideoUploadURL, r)
}

func (s *Service) GetFloodStats(r GetFloodStatsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetFloodStats, r)
}

func (s *Service) GetRejectionReason(r GetRejectionReasonRequest) ([]byte, error) {
	return s.Client.Do(MethodGetRejectionReason, r)
}

func (s *Service) CreateTargetGroup(r CreateTargetGroupRequest) ([]byte, error) {
	return s.Client.Do(MethodCreateTargetGroup, r)
}

func (s *Service) UpdateTargetGroup(r UpdateTargetGroupRequest) ([]byte, error) {
	return s.Client.Do(MethodUpdateTargetGroup, r)
}

func (s *Service) DeleteTargetGroup(r DeleteTargetGroupRequest) ([]byte, error) {
	return s.Client.Do(MethodDeleteTargetGroup, r)
}

func (s *Service) GetTargetGroups(r GetTargetGroupsRequest) ([]byte, error) {
	return s.Client.Do(MethodGetTargetGroups, r)
}

func (s *Service) ImportTargetContacts(r ImportTargetContactsRequest) ([]byte, error) {
	return s.Client.Do(MethodImportTargetContacts, r)
}
