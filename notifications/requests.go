package notifications

type GetRequest struct {
	Count     int    `param:"count"`
	StartFrom string `param:"start_from"`
	Filters   string `param:"filters"`
	StartTime int    `param:"start_time"`
	EndTime   int    `param:"end_time"`
}

type MarkAsViewedRequest struct {
}
