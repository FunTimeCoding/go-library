package team_map

import "github.com/funtimecoding/go-library/pkg/opsgenie/team"
import raw "github.com/opsgenie/opsgenie-go-sdk-v2/team"

func (m *Map) Add(
	identifier string,
	name string,
	description string,
) {
	t := team.New(
		&raw.ListedTeams{
			TeamMeta:    raw.TeamMeta{Id: identifier, Name: name},
			Description: description,
		},
	)
	m.Teams = append(m.Teams, t)
	m.TeamMap[identifier] = t
}
