package keepass

import "github.com/tobischo/gokeepasslib/v3"

func (c *Client) ByTitle(s string) *gokeepasslib.Entry {
	for _, groupList := range c.database.Content.Root.Groups {
		for _, entry := range groupList.Entries {
			if entry.GetTitle() == s {
				return &entry
			}
		}

		for _, group := range groupList.Groups {
			for _, entry := range group.Entries {
				if entry.GetTitle() == s {
					return &entry
				}
			}
		}
	}

	return nil
}
