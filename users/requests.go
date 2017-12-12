package users

const (
	EnumNominative        = "nom"
	EnumGenitive          = "gen"
	EnumDative            = "dat"
	EnumAccusative        = "acc"
	EnumInstrumental      = "ins"
	EnumPrepositional     = "abl"
	EnumByRating          = 0
	EnumByDateRegistered  = 1
	EnumAny               = 0
	EnumFemale            = 1
	EnumMale              = 2
	EnumNotSpecified      = 0
	EnumNotMarried        = 1
	EnumRelationship      = 2
	EnumEngaged           = 3
	EnumMarried           = 4
	EnumComplicated       = 5
	EnumActivelySearching = 6
	EnumInLove            = 7
	EnumPorn              = "porn"
	EnumSpam              = "spam"
	EnumInsult            = "insult"
	EnumAdvertisment      = "advertisment"
	Enum300M              = 1
	Enum2400M             = 2
	Enum18Km              = 3
	Enum150Km             = 4
)

type GetRequest struct {
	UserIds  string `param:"user_ids"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type SearchRequest struct {
	Q                 string `param:"q"`
	Sort              int    `param:"sort"`
	Offset            int    `param:"offset"`
	Count             int    `param:"count"`
	Fields            string `param:"fields"`
	City              int    `param:"city"`
	Country           int    `param:"country"`
	Hometown          string `param:"hometown"`
	UniversityCountry int    `param:"university_country"`
	University        int    `param:"university"`
	UniversityYear    int    `param:"university_year"`
	UniversityFaculty int    `param:"university_faculty"`
	UniversityChair   int    `param:"university_chair"`
	Sex               int    `param:"sex"`
	Status            int    `param:"status"`
	AgeFrom           int    `param:"age_from"`
	AgeTo             int    `param:"age_to"`
	BirthDay          int    `param:"birth_day"`
	BirthMonth        int    `param:"birth_month"`
	BirthYear         int    `param:"birth_year"`
	Online            bool   `param:"online"`
	HasPhoto          bool   `param:"has_photo"`
	SchoolCountry     int    `param:"school_country"`
	SchoolCity        int    `param:"school_city"`
	SchoolClass       int    `param:"school_class"`
	School            int    `param:"school"`
	SchoolYear        int    `param:"school_year"`
	Religion          string `param:"religion"`
	Interests         string `param:"interests"`
	Company           string `param:"company"`
	Position          string `param:"position"`
	GroupId           int    `param:"group_id"`
	FromList          string `param:"from_list"`
}

type IsAppUserRequest struct {
	UserId int `param:"user_id"`
}

type GetSubscriptionsRequest struct {
	UserId   int    `param:"user_id"`
	Extended bool   `param:"extended"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
	Fields   string `param:"fields"`
}

type GetFollowersRequest struct {
	UserId   int    `param:"user_id"`
	Offset   int    `param:"offset"`
	Count    int    `param:"count"`
	Fields   string `param:"fields"`
	NameCase string `param:"name_case"`
}

type ReportRequest struct {
	UserId  int    `param:"user_id"`
	Type    string `param:"type"`
	Comment string `param:"comment"`
}

type GetNearbyRequest struct {
	Latitude  float64 `param:"latitude"`
	Longitude float64 `param:"longitude"`
	Accuracy  int     `param:"accuracy"`
	Timeout   int     `param:"timeout"`
	Radius    int     `param:"radius"`
	Fields    string  `param:"fields"`
	NameCase  string  `param:"name_case"`
}
