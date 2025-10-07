package build

import (
	"github.com/funtimecoding/go-library/pkg/debian/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"strings"
)

func DebianPackages() []string {
	var result []string

	for _, d := range system.Files(system.WorkingDirectory()) {
		if !strings.HasSuffix(d, constant.PackageExtension) {
			continue
		}

		result = append(result, d)
	}

	return result
}
