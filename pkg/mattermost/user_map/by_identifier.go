package user_map

import "github.com/mattermost/mattermost-server/v6/model"

func (m *Map) ByIdentifier(identifier string) *model.User {
	return m.byIdentifier[identifier]
}
