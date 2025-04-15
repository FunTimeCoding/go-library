package issue

import "github.com/funtimecoding/go-library/pkg/strings/contains"

func (i *Issue) HasAnyLabel(s ...string) bool {
	return contains.Any(i.Raw.Fields.Labels, s)
}
