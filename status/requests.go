package status

type GetRequest struct {
	UserId  int `param:"user_id"`
	GroupId int `param:"group_id"`
}

type SetRequest struct {
	Text    string `param:"text"`
	GroupId int    `param:"group_id"`
}
