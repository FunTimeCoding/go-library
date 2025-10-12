package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
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
		c.basic.GetV2Path(
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
