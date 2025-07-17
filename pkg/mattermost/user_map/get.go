package user_map

import "github.com/mattermost/mattermost-server/v6/model"

func (m *Map) Get() map[string]*model.User {
	return m.byDirectory
}
