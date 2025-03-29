package opsgenie

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func (c *Client) ParseDescription(f constant.ParseDescription) {
	c.parseDescription = f
}
