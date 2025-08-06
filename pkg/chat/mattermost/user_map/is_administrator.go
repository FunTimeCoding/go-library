package user_map

import (
	"slices"

	"github.com/mattermost/mattermost/server/public/model"
)

func (m *Map) IsAdministrator(u *model.User) bool {
	return slices.Contains(m.administrator, u.Username)
}
