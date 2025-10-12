package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) Pages() []*page.Page {
	l := c.basic.Base().Copy().Path(constant.Page).Set(
		constant.BodyFormat,
		constant.StorageFormat,
	).String()
	var result *response.Pages
	notation.DecodeStrict(c.basic.GetV2(l), &result, false)

	return page.NewSlice(result.Results, c.host)
}
