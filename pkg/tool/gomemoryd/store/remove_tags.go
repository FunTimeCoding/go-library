package store

func (s *Store) RemoveTags(
	identifier int64,
	tags []string,
) error {
	for _, tag := range tags {
		_, e := s.database.Exec(
			`DELETE FROM memory_tag WHERE memory_identifier = ? AND tag = ?`,
			identifier,
			tag,
		)

		if e != nil {
			return e
		}
	}

	return nil
}
