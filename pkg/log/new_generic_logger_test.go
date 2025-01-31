package log

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestNewGenericLogger(t *testing.T) {
	assert.True(t, NewGenericLogger() != nil)
}
