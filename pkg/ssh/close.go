package ssh

import "github.com/funtimecoding/go-library/pkg/errors"

func (c *Client) Close() {
	if c.sftp != nil {
		errors.LogClose(c.sftp)
	}

	errors.LogClose(c.client)
}
