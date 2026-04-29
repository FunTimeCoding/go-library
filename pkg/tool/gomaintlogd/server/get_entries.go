package server

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/generated/server"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
)

func (s *Server) GetEntries(
	w http.ResponseWriter,
	_ *http.Request,
	v server.GetEntriesParams,
) {
	f := store.NewFilter()

	if v.System != nil {
		f.System = *v.System
	}

	if v.Service != nil {
		f.Service = *v.Service
	}

	if v.User != nil {
		f.User = *v.User
	}

	if v.Since != nil {
		f.Since = *v.Since
	}

	if v.Until != nil {
		f.Until = *v.Until
	}

	if v.Limit != nil {
		f.Limit = *v.Limit
	}

	entries, e := s.store.List(f)
	errors.PanicOnError(e)
	web.EncodeNotation(w, toResponse(entries))
}
