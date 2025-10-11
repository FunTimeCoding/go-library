package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) PageBySpaceAndName(
	spaceName string,
	name string,
) *page.Page {
	s := c.SpaceByName(spaceName)

	if s == nil {
		return nil
	}

	var result *response.Pages
	notation.DecodeStrict(
		c.basic.GetV2(
			fmt.Sprintf(
				"/pages?body-format=%s&space-id=%s&title=%s",
				constant.StorageFormat,
				s.Identifier,
				name,
			),
		),
		&result,
		false,
	)

	if len(result.Results) == 1 {
		return page.New(result.Results[0], c.host)
	}

	return nil
}
