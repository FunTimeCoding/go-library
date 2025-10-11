package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Page(identifier string) *page.Page {
	var result *response.Page
	notation.DecodeStrict(
		c.basic.GetV2(
			fmt.Sprintf(
				"/pages/%s?body-format=%s",
				identifier,
				constant.StorageFormat,
			),
		),
		&result,
		false,
	)

	return page.New(result, c.host)
}
