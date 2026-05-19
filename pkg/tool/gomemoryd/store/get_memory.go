package store

func (s *Store) GetMemory(identifier int64) (*Memory, error) {
	row := s.database.QueryRow(
		`SELECT identifier, name, content, description, type, created_at, updated_at, is_active
		FROM memory WHERE identifier = ?`,
		identifier,
	)
	m := &Memory{}
	var active int
	e := row.Scan(
		&m.Identifier,
		&m.Name,
		&m.Content,
		&m.Description,
		&m.Type,
		&m.CreatedAt,
		&m.UpdatedAt,
		&active,
	)

	if e != nil {
		return nil, e
	}

	m.IsActive = active == 1
	m.Tags = s.tagsForMemory(identifier)

	return m, nil
}
