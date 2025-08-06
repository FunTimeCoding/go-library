package user_map

import "fmt"

func (m *Map) Mention(name string) string {
	user := m.ByName(name)

	if user == nil {
		return ""
	}

	return fmt.Sprintf("@%s", user.Username)
}
