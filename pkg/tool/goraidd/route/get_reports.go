package route

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidd/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"os"
	"strings"
)

func (h *Router) GetReports(
	w http.ResponseWriter,
	_ *http.Request,
) {
	entries, e := os.ReadDir(h.outputPath)
	errors.PanicOnError(e)
	var result []generated.ReportResponse

	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(entry.Name(), ".html") {
			continue
		}

		info, f := entry.Info()
		errors.PanicOnError(f)
		result = append(
			result,
			generated.ReportResponse{
				FileName: entry.Name(),
				Time:     info.ModTime().Format("2006-01-02 15:04:05"),
				Size:     info.Size(),
			},
		)
	}

	web.EncodeNotation(w, result)
}
