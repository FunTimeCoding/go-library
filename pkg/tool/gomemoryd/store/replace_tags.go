package store

func (s *Store) ReplaceTags(
	identifier int64,
	tags []string,
) error {
	t, e := s.database.Begin()

	if e != nil {
		return e
	}

	defer rollback(t)
	_, e = t.Exec(
		`DELETE FROM memory_tag WHERE memory_identifier = ?`,
		identifier,
	)

	if e != nil {
		return e
	}

	for _, tag := range tags {
		_, e = t.Exec(
			`INSERT INTO memory_tag (memory_identifier, tag) VALUES (?, ?)`,
			identifier,
			tag,
		)

		if e != nil {
			return e
		}
	}

	return t.Commit()
}
