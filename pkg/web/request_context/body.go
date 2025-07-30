package request_context

import "github.com/funtimecoding/go-library/pkg/system"

func (c *Context) Body() string {
	if c.bodyRead {
		return c.body
	}

	c.body = string(system.ReadAll(c.request.Body))
	c.bodyRead = true

	return c.body
}
