package issue

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Issue) Format(f *option.Format) string {
	return status.New(f).Integer(i.Identifier).String(
		i.State,
		i.Title,
	).RawList(i.Raw).Format()
}
