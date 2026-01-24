package server

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
	"strings"
)

func (s *Server) authorized(r *http.Request) bool {
	t := strings.TrimPrefix(
		r.Header.Get(web.Authorization),
		fmt.Sprintf("%s ", web.Bearer),
	)

	if t == "" {
		t = r.URL.Query().Get(constant.TokenParameter)
	}

	if s.tokenAuthentication && s.token != "" && t == s.token {
		return true
	}

	if s.openAuthentication && t != "" && s.validateOpenToken(t) {
		return true
	}

	fmt.Printf("Unauthorized token: %s\n", t)

	return false
}
