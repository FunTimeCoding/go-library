package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func (c *Client) PagesByLabel(labelIdentifier string) []*page.Page {
	var result []*page.Page

	fmt.Println(
		c.basic.GetV2(
			fmt.Sprintf("/labels/%s/pages", labelIdentifier),
		),
	)

	return result
}
