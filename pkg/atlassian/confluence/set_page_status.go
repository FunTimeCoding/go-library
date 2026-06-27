package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_put"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) SetPageStatus(
	identifier string,
	status string,
) (*page.Page, error) {
	var current *page.Page
	var e error

	if status == constant.CurrentStatus {
		current, e = c.DraftOverlay(identifier)
	} else {
		current, e = c.Page(identifier)
	}

	if e != nil {
		return nil, e
	}

	version := current.Raw.Version.Number + 1

	if status == constant.DraftStatus {
		version = 1
	}

	body, f := c.basic.PutV2Path(
		fmt.Sprintf("%s/%s", constant.Page, identifier),
		page_put.NewWithStatus(
			identifier,
			current.Name,
			current.Raw.Body.Storage.Value,
			version,
			"",
			status,
		).Encode(),
	)

	if f != nil {
		return nil, f
	}

	var result *response.Page
	notation.MustDecode(body, &result, false)

	return page.New(result, c.host), nil
}
