package issue

import "github.com/google/go-github/v69/github"

func NewSlice(v []*github.Issue) []*Issue {
	var result []*Issue

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
