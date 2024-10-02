package issue

import "github.com/funtimecoding/go-library/pkg/console/description"

func (i *Issue) Meta() *description.Description {
	return description.New("Issue", "Issue")
}
