package task

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Todo) []*Task {
	var result []*Task

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
