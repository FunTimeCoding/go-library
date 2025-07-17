package user_map

import "github.com/mattermost/mattermost-server/v6/model"

func (m *Map) ByDirectory(name string) *model.User {
	return m.byDirectory[name]
}
