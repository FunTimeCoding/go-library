package convert

type IssueComment struct {
	Identifier string `json:"identifier"`
	Author     string `json:"author"`
	Body       string `json:"body"`
	Created    string `json:"created"`
	Updated    string `json:"updated,omitempty"`
}
