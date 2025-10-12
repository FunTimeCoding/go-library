package code

import "github.com/google/go-github/v75/github"

func NewSlice(v []*github.CodeResult) []*Code {
	var result []*Code

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
