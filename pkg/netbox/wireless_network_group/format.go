package wireless_network_group

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (g *Group) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(g.Identifier)
	}

	s.String(g.formatName(f)).RawList(g.Raw)

	return s.Format()
}
