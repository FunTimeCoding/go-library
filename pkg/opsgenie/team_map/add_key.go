package team_map

func (m *Map) AddKey(
	name string,
	key string,
) {
	for _, t := range m.Teams {
		if t.Name == name {
			m.KeyByNameMap[name] = key

			return
		}
	}

	panic("team not found")
}
