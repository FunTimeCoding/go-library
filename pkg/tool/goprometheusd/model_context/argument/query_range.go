package argument

type QueryRange struct {
	Query string  `json:"query"`
	Start string  `json:"start"`
	End   string  `json:"end"`
	Step  float64 `json:"step"`
}
