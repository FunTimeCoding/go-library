package authoring

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"os"
)

func (c *Client) ReadDirectory(path string) []os.FileInfo {
	result, e := c.client.ReadDir(path)
	errors.PanicOnError(e)

	return result
}
