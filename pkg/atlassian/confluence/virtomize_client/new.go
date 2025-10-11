package virtomize_client

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
	virtomize "github.com/virtomize/confluence-go-api"
)

func New(
	host string,
	user string,
	token string,
) *virtomize.API {
	result, e := virtomize.NewAPI(
		locator.New(host).Base(constant.OldBase).String(),
		user,
		token,
	)
	errors.PanicOnError(e)

	return result
}
