package prometheus

import (
	library "github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/prometheus/constant"
)

func (c *Client) AllMetrics() []string {
	return c.LabelValues(constant.Name, []string{}, library.StartOfTime)
}
