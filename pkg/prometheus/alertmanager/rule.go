package alertmanager

import "github.com/funtimecoding/go-library/pkg/prometheus/rule"

func (c *Client) Rule(name string) *rule.Rule {
	return c.Rules().Find(name)
}
