package newsfeed

const (
	EnumNominative    = "nom"
	EnumGenitive      = "gen"
	EnumDative        = "dat"
	EnumAccusative    = "acc"
	EnumInstrumental  = "ins"
	EnumPrepositional = "abl"
	EnumPostOnTheWall = "wall"
	EnumTagOnAPhoto   = "tag"
	EnumProfilePhoto  = "profilephoto"
	EnumVideo         = "video"
	EnumAudio         = "audio"
	EnumWall          = "wall"
	EnumTag           = "tag"
	EnumProfilephoto  = "profilephoto"
	EnumNote          = "note"
	EnumPhoto         = "photo"
	EnumPost          = "post"
	EnumTopic         = "topic"
)

type GetRequest struct {
	Filters      string `param:"filters"`
	ReturnBanned bool   `param:"return_banned"`
	StartTime    int    `param:"start_time"`
	EndTime      int    `param:"end_time"`
	MaxPhotos    int    `param:"max_photos"`
	SourceIds    string `param:"source_ids"`
	StartFrom    string `param:"start_from"`
	Count        int    `param:"count"`
	Fields       string `param:"fields"`
}

type GetRecommendedRequest struct {
	StartTime int    `param:"start_time"`
	EndTime   int    `param:"end_time"`
	MaxPhotos int    `param:"max_photos"`
	StartFrom string `param:"start_from"`
	Count     int    `param:"count"`
	Fields    string `param:"fields"`
}

type GetCommentsRequest struct {
	Count     int    `param:"count"`
	Filters   string `param:"filters"`
	Reposts   string `param:"reposts"`
	StartTime int    `param:"start_time"`
	EndTime   int    `param:"end_time"`
	StartFrom string `param:"start_from"`
	Fields    string `param:"fields"`
}

type GetMentionsRequest struct {
	OwnerId   int `param:"owner_id"`
	StartTime int `param:"start_time"`
	EndTime   int `param:"end_time"`
	Offset    int `param:"offset"`
	Count     int `param:"count"`
}

type GetBannedRequest struct {
	Extended bool   `param:"extended"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type AddBanRequest struct {
	UserIds  string `param:"user_ids"`
	GroupIds string `param:"group_ids"`
}

type DeleteBanRequest struct {
	UserIds  string `param:"user_ids"`
	GroupIds string `param:"group_ids"`
}

type IgnoreItemRequest struct {
	Type    string `param:"type"`
	OwnerId int    `param:"owner_id"`
	ItemId  int    `param:"item_id"`
}

type UnignoreItemRequest struct {
	Type    string `param:"type"`
	OwnerId int    `param:"owner_id"`
	ItemId  int    `param:"item_id"`
}

type SearchRequest struct {
	Q         string  `param:"q"`
	Extended  bool    `param:"extended"`
	Count     int     `param:"count"`
	Latitude  float64 `param:"latitude"`
	Longitude float64 `param:"longitude"`
	StartTime int     `param:"start_time"`
	EndTime   int     `param:"end_time"`
	StartFrom string  `param:"start_from"`
	Fields    string  `param:"fields"`
}

type GetListsRequest struct {
	ListIds  string `param:"list_ids"`
	Extended bool   `param:"extended"`
}

type SaveListRequest struct {
	ListId    int    `param:"list_id"`
	Title     string `param:"title"`
	SourceIds string `param:"source_ids"`
	NoReposts bool   `param:"no_reposts"`
}

type DeleteListRequest struct {
	ListId int `param:"list_id"`
}

type UnsubscribeRequest struct {
	Type    string `param:"type"`
	OwnerId int    `param:"owner_id"`
	ItemId  int    `param:"item_id"`
}

type GetSuggestedSourcesRequest struct {
	Offset  int    `param:"offset"`
	Count   int    `param:"count"`
	Shuffle bool   `param:"shuffle"`
	Fields  string `param:"fields"`
}
