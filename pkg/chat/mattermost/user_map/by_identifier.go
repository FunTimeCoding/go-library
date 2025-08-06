package user_map

import "github.com/mattermost/mattermost/server/public/model"

func (m *Map) ByIdentifier(identifier string) *model.User {
	return m.byIdentifier[identifier]
}
