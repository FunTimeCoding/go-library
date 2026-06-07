package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostListen(
	_ context.Context,
	r server.PostListenRequestObject,
) (server.PostListenResponseObject, error) {
	listening := r.Body.Listening != nil && *r.Body.Listening

	if e := s.service.SetListening(r.Body.Callsign, listening); e != nil {
		return server.PostListen500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostListen200Response{}, nil
}
