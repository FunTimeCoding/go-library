package file_report

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestReport(t *testing.T) {
	assert.True(t, New(strings.Alfa, nil) != nil)
}
