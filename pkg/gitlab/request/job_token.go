package request

import (
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"net/http"
)

func JobToken(
	r *http.Request,
	token string,
) {
	r.Header.Add(gitlab.JobTokenHeader, token)
}
