package route

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	generated "github.com/funtimecoding/go-library/pkg/tool/goraidparsed/server"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"log/slog"
	"net/http"
	"os"
	"os/exec"
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
	parseCommand := exec.Command(
		"python",
		parserScript,
		workdir,
		"-d",
		dateArgument,
	)
	var parseOutput bytes.Buffer
	parseCommand.Dir = h.parserPath
	parseCommand.Stdout = &parseOutput
	parseCommand.Stderr = &parseOutput

	if parseError := parseCommand.Run(); parseError != nil {
		slog.Error("parser failed", "output", parseOutput.String())
		errors.PanicOnError(parseError)
	}

	slog.Info("parser completed", "output", parseOutput.String())
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

	// Copy the template to work dir so Playwright can save alongside it
	importCommand := exec.Command(
		"python",
		tiddlerScript,
		"--html",
		h.templatePath,
		"--logfolder",
		workdir,
		"--download",
		workdir,
	)
	var tiddlerOutput bytes.Buffer
	importCommand.Dir = h.parserPath
	importCommand.Stdout = &tiddlerOutput
	importCommand.Stderr = &tiddlerOutput

	if tiddlerError := importCommand.Run(); tiddlerError != nil {
		slog.Error("tiddler generator failed", "output", tiddlerOutput.String())
		errors.PanicOnError(tiddlerError)
	}

	slog.Info("tiddler generator completed", "output", tiddlerOutput.String())
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
	errors.PanicOnError(os.Rename(reportSource, reportPath))
	w.Header().Set(constant.ContentType, constant.Object)
	errors.PanicOnError(json.NewEncoder(w).Encode(generated.GenerateResponse{
		Path:     reportPath,
		TidCount: tidCount,
	}))
}
