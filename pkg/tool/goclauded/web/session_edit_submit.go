package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/service/argument"
	"net/http"
)

func (s *Server) sessionEditSubmit(
	w http.ResponseWriter,
	r *http.Request,
) {
	identifier := r.PathValue(constant.Identifier)
	errors.PanicOnError(r.ParseForm())
	alias := r.FormValue(constant.Alias)
	description := r.FormValue(constant.Description)
	a := argument.NewEditSession()
	a.Alias = &alias
	a.Description = &description
	errors.PanicOnError(s.service.EditSession(identifier, a))
	http.Redirect(
		w,
		r,
		fmt.Sprintf("/sessions/%s", identifier),
		http.StatusSeeOther,
	)
}
