package netboot

import "fmt"

func Link(
	release string,
	architecture string,
) string {
	// noinspection HttpUrlsUsage
	return fmt.Sprintf(
		"http://ftp.debian.org/debian/dists/%s/main/installer-%s/current/images/netboot/netboot.tar.gz",
		release,
		architecture,
	)
}
