package argument

type QueryRules struct {
	Name   string  `json:"name"`
	Group  string  `json:"group"`
	Type   string  `json:"type"`
	State  string  `json:"state"`
	Limit  float64 `json:"limit"`
	Offset float64 `json:"offset"`
}
