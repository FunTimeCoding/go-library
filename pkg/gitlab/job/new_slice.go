package job

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Job) []*Job {
	var result []*Job

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
