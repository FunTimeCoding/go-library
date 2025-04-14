package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) PagesBySpaceName(n string) []*page.Page {
	var result *response.Pages
	notation.DecodeStrict(
		c.basic.GetV2(
			fmt.Sprintf(
				"/spaces/%s/pages?body-format=%s",
				c.SpaceByName(n).Identifier,
				constant.StorageFormat,
			),
		),
		&result,
		false,
	)

	return page.NewSlice(result.Results, c.host)
}
