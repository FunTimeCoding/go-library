package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/constant"
	"net/http"
	"strconv"
)

func (s *Server) delete(
	w http.ResponseWriter,
	r *http.Request,
) {
	id, e := strconv.ParseUint(r.URL.Query().Get(constant.Identifier), 10, 64)
	errors.PanicOnError(e)
	errors.PanicOnError(s.store.Delete(uint(id)))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
}
