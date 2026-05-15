package server

import (
	"encoding/json"
	"errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/netbox-community/go-netbox/v4"
	"net/http"
)

func (s *Server) captureDetail(
	w http.ResponseWriter,
	e error,
) {
	if f, okay := errors.AsType[netbox.GenericOpenAPIError](e); okay {
		var r response.Error

		if json.Unmarshal(f.Body(), &r) == nil && r.Detail != "" {
			s.reporter.CaptureException(e)
			http.Error(w, r.Detail, http.StatusBadRequest)

			return
		}

		if message := parseValidationErrors(f.Body()); message != "" {
			s.reporter.CaptureException(e)
			http.Error(w, message, http.StatusBadRequest)

			return
		}
	}

	s.reporter.CaptureException(e)
	http.Error(w, constant.InternalError, http.StatusInternalServerError)
}
