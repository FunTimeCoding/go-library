package opsgenie

import "github.com/funtimecoding/go-library/pkg/face"

func (c *Client) ShortAlert(f face.StringAlias) {
	c.shortAlert = f
}
