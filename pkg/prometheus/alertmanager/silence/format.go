package silence

import (
	"fmt"
	"github.com/docker/go-units"
	"github.com/funtimecoding/go-library/pkg/console/format"
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
	"github.com/funtimecoding/go-library/pkg/time"
)

func (s *Silence) Format(f *format.Settings) string {
	t := status.New(f).String(
		s.Identifier,
		s.formatRule(),
		s.Author,
		fmt.Sprintf("%s ago", units.HumanDuration(s.Age())),
	).Raw(s)

	if r := s.Remain(); r > 0 {
		t.String(fmt.Sprintf("%s remain", units.HumanDuration(r)))
	}

	t.String(s.End.Format(time.DateMinute))

	if f.HasTag(tag.State) {
		t.String(s.State)
	}

	return t.Format()
}
