package widgets

type GetCommentsRequest struct {
	WidgetApiId int    `param:"widget_api_id"`
	Url         string `param:"url"`
	PageId      string `param:"page_id"`
	Order       string `param:"order"`
	Fields      string `param:"fields"`
	Count       int    `param:"count"`
}

type GetPagesRequest struct {
	WidgetApiId int    `param:"widget_api_id"`
	Order       string `param:"order"`
	Period      string `param:"period"`
	Count       int    `param:"count"`
}
