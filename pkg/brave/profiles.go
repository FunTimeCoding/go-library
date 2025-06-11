package brave

import (
	"github.com/funtimecoding/go-library/pkg/brave/constant"
	"github.com/funtimecoding/go-library/pkg/system"
	"path/filepath"
	"strings"
)

func Profiles() []string {
	var result []string

	for _, f := range system.ReadDirectory(
		filepath.Join(system.Home(), constant.BraveSettings),
	) {
		if !f.IsDir() {
			continue
		}

		if f.Name() == constant.DefaultProfile ||
			strings.HasPrefix(f.Name(), constant.ProfilePrefix) {
			result = append(result, f.Name())
		}
	}

	return result
}
