package ssh

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"io/fs"
)

func (c *Client) ListDirectory(path string) []fs.FileInfo {
	entries, e := c.sftpClient().ReadDir(path)
	errors.PanicOnError(e)

	return entries
}
