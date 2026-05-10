package response

type JobDetail struct {
	Details []Job            `json:"info"`
	Return  []map[string]any `json:"return"`
}
