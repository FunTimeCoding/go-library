package response

type RunSummary struct {
	ID                  uint   `json:"id"`
	CreatedAt           string `json:"created_at"`
	Playbook            string `json:"playbook"`
	TriggerSource       string `json:"trigger_source"`
	DurationMillisecond int64  `json:"duration_ms"`
	Status              string `json:"status"`
	GitHead             string `json:"git_head"`
}
