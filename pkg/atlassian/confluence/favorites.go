package confluence

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func (c *Client) Favorites() []*page.Page {
	var result []*page.Page

	for _, r := range c.Search("favorite=currentUser()") {
		if r.Raw.Type == constant.PageType {
			result = append(result, c.Page(r.Raw.Id))
		}
	}

	return result
}
