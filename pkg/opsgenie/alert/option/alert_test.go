package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestAlert(t *testing.T) {
	assert.True(
		t,
		New(
			nil,
			nil,
			strings.Alfa,
			nil,
			nil,
			nil,
		) != nil,
	)
}
