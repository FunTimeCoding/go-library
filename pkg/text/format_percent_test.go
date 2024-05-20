package text

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestPercent(t *testing.T) {
	assert.String(t, "A (1%)", FormatPercent("A", 1.2))
}
