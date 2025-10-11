package response

type Labels struct {
	Results []*LabelResult `json:"results"`
	Links   Links          `json:"_links"`
}
