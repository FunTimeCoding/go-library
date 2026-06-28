package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"

func (c *Client) draftPagesBySpace(
	spaceIdentifier string,
) ([]*page.Page, error) {
	return c.DraftPages()
}
