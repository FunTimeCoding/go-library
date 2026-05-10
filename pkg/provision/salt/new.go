package salt

import (
	"context"
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	host string,
	port int,
	user string,
	password string,
	o ...Option,
) *Client {
	result := &Client{context: context.Background()}

	for _, f := range o {
		f(result)
	}

	options := []salt.ClientOption{
		salt.WithEndpoint(locator.New(host).Port(port).Insecure().String()),
		salt.WithUsername(user),
		salt.WithPassword(password),
	}

	if result.authentication != "" {
		options = append(
			options,
			salt.WithAuthBackend(result.authentication),
		)
	}

	result.client = salt.NewClient(options...)
	errors.PanicOnError(result.client.Login(result.context))

	return result
}
