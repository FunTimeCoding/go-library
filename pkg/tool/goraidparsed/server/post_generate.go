package server

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/run"
	"github.com/funtimecoding/go-library/pkg/tool/goraidparsed/generated/server"
	"github.com/funtimecoding/go-library/pkg/web"
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
	var body server.PostGenerateJSONRequestBody
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
		s.reporter.CaptureWithContext(
			parser.Error,
			"parser",
			map[string]any{
				"output": parser.OutputString,
				"stderr": parser.ErrorString,
			},
		)
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
		s.reporter.CaptureWithContext(
			t.Error,
			"tiddler_generator",
			map[string]any{
				"output": t.OutputString,
				"stderr": t.ErrorString,
			},
		)
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
		server.GenerateResponse{Path: reportPath, TidCount: tidCount},
	)
}
