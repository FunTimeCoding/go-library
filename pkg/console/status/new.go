package status

import "github.com/funtimecoding/go-library/pkg/console/format"

func New(s format.Settings) *Status {
	return &Status{
		color:    s.UseColor,
		raw:      s.ShowRaw,
		extended: s.ShowExtended,
	}
}
