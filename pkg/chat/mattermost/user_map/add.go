package user_map

import "github.com/mattermost/mattermost/server/public/model"

func (m *Map) Add(u *model.User) *model.User {
	var name string

	if m.userAlias != nil {
		name = m.userAlias(u.Username)
	}

	if name == "" {
		name = u.Username
	}

	if _, okay := m.byName[name]; !okay {
		m.byName[name] = u
	}

	if _, okay := m.byIdentifier[u.Id]; !okay {
		m.byIdentifier[u.Id] = u
	}

	return u
}
