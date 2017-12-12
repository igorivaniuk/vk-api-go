package friends

const (
	EnumName          = "name"
	EnumHints         = "hints"
	EnumNominative    = "nom"
	EnumGenitive      = "gen"
	EnumDative        = "dat"
	EnumAccusative    = "acc"
	EnumInstrumental  = "ins"
	EnumPrepositional = "abl"
	EnumDate          = 0
	EnumMutual        = 1
)

type GetRequest struct {
	UserId   int    `param:"user_id"`
	Order    string `param:"order"`
	ListId   int    `param:"list_id"`
	Count    int    `param:"count"`
	Offset   int    `param:"offset"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type GetOnlineRequest struct {
	UserId       int    `param:"user_id"`
	ListId       int    `param:"list_id"`
	OnlineMobile bool   `param:"online_mobile"`
	Order        string `param:"order"`
	Count        int    `param:"count"`
	Offset       int    `param:"offset"`
}

type GetMutualRequest struct {
	SourceUid  int    `param:"source_uid"`
	TargetUid  int    `param:"target_uid"`
	TargetUids string `param:"target_uids"`
	Order      string `param:"order"`
	Count      int    `param:"count"`
	Offset     int    `param:"offset"`
}

type GetRecentRequest struct {
	Count int `param:"count"`
}

type GetRequestsRequest struct {
	Offset     int  `param:"offset"`
	Count      int  `param:"count"`
	Extended   bool `param:"extended"`
	NeedMutual bool `param:"need_mutual"`
	Out        bool `param:"out"`
	Sort       int  `param:"sort"`
	Suggested  bool `param:"suggested"`
}

type AddRequest struct {
	UserId int    `param:"user_id"`
	Text   string `param:"text"`
	Follow bool   `param:"follow"`
}

type EditRequest struct {
	UserId  int    `param:"user_id"`
	ListIds string `param:"list_ids"`
}

type DeleteRequest struct {
	UserId int `param:"user_id"`
}

type GetListsRequest struct {
	UserId       int  `param:"user_id"`
	ReturnSystem bool `param:"return_system"`
}

type AddListRequest struct {
	Name    string `param:"name"`
	UserIds string `param:"user_ids"`
}

type EditListRequest struct {
	Name          string `param:"name"`
	ListId        int    `param:"list_id"`
	UserIds       string `param:"user_ids"`
	AddUserIds    string `param:"add_user_ids"`
	DeleteUserIds string `param:"delete_user_ids"`
}

type DeleteListRequest struct {
	ListId int `param:"list_id"`
}

type GetAppUsersRequest struct {
}

type GetByPhonesRequest struct {
	Phones string `param:"phones"`
	Fields string `param:"fields"`
}

type DeleteAllRequestsRequest struct {
}

type GetSuggestionsRequest struct {
	Filter   string `param:"filter"`
	Count    int    `param:"count"`
	Offset   int    `param:"offset"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type AreFriendsRequest struct {
	UserIds  string `param:"user_ids"`
	NeedSign bool   `param:"need_sign"`
}

type GetAvailableForCallRequest struct {
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type SearchRequest struct {
	UserId   int    `param:"user_id"`
	Q        string `param:"q"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
}
