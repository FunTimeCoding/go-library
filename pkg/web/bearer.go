package web

import (
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"net/http"
)

func Bearer(
	r *http.Request,
	token string,
) {
	r.Header.Set(AuthorizationHeader, key_value.Space(BearerPrefix, token))
}
