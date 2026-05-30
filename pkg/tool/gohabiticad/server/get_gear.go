package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) GetGear(
	_ context.Context,
	_ server.GetGearRequestObject,
) (server.GetGearResponseObject, error) {
	result, e := s.habitica.UserGear()

	if e != nil {
		return server.GetGear500Response{}, nil
	}

	return server.GetGear200JSONResponse(
		*convert.Gear(result),
	), nil
}
