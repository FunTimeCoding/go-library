package store

func (s *Store) EnsurePushCollection(name string) {
	_, e := s.database.Exec(
		`INSERT INTO collection (name, path, pattern, type)
		VALUES (?, '', '', 'push')
		ON CONFLICT(name) DO NOTHING`,
		name,
	)

	if e != nil {
		panic(e)
	}
}
