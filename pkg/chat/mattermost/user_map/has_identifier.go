package user_map

func (m *Map) HasIdentifier(identifier string) bool {
	if _, okay := m.byIdentifier[identifier]; okay {
		return true
	}

	return false
}
