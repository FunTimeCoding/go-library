package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
)

func (s *Server) PostEditSession(
	_ context.Context,
	r server.PostEditSessionRequestObject,
) (server.PostEditSessionResponseObject, error) {
	a := argument.NewEditSession()
	a.Alias = r.Body.Name
	a.Description = r.Body.Description
	a.Topic = r.Body.Topic
	a.Files = r.Body.Files

	if e := s.service.EditSession(r.Body.Session, a); e != nil {
		return server.PostEditSession500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostEditSession200Response{}, nil
}
