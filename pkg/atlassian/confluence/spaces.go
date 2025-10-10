package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"
	"github.com/funtimecoding/go-library/pkg/notation"
	"strings"
)

func (c *Client) Spaces() []*space.Space {
	var result []*response.Space
	path := "/spaces"
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-space/#api-spaces-get
	path += "?status=current"

	for {
		var s *response.Spaces
		notation.DecodeStrict(c.basic.GetV2(path), &s, false)
		result = append(result, s.Results...)

		if s.Links.Next == "" {
			break
		}

		path = strings.TrimPrefix(s.Links.Next, constant.Base)
	}

	return space.Sort(space.NewSlice(result))
}
