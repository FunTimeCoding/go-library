package system_number

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (n *Number) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(n.Identifier)
	}

	s.String(n.formatName(f)).RawList(n.Raw)

	return s.Format()
}
