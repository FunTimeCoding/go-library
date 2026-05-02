package cron_job

import "k8s.io/api/batch/v1"

func New(
	v *v1.CronJob,
	cluster string,
) *Job {
	return &Job{Name: v.Name, Cluster: cluster, Raw: v}
}
