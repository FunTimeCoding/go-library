package user

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (u *User) Format(f *option.Format) string {
	result := status.New(f).String(
		u.Name,
		u.Identifier,
	).Raw(u.Raw)

	if u.FullName != "" {
		result.String(u.FullName)
	}

	return result.Format()
}
