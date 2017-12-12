package fave

type GetUsersRequest struct {
	Offset int `param:"offset"`
	Count  int `param:"count"`
}

type GetPhotosRequest struct {
	Offset     int  `param:"offset"`
	Count      int  `param:"count"`
	PhotoSizes bool `param:"photo_sizes"`
}

type GetPostsRequest struct {
	Offset   int  `param:"offset"`
	Count    int  `param:"count"`
	Extended bool `param:"extended"`
}

type GetVideosRequest struct {
	Offset   int  `param:"offset"`
	Count    int  `param:"count"`
	Extended bool `param:"extended"`
}

type GetLinksRequest struct {
	Offset int `param:"offset"`
	Count  int `param:"count"`
}

type GetMarketItemsRequest struct {
	Count    int  `param:"count"`
	Extended bool `param:"extended"`
}

type AddUserRequest struct {
	UserId int `param:"user_id"`
}

type RemoveUserRequest struct {
	UserId int `param:"user_id"`
}

type AddGroupRequest struct {
	GroupId int `param:"group_id"`
}

type RemoveGroupRequest struct {
	GroupId int `param:"group_id"`
}

type AddLinkRequest struct {
	Link string `param:"link"`
	Text string `param:"text"`
}

type RemoveLinkRequest struct {
	LinkId string `param:"link_id"`
}
