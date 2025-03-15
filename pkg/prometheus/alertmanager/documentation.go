package alertmanager

import "github.com/funtimecoding/go-library/pkg/prometheus/alertmanager/constant"

func (c *Client) Documentation(name string) string {
	r := c.Rules()

	if r == nil {
		return constant.None
	}

	result := r.FindDocumentation(name)

	if result == "" {
		result = constant.None
	}

	return result
}
