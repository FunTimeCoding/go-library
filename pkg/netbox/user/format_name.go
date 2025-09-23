package user

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (u *User) formatName(f *option.Format) string {
	if f.UseColor {
		return console.Cyan("%s", u.Name)
	}

	return u.Name
}
