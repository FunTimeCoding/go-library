package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetTag(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetTagParams,
) {
	collection := ""

	if v.Collection != nil {
		collection = *v.Collection
	}

	web.EncodeNotation(
		w,
		map[string]string{
			"source_type": s.service.GetSourceType(collection, v.Path),
		},
	)
}
