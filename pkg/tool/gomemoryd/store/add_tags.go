package store

func (s *Store) AddTags(
	identifier int64,
	tags []string,
) error {
	for _, tag := range tags {
		_, e := s.database.Exec(
			`INSERT OR IGNORE INTO memory_tag (memory_identifier, tag) VALUES (?, ?)`,
			identifier,
			tag,
		)

		if e != nil {
			return e
		}
	}

	return nil
}
