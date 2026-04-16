package virtual_machine

func (m *Machine) RemoveTag(s string) []string {
	var tags []string

	for _, t := range m.Tags {
		if t != s {
			tags = append(tags, t)
		}
	}

	m.Tags = tags

	return m.Tags
}
