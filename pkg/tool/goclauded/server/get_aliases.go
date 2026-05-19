package server

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetAliases(
	w http.ResponseWriter,
	_ *http.Request,
) {
	aliases := s.service.Store.ListAliases()
	var entries []server.AliasEntry

	for _, a := range aliases {
		entries = append(
			entries,
			server.AliasEntry{Session: a.SessionIdentifier, Name: a.Name},
		)
	}

	if entries == nil {
		entries = []server.AliasEntry{}
	}

	web.EncodeNotation(
		w,
		server.AliasListResponse{Aliases: entries},
	)
}
