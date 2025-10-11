package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) GraphLink(expression string) string {
	var host string

	if c.alternateGraphHost != "" {
		host = c.alternateGraphHost
	} else {
		host = c.host
	}

	return locator.New(
		host,
	).Path(constant.Graph).Set(constant.Graph0Expression, expression).String()
}
