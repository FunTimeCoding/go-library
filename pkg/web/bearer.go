package web

import (
	"fmt"
	"net/http"
)

func Bearer(
	r *http.Request,
	token string,
) {
	r.Header.Set(AuthorizationHeader, fmt.Sprintf("Bearer %s", token))
}
