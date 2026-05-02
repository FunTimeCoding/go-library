package gofix

type results struct {
	entries []result
}

func (r *results) add(path string, message string) {
	r.entries = append(
		r.entries,
		result{path: path, message: message},
	)
}

func (r *results) addBlocked(path string, message string) {
	r.entries = append(
		r.entries,
		result{path: path, message: message, blocked: true},
	)
}
