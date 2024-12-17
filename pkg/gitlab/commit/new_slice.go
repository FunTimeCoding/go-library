package commit

import "gitlab.com/gitlab-org/api/client-go"

func NewSlice(v []*gitlab.Commit) []*Commit {
	var result []*Commit

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
