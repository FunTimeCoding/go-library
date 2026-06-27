package mock_client

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func (c *Client) CreatePage(
	spaceIdentifier string,
	parentIdentifier string,
	title string,
	markdown string,
) (*page.Page, error) {
	return c.createWithStatus(
		spaceIdentifier,
		parentIdentifier,
		title,
		markdown,
		constant.CurrentStatus,
	)
}
