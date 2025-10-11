package web

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func Bearer(
	r *http.Request,
	token string,
) {
	r.Header.Set(
		constant.Authorization,
		key_value.Space(constant.Bearer, token),
	)
}
