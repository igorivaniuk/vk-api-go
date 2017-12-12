package likes

const (
	EnumPost         = "post"
	EnumComment      = "comment"
	EnumPhoto        = "photo"
	EnumAudio        = "audio"
	EnumVideo        = "video"
	EnumNote         = "note"
	EnumPhotoComment = "photo_comment"
	EnumVideoComment = "video_comment"
	EnumTopicComment = "topic_comment"
	EnumSitepage     = "sitepage"
	EnumLikes        = "likes"
	EnumCopies       = "copies"
)

type GetListRequest struct {
	Type        string `param:"type"`
	OwnerId     int    `param:"owner_id"`
	ItemId      int    `param:"item_id"`
	PageUrl     string `param:"page_url"`
	Filter      string `param:"filter"`
	FriendsOnly bool   `param:"friends_only"`
	Extended    bool   `param:"extended"`
	Offset      int    `param:"offset"`
	Count       int    `param:"count"`
	SkipOwn     bool   `param:"skip_own"`
}

type AddRequest struct {
	Type      string `param:"type"`
	OwnerId   int    `param:"owner_id"`
	ItemId    int    `param:"item_id"`
	AccessKey string `param:"access_key"`
}

type DeleteRequest struct {
	Type    string `param:"type"`
	OwnerId int    `param:"owner_id"`
	ItemId  int    `param:"item_id"`
}

type IsLikedRequest struct {
	UserId  int    `param:"user_id"`
	Type    string `param:"type"`
	OwnerId int    `param:"owner_id"`
	ItemId  int    `param:"item_id"`
}
