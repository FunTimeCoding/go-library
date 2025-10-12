package confluence

import (
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
			c.basic.Base().Copy().Path(constant.Page).Set(
				constant.BodyFormat,
				constant.StorageFormat,
			).Set(
				constant.SpaceIdentifier,
				s.Identifier,
			).Set(constant.Title, name).String(),
		),
		&result,
		false,
	)

	if len(result.Results) == 1 {
		return page.New(result.Results[0], c.host)
	}

	return nil
}
