package device_role

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (d *DeviceRole) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(d.Identifier)
	}

	s.String(d.formatName(f)).RawList(d.Raw)

	return s.Format()
}
