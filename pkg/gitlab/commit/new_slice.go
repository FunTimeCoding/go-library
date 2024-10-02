package commit

import "github.com/xanzy/go-gitlab"

func NewSlice(v []*gitlab.Commit) []*Commit {
	var result []*Commit

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
