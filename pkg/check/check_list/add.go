package check_list

import "github.com/funtimecoding/go-library/pkg/check/check_entry"

func (l *List) Add(v *check_entry.Entry) {
	l.entries = append(l.entries, v)
}
