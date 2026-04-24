package issue

import "github.com/google/go-github/v85/github"

func NewSlice(v []*github.Issue) []*Issue {
	var result []*Issue

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
