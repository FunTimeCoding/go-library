package user_map

func (m *Map) HasName(name string) bool {
	if _, okay := m.byName[name]; okay {
		return true
	}

	return false
}
