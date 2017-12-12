package account

const (
	EnumEmail             = "email"
	EnumPhone             = "phone"
	EnumTwitter           = "twitter"
	EnumFacebook          = "facebook"
	EnumOdnoklassniki     = "odnoklassniki"
	EnumInstagram         = "instagram"
	EnumGoogle            = "google"
	EnumUndefined         = 0
	EnumFemale            = 1
	EnumMale              = 2
	EnumSingle            = 1
	EnumRelationship      = 2
	EnumEngaged           = 3
	EnumMarried           = 4
	EnumComplicated       = 5
	EnumActivelySearching = 6
	EnumInLove            = 7
	EnumNotSpecified      = 0
	EnumShow              = 1
	EnumHideYear          = 2
	EnumHide              = 0
)

type GetCountersRequest struct {
	Filter string `param:"filter"`
}

type SetNameInMenuRequest struct {
	UserId int    `param:"user_id"`
	Name   string `param:"name"`
}

type SetOnlineRequest struct {
	Voip bool `param:"voip"`
}

type SetOfflineRequest struct {
}

type LookupContactsRequest struct {
	Contacts  string `param:"contacts"`
	Service   string `param:"service"`
	Mycontact string `param:"mycontact"`
	ReturnAll bool   `param:"return_all"`
	Fields    string `param:"fields"`
}

type RegisterDeviceRequest struct {
	Token         string `param:"token"`
	DeviceModel   string `param:"device_model"`
	DeviceYear    int    `param:"device_year"`
	DeviceId      string `param:"device_id"`
	SystemVersion string `param:"system_version"`
	Settings      string `param:"settings"`
}

type UnregisterDeviceRequest struct {
	DeviceId string `param:"device_id"`
}

type SetSilenceModeRequest struct {
	DeviceId string `param:"device_id"`
	Time     int    `param:"time"`
	PeerId   int    `param:"peer_id"`
	Sound    int    `param:"sound"`
}

type GetPushSettingsRequest struct {
	DeviceId string `param:"device_id"`
}

type SetPushSettingsRequest struct {
	DeviceId string `param:"device_id"`
	Settings string `param:"settings"`
	Key      string `param:"key"`
	Value    string `param:"value"`
}

type GetAppPermissionsRequest struct {
	UserId int `param:"user_id"`
}

type GetActiveOffersRequest struct {
	Count int `param:"count"`
}

type BanUserRequest struct {
	UserId int `param:"user_id"`
}

type UnbanUserRequest struct {
	UserId int `param:"user_id"`
}

type GetBannedRequest struct {
	Offset int `param:"offset"`
	Count  int `param:"count"`
}

type GetInfoRequest struct {
	Fields string `param:"fields"`
}

type SetInfoRequest struct {
	Name  string `param:"name"`
	Value string `param:"value"`
}

type ChangePasswordRequest struct {
	RestoreSid         string `param:"restore_sid"`
	ChangePasswordHash string `param:"change_password_hash"`
	OldPassword        string `param:"old_password"`
	NewPassword        string `param:"new_password"`
}

type GetProfileInfoRequest struct {
}

type SaveProfileInfoRequest struct {
	FirstName         string `param:"first_name"`
	LastName          string `param:"last_name"`
	MaidenName        string `param:"maiden_name"`
	ScreenName        string `param:"screen_name"`
	CancelRequestId   int    `param:"cancel_request_id"`
	Sex               int    `param:"sex"`
	Relation          int    `param:"relation"`
	RelationPartnerId int    `param:"relation_partner_id"`
	Bdate             string `param:"bdate"`
	BdateVisibility   int    `param:"bdate_visibility"`
	HomeTown          string `param:"home_town"`
	CountryId         int    `param:"country_id"`
	CityId            int    `param:"city_id"`
	Status            string `param:"status"`
}
