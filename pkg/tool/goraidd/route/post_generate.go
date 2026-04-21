package route

import (
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"path/filepath"
	"strings"
)

func (h *Router) PostGenerate(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body generated.PostGenerateJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	var files []string

	for _, fileName := range body.FileNames {
		base := filepath.Base(fileName)
		timestamp := strings.TrimSuffix(base, filepath.Ext(base))
		notationName := timestamp + "_detailed_wvw_kill.json"
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
