package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetVersions(
	w http.ResponseWriter,
	_ *http.Request,
	params server.GetVersionsParams,
) {
	limit := 50

	if params.Limit != nil {
		limit = *params.Limit
	}

	offset := 0

	if params.Offset != nil {
		offset = *params.Offset
	}

	versions, e := s.service.VersionsSince(params.Since, limit, offset)

	if e != nil {
		http.Error(w, e.Error(), http.StatusInternalServerError)

		return
	}

	var entries []server.VersionEntry

	for _, v := range versions {
		entries = append(
			entries,
			server.VersionEntry{
				Identifier:       int(v.Identifier),
				MemoryIdentifier: int(v.MemoryIdentifier),
				Name:             v.Name,
				Description:      v.Description,
				ChangedAt:        v.ChangedAt,
				ChangeType:       v.ChangeType,
				Source:           v.Source,
			},
		)
	}

	if entries == nil {
		entries = []server.VersionEntry{}
	}

	web.EncodeNotation(w, entries)
}
