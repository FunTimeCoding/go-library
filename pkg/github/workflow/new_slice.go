package workflow

import "github.com/google/go-github/v69/github"

func NewSlice(v []*github.Workflow) []*Workflow {
	var result []*Workflow

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
