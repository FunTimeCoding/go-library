package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/content"

func (c *Client) Favourites() []*content.Content {
	return c.Search("favorite=currentUser()")
}
