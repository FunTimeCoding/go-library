package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetMetadata(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetMetadataParams,
) {
	key := ""

	if v.Key != nil {
		key = *v.Key
	}

	if key != "" {
		web.EncodeNotation(
			w,
			s.service.CollectionFacetsForKey(v.Collection, key),
		)

		return
	}

	web.EncodeNotation(w, s.service.CollectionFacets(v.Collection, nil))
}
