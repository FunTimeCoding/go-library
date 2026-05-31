package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/goalertlogd/generated/server"
)

func (s *Server) GetAlerts(
	_ context.Context,
	r server.GetAlertsRequestObject,
) (server.GetAlertsResponseObject, error) {
	records, e := s.store.ByName(r.Params.Name)

	if e != nil {
		return server.GetAlerts500JSONResponse(
			*s.captureFail(e, "failed to query alerts"),
		), nil
	}

	return server.GetAlerts200JSONResponse(toResponse(records)), nil
}
