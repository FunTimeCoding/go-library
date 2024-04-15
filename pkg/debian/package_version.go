package debian

import "fmt"

func PackageVersion(
	name string,
	version string,
	packageVersion int,
	architecture string,
) string {
	return fmt.Sprintf(
		"%s_%s-%d_%s",
		name,
		version,
		packageVersion,
		architecture,
	)
}
