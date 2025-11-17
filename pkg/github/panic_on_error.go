package github

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/google/go-github/v79/github"
)

func panicOnError(
	r *github.Response,
	e error,
) {
	errors.PanicOnWebError(r.Response, e)
}
