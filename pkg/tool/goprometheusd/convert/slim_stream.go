package convert

type SlimStream struct {
	Metric map[string]string `json:"metric"`
	Values []SlimPoint       `json:"values"`
}
