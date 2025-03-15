package alertmanager

import "github.com/funtimecoding/go-library/pkg/prometheus/rule/rule_list"

func (c *Client) Rules() *rule_list.List {
	if c.rules != nil {
		return c.rules
	}

	c.rules = c.prometheus.Rules()

	return c.rules
}
