package runner

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Runner) []*Runner {
	var result []*Runner

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
