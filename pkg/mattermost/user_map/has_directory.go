package user_map

func (m *Map) HasDirectory(name string) bool {
	if _, okay := m.byDirectory[name]; okay {
		return true
	}

	return false
}
