package run

import "github.com/funtimecoding/go-library/pkg/github/repository"

func (r *Run) Repository() *repository.Repository {
	return repository.New(r.Raw.Repository)
}
