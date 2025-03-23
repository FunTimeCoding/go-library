package user_map

import "github.com/funtimecoding/go-library/pkg/opsgenie/user"

func (m *Map) ByIdentifier(identifier string) *user.User {
	if v, okay := m.UserMap[identifier]; okay {
		return v
	}

	return nil
}
