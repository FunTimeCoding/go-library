package response

type Pages struct {
	Results []*Page `json:"results"`
	Links   Links   `json:"_links"`
}
