package job

import batch "k8s.io/api/batch/v1"

func NewSlice(
	v []batch.Job,
	cluster string,
) []*Job {
	var result []*Job

	for _, p := range v {
		result = append(result, New(&p, cluster))
	}

	return result
}
