package wall

const (
	EnumChronological        = "asc"
	EnumReverseChronological = "desc"
	EnumSpam                 = 0
	EnumChildPornography     = 1
	EnumExtremism            = 2
	EnumViolence             = 3
	EnumDrugPropaganda       = 4
	EnumAdultMaterial        = 5
	EnumInsultAbuse          = 6
)

type GetByIdRequest struct {
	Posts            string `param:"posts"`
	Extended         bool   `param:"extended"`
	CopyHistoryDepth int    `param:"copy_history_depth"`
	Fields           string `param:"fields"`
}

type PostRequest struct {
	OwnerId     int     `param:"owner_id"`
	FriendsOnly bool    `param:"friends_only"`
	FromGroup   bool    `param:"from_group"`
	Message     string  `param:"message"`
	Attachments string  `param:"attachments"`
	Services    string  `param:"services"`
	Signed      bool    `param:"signed"`
	PublishDate int     `param:"publish_date"`
	Lat         float64 `param:"lat"`
	Long        float64 `param:"long"`
	PlaceId     int     `param:"place_id"`
	PostId      int     `param:"post_id"`
	Guid        string  `param:"guid"`
	MarkAsAds   bool    `param:"mark_as_ads"`
}

type RepostRequest struct {
	Object    string `param:"object"`
	Message   string `param:"message"`
	GroupId   int    `param:"group_id"`
	MarkAsAds bool   `param:"mark_as_ads"`
}

type GetRepostsRequest struct {
	OwnerId int `param:"owner_id"`
	PostId  int `param:"post_id"`
	Offset  int `param:"offset"`
	Count   int `param:"count"`
}

type EditRequest struct {
	OwnerId     int     `param:"owner_id"`
	PostId      int     `param:"post_id"`
	FriendsOnly bool    `param:"friends_only"`
	Message     string  `param:"message"`
	Attachments string  `param:"attachments"`
	Services    string  `param:"services"`
	Signed      bool    `param:"signed"`
	PublishDate int     `param:"publish_date"`
	Lat         float64 `param:"lat"`
	Long        float64 `param:"long"`
	PlaceId     int     `param:"place_id"`
	MarkAsAds   bool    `param:"mark_as_ads"`
}

type DeleteRequest struct {
	OwnerId int `param:"owner_id"`
	PostId  int `param:"post_id"`
}

type RestoreRequest struct {
	OwnerId int `param:"owner_id"`
	PostId  int `param:"post_id"`
}

type PinRequest struct {
	OwnerId int `param:"owner_id"`
	PostId  int `param:"post_id"`
}

type UnpinRequest struct {
	OwnerId int `param:"owner_id"`
	PostId  int `param:"post_id"`
}

type GetCommentsRequest struct {
	OwnerId        int    `param:"owner_id"`
	PostId         int    `param:"post_id"`
	NeedLikes      bool   `param:"need_likes"`
	StartCommentId int    `param:"start_comment_id"`
	Offset         int    `param:"offset"`
	Count          int    `param:"count"`
	Sort           string `param:"sort"`
	PreviewLength  int    `param:"preview_length"`
	Extended       bool   `param:"extended"`
}

type CreateCommentRequest struct {
	OwnerId        int    `param:"owner_id"`
	PostId         int    `param:"post_id"`
	FromGroup      int    `param:"from_group"`
	Message        string `param:"message"`
	ReplyToComment int    `param:"reply_to_comment"`
	Attachments    string `param:"attachments"`
	StickerId      int    `param:"sticker_id"`
	Guid           string `param:"guid"`
}

type EditCommentRequest struct {
	OwnerId     int    `param:"owner_id"`
	CommentId   int    `param:"comment_id"`
	Message     string `param:"message"`
	Attachments string `param:"attachments"`
}

type DeleteCommentRequest struct {
	OwnerId   int `param:"owner_id"`
	CommentId int `param:"comment_id"`
}

type RestoreCommentRequest struct {
	OwnerId   int `param:"owner_id"`
	CommentId int `param:"comment_id"`
}

type ReportPostRequest struct {
	OwnerId int `param:"owner_id"`
	PostId  int `param:"post_id"`
	Reason  int `param:"reason"`
}

type ReportCommentRequest struct {
	OwnerId   int `param:"owner_id"`
	CommentId int `param:"comment_id"`
	Reason    int `param:"reason"`
}
