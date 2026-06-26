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
	fetchStatus := ""

	if status == constant.CurrentStatus {
		fetchStatus = constant.DraftStatus
	}

	current, e := c.pageWithStatus(identifier, fetchStatus)

	if e != nil {
		return nil, e
	}

	body, f := c.basic.PutV2Path(
		fmt.Sprintf("%s/%s", constant.Page, identifier),
		page_put.NewWithStatus(
			identifier,
			current.Name,
			current.Raw.Body.Storage.Value,
			current.Raw.Version.Number+1,
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
