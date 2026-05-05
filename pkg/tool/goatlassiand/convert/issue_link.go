package convert

type IssueLink struct {
	Identifier string `json:"identifier"`
	Direction  string `json:"direction"`
	Type       string `json:"type"`
	Key        string `json:"key"`
	Summary    string `json:"summary"`
	Status     string `json:"status"`
}
