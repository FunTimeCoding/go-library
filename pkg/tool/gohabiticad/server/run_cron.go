package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) RunCron(
	w http.ResponseWriter,
	_ *http.Request,
) {
	web.EncodeNotation(w, convert.CronResult(s.habitica.MustCron()))
}
