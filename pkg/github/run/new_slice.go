package run

import "github.com/google/go-github/v69/github"

func NewSlice(v []*github.WorkflowRun) []*Run {
	var result []*Run

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
