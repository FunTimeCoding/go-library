package server

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/server"
)

func (s *Server) EquipGear(
	_ context.Context,
	r server.EquipGearRequestObject,
) (server.EquipGearResponseObject, error) {
	e := s.habitica.Equip(r.Key)

	if e != nil {
		return server.EquipGear400JSONResponse(
			*s.captureFail(e, "failed to equip gear"),
		), nil
	}

	result, f := s.habitica.UserGear()

	if f != nil {
		return server.EquipGear500JSONResponse(
			*s.captureFail(f, "failed to fetch gear after equip"),
		), nil
	}

	return server.EquipGear200JSONResponse(
		*convert.Gear(result),
	), nil
}
