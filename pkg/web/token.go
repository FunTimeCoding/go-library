package web

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Token(
	r *http.Request,
	token string,
) {
	r.Header.Set(
		constant.AuthorizationHeader,
		key_value.Space(constant.TokenPrefix, token),
	)
}
