package job

import batch "k8s.io/api/batch/v1"

func New(
	v *batch.Job,
	cluster string,
) *Job {
	return &Job{Name: v.Name, Cluster: cluster, Raw: v}
}
