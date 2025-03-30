package response

type Spaces struct {
	Results []*Space `json:"results"`
	Links   Links    `json:"_links"`
}
