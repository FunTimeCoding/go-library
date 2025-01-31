package systemd

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestSystemd(t *testing.T) {
	assert.True(t, New(nil) != nil)
}
