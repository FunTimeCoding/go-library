package conversations

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	goclauded "github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"maragu.dev/gomponents"
	"maragu.dev/gomponents/html"
	"net/http"
)

func (s *Server) renameForm(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue("id")
	session := s.claude.Resolve(identifier)
	current := session.Slug

	if a := s.store.GetAlias(session.Identifier); a != "" {
		current = a
	}

	w.Header().Set(constant.ContentType, "text/html; charset=utf-8")
	errors.PanicOnError(
		html.Form(
			gomponents.Attr(
				"hx-post",
				fmt.Sprintf("/conversations/%s/rename", identifier),
			),
			gomponents.Attr(
				"hx-target",
				fmt.Sprintf("#entry-%s", identifier),
			),
			gomponents.Attr("hx-swap", "outerHTML"),
			html.Input(
				html.Class("rename-input"),
				html.Type("text"),
				html.Name(goclauded.Alias),
				html.Value(current),
				gomponents.Attr("autofocus", ""),
				gomponents.Attr("onclick", "event.stopPropagation()"),
			),
		).Render(w),
	)
}
