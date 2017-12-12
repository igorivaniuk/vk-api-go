package orders

const (
	EnumCancel = "cancel"
	EnumCharge = "charge"
	EnumRefund = "refund"
)

type GetRequest struct {
	Count    int  `param:"count"`
	TestMode bool `param:"test_mode"`
}

type GetByIdRequest struct {
	OrderId  int    `param:"order_id"`
	OrderIds string `param:"order_ids"`
	TestMode bool   `param:"test_mode"`
}

type ChangeStateRequest struct {
	OrderId    int    `param:"order_id"`
	Action     string `param:"action"`
	AppOrderId int    `param:"app_order_id"`
	TestMode   bool   `param:"test_mode"`
}

type GetAmountRequest struct {
	UserId int    `param:"user_id"`
	Votes  string `param:"votes"`
}
