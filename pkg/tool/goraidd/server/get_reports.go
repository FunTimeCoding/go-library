package server

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/time"
	"github.com/funtimecoding/go-library/pkg/tool/goraidd/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"net/http"
	"strings"
)

func (s *Server) GetReports(
	w http.ResponseWriter,
	_ *http.Request,
) {
	var result []server.ReportResponse

	for _, e := range system.ReadDirectory(s.outputPath) {
		if e.IsDir() || !strings.HasSuffix(
			e.Name(),
			constant.HypertextExtension,
		) {
			continue
		}

		i, f := e.Info()
		errors.PanicOnError(f)
		result = append(
			result,
			server.ReportResponse{
				FileName: e.Name(),
				Time:     i.ModTime().Format(time.DateSecond),
				Size:     i.Size(),
			},
		)
	}

	web.EncodeNotation(w, result)
}
