package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v80/github"
)

func panicOnError(
	r *github.Response,
	e error,
) {
	if r == nil {
		errors.PanicOnError(e)

		return
	}

	errors.PanicOnWebError(r.Response, e)
}
