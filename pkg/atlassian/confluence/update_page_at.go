package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_put"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) UpdatePageAt(
	identifier string,
	title string,
	markdown string,
	version int,
	message string,
) (*page.Page, error) {
	body, e := c.basic.PutV2Path(
		fmt.Sprintf("%s/%s", constant.Page, identifier),
		page_put.New(
			identifier,
			title,
			page.ToStorage(markdown),
			version,
			message,
		).Encode(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Page
	notation.MustDecode(body, &result, false)

	return page.New(result, c.host), nil
}
