package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) ChildPages(
	space string,
	name string,
) []*page.Page {
	parent := c.PageBySpaceAndName(space, name)
	var result []*page.Page

	if parent == nil {
		return result
	}

	var children *response.Pages
	notation.DecodeStrict(
		c.basic.GetV2(
			c.basic.Base().Copy().Path(
				"%s/%s%s",
				constant.Page,
				parent.Identifier,
				constant.Children,
			).String(),
		),
		&children,
		false,
	)

	for _, p := range children.Results {
		result = append(result, c.Page(p.Id))
	}

	return result
}
