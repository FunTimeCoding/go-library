package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
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
			fmt.Sprintf(
				"/pages/%s/direct-children",
				parent.Identifier,
			),
		),
		&children,
		false,
	)

	for _, p := range children.Results {
		result = append(result, c.Page(p.Id))
	}

	return result
}
