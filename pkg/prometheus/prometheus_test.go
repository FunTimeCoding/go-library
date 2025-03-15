package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web"
	"testing"
)

func TestPrometheus(t *testing.T) {
	assert.True(
		t,
		New(web.Localhost, 0, false, "", "") != nil,
	)
}
