package route

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/run"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (h *Router) PostGenerate(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body generated.PostGenerateJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	date := time.Now()

	if body.Date != nil {
		date = *body.Date
	}

	h.logger.Structured("generate_request", "files", body.Files)
	workdir, e := os.MkdirTemp("", "goraidparsed-")
	errors.PanicOnError(e)
	defer func() { errors.PanicOnError(os.RemoveAll(workdir)) }()

	for _, file := range body.Files {
		system.Stat(file)
		errors.PanicOnError(
			os.Symlink(file, filepath.Join(workdir, filepath.Base(file))),
		)
	}

	dateArgument := date.Format("2006-01-02T15:04:05")
	parserScript := filepath.Join(
		h.parserPath,
		"TW5_parse_top_stats_detailed.py",
	)
	parser := run.New().NoPanic()
	parser.Directory = h.parserPath
	parser.Start("python", parserScript, workdir, "-d", dateArgument)

	if parser.Error != nil {
		if h.hub != nil {
			h.hub.WithScope(
				func(s *sentry.Scope) {
					s.SetContext(
						"parser",
						map[string]any{
							"output": parser.OutputString,
							"stderr": parser.ErrorString,
						},
					)
					h.hub.CaptureException(parser.Error)
				},
			)
		}

		errors.PanicOnError(parser.Error)
	}

	tidCount := 0
	entries := system.ReadDirectory(workdir)

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".tid") {
			tidCount++
		}
	}

	tiddlerScript := filepath.Join(
		h.parserPath,
		"tiddler_import.py",
	)
	t := run.New().NoPanic()
	t.Directory = h.parserPath
	t.Start(
		"python",
		tiddlerScript,
		"--html",
		h.templatePath,
		"--logfolder",
		workdir,
		"--download",
		workdir,
	)

	if t.Error != nil {
		if h.hub != nil {
			h.hub.WithScope(
				func(s *sentry.Scope) {
					s.SetContext(
						"tiddler_generator",
						map[string]any{
							"output": t.OutputString,
							"stderr": t.ErrorString,
						},
					)
					h.hub.CaptureException(t.Error)
				},
			)
		}

		errors.PanicOnError(t.Error)
	}

	reportName := fmt.Sprintf(
		"ONYX Log %s.html",
		date.Format("2006-01-02"),
	)
	var reportSource string

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), constant.HypertextExtension) &&
			entry.Name() != filepath.Base(h.templatePath) {
			reportSource = filepath.Join(workdir, entry.Name())

			break
		}
	}

	// Re-scan after tiddler generation for the output HTML
	if reportSource == "" {
		updated := system.ReadDirectory(workdir)

		for _, entry := range updated {
			if strings.HasSuffix(entry.Name(), constant.HypertextExtension) &&
				entry.Name() != filepath.Base(h.templatePath) {
				reportSource = filepath.Join(workdir, entry.Name())

				break
			}
		}
	}

	reportPath := filepath.Join(h.outputPath, reportName)
	system.Move(reportSource, reportPath)
	web.EncodeNotation(
		w,
		generated.GenerateResponse{Path: reportPath, TidCount: tidCount},
	)
}
