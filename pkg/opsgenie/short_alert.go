package opsgenie

import "github.com/funtimecoding/go-library/pkg/opsgenie/constant"

func (c *Client) ShortAlert(f constant.StringAlias) {
	c.shortAlert = f
}
