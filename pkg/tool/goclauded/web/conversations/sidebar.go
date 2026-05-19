package conversations

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	goclauded "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"maragu.dev/gomponents"
	"net/http"
	"strconv"
)

func (s *Server) sidebar(
	w http.ResponseWriter,
	r *http.Request,
) {
	skip := 0

	if v := r.URL.Query().Get(goclauded.Offset); v != "" {
		if n, e := strconv.Atoi(v); e == nil && n > 0 {
			skip = n
		}
	}

	sessions := s.claude.Sessions()
	aliases := s.store.AllAliases()
	limit := 30

	if skip >= len(sessions) {
		return
	}

	sessions = sessions[skip:]
	hasMore := len(sessions) > limit

	if hasMore {
		sessions = sessions[:limit]
	}

	var entries []gomponents.Node

	for i := range sessions {
		entries = append(entries, entry(sessions[i], aliases))
	}

	if hasMore {
		entries = append(entries, sentinel(skip+limit))
	}

	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(gomponents.Group(entries).Render(w))
}
