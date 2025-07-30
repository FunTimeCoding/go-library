package user_map

import "github.com/mattermost/mattermost/server/public/model"

func (m *Map) Get() map[string]*model.User {
	return m.byDirectory
}
