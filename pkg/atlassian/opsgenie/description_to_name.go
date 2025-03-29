package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (c *Client) DescriptionToName(f constant.StringAlias) {
	c.descriptionToName = f
}
