package server

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/run"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/getsentry/sentry-go"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func (s *Server) PostGenerate(
	w http.ResponseWriter,
	r *http.Request,
) {
	var body generated.PostGenerateJSONRequestBody
	errors.PanicOnError(json.NewDecoder(r.Body).Decode(&body))
	date := time.Now()

	if body.Date != nil {
		date = *body.Date
	}

	s.logger.Structured("generate_request", "files", body.Files)
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
		s.parserPath,
		"TW5_parse_top_stats_detailed.py",
	)
	parser := run.New().NoPanic()
	parser.Directory = s.parserPath
	parser.Start("python", parserScript, workdir, "-d", dateArgument)

	if parser.Error != nil {
		if s.hub != nil {
			s.hub.WithScope(
				func(c *sentry.Scope) {
					c.SetContext(
						"parser",
						map[string]any{
							"output": parser.OutputString,
							"stderr": parser.ErrorString,
						},
					)
					s.hub.CaptureException(parser.Error)
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
		s.parserPath,
		"tiddler_import.py",
	)
	t := run.New().NoPanic()
	t.Directory = s.parserPath
	t.Start(
		"python",
		tiddlerScript,
		"--html",
		s.templatePath,
		"--logfolder",
		workdir,
		"--download",
		workdir,
	)

	if t.Error != nil {
		if s.hub != nil {
			s.hub.WithScope(
				func(c *sentry.Scope) {
					c.SetContext(
						"tiddler_generator",
						map[string]any{
							"output": t.OutputString,
							"stderr": t.ErrorString,
						},
					)
					s.hub.CaptureException(t.Error)
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
			entry.Name() != filepath.Base(s.templatePath) {
			reportSource = filepath.Join(workdir, entry.Name())

			break
		}
	}

	// Re-scan after tiddler generation for the output HTML
	if reportSource == "" {
		updated := system.ReadDirectory(workdir)

		for _, entry := range updated {
			if strings.HasSuffix(entry.Name(), constant.HypertextExtension) &&
				entry.Name() != filepath.Base(s.templatePath) {
				reportSource = filepath.Join(workdir, entry.Name())

				break
			}
		}
	}

	reportPath := filepath.Join(s.outputPath, reportName)
	system.Move(reportSource, reportPath)
	web.EncodeNotation(
		w,
		generated.GenerateResponse{Path: reportPath, TidCount: tidCount},
	)
}
