package system

import (
	"github.com/funtimecoding/go-library/pkg/system/constant"
	"strings"
)

func Expand(path string) string {
	if strings.HasPrefix(path, constant.Tilde) {
		return strings.Replace(path, constant.Tilde, Home(), 1)
	}

	return path
}
