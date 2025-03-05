package branch

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Branch) []*Branch {
	var result []*Branch

	for _, s := range v {
		result = append(result, New(s))
	}

	return result
}
