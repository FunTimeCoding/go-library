package custom_field

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (f *Field) Format(o *option.Format) string {
	s := status.New(o)

	if o.HasTag(tag.Identifier) {
		s.Integer32(f.Identifier)
	}

	s.String(f.formatName(o)).RawList(f.Raw)

	return s.Format()
}
