package runner

import "github.com/xanzy/go-gitlab"

func NewSlice(v []*gitlab.Runner) []*Runner {
	var result []*Runner

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
