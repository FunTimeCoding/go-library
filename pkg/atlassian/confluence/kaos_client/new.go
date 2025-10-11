package kaos_client

import (
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	host string,
	user string,
	token string,
) *kaos.API {
	result, e := kaos.NewAPI(
		locator.New(host).Base(constant.Wiki).String(),
		&Credential{User: user, Token: token},
	)
	errors.PanicOnError(e)

	return result
}
