package prometheus

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/client"
	"github.com/funtimecoding/go-library/pkg/prometheus/round_tripper"
	"github.com/funtimecoding/go-library/pkg/web"
)

func New(
	host string,
	port int,
	secure bool,
	user string,
	password string,
) *Client {
	errors.FatalOnEmpty(host, "host")

	return &Client{
		context: context.Background(),
		client: client.New(
			web.Link(host, port, secure),
			round_tripper.New(user, password),
		),
		host: host,
	}
}
