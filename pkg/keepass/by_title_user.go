package keepass

import "github.com/tobischo/gokeepasslib/v3"

func (c *Client) ByTitleUser(
	title string,
	user string,
) *gokeepasslib.Entry {
	return ByTitleUserRecursive(c.database.Content.Root.Groups, title, user)
}
