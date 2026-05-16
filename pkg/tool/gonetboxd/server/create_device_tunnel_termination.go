package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/netbox/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/convert"
	"github.com/funtimecoding/go-library/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) CreateDeviceTunnelTermination(
	w http.ResponseWriter,
	q *http.Request,
	name string,
) {
	var body server.CreateTunnelTerminationRequest

	if e := json.NewDecoder(q.Body).Decode(&body); e != nil {
		web.InvalidRequestBody(w)

		return
	}

	t, e := s.client.TunnelByName(body.Tunnel)

	if e != nil {
		s.captureDetail(w, e)

		return
	}

	d, f := s.client.DeviceByNameStrict(name)

	if f != nil {
		s.captureDetail(w, f)

		return
	}

	i, g := s.client.DeviceInterfaceByNameStrict(d, body.Interface)

	if g != nil {
		s.captureDetail(w, g)

		return
	}

	tt, h := s.client.CreateTunnelTermination(
		t,
		constant.InterfaceAddress,
		int64(i.Identifier),
		body.Role,
	)

	if h != nil {
		s.captureDetail(w, h)

		return
	}

	web.ObjectHeader(w)
	w.WriteHeader(http.StatusCreated)
	web.Encode(w, convert.TunnelTermination(tt))
}
