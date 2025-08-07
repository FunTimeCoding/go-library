package prometheus

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/web/constant"
	"testing"
)

func TestPrometheus(t *testing.T) {
	assert.True(
		t,
		New(
			constant.Localhost,
			0,
			false,
			"",
			"",
			"",
		) != nil,
	)
}
