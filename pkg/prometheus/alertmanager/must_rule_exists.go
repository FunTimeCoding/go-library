package alertmanager

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustRuleExists(name string) bool {
	result, e := c.RuleExists(name)
	errors.PanicOnError(e)

	return result
}
