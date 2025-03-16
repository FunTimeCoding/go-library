package response

type LabelValues struct {
	Status string   `json:"status"`
	Data   []string `json:"data"`
}
