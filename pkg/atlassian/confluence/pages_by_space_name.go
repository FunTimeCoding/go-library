package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"

func (c *Client) PagesBySpaceName(n string) []*page.Page {
	return c.PagesBySpace(c.SpaceByName(n).Identifier)
}
