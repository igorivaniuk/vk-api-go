package notes

type GetRequest struct {
	NoteIds string `param:"note_ids"`
	UserId  int    `param:"user_id"`
	Count   int    `param:"count"`
}

type GetByIdRequest struct {
	NoteId  int `param:"note_id"`
	OwnerId int `param:"owner_id"`
}

type AddRequest struct {
	Title          string `param:"title"`
	Text           string `param:"text"`
	PrivacyView    string `param:"privacy_view"`
	PrivacyComment string `param:"privacy_comment"`
}

type EditRequest struct {
	NoteId         int    `param:"note_id"`
	Title          string `param:"title"`
	Text           string `param:"text"`
	PrivacyView    string `param:"privacy_view"`
	PrivacyComment string `param:"privacy_comment"`
}

type DeleteRequest struct {
	NoteId int `param:"note_id"`
}

type GetCommentsRequest struct {
	NoteId  int `param:"note_id"`
	OwnerId int `param:"owner_id"`
	Count   int `param:"count"`
}

type CreateCommentRequest struct {
	NoteId  int    `param:"note_id"`
	OwnerId int    `param:"owner_id"`
	ReplyTo int    `param:"reply_to"`
	Message string `param:"message"`
	Guid    string `param:"guid"`
}

type EditCommentRequest struct {
	CommentId int    `param:"comment_id"`
	OwnerId   int    `param:"owner_id"`
	Message   string `param:"message"`
}

type DeleteCommentRequest struct {
	CommentId int `param:"comment_id"`
	OwnerId   int `param:"owner_id"`
}

type RestoreCommentRequest struct {
	CommentId int `param:"comment_id"`
	OwnerId   int `param:"owner_id"`
}
