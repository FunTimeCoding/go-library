package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"

func (c *Client) Delete(pageIdentifier string) {
	c.basic.DeleteV2(
		c.basic.Base().Copy().Path(
			"%s/%s",
			constant.Page,
			pageIdentifier,
		).String(),
	)
}
