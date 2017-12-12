package storage

type GetRequest struct {
	Key    string `param:"key"`
	Keys   string `param:"keys"`
	UserId int    `param:"user_id"`
}

type SetRequest struct {
	Key    string `param:"key"`
	Value  string `param:"value"`
	UserId int    `param:"user_id"`
}

type GetKeysRequest struct {
	UserId int `param:"user_id"`
	Count  int `param:"count"`
}
