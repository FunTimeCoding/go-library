package printer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestPrinter(t *testing.T) {
	assert.NotNil(t, New(nil))
}
