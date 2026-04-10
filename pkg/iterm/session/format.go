package session

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (s *Session) Format(f *option.Format) string {
	t := status.New(f)

	if f.HasTag(tag.Identifier) {
		t.String(s.Identifier)
	}

	t.String(s.formatTabTitle(f), s.formatJob(f))

	return t.Format()
}
