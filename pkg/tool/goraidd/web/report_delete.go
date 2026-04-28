package web

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"net/http"
	"path/filepath"
)

func (s *Server) reportDelete(
	w http.ResponseWriter,
	r *http.Request,
) {
	system.RemoveFile(
		filepath.Join(s.outputPath, r.PathValue("fileName")),
	)
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}
