package cron_job

import batch "k8s.io/api/batch/v1"

func NewSlice(
	v []batch.CronJob,
	cluster string,
) []*Job {
	var result []*Job

	for _, p := range v {
		result = append(result, New(&p, cluster))
	}

	return result
}
