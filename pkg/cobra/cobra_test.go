package cobra

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestCobra(t *testing.T) {
	assert.True(t, New("test") != nil)
}
