package image

import (
	"github.com/funtimecoding/go-library/pkg/console/status"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (i *Image) Format(f *option.Format) string {
	s := status.New(f).String(
		i.formatDigest(f),
		i.formatTags(f),
	).RawList(i.Raw)

	return s.Format()
}
