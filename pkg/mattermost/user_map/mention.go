package user_map

import "fmt"

func (m *Map) Mention(name string) string {
	user := m.ByDirectory(name)

	if user == nil {
		return ""
	}

	return fmt.Sprintf("@%s", user.Username)
}
