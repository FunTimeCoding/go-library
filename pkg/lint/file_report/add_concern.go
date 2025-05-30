package file_report

import "github.com/funtimecoding/go-library/pkg/lint/concern"

func (r *Report) AddConcern(
	key string,
	text string,
	path string,
	line int,
	lineText string,
) {
	r.Add(concern.New(key, text, path, line, lineText))
}
