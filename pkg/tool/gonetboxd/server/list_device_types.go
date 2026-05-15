package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListDeviceTypes(
	w http.ResponseWriter,
	_ *http.Request,
) {
	types, e := s.client.DeviceTypes()

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	result := make([]server.DeviceType, 0, len(types))

	for _, t := range types {
		entry := server.DeviceType{
			Identifier: t.Identifier,
			Model:      t.Model,
		}

		if t.Raw.Manufacturer.Name != "" {
			entry.Manufacturer = &t.Raw.Manufacturer.Name
		}

		result = append(result, entry)
	}

	web.EncodeNotation(w, result)
}
