package lint

import (
	"github.com/funtimecoding/go-library/pkg/system/vfs"
	"strings"
)

func packageNameOf(
	v *vfs.VFS,
	path string,
) string {
	for _, line := range strings.Split(v.Read(path), "\n") {
		if strings.HasPrefix(line, "package ") {
			return strings.TrimPrefix(line, "package ")
		}
	}

	return ""
}
