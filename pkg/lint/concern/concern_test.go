package concern

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestConcern(t *testing.T) {
	assert.NotNil(
		t,
		New(
			strings.Alfa,
			strings.Bravo,
			strings.Charlie,
			1,
			strings.Charlie,
		),
	)
}
