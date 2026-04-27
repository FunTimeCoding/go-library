package habitica

import "github.com/funtimecoding/go-library/pkg/tool/gohabd/habitica/response"

type CronResult struct {
	RolledOver bool            `json:"rolled_over"`
	LastCron   string          `json:"last_cron"`
	Before     *response.Stats `json:"before,omitempty"`
	After      *response.Stats `json:"after,omitempty"`
}
