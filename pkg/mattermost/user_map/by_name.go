package user_map

import "github.com/mattermost/mattermost/server/public/model"

func (m *Map) ByName(name string) *model.User {
	return m.byName[name]
}
