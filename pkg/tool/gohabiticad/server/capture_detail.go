package server

import (
	"errors"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
	"github.com/funtimecoding/go-library/pkg/web/detail_error"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if d, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, d.Detail)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
