package video

const (
	EnumDuration           = 1
	EnumRelevance          = 2
	EnumDateAdded          = 0
	EnumOldestCommentFirst = "asc"
	EnumNewestCommentFirst = "desc"
	EnumSpam               = 0
	EnumChildPornography   = 1
	EnumExtremism          = 2
	EnumViolence           = 3
	EnumDrugPropaganda     = 4
	EnumAdultMaterial      = 5
	EnumInsultAbuse        = 6
)

type GetRequest struct {
	OwnerId  int    `param:"owner_id"`
	Videos   string `param:"videos"`
	AlbumId  int    `param:"album_id"`
	Count    int    `param:"count"`
	Offset   int    `param:"offset"`
	Extended bool   `param:"extended"`
}

type EditRequest struct {
	OwnerId        int    `param:"owner_id"`
	VideoId        int    `param:"video_id"`
	Name           string `param:"name"`
	Desc           string `param:"desc"`
	PrivacyView    string `param:"privacy_view"`
	PrivacyComment string `param:"privacy_comment"`
	NoComments     bool   `param:"no_comments"`
	Repeat         bool   `param:"repeat"`
}

type AddRequest struct {
	TargetId int `param:"target_id"`
	VideoId  int `param:"video_id"`
	OwnerId  int `param:"owner_id"`
}

type SaveRequest struct {
	Name           string `param:"name"`
	Description    string `param:"description"`
	IsPrivate      bool   `param:"is_private"`
	Wallpost       bool   `param:"wallpost"`
	Link           string `param:"link"`
	GroupId        int    `param:"group_id"`
	AlbumId        int    `param:"album_id"`
	PrivacyView    string `param:"privacy_view"`
	PrivacyComment string `param:"privacy_comment"`
	NoComments     bool   `param:"no_comments"`
	Repeat         bool   `param:"repeat"`
}

type DeleteRequest struct {
	VideoId  int `param:"video_id"`
	OwnerId  int `param:"owner_id"`
	TargetId int `param:"target_id"`
}

type RestoreRequest struct {
	VideoId int `param:"video_id"`
	OwnerId int `param:"owner_id"`
}

type SearchRequest struct {
	Q         string `param:"q"`
	Sort      int    `param:"sort"`
	Hd        int    `param:"hd"`
	Adult     bool   `param:"adult"`
	Filters   string `param:"filters"`
	SearchOwn bool   `param:"search_own"`
	Offset    int    `param:"offset"`
	Longer    int    `param:"longer"`
	Shorter   int    `param:"shorter"`
	Count     int    `param:"count"`
	Extended  bool   `param:"extended"`
}

type GetUserVideosRequest struct {
	UserId   int  `param:"user_id"`
	Offset   int  `param:"offset"`
	Count    int  `param:"count"`
	Extended bool `param:"extended"`
}

type GetAlbumsRequest struct {
	OwnerId  int  `param:"owner_id"`
	Offset   int  `param:"offset"`
	Count    int  `param:"count"`
	Extended bool `param:"extended"`
}

type GetAlbumByIdRequest struct {
	OwnerId int `param:"owner_id"`
	AlbumId int `param:"album_id"`
}

type AddAlbumRequest struct {
	GroupId int    `param:"group_id"`
	Title   string `param:"title"`
	Privacy string `param:"privacy"`
}

type EditAlbumRequest struct {
	GroupId int    `param:"group_id"`
	AlbumId int    `param:"album_id"`
	Title   string `param:"title"`
	Privacy string `param:"privacy"`
}

type DeleteAlbumRequest struct {
	GroupId int `param:"group_id"`
	AlbumId int `param:"album_id"`
}

type ReorderAlbumsRequest struct {
	OwnerId int `param:"owner_id"`
	AlbumId int `param:"album_id"`
	Before  int `param:"before"`
	After   int `param:"after"`
}

