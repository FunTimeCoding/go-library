package argument

type Explain struct {
	SQL     string `json:"sql"`
	Analyze bool   `json:"analyze"`
}
