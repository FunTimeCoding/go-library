package user_map

func (m *Map) Count() int {
	return len(m.byDirectory)
}
