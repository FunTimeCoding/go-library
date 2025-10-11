package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) PagesByLabel(labelIdentifier string) []*page.Page {
	var result *response.Pages
	notation.DecodeStrict(
		c.basic.GetV2(
			fmt.Sprintf(
				"/labels/%s/pages?body-format=%s",
				labelIdentifier,
				constant.StorageFormat,
			),
		),
		&result,
		false,
	)

	return page.NewSlice(result.Results, c.host)
}
