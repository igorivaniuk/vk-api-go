package board

const (
	EnumUpdatedDesc          = 1
	EnumCreatedDesc          = 2
	EnumUpdatedAsc           = -1
	EnumCreatedAsc           = -2
	EnumAsByAdministrator    = 0
	EnumFirst                = 1
	EnumLast                 = 2
	EnumNone                 = 0
	EnumChronological        = "asc"
	EnumReverseChronological = "desc"
)

type GetTopicsRequest struct {
	GroupId       int    `param:"group_id"`
	TopicIds      string `param:"topic_ids"`
	Order         int    `param:"order"`
	Offset        int    `param:"offset"`
	Count         int    `param:"count"`
	Extended      bool   `param:"extended"`
	Preview       int    `param:"preview"`
	PreviewLength int    `param:"preview_length"`
}

type GetCommentsRequest struct {
	GroupId        int    `param:"group_id"`
	TopicId        int    `param:"topic_id"`
	NeedLikes      bool   `param:"need_likes"`
	StartCommentId int    `param:"start_comment_id"`
	Offset         int    `param:"offset"`
	Count          int    `param:"count"`
	Extended       bool   `param:"extended"`
	Sort           string `param:"sort"`
}

type AddTopicRequest struct {
	GroupId     int    `param:"group_id"`
	Title       string `param:"title"`
	Text        string `param:"text"`
	FromGroup   bool   `param:"from_group"`
	Attachments string `param:"attachments"`
}

type CreateCommentRequest struct {
	GroupId     int    `param:"group_id"`
	TopicId     int    `param:"topic_id"`
	Message     string `param:"message"`
	Attachments string `param:"attachments"`
	FromGroup   bool   `param:"from_group"`
	StickerId   int    `param:"sticker_id"`
	Guid        string `param:"guid"`
}

type DeleteTopicRequest struct {
	GroupId int `param:"group_id"`
	TopicId int `param:"topic_id"`
}

type EditTopicRequest struct {
	GroupId int    `param:"group_id"`
	TopicId int    `param:"topic_id"`
	Title   string `param:"title"`
}

type EditCommentRequest struct {
	GroupId     int    `param:"group_id"`
	TopicId     int    `param:"topic_id"`
	CommentId   int    `param:"comment_id"`
	Message     string `param:"message"`
	Attachments string `param:"attachments"`
}

type RestoreCommentRequest struct {
	GroupId   int `param:"group_id"`
	TopicId   int `param:"topic_id"`
	CommentId int `param:"comment_id"`
}

type DeleteCommentRequest struct {
	GroupId   int `param:"group_id"`
	TopicId   int `param:"topic_id"`
	CommentId int `param:"comment_id"`
}

type OpenTopicRequest struct {
	GroupId int `param:"group_id"`
	TopicId int `param:"topic_id"`
}

type CloseTopicRequest struct {
	GroupId int `param:"group_id"`
	TopicId int `param:"topic_id"`
}

type FixTopicRequest struct {
	GroupId int `param:"group_id"`
	TopicId int `param:"topic_id"`
}

type UnfixTopicRequest struct {
	GroupId int `param:"group_id"`
	TopicId int `param:"topic_id"`
}
