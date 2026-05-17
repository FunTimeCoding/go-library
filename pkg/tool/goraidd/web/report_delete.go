package web

import (
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/constant"
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
	http.Redirect(w, r, constant.ReportsPath, http.StatusSeeOther)
}
