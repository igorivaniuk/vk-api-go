package secure

type CheckTokenRequest struct {
	Token string `param:"token"`
	Ip    string `param:"ip"`
}
