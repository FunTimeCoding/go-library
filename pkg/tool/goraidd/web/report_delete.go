package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
	"os"
	"path/filepath"
)

func (s *Server) reportDelete(
	w http.ResponseWriter,
	r *http.Request,
) {
	name := r.PathValue("fileName")
	errors.PanicOnError(os.Remove(filepath.Join(s.outputPath, name)))
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}
