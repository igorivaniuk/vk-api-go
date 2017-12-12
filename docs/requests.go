package docs

type GetRequest struct {
	Count   int `param:"count"`
	Offset  int `param:"offset"`
	OwnerId int `param:"owner_id"`
}

type GetByIdRequest struct {
	Docs string `param:"docs"`
}

type GetUploadServerRequest struct {
	GroupId int `param:"group_id"`
}

type GetWallUploadServerRequest struct {
	GroupId int `param:"group_id"`
}

type SaveRequest struct {
	File  string `param:"file"`
	Title string `param:"title"`
	Tags  string `param:"tags"`
}

type DeleteRequest struct {
	OwnerId int `param:"owner_id"`
	DocId   int `param:"doc_id"`
}

type AddRequest struct {
	OwnerId   int    `param:"owner_id"`
	DocId     int    `param:"doc_id"`
	AccessKey string `param:"access_key"`
}

type GetTypesRequest struct {
	OwnerId int `param:"owner_id"`
}

type SearchRequest struct {
	Q      string `param:"q"`
	Count  int    `param:"count"`
	Offset int    `param:"offset"`
}

type EditRequest struct {
	OwnerId int    `param:"owner_id"`
	DocId   int    `param:"doc_id"`
	Title   string `param:"title"`
	Tags    string `param:"tags"`
}
