package report

import "fmt"

func (r *Report) AddNote(
	s string,
	a ...any,
) {
	r.Note = append(r.Note, fmt.Sprintf(s, a...))
}
