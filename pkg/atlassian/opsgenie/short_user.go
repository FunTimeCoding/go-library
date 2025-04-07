package opsgenie

import "github.com/funtimecoding/go-library/pkg/face"

func (c *Client) ShortUser(f face.StringAlias) {
	c.shortUser = f
}
