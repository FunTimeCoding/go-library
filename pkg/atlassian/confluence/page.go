package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"

func (c *Client) Page(identifier string) (*page.Page, error) {
	return c.pageWithStatus(identifier, "")
}
