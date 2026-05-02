package gofix

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/errors"
	"io"
	"os"
	"testing"
)

func TestPrintResults(t *testing.T) {
	t.Run(
		"VerboseApplied",
		func(t *testing.T) {
			entries := []result{
				{Path: "pkg/foo/bar.go", Message: "de-aliased sentry → constant"},
				{Path: "pkg/foo/baz.go", Message: "renamed dirName → directoryName (4 references)"},
				{Path: "pkg/foo/baz.go", Message: "formatted call (line 42)"},
			}
			output := captureStdout(func() {
				hasBlocked := printResults(entries, false)
				assert.False(t, hasBlocked)
			})
			assert.String(
				t,
				"pkg/foo/bar.go: de-aliased sentry → constant\npkg/foo/baz.go: renamed dirName → directoryName (4 references)\npkg/foo/baz.go: formatted call (line 42)\n",
				output,
			)
		},
	)
	t.Run(
		"SummaryDeduplicates",
		func(t *testing.T) {
			entries := []result{
				{Path: "pkg/foo/bar.go", Message: "de-aliased sentry → constant"},
				{Path: "pkg/foo/baz.go", Message: "renamed dirName → directoryName (4 references)"},
				{Path: "pkg/foo/baz.go", Message: "formatted call (line 42)"},
			}
			output := captureStdout(func() {
				hasBlocked := printResults(entries, true)
				assert.False(t, hasBlocked)
			})
			assert.String(
				t,
				"pkg/foo/bar.go\npkg/foo/baz.go\n",
				output,
			)
		},
	)
	t.Run(
		"BlockedAlwaysDetailed",
		func(t *testing.T) {
			entries := []result{
				{Path: "pkg/foo/bar.go", Message: "de-aliased h → helper"},
				{Path: "pkg/foo/baz.go", Message: "cannot de-alias mark → server (local collision with \"server\")", Blocked: true},
			}
			output := captureStdout(func() {
				hasBlocked := printResults(entries, true)
				assert.True(t, hasBlocked)
			})
			assert.String(
				t,
				"pkg/foo/bar.go\npkg/foo/baz.go: cannot de-alias mark → server (local collision with \"server\")\n",
				output,
			)
		},
	)
	t.Run(
		"BlockedVerbose",
		func(t *testing.T) {
			entries := []result{
				{Path: "pkg/foo/bar.go", Message: "de-aliased h → helper"},
				{Path: "pkg/foo/baz.go", Message: "cannot rename dirName (collision)", Blocked: true},
			}
			output := captureStdout(func() {
				hasBlocked := printResults(entries, false)
				assert.True(t, hasBlocked)
			})
			assert.String(
				t,
				"pkg/foo/bar.go: de-aliased h → helper\npkg/foo/baz.go: cannot rename dirName (collision)\n",
				output,
			)
		},
	)
	t.Run(
		"EmptyNoOutput",
		func(t *testing.T) {
			output := captureStdout(func() {
				hasBlocked := printResults(nil, false)
				assert.False(t, hasBlocked)
			})
			assert.String(t, "", output)
		},
	)
}

func captureStdout(f func()) string {
	original := os.Stdout
	reader, writer, e := os.Pipe()
	errors.PanicOnError(e)
	os.Stdout = writer
	f()
	errors.PanicClose(writer)
	os.Stdout = original
	captured, e := io.ReadAll(reader)
	errors.PanicOnError(e)

	return string(captured)
}
