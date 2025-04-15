package issue

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Issue) []*Issue {
	var result []*Issue

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
