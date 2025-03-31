package user

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (u *User) Format(f *option.Format) string {
	return status.New(f).String(u.Name).Format()
}
