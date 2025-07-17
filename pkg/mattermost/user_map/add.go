package user_map

import "github.com/mattermost/mattermost-server/v6/model"

func (m *Map) Add(u *model.User) *model.User {
	var name string

	for k, v := range MattermostToDirectory {
		if k == u.Username {
			name = v

			break
		}
	}

	if name == "" {
		name = u.Username
	}

	if _, found := m.byDirectory[name]; !found {
		m.byDirectory[name] = u
	}

	if _, found := m.byIdentifier[u.Id]; !found {
		m.byIdentifier[u.Id] = u
	}

	return u
}
