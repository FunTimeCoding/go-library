package mock_client

func (c *Client) GroupTabs(
	tabIdentifiers []int,
	groupIdentifier int,
	title string,
	color string,
) (int, error) {
	id := groupIdentifier

	if id == 0 {
		c.groupID++
		id = c.groupID
	}

	c.groups[id] = &group{title: title, color: color}

	for _, tabID := range tabIdentifiers {
		for i, t := range c.tabs {
			if t.Identifier == tabID {
				c.tabs[i].GroupId = id
			}
		}
	}

	return id, nil
}
