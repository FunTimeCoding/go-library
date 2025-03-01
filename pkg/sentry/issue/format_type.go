package issue

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/sentry/constant"
)

func (i *Issue) formatType(s *format.Settings) string {
	if s.UseColor {
		switch i.Type {
		case constant.ErrorType:
			return console.Red(i.Type)
		}
	}

	return i.Type
}
