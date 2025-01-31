package optimize_settings

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSettings(t *testing.T) {
	assert.True(t, New() != nil)
}
