package dash

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestToUnderscore(t *testing.T) {
	assert.String(t, "a_b_c", ToUnderscore("a-b-c"))
}
