package alertmanager

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) MustDeleteSilence(identifier string) {
	errors.PanicOnError(c.DeleteSilence(identifier))
}
