package issue

import "github.com/funtimecoding/go-library/pkg/strings/contains"

func (i *Issue) HasLabels(s ...string) bool {
	return contains.All(i.Raw.Fields.Labels, s)
}
