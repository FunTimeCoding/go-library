package issue

import "github.com/xanzy/go-gitlab"

func NewSlice(v []*gitlab.Issue) []*Issue {
	var result []*Issue

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
