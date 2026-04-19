package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"net/http"
	"path/filepath"
	"strings"
)

func (s *Server) generate(
	w http.ResponseWriter,
	r *http.Request,
) {
	errors.PanicOnError(r.ParseForm())
	fileNames := r.Form["fileNames"]
	var files []string

	for _, fileName := range fileNames {
		base := filepath.Base(fileName)
		timestamp := strings.TrimSuffix(base, filepath.Ext(base))
		notationName := timestamp + "_wvw_kill.json"
		files = append(
			files,
			filepath.Join(s.eliteInsightsPath, notationName),
		)
	}

	s.parser.Generate(files, nil)
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}
