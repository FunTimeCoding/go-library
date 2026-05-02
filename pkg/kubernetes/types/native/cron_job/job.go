package cron_job

import "k8s.io/api/batch/v1"

type Job struct {
	Name    string
	Cluster string
	Raw     *v1.CronJob
}
