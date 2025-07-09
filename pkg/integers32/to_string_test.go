package integers32

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToString(t *testing.T) {
	assert.String(t, "1", ToString(1))
}
