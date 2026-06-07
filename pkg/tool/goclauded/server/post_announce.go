package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
)

func (s *Server) PostAnnounce(
	_ context.Context,
	r server.PostAnnounceRequestObject,
) (server.PostAnnounceResponseObject, error) {
	var files string

	if r.Body.Files != nil {
		files = join.NewLine(*r.Body.Files)
	}

	identifier := s.service.ResolveByCallsign(r.Body.Callsign)

	if e := s.service.Announce(
		identifier,
		r.Body.Callsign,
		r.Body.Topic,
		files,
	); e != nil {
		return server.PostAnnounce500JSONResponse(
			*s.captureFail(e, constant.UnexpectedError),
		), nil
	}

	return server.PostAnnounce200Response{}, nil
}
