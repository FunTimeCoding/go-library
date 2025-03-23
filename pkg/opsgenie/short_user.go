package opsgenie

import "github.com/funtimecoding/go-library/pkg/opsgenie/constant"

func (c *Client) ShortUser(f constant.StringAlias) {
	c.shortUser = f
}
