package event

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (e *Event) Format(f *option.Format) string {
	return status.New(f).String(
		e.Name,
		e.Cluster,
		e.Type,
	).RawList(e.Raw).Format()
}
