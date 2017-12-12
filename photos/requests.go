package photos

const (
	EnumOldFirst = "asc"
	EnumNewFirst = "desc"
)

type SaveRequest struct {
	AlbumId    int     `param:"album_id"`
	GroupId    int     `param:"group_id"`
	Server     int     `param:"server"`
	PhotosList string  `param:"photos_list"`
	Hash       string  `param:"hash"`
	Latitude   float64 `param:"latitude"`
	Longitude  float64 `param:"longitude"`
	Caption    string  `param:"caption"`
}

type CopyRequest struct {
	OwnerId   int    `param:"owner_id"`
	PhotoId   int    `param:"photo_id"`
	AccessKey string `param:"access_key"`
}

type EditRequest struct {
	OwnerId      int     `param:"owner_id"`
	PhotoId      int     `param:"photo_id"`
	Caption      string  `param:"caption"`
	Latitude     float64 `param:"latitude"`
	Longitude    float64 `param:"longitude"`
	PlaceStr     string  `param:"place_str"`
	FoursquareId string  `param:"foursquare_id"`
	DeletePlace  bool    `param:"delete_place"`
}

type MoveRequest struct {
	OwnerId       int `param:"owner_id"`
	TargetAlbumId int `param:"target_album_id"`
	PhotoId       int `param:"photo_id"`
}

type MakeCoverRequest struct {
	OwnerId int `param:"owner_id"`
	PhotoId int `param:"photo_id"`
	AlbumId int `param:"album_id"`
}

type ReorderAlbumsRequest struct {
	OwnerId int `param:"owner_id"`
	AlbumId int `param:"album_id"`
	Before  int `param:"before"`
	After   int `param:"after"`
}

type ReorderPhotosRequest struct {
	OwnerId int `param:"owner_id"`
	PhotoId int `param:"photo_id"`
	Before  int `param:"before"`
	After   int `param:"after"`
}

type GetAllRequest struct {
	OwnerId         int  `param:"owner_id"`
	Extended        bool `param:"extended"`
	Offset          int  `param:"offset"`
	Count           int  `param:"count"`
	PhotoSizes      bool `param:"photo_sizes"`
	NoServiceAlbums bool `param:"no_service_albums"`
	NeedHidden      bool `param:"need_hidden"`
	SkipHidden      bool `param:"skip_hidden"`
}

type GetUserPhotosRequest struct {
	UserId   int    `param:"user_id"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
	Extended bool   `param:"extended"`
	Sort     string `param:"sort"`
}

type DeleteAlbumRequest struct {
	AlbumId int `param:"album_id"`
	GroupId int `param:"group_id"`
}

type DeleteRequest struct {
	OwnerId int `param:"owner_id"`
	PhotoId int `param:"photo_id"`
}

type RestoreRequest struct {
	OwnerId int `param:"owner_id"`
	PhotoId int `param:"photo_id"`
}

type ConfirmTagRequest struct {
	OwnerId int    `param:"owner_id"`
	PhotoId string `param:"photo_id"`
	TagId   int    `param:"tag_id"`
}

type GetCommentsRequest struct {
	OwnerId        int    `param:"owner_id"`
	PhotoId        int    `param:"photo_id"`
	NeedLikes      bool   `param:"need_likes"`
	StartCommentId int    `param:"start_comment_id"`
	Offset         int    `param:"offset"`
	Count          int    `param:"count"`
	Sort           string `param:"sort"`
	AccessKey      string `param:"access_key"`
	Extended       bool   `param:"extended"`
	Fields         string `param:"fields"`
}

type GetAllCommentsRequest struct {
	OwnerId   int  `param:"owner_id"`
	AlbumId   int  `param:"album_id"`
	NeedLikes bool `param:"need_likes"`
	Offset    int  `param:"offset"`
	Count     int  `param:"count"`
}

type CreateCommentRequest struct {
	OwnerId        int    `param:"owner_id"`
	PhotoId        int    `param:"photo_id"`
	Message        string `param:"message"`
	Attachments    string `param:"attachments"`
	FromGroup      bool   `param:"from_group"`
	ReplyToComment int    `param:"reply_to_comment"`
	StickerId      int    `param:"sticker_id"`
	AccessKey      string `param:"access_key"`
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
	OwnerId   int    `param:"owner_id"`
	PhotoId   int    `param:"photo_id"`
	AccessKey string `param:"access_key"`
}

type PutTagRequest struct {
	OwnerId int     `param:"owner_id"`
	PhotoId int     `param:"photo_id"`
	UserId  int     `param:"user_id"`
	X       float64 `param:"x"`
	Y       float64 `param:"y"`
	X2      float64 `param:"x2"`
	Y2      float64 `param:"y2"`
}

type RemoveTagRequest struct {
	OwnerId int `param:"owner_id"`
	PhotoId int `param:"photo_id"`
	TagId   int `param:"tag_id"`
}

type GetNewTagsRequest struct {
	Offset int `param:"offset"`
	Count  int `param:"count"`
}
