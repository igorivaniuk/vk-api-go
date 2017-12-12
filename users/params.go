package users

const (
	NameCaseNom = "nom"
	NameCaseGen = "gen"
	NameCaseDat = "dat"
	NameCaseAcc = "acc"
	NameCaseInt = "ins"
	NameCaseAbl = "abl"
)

const (
	FieldsAll = `photo_id,verified,sex,bdate,city,country,home_town,has_photo,photo_50,photo_100,photo_200_orig,photo_200,photo_400_orig,photo_max,photo_max_orig,online,domain,has_mobile,contacts,site,education,universities,schools,status,last_seen,followers_count,common_count,occupation,nickname,relatives,relation,personal,connections,exports,wall_comments,activities,interests,music,movies,tv,books,games,about,quotes,can_post,can_see_all_posts,can_see_audio,can_write_private_message,can_send_friend_request,is_favorite,is_hidden_from_feed,timezone,screen_name,maiden_name,crop_photo,is_friend,friend_status,career,military,blacklisted,blacklisted_by_me`
)

const (
	ReportTypePorn         = "porn"
	ReportTypeSpam         = "spam"
	ReportTypeInsult       = "insult"
	ReportTypeAdvertisment = "advertisment"
)

type GetRequest struct {
	UserIds  string `param:"user_ids"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type GetFollowersRequest struct {
	UserId   int    `param:"user_id"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type GetNearbyRequest struct {
	Latitude        float64 `param:"latitude"`
	Longitude       float64 `param:"longitude"`
	Accuracy        int     `param:"accuracy"`
	Timeout         int     `param:"timeout"`
	Radius          int     `param:"radius"`
	Fields          string  `param:"fields"`
	NameCase        string  `param:"name_case"`
	NeedDescription int     `param:"need_description"`
}

type GetSubscriptionsRequest struct {
	UserId   int    `param:"user_id"`
	Extended int    `param:"extended"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
	Fields   string `param:"fields"`
}

type IsAppUserRequest struct {
	UserId int `param:"user_id"`
}

type ReportRequest struct {
	UserId  int    `param:"user_id"`
	Type    string `param:"type"`
	Comment string `param:"comment"`
}

type SearchRequest struct {
	Q                 string `json:"q"`
	Sort              int    `json:"sort"`
	Offset            int    `param:"offset"`
	Count             int    `param:"count"`
	Fields            string `param:"fields"`
	City              int    `json:"city"`
	Country           int    `json:"country"`
	Hometown          string `json:"hometown"`
	UniversityCountry int    `json:"university_country"`
	University        int    `json:"university"`
	UniversityYear    int    `json:"university_year"`
	UniversityFaculty int    `json:"university_faculty"`
	UniversityChair   int    `json:"university_chair"`
	Sex               int    `json:"sex"`
	Status            int    `json:"status"`
	AgeFrom           int    `json:"age_from"`
	AgeTo             int    `json:"age_to"`
	BirthDay          int    `json:"birth_day"`
	BirthMonth        int    `json:"birth_month"`
	BirthYear         int    `json:"birth_year"`
	Online            bool   `json:"online"`
	HasPhoto          bool   `json:"has_photo"`
	SchoolCountry     int    `json:"school_country"`
	SchoolCity        int    `json:"school_city"`
	SchoolClass       int    `json:"school_class"`
	School            int    `json:"school"`
	SchoolYear        int    `json:"school_year"`
	Religion          string `json:"religion"`
	Interests         string `json:"interests"`
	Company           string `json:"company"`
	Position          string `json:"position"`
	GroupID           int    `json:"group_id"`
	FromList          string `json:"from_list"`
}
