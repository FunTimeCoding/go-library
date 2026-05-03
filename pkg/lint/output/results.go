package output

import "strings"

type Results struct {
	workingDirectory string
	Entries          []Result
}

func (r *Results) Add(
	path string,
	message string,
) {
	r.Entries = append(
		r.Entries,
		Result{Path: r.Relativize(path), Message: message},
	)
}

func (r *Results) AddBlocked(
	path string,
	message string,
) {
	r.Entries = append(
		r.Entries,
		Result{Path: r.Relativize(path), Message: message, Blocked: true},
	)
}

func (r *Results) Relativize(path string) string {
	return strings.TrimPrefix(path, r.workingDirectory)
}
