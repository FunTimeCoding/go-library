package opsgenie

import "github.com/funtimecoding/go-library/pkg/opsgenie/team_map"

func (c *Client) TeamMap() *team_map.Map {
	if c.teamMap == nil {
		c.teamMap = team_map.New(c.Teams())
		c.teamMap.TagsToName(c.tagsToTeam)
	}

	return c.teamMap
}
