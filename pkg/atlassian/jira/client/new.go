package client

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	t jira.BasicAuthTransport,
	host string,
) *jira.Client {
	result, e := jira.NewClient(t.Client(), locator.New(host).String())
	errors.PanicOnError(e)

	return result
}
