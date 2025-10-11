package web

import "github.com/funtimecoding/go-library/pkg/web/constant"

func SchemePrefix(secure bool) string {
	if secure {
		return constant.SecurePrefix
	}

	return constant.InsecurePrefix
}
