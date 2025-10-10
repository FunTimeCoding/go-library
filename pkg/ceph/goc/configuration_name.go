package goc

import (
	"github.com/funtimecoding/go-library/pkg/strings/split"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/join"
	"strings"
)

func configurationName(
	base string,
	selected string,
) string {
	var result string

	for _, f := range system.Files(join.Absolute(base, selected)) {
		if strings.HasSuffix(f, keyringSuffix) {
			return split.Dot(f)[2]
		}
	}

	return result
}
