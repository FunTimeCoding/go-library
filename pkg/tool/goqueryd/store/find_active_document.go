package store

func (s *Store) findActiveDocument(
	collection string,
	path string,
) *activeDocument {
	row := s.database.QueryRow(
		"SELECT identifier, hash, title FROM document WHERE collection = ? AND path = ? AND active = 1",
		collection,
		path,
	)
	var d activeDocument
	e := row.Scan(&d.identifier, &d.hash, &d.title)

	if e != nil {
		return nil
	}

	return &d
}
