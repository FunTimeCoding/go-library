package team_map

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"

func (m *Map) ByName(s string) *team.Team {
	for _, v := range m.TeamMap {
		if v.Name == s {
			return v
		}
	}

	return nil
}
