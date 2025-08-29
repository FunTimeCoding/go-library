package tenant

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (t *Tenant) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(t.Identifier)
	}

	s.String(t.formatName(f), t.Group).RawList(t.Raw)

	return s.Format()
}
