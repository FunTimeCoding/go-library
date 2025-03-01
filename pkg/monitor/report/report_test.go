package report

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestReport(t *testing.T) {
	assert.True(t, New() != nil)
}
