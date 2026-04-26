package route

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/system/run"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/go-library/pkg/web"
	"log/slog"
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

	slog.Info("generate request", "files", body.Files)
	workdir, e := os.MkdirTemp("", "goraidparsed-")
	errors.PanicOnError(e)
	slog.Info("workdir created", "path", workdir)
	defer func() { errors.PanicOnError(os.RemoveAll(workdir)) }()

	for _, file := range body.Files {
		_, statError := os.Stat(file)
		exists := statError == nil
		slog.Info("symlink target", "path", file, "exists", exists)
		target := filepath.Join(workdir, filepath.Base(file))
		errors.PanicOnError(os.Symlink(file, target))
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
		slog.Error(
			"parser failed",
			"output",
			parser.OutputString,
			"stderr",
			parser.ErrorString,
		)
		errors.PanicOnError(parser.Error)
	}

	slog.Info("parser completed", "output", parser.OutputString)
	tidCount := 0
	entries, f := os.ReadDir(workdir)
	errors.PanicOnError(f)

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
		slog.Error(
			"tiddler generator failed",
			"output",
			t.OutputString,
			"stderr",
			t.ErrorString,
		)
		errors.PanicOnError(t.Error)
	}

	slog.Info("tiddler generator completed", "output", t.OutputString)
	reportName := fmt.Sprintf(
		"ONYX Log %s.html",
		date.Format("2006-01-02"),
	)
	var reportSource string

	for _, entry := range entries {
		if strings.HasSuffix(entry.Name(), ".html") &&
			entry.Name() != filepath.Base(h.templatePath) {
			reportSource = filepath.Join(workdir, entry.Name())

			break
		}
	}

	// Re-scan after tiddler generation for the output HTML
	if reportSource == "" {
		updated, g := os.ReadDir(workdir)
		errors.PanicOnError(g)

		for _, entry := range updated {
			if strings.HasSuffix(entry.Name(), ".html") &&
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
