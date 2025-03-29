package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (c *Client) ShortUser(f constant.StringAlias) {
	c.shortUser = f
}
