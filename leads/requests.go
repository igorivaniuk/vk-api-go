package leads

const (
	EnumStart            = 0
	EnumFinish           = 1
	EnumBlockingUsers    = 2
	EnumStartInTestMode  = 3
	EnumFinishInTestMode = 4
)

type CompleteRequest struct {
	VkSid   string `param:"vk_sid"`
	Secret  string `param:"secret"`
	Comment string `param:"comment"`
}

type StartRequest struct {
	LeadId int    `param:"lead_id"`
	Secret string `param:"secret"`
}

type GetStatsRequest struct {
	LeadId    int    `param:"lead_id"`
	Secret    string `param:"secret"`
	DateStart string `param:"date_start"`
	DateEnd   string `param:"date_end"`
}

type GetUsersRequest struct {
	OfferId int    `param:"offer_id"`
	Secret  string `param:"secret"`
	Offset  int    `param:"offset"`
	Count   int    `param:"count"`
	Status  int    `param:"status"`
	Reverse bool   `param:"reverse"`
}

type CheckUserRequest struct {
	LeadId     int    `param:"lead_id"`
	TestResult int    `param:"test_result"`
	Age        int    `param:"age"`
	Country    string `param:"country"`
}

type MetricHitRequest struct {
	Data string `param:"data"`
}
