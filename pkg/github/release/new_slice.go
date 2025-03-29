package release

import "github.com/google/go-github/v70/github"

func NewSlice(v []*github.RepositoryRelease) []*Release {
	var result []*Release

	for _, element := range v {
		result = append(result, New(element))
	}

	return result
}
