package task

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Todo) []*Task {
	var result []*Task

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
