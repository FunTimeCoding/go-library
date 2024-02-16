package issue

import "github.com/google/go-github/v59/github"

func NewSlice(input []*github.Issue) []*Issue {
	var result []*Issue

	for _, element := range input {
		result = append(result, New(element))
	}

	return result
}
