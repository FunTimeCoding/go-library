package ssh

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) RunStrict(command string) string {
	result := c.Run(command)
	errors.PanicOnError(result.Error)

	return result.OutputString
}
