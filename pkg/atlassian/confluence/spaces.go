package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/go-library/pkg/notation"
	"strings"
)

func (c *Client) Spaces() []*space.Space {
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-space/#api-spaces-get
	l := c.basic.Base().Copy().Path(constant.Space).Set(
		constant.Status,
		constant.CurrentStatus,
	).String()
	var result []*response.Space

	for {
		var s *response.Spaces
		notation.DecodeStrict(c.basic.GetV2(l), &s, false)
		result = append(result, s.Results...)

		if s.Links.Next == "" {
			break
		}

		l = c.basic.Base().Copy().Path(
			strings.TrimPrefix(s.Links.Next, constant.Base),
		).String()
	}

	return space.Sort(space.NewSlice(result))
}
