package virtomize_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	virtomize "github.com/virtomize/confluence-go-api"
)

func New(
	host string,
	user string,
	token string,
) *virtomize.API {
	result, e := virtomize.NewAPI(
		fmt.Sprintf("https://%s/wiki/rest/api", host),
		user,
		token,
	)
	errors.PanicOnError(e)

	return result
}
