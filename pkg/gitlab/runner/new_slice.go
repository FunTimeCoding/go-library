package runner

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Runner) []*Runner {
	var result []*Runner

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
