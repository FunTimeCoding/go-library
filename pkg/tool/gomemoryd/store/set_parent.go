package store

func (s *Store) SetParent(
	identifier int64,
	parentIdentifier *int64,
) error {
	_, e := s.database.Exec(
		`UPDATE memory SET parent_identifier = ? WHERE identifier = ?`,
		parentIdentifier,
		identifier,
	)

	return e
}
