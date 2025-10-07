package debian

import (
	"github.com/funtimecoding/go-library/pkg/debian/netboot"
	"github.com/funtimecoding/go-library/pkg/debian/preseed"
	"github.com/funtimecoding/go-library/pkg/debian/release"
	"github.com/funtimecoding/go-library/pkg/system"
)

func (c *Client) DownloadNetboot(
	r *release.Release,
	architecture string,
) {
	inputArchive := system.Join(c.workDirectory, netboot.Name(r.Codename))
	netboot.Download(r.Codename, architecture, inputArchive)
	inputConfiguration := system.Join(
		c.workDirectory,
		preseed.Name(r.Codename),
	)
	preseed.Download(r.Codename, inputConfiguration)
	temporary := system.Join(c.workDirectory, "patched_netboot")

	if system.FileExists(temporary) {
		system.Remove(temporary)
	}

	PrepareNetBoot(
		inputArchive,
		system.Join(c.workDirectory, "patched_netboot.tar.gz"),
		inputConfiguration,
		temporary,
	)
	system.Remove(temporary)
}
