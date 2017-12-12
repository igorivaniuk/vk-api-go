package search

type GetHintsRequest struct {
	Q            string `param:"q"`
	Limit        int    `param:"limit"`
	Filters      string `param:"filters"`
	SearchGlobal bool   `param:"search_global"`
}
