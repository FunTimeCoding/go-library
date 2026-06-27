package mock_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func (c *Client) SetPageStatus(
	identifier string,
	status string,
) (*page.Page, error) {
	e, okay := c.pages[identifier]

	if !okay || e.deleted {
		return nil, fmt.Errorf("page not found: %s", identifier)
	}

	if status == constant.DraftStatus {
		if e.page == nil {
			return nil, fmt.Errorf("page not found: %s", identifier)
		}

		draft := *e.page
		draft.Status = constant.DraftStatus
		draft.Version.Number = 1
		e.draft = &draft

		return toPage(e.draft), nil
	}

	if e.draft != nil {
		e.page = e.draft
		e.page.Status = constant.CurrentStatus
		e.page.Version.Number++
		e.draft = nil

		return toPage(e.page), nil
	}

	if e.page != nil && e.page.Status == constant.DraftStatus {
		e.page.Status = constant.CurrentStatus

		return toPage(e.page), nil
	}

	return nil, fmt.Errorf("no draft to publish: %s", identifier)
}
