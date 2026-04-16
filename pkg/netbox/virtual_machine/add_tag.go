package virtual_machine

func (m *Machine) AddTag(s string) []string {
	m.Tags = append(m.Tags, s)

	return m.Tags
}
