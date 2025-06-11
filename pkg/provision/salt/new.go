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
	user string,
	password string,
) *Client {
	x := context.Background()
	c := salt.NewClient(
		salt.WithEndpoint(
			fmt.Sprintf("%s://%s", web.InsecureScheme, host),
		),
		salt.WithUsername(user),
		salt.WithPassword(password),
		salt.WithAuthBackend("pam"),
		salt.WithInsecure(),
	)
	errors.PanicOnError(c.Login(x))

	return &Client{context: x, client: c}
}
