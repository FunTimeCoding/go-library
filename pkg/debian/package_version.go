package debian

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/system"
)

func PackageVersion(
	name string,
	version string,
	packageVersion int,
	architecture string,
) string {
	system.ValidateArchitecture(architecture)

	return fmt.Sprintf(
		"%s_%s-%d_%s",
		name,
		version,
		packageVersion,
		architecture,
	)
}
