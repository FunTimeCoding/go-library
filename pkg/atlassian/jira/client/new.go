package client

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New(
	t jira.BasicAuthTransport,
	locator string,
) *jira.Client {
	result, e := jira.NewClient(t.Client(), locator)
	errors.PanicOnError(e)

	return result
}
