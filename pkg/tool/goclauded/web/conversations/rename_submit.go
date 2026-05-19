package conversations

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	goclauded "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"net/http"
)

func (s *Server) renameSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue("id")
	errors.PanicOnError(r.ParseForm())
	alias := r.FormValue(goclauded.Alias)
	session := s.claude.Resolve(identifier)

	if session.Identifier == "" {
		return
	}

	s.store.SetAlias(session.Identifier, alias)
	aliases := s.store.AllAliases()
	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(entry(session, aliases).Render(w))
}
