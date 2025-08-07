package salt

import (
	"context"
	"fmt"
	"github.com/daixijun/go-salt/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
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
			fmt.Sprintf("%s://%s:%d", constant.InsecureScheme, host, port),
		),
		salt.WithUsername(user),
		salt.WithPassword(password),
	)
	errors.PanicOnError(c.Login(x))

	return &Client{context: x, client: c}
}
