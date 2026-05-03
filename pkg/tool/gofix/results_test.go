package gofix

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"testing"
)

func TestRelativize(t *testing.T) {
	t.Run(
		"StripsPrefix",
		func(t *testing.T) {
			r := output.NewResultsWithDirectory("/Users/shiin/src/go-mint/")
			assert.String(
				t,
				"pkg/foo/bar.go",
				r.Relativize("/Users/shiin/src/go-mint/pkg/foo/bar.go"),
			)
		},
	)
	t.Run(
		"NonMatchingPassesThrough",
		func(t *testing.T) {
			r := output.NewResultsWithDirectory("/Users/shiin/src/go-mint/")
			assert.String(
				t,
				"/other/path/file.go",
				r.Relativize("/other/path/file.go"),
			)
		},
	)
	t.Run(
		"EmptyPrefixPassesThrough",
		func(t *testing.T) {
			r := output.NewResultsWithDirectory("")
			assert.String(
				t,
				"/some/path/file.go",
				r.Relativize("/some/path/file.go"),
			)
		},
	)
}
