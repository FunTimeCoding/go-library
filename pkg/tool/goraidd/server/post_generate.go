package server

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"path/filepath"
	"strings"
)

func (s *Server) PostGenerate(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body server.PostGenerateJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	var files []string

	for _, fileName := range body.FileNames {
		base := filepath.Base(fileName)
		timestamp := strings.TrimSuffix(base, filepath.Ext(base))
		notationName := join.Empty(timestamp, constant.DetailedWvWKillSuffix)
		files = append(
			files,
			filepath.Join(s.elitePath, notationName),
		)
	}

	result := s.parser.Generate(files, nil)
	web.ObjectHeader(w)
	_, e := w.Write([]byte(result))
	errors.PanicOnError(e)
}
