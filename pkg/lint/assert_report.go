package lint

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/lint/concern"
	"github.com/funtimecoding/go-library/pkg/lint/file_report"
	"testing"
)

func assertReport(
	t *testing.T,
	path string,
	concerns bool,
	c []*concern.Concern,
	fixed string,
	l *file_report.Report,
) {
	t.Helper()
	assert.String(t, path, l.FilePath)
	assert.Boolean(t, concerns, l.HasConcerns())
	assert.Any(t, c, l.Concerns)
	assert.String(t, fixed, l.Fixed)
}
