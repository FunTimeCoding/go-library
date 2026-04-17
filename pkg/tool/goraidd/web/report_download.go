package web

import (
	"net/http"
	"path/filepath"
)

func (s *Server) reportDownload(
	w http.ResponseWriter,
	r *http.Request,
) {
	fileName := r.PathValue("fileName")
	path := filepath.Join(s.outputPath, fileName)
	http.ServeFile(w, r, path)
}
