package issue

type Response struct {
	Expands    []string `json:"_expands"`
	Size       int      `json:"size"`
	Start      int      `json:"start"`
	Limit      int      `json:"limit"`
	IsLastPage bool     `json:"isLastPage"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
	} `json:"_links"`
	Values []Values `json:"values"`
}
