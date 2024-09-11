package issue

import (
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
)

func (i *Issue) Format(s *format.Settings) string {
	return status.New(s).Integer(i.Identifier).String(
		i.State,
		i.Title,
	).Raw(i.Raw).Format()
}
