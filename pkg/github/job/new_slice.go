package job

import "github.com/google/go-github/v69/github"

func NewSlice(v []*github.WorkflowJob) []*Job {
	var result []*Job

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
