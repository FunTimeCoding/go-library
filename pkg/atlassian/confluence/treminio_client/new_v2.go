package treminio_client

import (
	"fmt"
	"github.com/ctreminiom/go-atlassian/v2/confluence/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func NewV2(
	host string,
	user string,
	token string,
) *v2.Client {
	result, e := v2.New(
		http.DefaultClient,
		fmt.Sprintf("https://%s", host),
	)
	errors.PanicOnError(e)
	result.Auth.SetBasicAuth(user, token)

	return result
}
