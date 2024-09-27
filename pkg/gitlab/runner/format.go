package runner

import (
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/strings/join"
)

func (r *Runner) Format(f *format.Settings) string {
	s := status.New(f).Integer(r.Identifier).String(
		r.formatName(),
		r.formatDescription(),
		r.hostname(),
		r.Type,
		r.formatShared(),
		r.formatStatus(f),
	).Raw(r.Raw)

	if v := r.validate(); len(v) > 0 {
		concerns := join.Comma(v)

		if f.UseColor {
			concerns = console.Yellow(concerns)
		}

		s.String(concerns)
	}

	return s.Format()
}
