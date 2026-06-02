package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/prometheus/rule/rule_list"
)

func (c *Client) MustRules() *rule_list.List {
	result, e := c.Rules()
	errors.PanicOnError(e)

	return result
}