type ReorderVideosRequest struct {
	TargetId      int `param:"target_id"`
	AlbumId       int `param:"album_id"`
	OwnerId       int `param:"owner_id"`
	VideoId       int `param:"video_id"`
	BeforeOwnerId int `param:"before_owner_id"`
	BeforeVideoId int `param:"before_video_id"`
	AfterOwnerId  int `param:"after_owner_id"`
	AfterVideoId  int `param:"after_video_id"`
}

type AddToAlbumRequest struct {
	TargetId int    `param:"target_id"`
	AlbumId  int    `param:"album_id"`
	AlbumIds string `param:"album_ids"`
	OwnerId  int    `param:"owner_id"`
	VideoId  int    `param:"video_id"`
}

type RemoveFromAlbumRequest struct {
	TargetId int    `param:"target_id"`
	AlbumId  int    `param:"album_id"`
	AlbumIds string `param:"album_ids"`
	OwnerId  int    `param:"owner_id"`
	VideoId  int    `param:"video_id"`
}

type GetAlbumsByVideoRequest struct {
	TargetId int  `param:"target_id"`
	OwnerId  int  `param:"owner_id"`
	VideoId  int  `param:"video_id"`
	Extended bool `param:"extended"`
}

type GetCommentsRequest struct {
	OwnerId        int    `param:"owner_id"`
	VideoId        int    `param:"video_id"`
	NeedLikes      bool   `param:"need_likes"`
	StartCommentId int    `param:"start_comment_id"`
	Offset         int    `param:"offset"`
	Count          int    `param:"count"`
	Sort           string `param:"sort"`
	Extended       bool   `param:"extended"`
}

type CreateCommentRequest struct {
	OwnerId        int    `param:"owner_id"`
	VideoId        int    `param:"video_id"`
	Message        string `param:"message"`
	Attachments    string `param:"attachments"`
	FromGroup      bool   `param:"from_group"`
	ReplyToComment int    `param:"reply_to_comment"`
	StickerId      int    `param:"sticker_id"`
	Guid           string `param:"guid"`
}

type DeleteCommentRequest struct {
	OwnerId   int `param:"owner_id"`
	CommentId int `param:"comment_id"`
}

type RestoreCommentRequest struct {
	OwnerId   int `param:"owner_id"`
	CommentId int `param:"comment_id"`
}

type EditCommentRequest struct {
	OwnerId     int    `param:"owner_id"`
	CommentId   int    `param:"comment_id"`
	Message     string `param:"message"`
	Attachments string `param:"attachments"`
}

type GetTagsRequest struct {
	OwnerId int `param:"owner_id"`
	VideoId int `param:"video_id"`
}

type PutTagRequest struct {
	UserId     int    `param:"user_id"`
	OwnerId    int    `param:"owner_id"`
	VideoId    int    `param:"video_id"`
	TaggedName string `param:"tagged_name"`
}

type RemoveTagRequest struct {
	TagId   int `param:"tag_id"`
	OwnerId int `param:"owner_id"`
	VideoId int `param:"video_id"`
}

type GetNewTagsRequest struct {
	Offset int `param:"offset"`
	Count  int `param:"count"`
}

type ReportRequest struct {
	OwnerId     int    `param:"owner_id"`
	VideoId     int    `param:"video_id"`
	Reason      int    `param:"reason"`
	Comment     string `param:"comment"`
	SearchQuery string `param:"search_query"`
}

type ReportCommentRequest struct {
	OwnerId   int `param:"owner_id"`
	CommentId int `param:"comment_id"`
	Reason    int `param:"reason"`
}

type GetCatalogRequest struct {
	Count      int    `param:"count"`
	ItemsCount int    `param:"items_count"`
	From       string `param:"from"`
	Filters    string `param:"filters"`
	Extended   bool   `param:"extended"`
}

type GetCatalogSectionRequest struct {
	SectionId string `param:"section_id"`
	From      string `param:"from"`
	Count     int    `param:"count"`
	Extended  bool   `param:"extended"`
}

type HideCatalogSectionRequest struct {
	SectionId int `param:"section_id"`
}
