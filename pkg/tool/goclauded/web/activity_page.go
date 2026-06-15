package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"net/http"
)

func (s *Server) activityPage(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(
		s.activitySection(r.URL.Query()[constant.Kind]).Render(w),
	)
}
