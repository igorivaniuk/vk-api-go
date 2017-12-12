package groups

const (
	EnumIdAsc                     = "id_asc"
	EnumIdDesc                    = "id_desc"
	EnumTimeAsc                   = "time_asc"
	EnumTimeDesc                  = "time_desc"
	EnumFriends                   = "friends"
	EnumUnsure                    = "unsure"
	EnumGroup                     = "group"
	EnumPage                      = "page"
	EnumEvent                     = "event"
	EnumDefault                   = 0
	EnumGrowth                    = 1
	EnumAttendance                = 2
	EnumLikes                     = 3
	EnumComments                  = 4
	EnumEntries                   = 5
	EnumNominative                = "nom"
	EnumGenitive                  = "gen"
	EnumDative                    = "dat"
	EnumAccusative                = "acc"
	EnumInstrumental              = "ins"
	EnumPrepositional             = "abl"
	EnumOther                     = 0
	EnumSpam                      = 1
	EnumVerbalAbuse               = 2
	EnumStrongLanguage            = 3
	EnumIrrelevantMessages        = 4
	EnumPublic                    = "public"
	EnumPlaceOrBusiness           = 1
	EnumCompanyOrWebsite          = 2
	EnumPersonOrGroup             = 3
	EnumProductOrArt              = 4
	EnumOpen                      = 0
	EnumClosed                    = 1
	EnumPrivate                   = 2
	EnumAuto                      = 1
	EnumActivityHolidays          = 2
	EnumBusiness                  = 3
	EnumPets                      = 4
	EnumHealth                    = 5
	EnumDatingAndCommunication    = 6
	EnumGames                     = 7
	EnumIt                        = 8
	EnumCinema                    = 9
	EnumBeautyAndFashion          = 10
	EnumCooking                   = 11
	EnumArtAndCulture             = 12
	EnumLiterature                = 13
	EnumMobileServicesAndInternet = 14
	EnumMusic                     = 15
	EnumScienceAndTechnology      = 16
	EnumRealEstate                = 17
	EnumNewsAndMedia              = 18
	EnumSecurity                  = 19
	EnumEducation                 = 20
	EnumHomeAndRenovations        = 21
	EnumPolitics                  = 22
	EnumFood                      = 23
	EnumIndustry                  = 24
	EnumTravel                    = 25
	EnumWork                      = 26
	EnumEntertainment             = 27
	EnumReligion                  = 28
	EnumFamily                    = 29
	EnumSports                    = 30
	EnumInsurance                 = 31
	EnumTelevision                = 32
	EnumGoodsAndServices          = 33
	EnumHobbies                   = 34
	EnumFinance                   = 35
	EnumPhoto                     = 36
	EnumEsoterics                 = 37
	EnumElectronicsAndAppliances  = 38
	EnumErotic                    = 39
	EnumHumor                     = 40
	EnumSocietyHumanities         = 41
	EnumDesignAndGraphics         = 42
	EnumDisabled                  = 0
	EnumOpen                      = 1
	EnumLimited                   = 2
	EnumClosed                    = 3
	EnumUnlimited                 = 1
	Enum16Plus                    = 2
	Enum18Plus                    = 3
	EnumRussianRubles             = 643
	EnumUkrainianHryvnia          = 980
	EnumKazakhTenge               = 398
	EnumEuro                      = 978
	EnumUsDollars                 = 840
	EnumModerator                 = "moderator"
	EnumEditor                    = "editor"
	EnumAdministrator             = "administrator"
)

type IsMemberRequest struct {
	GroupId  string `param:"group_id"`
	UserId   int    `param:"user_id"`
	UserIds  string `param:"user_ids"`
	Extended bool   `param:"extended"`
}

type GetByIdRequest struct {
	GroupIds string `param:"group_ids"`
	GroupId  string `param:"group_id"`
	Fields   string `param:"fields"`
}

