package check_list

import "github.com/funtimecoding/go-library/pkg/check/check_entry"

func (l *List) All() []*check_entry.Entry {
	return l.entries
}
