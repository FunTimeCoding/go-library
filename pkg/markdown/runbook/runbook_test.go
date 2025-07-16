package runbook

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestRunbook(t *testing.T) {
	assert.True(t, New(nil) != nil)
}
