package atl

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/client"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func New(
	host string,
	port int,
	insecure bool,
) *Client {
	l := locator.New(host).Port(port)

	if insecure {
		l.Insecure()
	}

	c, e := client.NewClientWithResponses(l.String())
	errors.PanicOnError(e)

	return &Client{context: context.Background(), client: c}
}
