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

	if _, okay := m.byDirectory[name]; !okay {
		m.byDirectory[name] = u
	}

	if _, okay := m.byIdentifier[u.Id]; !okay {
		m.byIdentifier[u.Id] = u
	}

	return u
}