type GetRequest struct {
	UserId   int    `param:"user_id"`
	Extended bool   `param:"extended"`
	Filter   string `param:"filter"`
	Fields   string `param:"fields"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
}

type GetMembersRequest struct {
	GroupId string `param:"group_id"`
	Sort    string `param:"sort"`
	Offset  int    `param:"offset"`
	Count   int    `param:"count"`
	Fields  string `param:"fields"`
	Filter  string `param:"filter"`
}

type JoinRequest struct {
	GroupId int    `param:"group_id"`
	NotSure string `param:"not_sure"`
}

type LeaveRequest struct {
	GroupId int `param:"group_id"`
}

type SearchRequest struct {
	Q         string `param:"q"`
	Type      string `param:"type"`
	CountryId int    `param:"country_id"`
	CityId    int    `param:"city_id"`
	Future    bool   `param:"future"`
	Market    bool   `param:"market"`
	Sort      int    `param:"sort"`
	Offset    int    `param:"offset"`
	Count     int    `param:"count"`
}

type GetCatalogRequest struct {
	CategoryId    int `param:"category_id"`
	SubcategoryId int `param:"subcategory_id"`
}

type GetCatalogInfoRequest struct {
	Extended      bool `param:"extended"`
	Subcategories bool `param:"subcategories"`
}

type GetInvitesRequest struct {
	Offset   int  `param:"offset"`
	Count    int  `param:"count"`
	Extended bool `param:"extended"`
}

type GetInvitedUsersRequest struct {
	GroupId  int    `param:"group_id"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type BanUserRequest struct {
	GroupId        int    `param:"group_id"`
	UserId         int    `param:"user_id"`
	EndDate        int    `param:"end_date"`
	Reason         int    `param:"reason"`
	Comment        string `param:"comment"`
	CommentVisible bool   `param:"comment_visible"`
}

type UnbanUserRequest struct {
	GroupId int `param:"group_id"`
	UserId  int `param:"user_id"`
}

type GetBannedRequest struct {
	GroupId int    `param:"group_id"`
	Offset  int    `param:"offset"`
	Count   int    `param:"count"`
	Fields  string `param:"fields"`
	UserId  int    `param:"user_id"`
}

type CreateRequest struct {
	Title          string `param:"title"`
	Description    string `param:"description"`
	Type           string `param:"type"`
	PublicCategory int    `param:"public_category"`
	Subtype        int    `param:"subtype"`
}

type EditRequest struct {
	GroupId           int    `param:"group_id"`
	Title             string `param:"title"`
	Description       string `param:"description"`
	ScreenName        string `param:"screen_name"`
	Access            int    `param:"access"`
	Website           string `param:"website"`
	Subject           string `param:"subject"`
	Email             string `param:"email"`
	Phone             string `param:"phone"`
	Rss               string `param:"rss"`
	EventStartDate    int    `param:"event_start_date"`
	EventFinishDate   int    `param:"event_finish_date"`
	EventGroupId      int    `param:"event_group_id"`
	PublicCategory    int    `param:"public_category"`
	PublicSubcategory int    `param:"public_subcategory"`
	PublicDate        string `param:"public_date"`
	Wall              int    `param:"wall"`
	Topics            int    `param:"topics"`
	Photos            int    `param:"photos"`
	Video             int    `param:"video"`
	Audio             int    `param:"audio"`
	Links             bool   `param:"links"`
	Events            bool   `param:"events"`
	Places            bool   `param:"places"`
	Contacts          bool   `param:"contacts"`
	Docs              int    `param:"docs"`
	Wiki              int    `param:"wiki"`
	Messages          bool   `param:"messages"`
	AgeLimits         int    `param:"age_limits"`
	Market            bool   `param:"market"`
	MarketComments    bool   `param:"market_comments"`
	MarketCountry     string `param:"market_country"`
	MarketCity        string `param:"market_city"`
	MarketCurrency    int    `param:"market_currency"`
	MarketContact     int    `param:"market_contact"`
	MarketWiki        int    `param:"market_wiki"`
	ObsceneFilter     bool   `param:"obscene_filter"`
	ObsceneStopwords  bool   `param:"obscene_stopwords"`
	ObsceneWords      string `param:"obscene_words"`
}

type EditPlaceRequest struct {
	GroupId   int     `param:"group_id"`
	Title     string  `param:"title"`
	Address   string  `param:"address"`
	CountryId int     `param:"country_id"`
	CityId    int     `param:"city_id"`
	Latitude  float64 `param:"latitude"`
	Longitude float64 `param:"longitude"`
}

type GetSettingsRequest struct {
	GroupId int `param:"group_id"`
}

type GetRequestsRequest struct {
	GroupId int    `param:"group_id"`
	Offset  int    `param:"offset"`
	Count   int    `param:"count"`
	Fields  string `param:"fields"`
}

type EditManagerRequest struct {
	GroupId         int    `param:"group_id"`
	UserId          int    `param:"user_id"`
	Role            string `param:"role"`
	IsContact       bool   `param:"is_contact"`
	ContactPosition string `param:"contact_position"`
	ContactPhone    string `param:"contact_phone"`
	ContactEmail    string `param:"contact_email"`
}

type InviteRequest struct {
	GroupId int `param:"group_id"`
	UserId  int `param:"user_id"`
}

type AddLinkRequest struct {
	GroupId int    `param:"group_id"`
	Link    string `param:"link"`
	Text    string `param:"text"`
}

type DeleteLinkRequest struct {
	GroupId int `param:"group_id"`
	LinkId  int `param:"link_id"`
}

type EditLinkRequest struct {
	GroupId int    `param:"group_id"`
	LinkId  int    `param:"link_id"`
	Text    string `param:"text"`
}

type ReorderLinkRequest struct {
	GroupId int `param:"group_id"`
	LinkId  int `param:"link_id"`
	After   int `param:"after"`
}

type RemoveUserRequest struct {
	GroupId int `param:"group_id"`
	UserId  int `param:"user_id"`
}

type ApproveRequestRequest struct {
	GroupId int `param:"group_id"`
	UserId  int `param:"user_id"`
}

type GetCallbackConfirmationCodeRequest struct {
	GroupId int `param:"group_id"`
}

type GetCallbackServerSettingsRequest struct {
	GroupId int `param:"group_id"`
}

type GetCallbackSettingsRequest struct {
	GroupId int `param:"group_id"`
}

type SetCallbackServerRequest struct {
	GroupId   int    `param:"group_id"`
	ServerUrl string `param:"server_url"`
}

type SetCallbackServerSettingsRequest struct {
	GroupId   int    `param:"group_id"`
	SecretKey string `param:"secret_key"`
}

type SetCallbackSettingsRequest struct {
	GroupId              int  `param:"group_id"`
	MessageNew           bool `param:"message_new"`
	MessageReply         bool `param:"message_reply"`
	MessageAllow         bool `param:"message_allow"`
	MessageDeny          bool `param:"message_deny"`
	PhotoNew             bool `param:"photo_new"`
	AudioNew             bool `param:"audio_new"`
	VideoNew             bool `param:"video_new"`
	WallReplyNew         bool `param:"wall_reply_new"`
	WallReplyEdit        bool `param:"wall_reply_edit"`
	WallReplyDelete      bool `param:"wall_reply_delete"`
	WallReplyRestore     bool `param:"wall_reply_restore"`
	WallPostNew          bool `param:"wall_post_new"`
	WallRepost           bool `param:"wall_repost"`
	BoardPostNew         bool `param:"board_post_new"`
	BoardPostEdit        bool `param:"board_post_edit"`
	BoardPostRestore     bool `param:"board_post_restore"`
	BoardPostDelete      bool `param:"board_post_delete"`
	PhotoCommentNew      bool `param:"photo_comment_new"`
	PhotoCommentEdit     bool `param:"photo_comment_edit"`
	PhotoCommentDelete   bool `param:"photo_comment_delete"`
	PhotoCommentRestore  bool `param:"photo_comment_restore"`
	VideoCommentNew      bool `param:"video_comment_new"`
	VideoCommentEdit     bool `param:"video_comment_edit"`
	VideoCommentDelete   bool `param:"video_comment_delete"`
	VideoCommentRestore  bool `param:"video_comment_restore"`
	MarketCommentNew     bool `param:"market_comment_new"`
	MarketCommentEdit    bool `param:"market_comment_edit"`
	MarketCommentDelete  bool `param:"market_comment_delete"`
	MarketCommentRestore bool `param:"market_comment_restore"`
	PollVoteNew          bool `param:"poll_vote_new"`
	GroupJoin            bool `param:"group_join"`
	GroupLeave           bool `param:"group_leave"`
}
