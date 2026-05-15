package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/common"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (s *Server) captureDetail(
	w http.ResponseWriter,
	e error,
) {
	if m, okay := common.ExtractMessage(e); okay {
		s.reporter.CaptureException(e)
		http.Error(w, m, http.StatusBadRequest)

		return
	}

	s.reporter.CaptureException(e)
	http.Error(w, constant.InternalError, http.StatusInternalServerError)
}
