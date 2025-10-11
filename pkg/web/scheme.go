package web

import "github.com/funtimecoding/go-library/pkg/web/constant"

func Scheme(secure bool) string {
	if secure {
		return constant.Secure
	}

	return constant.Insecure
}
