package convert

type SlimSample struct {
	Metric    map[string]string `json:"metric"`
	Value     string            `json:"value"`
	Timestamp string            `json:"timestamp"`
}
