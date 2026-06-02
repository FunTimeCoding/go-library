package alertmanager

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/silence"
)

func (c *Client) MustSilences(expired bool) []*silence.Silence {
	result, e := c.Silences(expired)
	errors.PanicOnError(e)

	return result
}
