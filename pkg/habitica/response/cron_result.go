package response

type CronResult struct {
	RolledOver bool   `json:"rolled_over"`
	LastCron   string `json:"last_cron"`
	Before     *Stats `json:"before,omitempty"`
	After      *Stats `json:"after,omitempty"`
}
