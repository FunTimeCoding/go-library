package gmail

import "github.com/funtimecoding/go-library/pkg/system"

func (c *Client) Load() *Client {
	if c.directory != "" {
		system.EnsurePathExists(c.directory)
	}

	c.credential = c.loadCredential()
	c.service = c.loadService()

	return c
}
