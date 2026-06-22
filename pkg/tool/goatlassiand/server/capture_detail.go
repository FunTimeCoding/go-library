package server

import (
	"errors"
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/generated/server"
	"github.com/funtimecoding/go-library/pkg/web/detail_error"
)

func (s *Server) captureDetail(e error) *server.ErrorResponse {
	if j, okay := errors.AsType[*jira.Error](e); okay {
		if len(j.ErrorMessages) > 0 {
			return s.captureFail(e, j.ErrorMessages[0])
		}
	}

	if d, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, d.Detail)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
