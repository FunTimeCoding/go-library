package silence

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (s *Silence) Format(f *format.Settings) string {
	t := status.New(f).String(
		s.Identifier,
		s.Match,
		fmt.Sprintf("%s ago", units.HumanDuration(s.Age())),
		s.End.Format(time.DateMinute),
		s.State,
	).Raw(s)

	return t.Format()
}
