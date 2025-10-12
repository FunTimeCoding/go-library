package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Page(identifier string) *page.Page {
	var result *response.Page
	notation.DecodeStrict(
		c.basic.GetV2(
			c.basic.Base().Copy().Path(
				"%s/%s",
				constant.Page,
				identifier,
			).Set(constant.BodyFormat, constant.StorageFormat).String(),
		),
		&result,
		false,
	)

	return page.New(result, c.host)
}
