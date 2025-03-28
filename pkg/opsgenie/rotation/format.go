package rotation

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/opsgenie/compact"
)

func (r *Rotation) Format(f *option.Format) string {
	s := status.New(f).String(
		r.Name,
		string(r.Type),
		r.Start.Format("Monday"),
		r.current(),
	).Raw(r.Raw)

	for _, participant := range r.Participants {
		s.Line("    Member: %s", compact.Mail(participant.Username))
	}

	return s.Format()
}
