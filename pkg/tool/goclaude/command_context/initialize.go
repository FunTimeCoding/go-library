package command_context

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Context) Initialize(
	host string,
	port int,
) {
	r, e := client.NewClientWithResponses(
		locator.New(host).Port(port).Insecure().String(),
	)
	errors.PanicOnError(e)
	c.client = r
}
