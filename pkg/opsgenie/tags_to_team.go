package opsgenie

import "github.com/funtimecoding/go-library/pkg/opsgenie/constant"

func (c *Client) TagsToTeam(f constant.SliceAlias) {
	c.tagsToTeam = f
}
