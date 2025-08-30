package data_source

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (s *Source) Format(f *option.Format) string {
	o := status.New(f)

	if f.HasTag(tag.Identifier) {
		o.Integer32(s.Identifier)
	}

	o.String(s.formatName(f)).RawList(s.Raw)

	return o.Format()
}
