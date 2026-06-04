package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/funtimecoding/go-library/pkg/web/palette"
	"net/http"
)

func (s *Server) paletteMemories(
	w http.ResponseWriter,
	_ *http.Request,
) {
	fragment := palette.SearchFragment(
		"/palette/memories/search",
		"Search memories...",
	)
	w.Header().Set(constant.ContentType, constant.MarkupUnicode)
	errors.PanicOnError(fragment.Render(w))
}
