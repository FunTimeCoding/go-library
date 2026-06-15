package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goqueryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetList(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetListParams,
) {
	var metadata map[string]string

	if v.Metadata != nil {
		metadata = *v.Metadata
	}

	if v.SourceType != nil && *v.SourceType != "" {
		if metadata == nil {
			metadata = map[string]string{}
		}

		metadata[constant.SourceType] = *v.SourceType
	}

	limit := 10

	if v.Limit != nil {
		limit = *v.Limit
	}

	offset := 0

	if v.Offset != nil {
		offset = *v.Offset
	}

	outcome, e := s.service.ListDocuments(
		v.Collection,
		metadata,
		limit,
		offset,
		v.Full != nil && *v.Full,
	)
	errors.PanicOnError(e)
	web.EncodeNotation(w, outcome)
}
