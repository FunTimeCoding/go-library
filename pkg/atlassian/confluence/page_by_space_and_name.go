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
) (*page.Page, error) {
	s, e := c.SpaceByName(spaceName)

	if e != nil {
		return nil, e
	}

	if s == nil {
		return nil, nil
	}

	body, f := c.basic.GetV2(
		c.basic.Base().Copy().Path(constant.Page).Set(
			constant.BodyFormat,
			constant.StorageFormat,
		).Set(
			constant.SpaceIdentifier,
			s.Identifier,
		).Set(constant.Title, name).String(),
	)

	if f != nil {
		return nil, f
	}

	var result *response.Pages
	notation.DecodeStrict(body, &result, false)

	if len(result.Results) == 1 {
		return page.New(result.Results[0], c.host), nil
	}

	return nil, nil
}
