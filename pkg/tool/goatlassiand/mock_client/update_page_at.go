package mock_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func (c *Client) UpdatePageAt(
	identifier string,
	title string,
	markdown string,
	version int,
	message string,
) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.deleted || e.page == nil {
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	e.page.Title = title
	e.page.Body.Storage.Value = page.ToStorage(markdown)
	e.page.Version.Number = version
	e.page.Version.Message = message
	e.page.Status = constant.CurrentStatus

	return toPage(e.page), nil
}
