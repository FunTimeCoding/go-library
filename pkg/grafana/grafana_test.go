package grafana

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestGrafana(t *testing.T) {
	assert.True(t, New() != nil)
}
