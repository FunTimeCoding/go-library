package repository

import "github.com/google/go-github/v76/github"

func NewSlice(v []*github.Repository) []*Repository {
	var result []*Repository

	for _, e := range v {
		result = append(result, New(e))
	}

	return result
}
