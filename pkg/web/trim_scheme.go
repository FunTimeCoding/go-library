package web

import (
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"strings"
)

func TrimScheme(s string) string {
	if strings.HasPrefix(s, constant.SecureSchemePrefix) {
		return strings.TrimPrefix(s, constant.SecureSchemePrefix)
	}

	if strings.HasPrefix(s, constant.InsecureSchemePrefix) {
		return strings.TrimPrefix(s, constant.InsecureSchemePrefix)
	}

	return s
}
