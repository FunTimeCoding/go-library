package convert

import (
	"github.com/funtimecoding/go-library/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "1", ToString(1))
}
