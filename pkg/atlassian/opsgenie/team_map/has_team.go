package team_map

func (m *Map) HasTeam(name string) bool {
	for _, t := range m.Teams {
		if t.Name == name {
			return true
		}
	}

	return false
}
