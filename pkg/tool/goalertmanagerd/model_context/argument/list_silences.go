package argument

type ListSilences struct {
	Expired bool    `json:"expired"`
	Filter  string  `json:"filter"`
	Limit   float64 `json:"limit"`
	Offset  float64 `json:"offset"`
}
