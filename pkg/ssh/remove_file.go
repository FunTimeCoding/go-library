package ssh

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) RemoveFile(path string) {
	errors.PanicOnError(c.sftpClient().Remove(path))
}
