package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gopgd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) ListInstances(
	w http.ResponseWriter,
	_ *http.Request,
) {
	var result []server.Instance

	for _, i := range s.store.Instances() {
		result = append(
			result,
			server.Instance{
				Name:     i.Name,
				Host:     i.Host,
				Port:     i.Port,
				Database: i.Database,
			},
		)
	}

	web.EncodeNotation(w, result)
}
