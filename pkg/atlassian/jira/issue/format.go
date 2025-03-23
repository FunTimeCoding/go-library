package issue

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Issue) Format(f *option.Format) string {
	var description string

	if i.Description == "" {
		description = "No description"
	} else {
		description = i.Description
	}

	s := status.New(f).String(i.Key, i.Summary, description).Raw(i)
	s.Line("  %s", i.Link)

	return s.Format()
}
