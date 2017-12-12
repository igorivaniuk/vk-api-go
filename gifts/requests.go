package gifts

type GetRequest struct {
	UserId int `param:"user_id"`
	Count  int `param:"count"`
	Offset int `param:"offset"`
}
