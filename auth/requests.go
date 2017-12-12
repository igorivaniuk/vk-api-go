package auth

const (
	EnumUndefined = 0
	EnumFemale    = 1
	EnumMale      = 2
)

type CheckPhoneRequest struct {
	Phone        string `param:"phone"`
	ClientId     int    `param:"client_id"`
	ClientSecret string `param:"client_secret"`
	AuthByPhone  bool   `param:"auth_by_phone"`
}

type SignupRequest struct {
	FirstName    string `param:"first_name"`
	LastName     string `param:"last_name"`
	Birthday     string `param:"birthday"`
	ClientId     int    `param:"client_id"`
	ClientSecret string `param:"client_secret"`
	Phone        string `param:"phone"`
	Password     string `param:"password"`
	TestMode     bool   `param:"test_mode"`
	Voice        bool   `param:"voice"`
	Sex          int    `param:"sex"`
	Sid          string `param:"sid"`
}

type ConfirmRequest struct {
	ClientId     int    `param:"client_id"`
	ClientSecret string `param:"client_secret"`
	Phone        string `param:"phone"`
	Code         string `param:"code"`
	Password     string `param:"password"`
	TestMode     bool   `param:"test_mode"`
	Intro        int    `param:"intro"`
}

type RestoreRequest struct {
	Phone    string `param:"phone"`
	LastName string `param:"last_name"`
}
