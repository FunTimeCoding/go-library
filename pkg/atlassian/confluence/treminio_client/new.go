package treminio_client

import (
	"github.com/ctreminiom/go-atlassian/v2/confluence"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	"net/http"
)

func New(
	host string,
	user string,
	token string,
) *confluence.Client {
	result, e := confluence.New(
		http.DefaultClient,
		locator.New(host).String(),
	)
	errors.PanicOnError(e)
	result.Auth.SetBasicAuth(user, token)

	return result
}
