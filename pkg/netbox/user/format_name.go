package user

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (u *User) formatName(f *option.Format) string {
	result := u.Name

	if f.UseColor {
		result = console.Cyan(result)
	}

	return result
}
