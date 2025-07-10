package version

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/go_mod"
	"github.com/funtimecoding/go-library/pkg/system"
)

func checkMod(
	path string,
	display string,
	installedVersion string,
	root string,
) {
	p := system.Join(path, go_mod.ModFile)

	if system.FileExists(p) {
		v := readVersion(p)

		if v != installedVersion {
			fmt.Printf("%s %s\n", system.Join(root, display), v)
		}
	}
}
