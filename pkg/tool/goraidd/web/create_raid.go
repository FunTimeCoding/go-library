package web

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
)

func (s *Server) createRaid(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	filenames := r.Form["fileNames"]
	id := s.store.CreateRaid(filenames)
	http.Redirect(
		w,
		r,
		fmt.Sprintf("/raids/%d", id),
		http.StatusSeeOther,
	)
}
