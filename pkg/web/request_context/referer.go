package request_context

import "github.com/funtimecoding/go-library/pkg/web/location"

func (c *Context) Referer() string {
	result := c.request.Referer()

	if result == "" {
		return location.Root
	}

	return result
}
