package response

type Search struct {
	Results []*Result `json:"results"`
	Start   int       `json:"start"`
	Limit   int       `json:"limit"`
	Size    int       `json:"size"`
	Links   Links     `json:"_links"`
}
