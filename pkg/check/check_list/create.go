package check_list

import "github.com/funtimecoding/go-library/pkg/check/check_entry"

func (l *List) Create(
	level string,
	text string,
) {
	l.Add(check_entry.New(level, text))
}
