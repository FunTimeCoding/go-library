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
) *Client {
	x := context.Background()
	c := salt.NewClient(
		salt.WithEndpoint(locator.New(host).Port(port).Insecure().String()),
		salt.WithUsername(user),
		salt.WithPassword(password),
	)
	errors.PanicOnError(c.Login(x))

	return &Client{context: x, client: c}
}
