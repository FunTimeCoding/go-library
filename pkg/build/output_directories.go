package build

import (
	"github.com/funtimecoding/go-library/pkg/strings/contains"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/constant"
)

func OutputDirectories() []string {
	var result []string

	for _, element := range system.Subdirectories(constant.Temporary) {
		if contains.Any(
			constant.SystemArchitectures,
			system.Subdirectories(system.Join(constant.Temporary, element)),
		) {
			result = append(result, element)
		}
	}

	return result
}
