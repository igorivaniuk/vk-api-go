package pages

const (
	EnumManagers = 0
	EnumMembers  = 1
	EnumAll      = 2
)

type GetRequest struct {
	OwnerId     int    `param:"owner_id"`
	PageId      int    `param:"page_id"`
	Global      bool   `param:"global"`
	SitePreview bool   `param:"site_preview"`
	Title       string `param:"title"`
	NeedSource  bool   `param:"need_source"`
	NeedHtml    bool   `param:"need_html"`
}

type SaveRequest struct {
	Text    string `param:"text"`
	PageId  int    `param:"page_id"`
	GroupId int    `param:"group_id"`
	UserId  int    `param:"user_id"`
	Title   string `param:"title"`
}

type SaveAccessRequest struct {
	PageId  int `param:"page_id"`
	GroupId int `param:"group_id"`
	UserId  int `param:"user_id"`
	View    int `param:"view"`
	Edit    int `param:"edit"`
}

type GetHistoryRequest struct {
	PageId  int `param:"page_id"`
	GroupId int `param:"group_id"`
	UserId  int `param:"user_id"`
}

type GetTitlesRequest struct {
	GroupId int `param:"group_id"`
}

type GetVersionRequest struct {
	VersionId int  `param:"version_id"`
	GroupId   int  `param:"group_id"`
	UserId    int  `param:"user_id"`
	NeedHtml  bool `param:"need_html"`
}

type ParseWikiRequest struct {
	Text    string `param:"text"`
	GroupId int    `param:"group_id"`
}

type ClearCacheRequest struct {
	Url string `param:"url"`
}
