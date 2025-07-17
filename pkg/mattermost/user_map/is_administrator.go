package user_map

import (
	"github.com/mattermost/mattermost-server/v6/model"
	"slices"
)

func (m *Map) IsAdministrator(u *model.User) bool {
	return slices.Contains(Administrators, u.Username)
}
