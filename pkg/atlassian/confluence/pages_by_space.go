package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/notation"
	"strings"
)

func (c *Client) PagesBySpace(identifier string) []*page.Page {
	var result []*response.Page
	path := fmt.Sprintf(
		"/spaces/%s/pages?body-format=%s",
		identifier,
		constant.StorageFormat,
	)
	// https://developer.atlassian.com/cloud/confluence/rest/v2/api-group-page/#api-spaces-id-pages-get
	path += "&status=current"

	for {
		var s *response.Pages
		notation.DecodeStrict(c.basic.GetV2(path), &s, false)
		result = append(result, s.Results...)

		if s.Links.Next == "" {
			break
		}

		path = strings.TrimPrefix(s.Links.Next, constant.Base)
	}

	return page.Sort(page.NewSlice(result, c.host))
}
