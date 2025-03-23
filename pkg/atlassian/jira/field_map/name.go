package field_map

func (m *Map) Name(identifier string) string {
	f := m.ByIdentifier(identifier)

	if f == nil {
		return Unknown
	}

	return f.Name
}
