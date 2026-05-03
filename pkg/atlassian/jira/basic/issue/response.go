package issue

type Response struct {
	Expands    []string        `json:"_expands"`
	Size       int             `json:"size"`
	Start      int             `json:"start"`
	Limit      int             `json:"limit"`
	IsLastPage bool            `json:"isLastPage"`
	Links      PaginationLinks `json:"_links"`
	Values     []Values        `json:"values"`
}
