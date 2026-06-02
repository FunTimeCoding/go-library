package paginate

func Slice[T any](v []T, limit int, offset int) []T {
	if offset > 0 && offset < len(v) {
		v = v[offset:]
	} else if offset >= len(v) {
		return nil
	}

	if limit > 0 && limit < len(v) {
		v = v[:limit]
	}

	return v
}
