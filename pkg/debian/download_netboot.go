package debian

import (
	"github.com/funtimecoding/go-library/pkg/debian/netboot"
	"github.com/funtimecoding/go-library/pkg/debian/preseed"
	"github.com/funtimecoding/go-library/pkg/debian/release"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
)

func (c *Client) DownloadNetboot(
	r *release.Release,
	architecture string,
) {
	inputArchive := join.Absolute(c.workDirectory, netboot.Name(r.Codename))
	netboot.Download(r.Codename, architecture, inputArchive)
	inputConfiguration := join.Absolute(
		c.workDirectory,
		preseed.Name(r.Codename),
	)
	preseed.Download(r.Codename, inputConfiguration)
	temporary := join.Absolute(c.workDirectory, "patched_netboot")

	if system.FileExists(temporary) {
		system.Remove(temporary)
	}

	PrepareNetBoot(
		inputArchive,
		join.Absolute(c.workDirectory, "patched_netboot.tar.gz"),
		inputConfiguration,
		temporary,
	)
	system.Remove(temporary)
}
