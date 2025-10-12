package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
	"strings"
)

func (c *Client) PagesBySpace(identifier string) []*page.Page {
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-page/#api-spaces-id-pages-get
	l := c.basic.Base().Copy().Path(
		"%s/%s%s",
		constant.Space,
		identifier,
		constant.Page,
	).Set(
		constant.BodyFormat,
		constant.StorageFormat,
	).Set(constant.Status, constant.CurrentStatus).String()
	var result []*response.Page

	for {
		var s *response.Pages
		notation.DecodeStrict(c.basic.GetV2(l), &s, false)
		result = append(result, s.Results...)

		if s.Links.Next == "" {
			break
		}

		l = c.basic.Base().Copy().Path(
			strings.TrimPrefix(s.Links.Next, constant.Base),
		).String()
	}

	return page.Sort(page.NewSlice(result, c.host))
}
