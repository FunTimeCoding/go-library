package service_client

import (
	"github.com/ctreminiom/go-atlassian/jira/sm"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New(
	locator string,
	user string,
	token string,
) *sm.Client {
	result, e := sm.New(nil, locator)
	errors.PanicOnError(e)
	result.Auth.SetBasicAuth(user, token)

	return result
}
