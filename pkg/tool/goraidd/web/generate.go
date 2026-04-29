package web

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
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
		notationName := join.Empty(timestamp, constant.DetailedWvWKillSuffix)
		files = append(
			files,
			filepath.Join(s.elitePath, notationName),
		)
	}

	s.parser.Generate(files, nil)
	http.Redirect(w, r, "/reports", http.StatusSeeOther)
}
