package task

import "github.com/xanzy/go-gitlab"

func NewSlice(v []*gitlab.Todo) []*Task {
	var result []*Task

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
