package issue

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Issue) []*Issue {
	var result []*Issue

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
