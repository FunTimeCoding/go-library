package manufacturer

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (m *Manufacturer) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(m.Identifier)
	}

	s.String(m.formatName(f)).RawList(m.Raw)

	return s.Format()
}
