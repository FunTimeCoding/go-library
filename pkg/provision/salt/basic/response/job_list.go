package response

type JobList struct {
	Return []map[string]Job `json:"return"`
}
