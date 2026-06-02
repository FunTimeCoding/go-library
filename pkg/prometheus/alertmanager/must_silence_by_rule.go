package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"
)

func (c *Client) MustSilenceByRule(name string) *silence.Silence {
	result, e := c.SilenceByRule(name)
	errors.PanicOnError(e)

	return result
}
