package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_put"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) UpdatePage(
	identifier string,
	title string,
	markdown string,
	message string,
) *page.Page {
	current := c.Page(identifier)
	var result *response.Page
	notation.DecodeStrict(
		c.basic.PutV2Path(
			fmt.Sprintf("%s/%s", constant.Page, identifier),
			page_put.New(
				identifier,
				title,
				page.ToStorage(markdown),
				current.Raw.Version.Number+1,
				message,
			).Encode(),
		),
		&result,
		false,
	)

	return page.New(result, c.host)
}
