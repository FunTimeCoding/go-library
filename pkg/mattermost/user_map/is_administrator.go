package user_map

import (
	"github.com/mattermost/mattermost/server/public/model"
	"slices"
)

func (m *Map) IsAdministrator(u *model.User) bool {
	return slices.Contains(m.administrator, u.Username)
}
