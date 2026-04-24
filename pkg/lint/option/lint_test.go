package option

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestLint(t *testing.T) {
	assert.NotNil(t, New("", false))
}
