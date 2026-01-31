package netboot

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestLink(t *testing.T) {
	assert.String(
		t,
		"http://ftp.debian.org/debian/dists/Alfa/main/installer-Bravo/current/images/netboot/netboot.tar.gz",
		Link(strings.Alfa, strings.Bravo),
	)
}
