package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/content"

func (c *Client) FavoritesKaos() []*content.Content {
	return c.SearchKaos("favorite=currentUser()")
}
