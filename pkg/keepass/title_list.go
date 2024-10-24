package keepass

func (c *Client) TitleList() []string {
	var result []string
	r := c.Root()

	for _, group := range r.Groups {
		for _, entry := range group.Entries {
			result = append(result, entry.GetTitle())
		}
	}

	return result
}
