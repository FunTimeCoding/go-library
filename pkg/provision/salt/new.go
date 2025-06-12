package salt

import (
	"context"
	"fmt"
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func New(
	host string,
	port int,
	user string,
	password string,
) *Client {
	x := context.Background()
	c := salt.NewClient(
		salt.WithEndpoint(
			fmt.Sprintf("%s://%s:%d", web.InsecureScheme, host, port),
		),
		salt.WithUsername(user),
		salt.WithPassword(password),
	)
	errors.PanicOnError(c.Login(x))

	return &Client{context: x, client: c}
}
