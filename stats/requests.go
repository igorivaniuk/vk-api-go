package stats

type GetRequest struct {
	GroupId  int    `param:"group_id"`
	AppId    int    `param:"app_id"`
	DateFrom string `param:"date_from"`
	DateTo   string `param:"date_to"`
}

type TrackVisitorRequest struct {
}

type GetPostReachRequest struct {
	OwnerId int `param:"owner_id"`
	PostId  int `param:"post_id"`
}
