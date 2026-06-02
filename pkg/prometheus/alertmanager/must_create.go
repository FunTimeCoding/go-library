package alertmanager

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustCreate(
	name string,
	instance string,
	summary string,
	description string,
	expression string,
) {
	errors.PanicOnError(
		c.Create(
			name,
			instance,
			summary,
			description,
			expression,
		),
	)
}
