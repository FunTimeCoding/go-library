package keepass

import "github.com/tobischo/gokeepasslib/v3"

func (c *Client) Root() *gokeepasslib.RootData {
	return c.database.Content.Root
}
