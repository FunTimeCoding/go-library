package version

import "github.com/funtimecoding/go-library/pkg/system"

func Check(
	root string,
	skip []string,
	installedVersion string,
) {
	for _, e := range system.ReadDirectory(root) {
		if !e.IsDir() || exclude(e.Name(), skip) {
			continue
		}

		depth1 := system.Join(root, e.Name())
		checkMod(
			depth1,
			system.Join(".", e.Name()),
			installedVersion,
			root,
		)

		for _, e2 := range system.ReadDirectory(depth1) {
			if !e2.IsDir() {
				continue
			}

			relative := system.Join(e.Name(), e2.Name())

			if exclude(e2.Name(), skip) || exclude(relative, skip) {
				continue
			}

			checkMod(
				system.Join(depth1, e2.Name()),
				system.Join(".", relative),
				installedVersion,
				root,
			)
		}
	}
}
