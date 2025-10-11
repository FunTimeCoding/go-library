package ssh

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) RunStrict(command string) string {
	r := c.Run(command)
	errors.PanicOnError(r.Error)

	return r.OutputString
}
