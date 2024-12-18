package request

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"net/http"
)

func JobToken(
	r *http.Request,
	token string,
) {
	r.Header.Add(constant.JobTokenHeader, token)
}
