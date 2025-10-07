package debian

import (
	"github.com/funtimecoding/go-library/pkg/debian/image"
	"github.com/funtimecoding/go-library/pkg/debian/image/checksum"
	"github.com/funtimecoding/go-library/pkg/debian/release"
)

func (c *Client) DownloadImage(
	r *release.Release,
	architecture string,
) {
	image.Download(r.Version(), architecture, c.workDirectory)
	checksum.Download(r.Version(), architecture, c.workDirectory)
}
