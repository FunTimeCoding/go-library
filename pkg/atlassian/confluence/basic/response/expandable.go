package response

type Expandable struct {
	Container   string `json:"container"`
	Metadata    string `json:"metadata"`
	Extensions  string `json:"extensions"`
	Operations  string `json:"operations"`
	Children    string `json:"children"`
	History     string `json:"history"`
	Ancestors   string `json:"ancestors"`
	Body        string `json:"body"`
	Version     string `json:"version"`
	Descendants string `json:"descendants"`
	Space       string `json:"space"`
}
