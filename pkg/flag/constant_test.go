package flag

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "filter", Filter)
	assert.String(t, "investigate", Investigate)
}
