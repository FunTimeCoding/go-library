package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gw2/constant"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"path/filepath"
	"strings"
)

func (h *Router) PostGenerate(
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
			filepath.Join(h.eliteInsightsPath, notationName),
		)
	}

	result := h.parser.Generate(files, nil)
	web.ObjectHeader(w)
	_, e := w.Write([]byte(result))
	errors.PanicOnError(e)
}
