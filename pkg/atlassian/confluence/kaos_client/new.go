package kaos_client

import (
	"fmt"
	kaos "github.com/essentialkaos/go-confluence/v6"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func New(
	host string,
	user string,
	token string,
) *kaos.API {
	result, e := kaos.NewAPI(
		fmt.Sprintf("https://%s/wiki", host),
		&Credential{User: user, Token: token},
	)
	errors.PanicOnError(e)

	return result
}
