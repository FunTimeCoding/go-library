package workflow

import "github.com/google/go-github/v83/github"

func NewSlice(v []*github.Workflow) []*Workflow {
	var result []*Workflow

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
