package response

type JobResult struct {
	Return  any    `json:"return"`
	Retcode int    `json:"retcode"`
	Success bool   `json:"success"`
	Out     string `json:"out,omitempty"`
}
