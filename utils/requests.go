package utils

const (
	EnumHour    = "hour"
	EnumDay     = "day"
	EnumWeek    = "week"
	EnumMonth   = "month"
	EnumForever = "forever"
)

type CheckLinkRequest struct {
	Url string `param:"url"`
}

type DeleteFromLastShortenedRequest struct {
	Key string `param:"key"`
}

type GetLastShortenedLinksRequest struct {
	Count  int `param:"count"`
	Offset int `param:"offset"`
}

type GetLinkStatsRequest struct {
	Key            string `param:"key"`
	AccessKey      string `param:"access_key"`
	Interval       string `param:"interval"`
	IntervalsCount int    `param:"intervals_count"`
	Extended       bool   `param:"extended"`
}

type GetShortLinkRequest struct {
	Url     string `param:"url"`
	Private bool   `param:"private"`
}

type ResolveScreenNameRequest struct {
	ScreenName string `param:"screen_name"`
}

type GetServerTimeRequest struct {
}
