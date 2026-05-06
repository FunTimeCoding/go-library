package response

type RunSummary struct {
	ID                  uint   `json:"id"`
	CreatedAt           string `json:"created_at"`
	TriggerSource       string `json:"trigger_source"`
	DurationMillisecond int64  `json:"duration_ms"`
	Status              string `json:"status"`
	GitHead             string `json:"git_head"`
}
