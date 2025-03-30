package treminio_client

import (
	"fmt"
	"github.com/ctreminiom/go-atlassian/v2/confluence"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func New(
	host string,
	user string,
	token string,
) *confluence.Client {
	result, e := confluence.New(
		http.DefaultClient,
		fmt.Sprintf("https://%s", host),
	)
	errors.PanicOnError(e)
	result.Auth.SetBasicAuth(user, token)

	return result
}
