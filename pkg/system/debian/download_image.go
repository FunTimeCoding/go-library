package debian

import (
	"github.com/funtimecoding/go-library/pkg/system/debian/image"
	"github.com/funtimecoding/go-library/pkg/system/debian/image/checksum"
	"github.com/funtimecoding/go-library/pkg/system/debian/release"
)

func (c *Client) DownloadImage(
	r *release.Release,
	architecture string,
) {
	image.Download(r.Version(), architecture, c.workDirectory)
	checksum.Download(r.Version(), architecture, c.workDirectory)
}
