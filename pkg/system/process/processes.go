package process

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/mitchellh/go-ps"
)

func (c *Client) Processes() []ps.Process {
	result, e := ps.Processes()
	errors.PanicOnError(e)

	return result
}
