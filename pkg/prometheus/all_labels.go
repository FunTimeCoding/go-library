package prometheus

import "github.com/funtimecoding/go-library/pkg/constant"

func (c *Client) AllLabels() []string {
	return c.MustLabelNames([]string{}, constant.StartOfTime).Values
}
