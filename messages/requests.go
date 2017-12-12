package messages

const (
	EnumChronological        = 1
	EnumReverseChronological = 0
	EnumPhoto                = "photo"
	EnumVideo                = "video"
	EnumDoc                  = "doc"
	EnumAudio                = "audio"
	EnumLink                 = "link"
	EnumMarket               = "market"
	EnumWall                 = "wall"
	EnumShare                = "share"
	EnumNominative           = "nom"
	EnumGenitive             = "gen"
	EnumDative               = "dat"
	EnumAccusative           = "acc"
	EnumInstrumental         = "ins"
	EnumPrepositional        = "abl"
)

type GetRequest struct {
	Out           bool `param:"out"`
	Offset        int  `param:"offset"`
	Count         int  `param:"count"`
	Filter        int  `param:"filter"`
	TimeOffset    int  `param:"time_offset"`
	PreviewLength int  `param:"preview_length"`
	LastMessageId int  `param:"last_message_id"`
}

type GetDialogsRequest struct {
	Offset         int  `param:"offset"`
	Count          int  `param:"count"`
	StartMessageId int  `param:"start_message_id"`
	PreviewLength  int  `param:"preview_length"`
	Unread         bool `param:"unread"`
	Important      bool `param:"important"`
	Unanswered     bool `param:"unanswered"`
}

type GetByIdRequest struct {
	MessageIds string `param:"message_ids"`
}

type SearchRequest struct {
	Q             string `param:"q"`
	PeerId        int    `param:"peer_id"`
	Date          int    `param:"date"`
	PreviewLength int    `param:"preview_length"`
	Offset        int    `param:"offset"`
	Count         int    `param:"count"`
}

type GetHistoryRequest struct {
	Offset         int `param:"offset"`
	Count          int `param:"count"`
	UserId         int `param:"user_id"`
	PeerId         int `param:"peer_id"`
	StartMessageId int `param:"start_message_id"`
	Rev            int `param:"rev"`
}

type GetHistoryAttachmentsRequest struct {
	PeerId     int    `param:"peer_id"`
	MediaType  string `param:"media_type"`
	StartFrom  string `param:"start_from"`
	Count      int    `param:"count"`
	PhotoSizes bool   `param:"photo_sizes"`
	Fields     string `param:"fields"`
}

type SendRequest struct {
	UserId          int     `param:"user_id"`
	RandomId        int     `param:"random_id"`
	PeerId          int     `param:"peer_id"`
	Domain          string  `param:"domain"`
	ChatId          int     `param:"chat_id"`
	UserIds         string  `param:"user_ids"`
	Message         string  `param:"message"`
	Lat             float64 `param:"lat"`
	Long            float64 `param:"long"`
	Attachment      string  `param:"attachment"`
	ForwardMessages string  `param:"forward_messages"`
	StickerId       int     `param:"sticker_id"`
	Notification    bool    `param:"notification"`
}

type DeleteRequest struct {
	MessageIds string `param:"message_ids"`
	Spam       bool   `param:"spam"`
}

type DeleteDialogRequest struct {
	UserId string `param:"user_id"`
	PeerId int    `param:"peer_id"`
	Offset int    `param:"offset"`
	Count  int    `param:"count"`
}

type RestoreRequest struct {
	MessageId int `param:"message_id"`
}

type MarkAsReadRequest struct {
	MessageIds     string `param:"message_ids"`
	PeerId         string `param:"peer_id"`
	StartMessageId int    `param:"start_message_id"`
}

type MarkAsImportantRequest struct {
	MessageIds string `param:"message_ids"`
	Important  int    `param:"important"`
}

type MarkAsImportantDialogRequest struct {
	PeerId    string `param:"peer_id"`
	Important int    `param:"important"`
}

type MarkAsUnansweredDialogRequest struct {
	PeerId    string `param:"peer_id"`
	Important int    `param:"important"`
}

type GetLongPollServerRequest struct {
	LpVersion int  `param:"lp_version"`
	NeedPts   bool `param:"need_pts"`
}

type GetLongPollHistoryRequest struct {
	Ts            int    `param:"ts"`
	Pts           int    `param:"pts"`
	PreviewLength int    `param:"preview_length"`
	Onlines       bool   `param:"onlines"`
	Fields        string `param:"fields"`
	EventsLimit   int    `param:"events_limit"`
	MsgsLimit     int    `param:"msgs_limit"`
	MaxMsgId      int    `param:"max_msg_id"`
}

type GetChatRequest struct {
	ChatId   int    `param:"chat_id"`
	ChatIds  string `param:"chat_ids"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type CreateChatRequest struct {
	UserIds string `param:"user_ids"`
	Title   string `param:"title"`
}

type EditChatRequest struct {
	ChatId int    `param:"chat_id"`
	Title  string `param:"title"`
}

type GetChatUsersRequest struct {
	ChatId   int    `param:"chat_id"`
	ChatIds  string `param:"chat_ids"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type SetActivityRequest struct {
	UserId string `param:"user_id"`
	Type   string `param:"type"`
	PeerId int    `param:"peer_id"`
}

type SearchDialogsRequest struct {
	Q      string `param:"q"`
	Limit  int    `param:"limit"`
	Fields string `param:"fields"`
}

type AddChatUserRequest struct {
	ChatId int `param:"chat_id"`
	UserId int `param:"user_id"`
}

type RemoveChatUserRequest struct {
	ChatId int    `param:"chat_id"`
	UserId string `param:"user_id"`
}

type GetLastActivityRequest struct {
	UserId int `param:"user_id"`
}

type SetChatPhotoRequest struct {
	File string `param:"file"`
}

type DeleteChatPhotoRequest struct {
	ChatId int `param:"chat_id"`
}

type DenyMessagesFromGroupRequest struct {
	GroupId int `param:"group_id"`
}

type AllowMessagesFromGroupRequest struct {
	GroupId int `param:"group_id"`
}

type IsMessagesFromGroupAllowedRequest struct {
	GroupId int `param:"group_id"`
	UserId  int `param:"user_id"`
}
