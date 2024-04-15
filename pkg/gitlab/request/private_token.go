package request

import (
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"net/http"
)

func PrivateToken(
	r *http.Request,
	token string,
) {
	r.Header.Add(gitlab.PrivateTokenHeader, token)
}
