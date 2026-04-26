package system

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestHostname(t *testing.T) {
	assert.True(t, len(Hostname()) > 0)
}
