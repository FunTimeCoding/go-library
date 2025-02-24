package build

import (
	"github.com/funtimecoding/go-library/pkg/debian"
	"github.com/funtimecoding/go-library/pkg/system"
	"strings"
)

func DebianPackages() []string {
	var result []string

	for _, element := range system.Files(system.WorkingDirectory()) {
		if !strings.HasSuffix(element, debian.PackageExtension) {
			continue
		}

		result = append(result, element)

		return result
	}

	return result
}
