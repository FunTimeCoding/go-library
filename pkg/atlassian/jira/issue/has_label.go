package issue

import "github.com/funtimecoding/go-library/pkg/strings/contains"

func (i *Issue) HasLabel(s ...string) bool {
	return contains.Any(i.Raw.Fields.Labels, s)
}
