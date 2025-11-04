package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func panicOnError(
	r *jira.Response,
	e error,
) {
	errors.PanicOnWebError(r.Response, e)
}
