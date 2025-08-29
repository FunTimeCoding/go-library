package virtual_chassis

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (c *Chassis) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(c.Identifier)
	}

	s.String(c.formatName(f)).RawList(c.Raw)

	return s.Format()
}
