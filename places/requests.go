package places

const (
	Enum100M = 1
	Enum800M = 2
	Enum6Km  = 3
	Enum50Km = 4
)

type AddRequest struct {
	Type      int     `param:"type"`
	Title     string  `param:"title"`
	Latitude  float64 `param:"latitude"`
	Longitude float64 `param:"longitude"`
	Country   int     `param:"country"`
	City      int     `param:"city"`
	Address   string  `param:"address"`
}

type GetByIdRequest struct {
	Places string `param:"places"`
}

type SearchRequest struct {
	Q         string  `param:"q"`
	City      int     `param:"city"`
	Latitude  float64 `param:"latitude"`
	Longitude float64 `param:"longitude"`
	Radius    int     `param:"radius"`
	Offset    int     `param:"offset"`
	Count     int     `param:"count"`
}

type CheckinRequest struct {
	PlaceId     int     `param:"place_id"`
	Text        string  `param:"text"`
	Latitude    float64 `param:"latitude"`
	Longitude   float64 `param:"longitude"`
	FriendsOnly bool    `param:"friends_only"`
	Services    string  `param:"services"`
}

type GetCheckinsRequest struct {
	Latitude    float64 `param:"latitude"`
	Longitude   float64 `param:"longitude"`
	Place       int     `param:"place"`
	UserId      int     `param:"user_id"`
	Offset      int     `param:"offset"`
	Count       int     `param:"count"`
	Timestamp   int     `param:"timestamp"`
	FriendsOnly bool    `param:"friends_only"`
	NeedPlaces  bool    `param:"need_places"`
}

type GetTypesRequest struct {
}
