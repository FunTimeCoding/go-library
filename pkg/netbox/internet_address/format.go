package internet_address

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (a *Address) Format(f *option.Format) string {
	return status.New(f).String(
		a.Name,
		a.formatObjectType(f),
	).RawList(a.Raw).Format()
}
