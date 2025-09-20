package cron_job

import batch "k8s.io/api/batch/v1"

type Job struct {
	Name    string
	Cluster string
	Raw     *batch.CronJob
}
