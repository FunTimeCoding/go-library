package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (c *Client) ShortAlert(f constant.StringAlias) {
	c.shortAlert = f
}
