package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (c *Client) TagToTeam(f constant.SliceAlias) {
	c.tagToTeam = f
}
