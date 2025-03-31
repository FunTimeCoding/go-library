package team_map

import "log"

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

	log.Panicf("team not found: %s", name)
}
