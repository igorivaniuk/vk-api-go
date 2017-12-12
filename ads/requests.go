package ads

const (
	EnumCommunity              = "community"
	EnumPost                   = "post"
	EnumApplication            = "application"
	EnumVideo                  = "video"
	EnumSite                   = "site"
	EnumAd                     = "ad"
	EnumCampaign               = "campaign"
	EnumClient                 = "client"
	EnumOffice                 = "office"
	EnumDay                    = "day"
	EnumMonth                  = "month"
	EnumOverall                = "overall"
	EnumImageAndText           = 1
	EnumBigImage               = 2
	EnumExclusiveFormat        = 3
	EnumCommunitySquareImage   = 4
	EnumSpecialAppFormat       = 7
	EnumSpecialCommunityFormat = 8
	EnumPostInCommunity        = 9
	EnumAppBoard               = 10
	EnumCountries              = "countries"
	EnumRegions                = "regions"
	EnumCities                 = "cities"
	EnumDistricts              = "districts"
	EnumStations               = "stations"
	EnumStreets                = "streets"
	EnumSchools                = "schools"
	EnumInterests              = "interests"
	EnumPositions              = "positions"
	EnumGroupTypes             = "group_types"
	EnumReligions              = "religions"
	EnumBrowsers               = "browsers"
	EnumRussian                = "ru"
	EnumUkrainian              = "ua"
	EnumEnglish                = "en"
)

type GetAccountsRequest struct {
}

type GetClientsRequest struct {
	AccountId int `param:"account_id"`
}

type CreateClientsRequest struct {
	AccountId int    `param:"account_id"`
	Data      string `param:"data"`
}

type UpdateClientsRequest struct {
	AccountId int    `param:"account_id"`
	Data      string `param:"data"`
}

type DeleteClientsRequest struct {
	AccountId int    `param:"account_id"`
	Ids       string `param:"ids"`
}

type GetCampaignsRequest struct {
	AccountId      int    `param:"account_id"`
	ClientId       int    `param:"client_id"`
	IncludeDeleted bool   `param:"include_deleted"`
	CampaignIds    string `param:"campaign_ids"`
}

type CreateCampaignsRequest struct {
	AccountId int    `param:"account_id"`
	Data      string `param:"data"`
}

type UpdateCampaignsRequest struct {
	AccountId int    `param:"account_id"`
	Data      string `param:"data"`
}

type DeleteCampaignsRequest struct {
	AccountId int    `param:"account_id"`
	Ids       string `param:"ids"`
}

type GetAdsRequest struct {
	AccountId      int    `param:"account_id"`
	ClientId       int    `param:"client_id"`
	IncludeDeleted bool   `param:"include_deleted"`
	CampaignIds    string `param:"campaign_ids"`
	AdIds          string `param:"ad_ids"`
	Limit          int    `param:"limit"`
	Offset         int    `param:"offset"`
}

type GetAdsLayoutRequest struct {
	AccountId      int    `param:"account_id"`
	ClientId       int    `param:"client_id"`
	IncludeDeleted bool   `param:"include_deleted"`
	CampaignIds    string `param:"campaign_ids"`
	AdIds          string `param:"ad_ids"`
	Limit          int    `param:"limit"`
	Offset         int    `param:"offset"`
}

type GetAdsTargetingRequest struct {
	AccountId      int    `param:"account_id"`
	ClientId       int    `param:"client_id"`
	IncludeDeleted bool   `param:"include_deleted"`
	CampaignIds    string `param:"campaign_ids"`
	AdIds          string `param:"ad_ids"`
	Limit          int    `param:"limit"`
	Offset         int    `param:"offset"`
}

type CreateAdsRequest struct {
	AccountId int    `param:"account_id"`
	Data      string `param:"data"`
}

type UpdateAdsRequest struct {
	AccountId int    `param:"account_id"`
	Data      string `param:"data"`
}

type DeleteAdsRequest struct {
	AccountId int    `param:"account_id"`
	Ids       string `param:"ids"`
}

type CheckLinkRequest struct {
	AccountId  int    `param:"account_id"`
	LinkType   string `param:"link_type"`
	LinkUrl    string `param:"link_url"`
	CampaignId int    `param:"campaign_id"`
}

type GetStatisticsRequest struct {
	AccountId int    `param:"account_id"`
	IdsType   string `param:"ids_type"`
	Ids       string `param:"ids"`
	Period    string `param:"period"`
	DateFrom  string `param:"date_from"`
	DateTo    string `param:"date_to"`
}

type GetDemographicsRequest struct {
	AccountId int    `param:"account_id"`
	IdsType   string `param:"ids_type"`
	Ids       string `param:"ids"`
	Period    string `param:"period"`
	DateFrom  string `param:"date_from"`
	DateTo    string `param:"date_to"`
}

type GetAdsPostsReachRequest struct {
	AccountId int    `param:"account_id"`
	AdsIds    string `param:"ads_ids"`
}

type GetBudgetRequest struct {
	AccountId int `param:"account_id"`
}

type GetOfficeUsersRequest struct {
	AccountId int `param:"account_id"`
}

type AddOfficeUsersRequest struct {
	AccountId int    `param:"account_id"`
	Data      string `param:"data"`
}

type RemoveOfficeUsersRequest struct {
	AccountId int    `param:"account_id"`
	Ids       string `param:"ids"`
}

type GetTargetingStatsRequest struct {
	AccountId  int    `param:"account_id"`
	Criteria   string `param:"criteria"`
	AdId       int    `param:"ad_id"`
	AdFormat   int    `param:"ad_format"`
	AdPlatform string `param:"ad_platform"`
	LinkUrl    string `param:"link_url"`
	LinkDomain string `param:"link_domain"`
}

type GetSuggestionsRequest struct {
	Section string `param:"section"`
	Ids     string `param:"ids"`
	Q       string `param:"q"`
	Country int    `param:"country"`
	Cities  string `param:"cities"`
	Lang    string `param:"lang"`
}

type GetCategoriesRequest struct {
	Lang string `param:"lang"`
}

type GetUploadURLRequest struct {
	AdFormat int `param:"ad_format"`
}

type GetVideoUploadURLRequest struct {
}

type GetFloodStatsRequest struct {
	AccountId int `param:"account_id"`
}

type GetRejectionReasonRequest struct {
	AccountId int `param:"account_id"`
	AdId      int `param:"ad_id"`
}

type CreateTargetGroupRequest struct {
	AccountId int    `param:"account_id"`
	ClientId  int    `param:"client_id"`
	Name      string `param:"name"`
	Domain    string `param:"domain"`
	Lifetime  int    `param:"lifetime"`
}

type UpdateTargetGroupRequest struct {
	AccountId     int    `param:"account_id"`
	ClientId      int    `param:"client_id"`
	TargetGroupId int    `param:"target_group_id"`
	Name          string `param:"name"`
	Domain        string `param:"domain"`
	Lifetime      int    `param:"lifetime"`
}

type DeleteTargetGroupRequest struct {
	AccountId     int `param:"account_id"`
	ClientId      int `param:"client_id"`
	TargetGroupId int `param:"target_group_id"`
}

type GetTargetGroupsRequest struct {
	AccountId int  `param:"account_id"`
	ClientId  int  `param:"client_id"`
	Extended  bool `param:"extended"`
}

type ImportTargetContactsRequest struct {
	AccountId     int    `param:"account_id"`
	ClientId      int    `param:"client_id"`
	TargetGroupId int    `param:"target_group_id"`
	Contacts      string `param:"contacts"`
}
