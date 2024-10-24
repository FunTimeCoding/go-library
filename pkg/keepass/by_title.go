package keepass

import "github.com/tobischo/gokeepasslib/v3"

func (c *Client) ByTitle(title string) *gokeepasslib.Entry {
	return ByTitleRecursive(c.database.Content.Root.Groups, title)
}
